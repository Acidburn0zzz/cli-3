// +build darwin

package installer

import (
	"os"
	"path/filepath"

	"github.com/ActiveState/cli/internal/appinfo"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/gobuffalo/packr"
)

const (
	contentsDir  = "/Applications/ActiveState Desktop.app/Contents"
	macOSDir     = "MacOS"
	resourcesDir = "Resources"
	iconBundle   = "state-tray.icns"
	infoFile     = "Info.plist"
)

// InstallSystemFiles installs files in the /Application directory
func InstallSystemFiles(installDir string) error {
	err := createAppDirStructure()
	if err != nil {
		return errs.Wrap(err, "Could not create app directory structure")
	}

	err = populateAppDir(installDir)
	if err != nil {
		return errs.Wrap(err, "Could not populate state-tray app directory")
	}

	return nil
}

func createAppDirStructure() error {
	err := fileutils.Mkdir(filepath.Join(contentsDir))
	if err != nil {
		return errs.Wrap(err, "Could not create Contents app dir")
	}

	err = fileutils.Mkdir(filepath.Join(contentsDir, macOSDir))
	if err != nil {
		return errs.Wrap(err, "Could not create MacOS app dir")
	}

	err = fileutils.Mkdir(filepath.Join(contentsDir, resourcesDir))
	if err != nil {
		return errs.Wrap(err, "Could not create Resources app dir")
	}

	return nil
}

func populateAppDir(installDir string) error {
	box := packr.NewBox("../../../../assets/macOS")
	err := fileutils.WriteFile(filepath.Join(contentsDir, resourcesDir, iconBundle), box.Bytes(iconBundle))
	if err != nil {
		return errs.Wrap(err, "Could not write icon file")
	}

	err = fileutils.WriteFile(filepath.Join(contentsDir, infoFile), box.Bytes(infoFile))
	if err != nil {
		return errs.Wrap(err, "Could not write info plist file")
	}

	trayInfo := appinfo.TrayApp()
	err = createNewSymlink(filepath.Join(installDir, filepath.Base(trayInfo.Exec())), filepath.Join(contentsDir, macOSDir, filepath.Base(trayInfo.Exec())))
	if err != nil {
		return errs.Wrap(err, "Could not create state-tray symlink")
	}

	return nil
}

func createNewSymlink(target, filename string) error {
	if fileutils.FileExists(filename) {
		if err := os.Remove(filename); err != nil {
			return errs.Wrap(err, "Could not delete existing symlink")
		}
	}

	err := os.Symlink(target, filename)
	if err != nil {
		return errs.Wrap(err, "Could not create symlink")
	}
	return nil
}
