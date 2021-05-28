package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/ActiveState/cli/internal/appinfo"
	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/events"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/internal/runbits"
	"github.com/ActiveState/cli/internal/subshell"
	"github.com/ActiveState/cli/internal/subshell/sscommon"
	"github.com/ActiveState/cli/internal/updater"
	"github.com/rollbar/rollbar-go"
)

func main() {
	var exitCode int
	defer func() {
		if runbits.HandlePanics() {
			exitCode = 1
		}
		events.WaitForEvents(1*time.Second, rollbar.Close)
		os.Exit(exitCode)
	}()

	verbose := os.Getenv("VERBOSE") != ""
	logging.CurrentHandler().SetVerbose(verbose)
	logging.SetupRollbar(constants.StateToolRollbarToken)

	if err := run(); err != nil {
		logging.Error(fmt.Sprintf("%s failed with error: %s", filepath.Base(os.Args[0]), errs.Join(err, ": ")))
		fmt.Println(errs.Join(err, ": ").Error())

		exitCode = 1
		return
	}
}

func removeOldStateToolEnvironmentSettings(cfg *config.Instance) error {
	isAdmin, err := osutils.IsWindowsAdmin()
	if err != nil {
		return errs.Wrap(err, "Could not determine if running as Windows administrator")
	}

	// remove shell file additions
	s := subshell.New(cfg)
	if err := s.CleanUserEnv(cfg, sscommon.InstallID, isAdmin); err != nil {
		return errs.Wrap(err, "Failed to remove environment variable changes")

	}

	if err := s.RemoveLegacyInstallPath(cfg); err != nil {
		return errs.Wrap(err, "Failed to remove legacy install path")
	}

	return nil
}

func run() error {
	// handle state export config --filter=dir (install scripts call this function to write the install-source file)
	if len(os.Args) == 4 && os.Args[1] == "export" && os.Args[2] == "config" && os.Args[3] == "--filter=dir" {
		cfg, err := config.Get()
		if err != nil {
			return errs.Wrap(err, "Failed to read configuration.")
		}
		fmt.Println(cfg.ConfigPath())
		return nil
	}

	if len(os.Args) < 1 || os.Args[1] != "_prepare" {
		fmt.Println("Sorry! This is a transitional tool that should have been replaced during the last update.   If you see this message, something must have gone wrong.  Re-trying to update now...")
	}

	up, err := updater.DefaultChecker.GetUpdateInfo("", "")
	if err != nil {
		return errs.Wrap(err, "Failed to check for latest update.")
	}

	cfg, err := config.Get()
	if err != nil {
		return errs.Wrap(err, "Failed to read configuration.")
	}
	if err := removeOldStateToolEnvironmentSettings(cfg); err != nil {
		return errs.Wrap(err, "failed to remove environment settings from old State Tool installation")
	}

	installTargetPath := filepath.Dir(appinfo.StateApp().Exec())
	err = up.InstallBlocking(installTargetPath)
	if err != nil {
		return errs.Wrap(err, "Failed to install multi-file update.")
	}

	return nil
}
