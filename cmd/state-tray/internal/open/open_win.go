//+build windows

package open

import (
	"os/exec"

	"github.com/ActiveState/cli/internal/locale"
)

// Prompt will open the a command prompt and execute the given command string
func Prompt(command string) error {
	// start will open an instance of the given executable. The first parameter
	// of start is the title, the second is the executable to start.
	// The /k flag tells the cmd prompt to stay open after the command
	// has completed execution.
	cmd := exec.Command("cmd.exe", "/c", "start", "", "cmd.exe", "/k", command)
	err := cmd.Run()
	if err != nil {
		return locale.WrapError(err, "err_open_prompt", "Could not open prompt")
	}

	return nil
}