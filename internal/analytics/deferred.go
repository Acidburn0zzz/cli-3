package analytics

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/osutils"
)

var deferAnalytics bool

type deferredData struct {
	Category   string
	Action     string
	Label      string
	Dimensions map[string]string
}

const deferredCfgKey = "deferrer_analytics"
const deferrerFileName = "deferrer"

func deferrerTimeStamp(cfg Configurable) string {
	return filepath.Join(cfg.ConfigPath(), deferrerFileName)
}

func deferrerTimeStampDayOld(cfg Configurable) bool {
	df := deferrerTimeStamp(cfg)
	stat, err := os.Stat(df)
	if err != nil {
		logging.Errorf("Could not stat file: %s, error: %v", df, err)
		return false
	}
	diff := time.Now().Sub(stat.ModTime())
	return diff > 24*time.Hour
}

func runNonDeferredStateToolCommand(cfg Configurable) error {
	exe, err := os.Executable()
	if err != nil {
		logging.Errorf("Could not determine State Tool executable: %v", err)
		exe = "state"
	}
	cmd := exec.Command(exe, "--version")
	cmd.SysProcAttr = osutils.SysProcAttrForNewProcessGroup()
	cmd.Env = append(os.Environ(), fmt.Sprintf("%s=true", constants.DisableUpdates))
	err = cmd.Start()
	if err != nil {
		return errs.Wrap(err, "Failed to run state --version in background")
	}
	err = cmd.Process.Release()
	if err != nil {
		return errs.Wrap(err, "Failed to release process resources for background process")
	}

	return nil
}

func SetDeferred(cfg Configurable, da bool) {
	deferAnalytics = da
	if deferAnalytics {
		/*
			if deferrerTimeStampDayOld(cfg) {
				err := runNonDeferredStateToolCommand(cfg)
				if err != nil {
					logging.Errorf("Failed to launch non-deferred State Tool command: %v", err)
				}
			}
		*/
		return
	}
	eventWaitGroup.Add(1)
	go func() {
		defer eventWaitGroup.Done()
		if err := sendDeferred(cfg, sendEvent); err != nil {
			logging.Errorf("Could not send deferred events: %v", err)
		}
	}()
}

type Configurable interface {
	Set(string, interface{})
	Save() error
	GetString(string) string
	ConfigPath() string
}

func deferEvent(cfg Configurable, category, action, label string, dimensions map[string]string) error {
	logging.Debug("Deferring: %s, %s, %s", category, action, label)
	deferred, err := loadDeferred(cfg)
	if err != nil {
		return errs.Wrap(err, "Could not load events on defer")
	}

	if !fileutils.FileExists(deferrerTimeStamp(cfg)) {
		err = fileutils.Touch(deferrerTimeStamp(cfg))
		if err != nil {
			logging.Errorf("Failed to create deferrer time stamp file: %v", err)
		}
	}

	deferred = append(deferred, deferredData{category, action, label, dimensions})
	if err := saveDeferred(cfg, deferred); err != nil {
		return errs.Wrap(err, "Could not save event on defer")
	}
	return nil
}

func sendDeferred(cfg Configurable, sender func(string, string, string, map[string]string) error) error {
	deferred, err := loadDeferred(cfg)
	if err != nil {
		return errs.Wrap(err, "Could not load events on send")
	}
	for n, event := range deferred {
		if err := sender(event.Category, event.Action, event.Label, event.Dimensions); err != nil {
			return errs.Wrap(err, "Could not send deferred event")
		}
		if err := saveDeferred(cfg, deferred[n+1:]); err != nil {
			return errs.Wrap(err, "Could not save deferred event on send")
		}
	}
	if err := cfg.Save(); err != nil { // the global viper instance is bugged, need to work around it for now -- https://www.pivotaltracker.com/story/show/175624789
		return locale.WrapError(err, "err_viper_write_send_defer", "Could not save configuration on send deferred")
	}

	// remove deferrer time stamp file
	err = os.Remove(deferrerTimeStamp(cfg))
	if err != nil && !os.IsNotExist(err) {
		logging.Errorf("Could not remove deferrer time stamp file: %v", err)
	}
	return nil
}

func saveDeferred(cfg Configurable, v []deferredData) error {
	s, err := json.Marshal(v)
	if err != nil {
		return errs.New("Could not serialize deferred analytics: %v, error: %v", v, err)
	}
	cfg.Set(deferredCfgKey, string(s))
	return nil
}

func loadDeferred(cfg Configurable) ([]deferredData, error) {
	v := cfg.GetString(deferredCfgKey)
	d := []deferredData{}
	if v != "" {
		err := json.Unmarshal([]byte(v), &d)
		if err != nil {
			return d, errs.Wrap(err, "Could not deserialize deferred analytics: %v", v)
		}
	}
	return d, nil
}
