package integration

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/strutils"
	"github.com/ActiveState/cli/internal/testhelpers/e2e"
	"github.com/ActiveState/cli/internal/testhelpers/tagsuite"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
)

type PushIntegrationTestSuite struct {
	tagsuite.Suite
	username string

	// some variables re-used between tests
	baseProject   string
	language      string
	extraPackage  string
	extraPackage2 string
}

func (suite *PushIntegrationTestSuite) SetupSuite() {
	suite.language = "perl@5.32.0"
	suite.baseProject = "ActiveState/Perl-5.32"
	suite.extraPackage = "JSON"
	suite.extraPackage2 = "DateTime"
	if runtime.GOOS == "darwin" {
		suite.language = "python3"
		suite.baseProject = "ActiveState-CLI/small-python"
		suite.extraPackage = "six@1.10.0"
	}
}

func (suite *PushIntegrationTestSuite) TestInitAndPush() {
	suite.OnlyRunForTags(tagsuite.Push)
	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	ts.LoginAsPersistentUser()
	username := "cli-integration-tests"
	pname := strutils.UUID()
	namespace := fmt.Sprintf("%s/%s", username, pname)
	cp := ts.Spawn(
		"init",
		namespace,
		suite.language,
		"--path", filepath.Join(ts.Dirs.Work, namespace),
		"--skeleton", "editor",
	)
	cp.ExpectExitCode(0)

	wd := filepath.Join(cp.WorkDirectory(), namespace)
	cp = ts.SpawnWithOpts(e2e.WithArgs("push"), e2e.WithWorkDirectory(wd))
	cp.ExpectLongString("Creating project")
	cp.ExpectLongString("Project created")
	cp.ExpectExitCode(0)

	// Check that languages were reset
	pjfilepath := filepath.Join(ts.Dirs.Work, namespace, constants.ConfigFileName)
	pjfile, err := projectfile.Parse(pjfilepath)
	suite.Require().NoError(err)
	if pjfile.Languages != nil {
		suite.FailNow("Expected languages to be nil, but got: %v", pjfile.Languages)
	}
	if pjfile.CommitID() == "" {
		suite.FailNow("commitID was not set after running push for project creation")
	}
	if pjfile.BranchName() == "" {
		suite.FailNow("branch was not set after running push for project creation")
	}

	// ensure that we are logged out
	cp = ts.Spawn("auth", "logout")
	cp.ExpectExitCode(0)

	cp = ts.SpawnWithOpts(e2e.WithArgs("install", suite.extraPackage), e2e.WithWorkDirectory(wd))
	switch runtime.GOOS {
	case "darwin":
		cp.ExpectRe("added|currently building", 60*time.Second) // while cold storage is off
		cp.Wait()
	default:
		cp.Expect("added", 60*time.Second)
		cp.ExpectExitCode(0)
	}

	pjfile, err = projectfile.Parse(pjfilepath)
	suite.Require().NoError(err)
	if !strings.Contains(pjfile.Project, fmt.Sprintf("/%s?", namespace)) {
		suite.FailNow("project field should include project (not headless): " + pjfile.Project)
	}

	ts.LoginAsPersistentUser()

	cp = ts.SpawnWithOpts(e2e.WithArgs("push", namespace), e2e.WithWorkDirectory(wd))
	cp.Expect("Pushing to project")
	cp.ExpectExitCode(0)
}

func (suite *PushIntegrationTestSuite) TestPush_HeadlessConvert() {
	suite.OnlyRunForTags(tagsuite.Push)
	// TODO: Re-enable and tweak test once https://www.pivotaltracker.com/story/show/178223554 is ready
	suite.T().Skip("This test is currently not relevant due to local environment mangement")
	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	ts.LoginAsPersistentUser()
	username := "cli-integration-tests"
	pname := strutils.UUID()
	pname2 := strutils.UUID()
	namespace := fmt.Sprintf("%s/%s", username, pname)
	namespace2 := fmt.Sprintf("%s/%s", username, pname2)
	cp := ts.Spawn(
		"init",
		namespace,
		suite.language,
		"--path", filepath.Join(ts.Dirs.Work, namespace),
		"--skeleton", "editor",
	)
	cp.ExpectExitCode(0)

	wd := filepath.Join(cp.WorkDirectory(), namespace)
	cp = ts.SpawnWithOpts(e2e.WithArgs("push"), e2e.WithWorkDirectory(wd))
	cp.ExpectLongString(fmt.Sprintf("Project created at https://%s/%s/%s", constants.PlatformURL, username, pname))
	cp.ExpectExitCode(0)

	// Check that languages were reset
	pjfilepath := filepath.Join(ts.Dirs.Work, namespace, constants.ConfigFileName)
	pjfile, err := projectfile.Parse(pjfilepath)
	suite.Require().NoError(err)
	if pjfile.Languages != nil {
		suite.FailNow("Expected languages to be nil, but got: %v", pjfile.Languages)
	}
	if pjfile.CommitID() == "" {
		suite.FailNow("commitID was not set after running push for project creation")
	}
	if pjfile.BranchName() == "" {
		suite.FailNow("branch was not set after running push for project creation")
	}

	// ensure that we are logged out
	cp = ts.Spawn("auth", "logout")
	cp.ExpectExitCode(0)

	cp = ts.SpawnWithOpts(e2e.WithArgs("install", suite.extraPackage), e2e.WithWorkDirectory(wd))
	cp.Expect("You're about to add packages as an anonymous user")
	cp.Expect("(Y/n)")
	cp.Send("y")
	switch runtime.GOOS {
	case "darwin":
		cp.ExpectRe("added|currently building", 60*time.Second) // while cold storage is off
		cp.Wait()
	default:
		cp.Expect("added", 60*time.Second)
		cp.ExpectExitCode(0)
	}

	pjfile, err = projectfile.Parse(pjfilepath)
	suite.Require().NoError(err)
	if !strings.Contains(pjfile.Project, "/commit/") {
		suite.FailNow("project field should be headless but isn't: " + pjfile.Project)
	}

	ts.LoginAsPersistentUser()

	cp = ts.SpawnWithOpts(e2e.WithArgs("push"), e2e.WithWorkDirectory(wd))
	cp.ExpectLongString("Who would you like the owner of this project to be?")
	cp.Send("")
	cp.ExpectLongString("What would you like the name of this project to be?")
	cp.SendUnterminated(string([]byte{0033, '[', 'B'})) // move cursor down, and then press enter
	cp.Expect("> Other")
	cp.Send("")
	cp.Expect(">")
	cp.Send(pname2.String())
	cp.Expect("Project created")
	cp.ExpectExitCode(0)

	pjfile, err = projectfile.Parse(pjfilepath)
	suite.Require().NoError(err)
	if !strings.Contains(pjfile.Project, fmt.Sprintf("/%s?", namespace2)) {
		suite.FailNow("project field should include project again: " + pjfile.Project)
	}
}

func (suite *PushIntegrationTestSuite) TestCarlisle() {
	suite.OnlyRunForTags(tagsuite.Push, tagsuite.Carlisle, tagsuite.Headless)
	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	username := "cli-integration-tests"
	pname := strutils.UUID()
	namespace := fmt.Sprintf("%s/%s", username, pname)

	cp := ts.SpawnWithOpts(
		e2e.WithArgs(
			"activate", suite.baseProject,
			"--path", filepath.Join(ts.Dirs.Work, namespace)),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false"),
	)
	cp.ExpectLongString("default project?")
	cp.Send("n")
	// The activestate.yaml on Windows runs custom activation to set shortcuts and file associations.
	if runtime.GOOS == "windows" {
		cp.Expect("Running Activation Events")
	} else {
		cp.Expect("You're Activated!")
	}
	cp.SendLine("exit")
	cp.ExpectExitCode(0)

	// ensure that we are logged out
	cp = ts.Spawn("auth", "logout")
	cp.ExpectExitCode(0)

	// anonymous commit
	wd := filepath.Join(cp.WorkDirectory(), namespace)
	cp = ts.SpawnWithOpts(e2e.WithArgs(
		"install", suite.extraPackage),
		e2e.WithWorkDirectory(wd),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false", "VERBOSE=true"))
	switch runtime.GOOS {
	case "darwin":
		cp.ExpectRe("added|currently building", 60*time.Second) // while cold storage is off
		cp.Wait()
	default:
		cp.Expect("added", 60*time.Second)
		cp.ExpectExitCode(0)
	}

	prj, err := project.FromPath(filepath.Join(wd, constants.ConfigFileName))
	suite.Require().NoError(err, "Could not parse project file")
	suite.Assert().False(prj.IsHeadless(), "project should NOT be headless: URL is %s", prj.URL())

	ts.LoginAsPersistentUser()

	cp = ts.SpawnWithOpts(e2e.WithArgs("push", namespace), e2e.WithWorkDirectory(wd))
	cp.Expect("Project created")
	cp.ExpectExitCode(0)
}

func (suite *PushIntegrationTestSuite) TestPush_Outdated() {
	suite.OnlyRunForTags(tagsuite.Push)
	projectLine := "project: https://platform.activestate.com/ActiveState-CLI/cli?branch=main&commitID="
	unPushedCommit := "882ae76e-fbb7-4989-acc9-9a8b87d49388"

	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	wd := filepath.Join(ts.Dirs.Work, namespace)
	pjfilepath := filepath.Join(ts.Dirs.Work, namespace, constants.ConfigFileName)
	err := fileutils.WriteFile(pjfilepath, []byte(projectLine+unPushedCommit))
	suite.Require().NoError(err)

	ts.LoginAsPersistentUser()
	cp := ts.SpawnWithOpts(e2e.WithArgs("push"), e2e.WithWorkDirectory(wd))
	cp.ExpectLongString("Your project has new changes available")
	cp.ExpectExitCode(1)
}

func (suite *PushIntegrationTestSuite) TestPush_AlreadyExists() {
	suite.OnlyRunForTags(tagsuite.Push)
	ts := e2e.New(suite.T(), false)
	defer ts.Close()
	ts.LoginAsPersistentUser()
	username := "cli-integration-tests"
	namespace := fmt.Sprintf("%s/%s", username, "Python3")
	cp := ts.Spawn(
		"init",
		namespace,
		"python3",
		"--path", filepath.Join(ts.Dirs.Work, namespace),
		"--skeleton", "editor",
	)
	cp.ExpectExitCode(0)
	wd := filepath.Join(cp.WorkDirectory(), namespace)
	cp = ts.SpawnWithOpts(e2e.WithArgs("push"), e2e.WithWorkDirectory(wd))
	cp.ExpectLongString(fmt.Sprintf("The project %s/%s already exists", username, "Python3"))
	cp.ExpectExitCode(1)
}

func TestPushIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(PushIntegrationTestSuite))
}
