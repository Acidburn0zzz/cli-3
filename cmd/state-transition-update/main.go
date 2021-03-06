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
	"github.com/ActiveState/cli/internal/machineid"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/internal/runbits/panics"
	"github.com/ActiveState/cli/internal/subshell"
	"github.com/ActiveState/cli/internal/subshell/sscommon"
	"github.com/ActiveState/cli/internal/updater"
	"github.com/rollbar/rollbar-go"
)

func main() {
	var exitCode int
	defer func() {
		if panics.HandlePanics() {
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
	switch {
	// handle state export config --filter=dir (install scripts call this function to write the install-source file)
	case len(os.Args) == 4 && os.Args[1] == "export" && os.Args[2] == "config" && os.Args[3] == "--filter=dir":
		return runExport()

	case len(os.Args) < 1 || os.Args[1] != "_prepare":
		return runPrepare()

	default:
		if err := runDefault(); err != nil {
			return err
		}

		fmt.Println("Please start a new shell to continue using the State Tool.")
		return nil
	}
}

func runExport() error {
	cfg, err := config.Get()
	if err != nil {
		return errs.Wrap(err, "Failed to read configuration.")
	}
	fmt.Println(cfg.ConfigPath())
	return nil
}

func runPrepare() error {
	fmt.Println("Sorry! This is a transitional tool that should have been replaced during the last update.   If you see this message, something must have gone wrong.  Re-trying to update now...")
	return nil
}

func runDefault() error {
	up, err := updater.DefaultChecker.GetUpdateInfo("", "")
	if err != nil {
		return errs.Wrap(err, "Failed to check for latest update.")
	}

	cfg, err := config.Get()
	if err != nil {
		return errs.Wrap(err, "Failed to read configuration.")
	}
	machineid.SetConfiguration(cfg)
	machineid.SetErrorLogger(logging.Error)
	logging.UpdateConfig(cfg)

	oldInfo, err := os.Stat(appinfo.StateApp().Exec())
	if err != nil {
		return errs.Wrap(err, "Failed to retrieve stat info for transitional State Tool executable.")
	}

	if err := removeOldStateToolEnvironmentSettings(cfg); err != nil {
		return errs.Wrap(err, "failed to remove environment settings from old State Tool installation")
	}

	err = up.InstallBlocking("")
	if err != nil {
		return errs.Wrap(err, "Failed to install multi-file update.")
	}

	logging.Debug("Multi-file State Tool is installed.")

	info, err := os.Stat(appinfo.StateApp().Exec())
	if err != nil {
		return errs.Wrap(err, "Failed to retrieve stat info for transitional State Tool executable.")
	}

	// if the transitional state tool has been replaced by the installer, we are done
	if oldInfo.ModTime() != info.ModTime() {
		return nil
	}

	logging.Debug("Removing transitional State Tool")
	// otherwise: remove the transitional State Tool
	if err := removeSelf(); err != nil {
		logging.Error("Failed to remove transitional State Tool: %s", errs.JoinMessage(err))
	}

	return nil
}
