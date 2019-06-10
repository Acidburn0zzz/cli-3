// +build !darwin

package runtime_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	rt "runtime"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/fileutils"
	hcMock "github.com/ActiveState/cli/pkg/platform/api/headchef/mock"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/platform/runtime"
	rtMock "github.com/ActiveState/cli/pkg/platform/runtime/mock"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
	"github.com/ActiveState/sysinfo"
)

type RuntimeDLTestSuite struct {
	suite.Suite

	project *project.Project
	dir     string

	hcMock *hcMock.Mock
	rtMock *rtMock.Mock
}

func (suite *RuntimeDLTestSuite) BeforeTest(suiteName, testName string) {
	projectURL := fmt.Sprintf("https://%s/%s/%s?commitID=%s", constants.PlatformURL, "string", "string", "imacommithash")
	pj := &projectfile.Project{Project: projectURL}
	suite.project = project.New(pj)

	var err error
	suite.dir, err = ioutil.TempDir("", "runtime-test")
	suite.Require().NoError(err)

	suite.hcMock = hcMock.Init()
	suite.rtMock = rtMock.Init()

	suite.rtMock.MockFullRuntime()

	pjfile := projectfile.Project{
		Project: projectURL,
	}
	pjfile.Persist()

	cachePath := config.CachePath()
	if fileutils.DirExists(cachePath) {
		err := os.RemoveAll(config.CachePath())
		suite.Require().NoError(err)
	}

	// Only linux is supported for now, so force it so we can run this test on mac
	// If we want to skip this on mac it should be skipped through build tags, in
	// which case this tweak is meaningless and only a convenience for when testing manually
	if rt.GOOS == "darwin" {
		model.OS = sysinfo.Linux
	}
}

func (suite *RuntimeDLTestSuite) AfterTest(suiteName, testName string) {
	suite.rtMock.Close()
	suite.hcMock.Close()

	err := os.RemoveAll(suite.dir)
	suite.Require().NoError(err)
}

func (suite *RuntimeDLTestSuite) TestGetRuntimeDL() {
	r := runtime.NewDownload(suite.project, suite.dir, suite.hcMock.Requester(hcMock.NoOptions))
	urls, fail := r.FetchArtifactURLs()
	suite.Require().NoError(fail.ToError())
	filenames, fail := r.Download(urls)
	suite.Require().NoError(fail.ToError())

	suite.Implements((*runtime.Downloader)(nil), r)
	suite.Contains(filenames, "python"+runtime.InstallerExtension)
	suite.Contains(filenames, "legacy-python"+runtime.InstallerExtension)

	for _, filename := range filenames {
		suite.FileExists(filepath.Join(suite.dir, filename))
	}
}

func (suite *RuntimeDLTestSuite) TestGetRuntimeDLNoArtifacts() {
	r := runtime.NewDownload(suite.project, suite.dir, suite.hcMock.Requester(hcMock.NoArtifacts))
	_, fail := r.FetchArtifactURLs()
	suite.Require().Error(fail.ToError())

	suite.Equal(runtime.FailNoArtifacts.Name, fail.Type.Name)
}

func (suite *RuntimeDLTestSuite) TestGetRuntimeDLInvalidArtifact() {
	r := runtime.NewDownload(suite.project, suite.dir, suite.hcMock.Requester(hcMock.InvalidArtifact))
	_, fail := r.FetchArtifactURLs()
	suite.Require().Error(fail.ToError())

	suite.Equal(runtime.FailNoValidArtifact.Name, fail.Type.Name)
}

func (suite *RuntimeDLTestSuite) TestGetRuntimeDLInvalidURL() {
	r := runtime.NewDownload(suite.project, suite.dir, suite.hcMock.Requester(hcMock.InvalidURL))
	urls, fail := r.FetchArtifactURLs()
	suite.Require().NoError(fail.ToError())
	_, fail = r.Download(urls)
	suite.Require().Error(fail.ToError())

	suite.Equal(model.FailSignS3URL.Name, fail.Type.Name)
}

func (suite *RuntimeDLTestSuite) TestGetRuntimeDLBuildFailure() {
	r := runtime.NewDownload(suite.project, suite.dir, suite.hcMock.Requester(hcMock.BuildFailure))
	_, fail := r.FetchArtifactURLs()
	suite.Require().Error(fail.ToError())

	suite.Equal(runtime.FailBuild.Name, fail.Type.Name)
}

func (suite *RuntimeDLTestSuite) TestGetRuntimeDLFailure() {
	r := runtime.NewDownload(suite.project, suite.dir, suite.hcMock.Requester(hcMock.RegularFailure))
	_, fail := r.FetchArtifactURLs()
	suite.Require().Error(fail.ToError())

	suite.Equal(failures.FailDeveloper.Name, fail.Type.Name)
}

func TestRuntimeDLSuite(t *testing.T) {
	suite.Run(t, new(RuntimeDLTestSuite))
}
