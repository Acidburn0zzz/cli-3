package integration

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/globaldefault"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/internal/testhelpers/e2e"
	"github.com/ActiveState/cli/internal/testhelpers/tagsuite"
	rt "github.com/ActiveState/cli/pkg/platform/runtime"
	"github.com/ActiveState/cli/pkg/platform/runtime/executor"
	"github.com/ActiveState/cli/pkg/platform/runtime/setup"
	"github.com/stretchr/testify/suite"
)

type PrepareIntegrationTestSuite struct {
	tagsuite.Suite
}

func (suite *PrepareIntegrationTestSuite) TestPrepare() {
	suite.OnlyRunForTags(tagsuite.Prepare)
	if !e2e.RunningOnCI() {
		suite.T().Skipf("Skipping TestPrepare when not running on CI or on MacOS, as it modifies PATH")
	}

	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	cp := ts.SpawnWithOpts(
		e2e.WithArgs("_prepare"),
		// e2e.AppendEnv(fmt.Sprintf("ACTIVESTATE_CLI_CONFIGDIR=%s", ts.Dirs.Work)),
	)
	cp.ExpectExitCode(0)

	isAdmin, err := osutils.IsWindowsAdmin()
	suite.Require().NoError(err, "Could not determine if we are a Windows Administrator")
	// For Windows Administrator users `state _prepare` is doing nothing now (because it doesn't make sense...)
	if isAdmin {
		return
	}
	suite.AssertConfig(filepath.Join(ts.Dirs.Cache, "bin"))
}

func (suite *PrepareIntegrationTestSuite) AssertConfig(target string) {
	if runtime.GOOS != "windows" {
		// Test bashrc
		homeDir, err := os.UserHomeDir()
		suite.Require().NoError(err)

		bashContents := fileutils.ReadFileUnsafe(filepath.Join(homeDir, ".bashrc"))
		suite.Contains(string(bashContents), constants.RCAppendDefaultStartLine, "bashrc should contain our RC Append Start line")
		suite.Contains(string(bashContents), constants.RCAppendDefaultStopLine, "bashrc should contain our RC Append Stop line")
		suite.Contains(string(bashContents), target, "bashrc should contain our target dir")
	} else {
		// Test registry
		out, err := exec.Command("reg", "query", `HKCU\Environment`, "/v", "Path").Output()
		suite.Require().NoError(err)
		suite.Contains(string(out), target, "Windows system PATH should contain our target dir")
	}
}

func (suite *PrepareIntegrationTestSuite) TestResetExecutors() {
	suite.OnlyRunForTags(tagsuite.Prepare)
	ts := e2e.New(suite.T(), true)
	err := ts.ClearCache()
	suite.Require().NoError(err)
	defer ts.Close()

	cp := ts.SpawnWithOpts(
		e2e.WithArgs("activate", "ActiveState-CLI/small-python", "--path", ts.Dirs.Work),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false"),
	)
	cp.ExpectLongString("default project?")
	cp.Send("y")
	cp.Expect("Downloading")
	cp.Expect("Installing")
	cp.Expect("activated state")

	cp.SendLine("exit")
	cp.ExpectExitCode(0)

	// Remove global executors
	cfg, err := config.NewWithDir(ts.Dirs.Config)
	suite.Require().NoError(err)
	globalExecDir := globaldefault.BinDir(cfg)
	os.RemoveAll(globalExecDir)

	// check existens of exec dir
	targetDir := rt.ProjectDirToTargetDir(ts.Dirs.Work, cfg.CachePath())
	projectExecDir := setup.ExecDir(targetDir)
	suite.DirExists(projectExecDir)

	cp = ts.Spawn("_prepare")
	cp.ExpectExitCode(0)

	suite.FileExists(filepath.Join(globalExecDir, executor.NameForExe("python3"+osutils.ExeExt)))
	suite.NoDirExists(projectExecDir)

	cp = ts.SpawnWithOpts(
		e2e.WithArgs("activate", "ActiveState-CLI/small-python", "--path", ts.Dirs.Work),
		e2e.AppendEnv("ACTIVESTATE_CLI_DISABLE_RUNTIME=false"),
	)
	cp.Expect("activated state")
	cp.SendLine("python3 --version")
	cp.Expect("Python 3.6.6")
	cp.SendLine("exit")
	cp.ExpectExitCode(0)

	// executor dir should be re-created
	suite.DirExists(projectExecDir)
}

func TestPrepareIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(PrepareIntegrationTestSuite))
}
