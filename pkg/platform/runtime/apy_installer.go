package runtime

import (
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/logging"
)

// InstallFromArchive will unpack the installer archive, locate the install script, and then use the installer
// script to install an ActivePython runtime to the configured runtime dir. Any failures
// during this process will result in a failed installation and the install-dir being removed.
func (m *MetaData) pythonRelocationDir() (string, error) {
	python, fail := locatePythonExecutable(m.Path)
	if fail != nil {
		return "", fail
	}

	prefix, fail := extractPythonRelocationPrefix(m.Path, python)
	if fail != nil {
		return "", fail
	}

	// relocate python
	return prefix, nil
}

// locatePythonExecutable will locate the path to the python binary in the runtime dir.
func locatePythonExecutable(installDir string) (string, error) {
	binPath := filepath.Join(installDir, "bin")
	python2 := filepath.Join(installDir, "bin", constants.ActivePython2Executable)
	python3 := filepath.Join(installDir, "bin", constants.ActivePython3Executable)

	var executable string
	var executablePath string
	if fileutils.FileExists(python3) {
		executable = constants.ActivePython3Executable
		executablePath = python3
	} else if fileutils.FileExists(python2) {
		executable = constants.ActivePython2Executable
		executablePath = python2
	} else {
		return "", FailRuntimeNoExecutable.New("installer_err_runtime_no_executable", binPath, constants.ActivePython2Executable, constants.ActivePython3Executable)
	}

	if !fileutils.IsExecutable(executablePath) {
		return "", FailRuntimeNotExecutable.New("installer_err_runtime_executable_not_exec", binPath, executable)
	}
	return executablePath, nil
}

// extractRelocationPrefix will extract the prefix that needs to be replaced for this installation.
func extractPythonRelocationPrefix(installDir string, python string) (string, error) {
	prefixBytes, err := exec.Command(python, "-c", "import activestate; print('\\n'.join(activestate.prefixes))").Output()
	logging.Debug("bin: %s", python)
	logging.Debug("OUTPUT: %s", string(prefixBytes))
	if err != nil {
		if _, isExitError := err.(*exec.ExitError); isExitError {
			logging.Errorf("obtaining relocation prefixes: %v : %s", err, string(prefixBytes))
			return "", FailRuntimeNoPrefixes.New("installer_err_fail_obtain_prefixes", installDir)
		}
		return "", FailRuntimeInvalid.Wrap(err)
	}
	if strings.TrimSpace(string(prefixBytes)) == "" {
		return "", FailRuntimeNoPrefixes.New("installer_err_fail_obtain_prefixes", installDir)
	}
	return strings.Split(string(prefixBytes), "\n")[0], nil
}
