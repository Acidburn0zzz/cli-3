package runbits

import (
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/platform/runtime"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/go-openapi/strfmt"
)

// RefreshRuntime should be called after runtime mutations.
func RefreshRuntime(auth *authentication.Auth, out output.Outputer, proj *project.Project, cachePath string, commitID strfmt.UUID, changed bool) error {
	rtMessages := DefaultRuntimeEventHandler(out)
	target := runtime.NewProjectTarget(proj, cachePath, &commitID)
	isCached := true
	isInitial := !fileutils.DirExists(target.Dir())
	rt, err := runtime.New(target)
	if err != nil {
		if runtime.IsNeedsUpdateError(err) {
			isCached = false
		} else {
			return locale.WrapError(err, "err_packages_update_runtime_init", "Could not initialize runtime.")
		}
	}

	if !changed && isCached {
		out.Print(locale.Tl("pkg_already_uptodate", "Requested dependencies are already configured and installed."))
		return nil
	}

	if !isCached {
		if isInitial {
			out.Notice(output.Heading(locale.Tl("install_runtime", "Installing Runtime")))
			out.Notice(locale.Tl("install_runtime_info", "Installing your runtime and dependencies."))
		} else {
			out.Notice(output.Heading(locale.Tl("update_runtime", "Updating Runtime")))
			out.Notice(locale.Tl("update_runtime_info", "Changes to your runtime may require some dependencies to be rebuilt."))
		}
		err := rt.Update(auth, rtMessages)
		if err != nil {
			return locale.WrapError(err, "err_packages_update_runtime_install", "Could not install dependencies.")
		}
	}

	return nil
}
