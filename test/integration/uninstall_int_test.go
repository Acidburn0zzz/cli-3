package integration

import (
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/internal/testhelpers/e2e"
	"github.com/ActiveState/cli/internal/testhelpers/tagsuite"
	"github.com/stretchr/testify/suite"
)

type UninstallIntegrationTestSuite struct {
	tagsuite.Suite
}

func (suite *UninstallIntegrationTestSuite) TestUninstall() {
	suite.OnlyRunForTags(tagsuite.Uninstall)
	ts := e2e.New(suite.T(), true)
	defer ts.Close()

	ts.UseDistinctStateExes()

	err := fileutils.Touch(filepath.Join(ts.Dirs.Config, "config.yaml"))
	suite.Require().NoError(err, "Could not create config file")

	cp := ts.SpawnCmdWithOpts(ts.SvcExe, e2e.WithArgs("start"))
	cp.ExpectExitCode(0)

	cp = ts.Spawn("clean", "uninstall")
	cp.Expect("You are about to remove")
	cp.SendLine("y")
	if runtime.GOOS == "windows" {
		cp.ExpectLongString("Deletion of State Tool has been scheduled.")
	} else {
		cp.ExpectLongString("Successfully removed State Tool and related files")
	}
	cp.ExpectExitCode(0)

	if runtime.GOOS == "windows" {
		// Allow time for spawned script to remove directories
		time.Sleep(500 * time.Millisecond)
	}

	if fileutils.DirExists(ts.Dirs.Cache) {
		suite.Fail("Cache dir should not exist after uninstall")
	}

	if fileutils.DirExists(ts.Dirs.Config) {
		suite.Fail("Config dir should not exist after uninstall")
	}

	if fileutils.FileExists(filepath.Join(ts.Dirs.Bin, "state"+osutils.ExeExt)) {
		suite.Fail("State tool executable should not exist after uninstall")
	}

	if fileutils.FileExists(ts.SvcExe) {
		suite.Fail("State service executable should not exist after uninstall")
	}

	if fileutils.FileExists(ts.TrayExe) {
		suite.Fail("State tray executable should not exist after uninstall")
	}
}

func TestUninstallIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(UninstallIntegrationTestSuite))
}
