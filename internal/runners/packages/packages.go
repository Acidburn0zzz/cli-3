package packages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/prompt"
	"github.com/ActiveState/cli/internal/runbits"
	"github.com/ActiveState/cli/pkg/cmdlets/auth"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/platform/runtime"
	"github.com/ActiveState/cli/pkg/project"
)

type PackageType int

const (
	Package PackageType = iota
	Bundle
)

func (pt PackageType) String() string {
	switch pt {
	case Package:
		return "package"
	case Bundle:
		return "bundle"
	}
	return ""
}

func (pt PackageType) Namespace() model.NamespacePrefix {
	switch pt {
	case Package:
		return model.PackageNamespacePrefix
	case Bundle:
		return model.BundlesNamespacePrefix
	}
	return ""
}

const latestVersion = "latest"

func executePackageOperation(pj *project.Project, out output.Outputer, authentication *authentication.Auth, prompt prompt.Prompter, language, name, version string, operation model.Operation, pt PackageType) error {
	isHeadless := pj.IsHeadless()
	if !isHeadless && !authentication.Authenticated() {
		anonymousOk, fail := prompt.Confirm(locale.Tl("continue_anon", "Continue Anonymously?"), locale.T("prompt_headless_anonymous"), true)
		if fail != nil {
			return locale.WrapInputError(fail.ToError(), "Authentication cancelled.")
		}
		isHeadless = anonymousOk
	}

	// Note: User also lands here if answering No to the question about anonymous commit.
	if !isHeadless {
		fail := auth.RequireAuthentication(locale.T("auth_required_activate"), out, prompt)
		if fail != nil {
			return fail.WithDescription("err_auth_required")
		}
	}

	if strings.ToLower(version) == latestVersion {
		version = ""
	}

	// Verify that the provided package actually exists (the vcs API doesn't care)
	var ingredient *model.IngredientAndVersion
	var err error
	if version == "" {
		ingredient, err = model.IngredientWithLatestVersion(language, name, pt.Namespace())
	} else {
		ingredient, err = model.IngredientByNameAndVersion(language, name, version, pt.Namespace())
	}
	if err != nil {
		return locale.WrapError(err, "package_ingredient_err", "Failed to resolve an ingredient named {{.V0}}.", name)
	}

	// Check if this is an addition or an update
	if operation == model.OperationAdded {
		req, err := model.GetRequirement(pj.CommitUUID(), ingredient.Namespace, *ingredient.Ingredient.Name)
		if err != nil {
			return errs.Wrap(err, "Could not get requirement")
		}
		if req != nil {
			operation = model.OperationUpdated
		}
	}

	parentCommitID := pj.CommitUUID()
	commitID, fail := model.CommitPackage(parentCommitID, operation, *ingredient.Ingredient.Name, ingredient.Namespace, version)
	if fail != nil {
		return locale.WrapError(fail.ToError(), fmt.Sprintf("err_%s_%s", pt.String(), operation))
	}

	revertCommit, err := model.GetRevertCommit(pj.CommitUUID(), commitID)
	if err != nil {
		return errs.Wrap(err, "Could not get revert commit to check if changes were indeed made")
	}

	orderChanged := len(revertCommit.Changeset) > 0

	logging.Debug("Order changed: %v", orderChanged)

	// Update project references to the new commit, if changes were indeed made (otherwise we effectively drop the new commit)
	if orderChanged {
		if !isHeadless {
			err := model.UpdateProjectBranchCommitByName(pj.Owner(), pj.Name(), commitID)
			if err != nil {
				return locale.WrapError(err, "err_package_"+string(operation))
			}
		}
		if fail := pj.Source().SetCommit(commitID.String(), isHeadless); fail != nil {
			return fail.WithDescription("err_package_update_pjfile")
		}
	} else {
		commitID = parentCommitID
	}

	if pt == Bundle {
		printBundleInfo(out, ingredient, version)
	}

	// Create runtime
	rt, err := runtime.NewRuntime(pj.Source().Path(), commitID, pj.Owner(), pj.Name(), runbits.NewRuntimeMessageHandler(out))
	if err != nil {
		return locale.WrapError(err, "err_packages_update_runtime_init", "Could not initialize runtime.")
	}

	if !orderChanged && rt.IsCachedRuntime() {
		out.Print(locale.Tl("pkg_already_uptodate", "Requested dependencies are already configured and installed."))
		return nil
	}

	// Update runtime
	if !rt.IsCachedRuntime() {
		out.Print(locale.Tl("update_runtime", "[NOTICE]Updating Runtime[/RESET]"))
		out.Print(locale.Tl("update_runtime_info", "Changes to your runtime may require some dependencies to be rebuilt."))
		_, _, fail := runtime.NewInstaller(rt).Install()
		if fail != nil {
			return locale.WrapError(fail, "err_packages_update_runtime_install", "Could not install dependencies.")
		}
	}

	// Print the result
	if version != "" {
		out.Print(locale.Tr(fmt.Sprintf("%s_version_%s", pt.String(), operation), *ingredient.Ingredient.Name, version))
	} else {
		out.Print(locale.Tr(fmt.Sprintf("%s_%s", pt.String(), operation), *ingredient.Ingredient.Name))
	}

	return nil
}

func splitNameAndVersion(input string) (string, string) {
	nameArg := strings.Split(input, "@")
	name := nameArg[0]
	version := ""
	if len(nameArg) == 2 {
		version = nameArg[1]
	}

	return name, version
}

func printBundleInfo(out output.Outputer, ingredient *model.IngredientAndVersion, version string) {
	out.Print("")
	if version == "" {
		out.Print(locale.Tl("bundle_no_version", "No bundle version specified, choosing version {{.V0}}", ingredient.Version))
		out.Print("")
	}

	var count int
	var dependencies []string
	for _, dep := range ingredient.LatestVersion.DependencySets {
		if dep.OriginalRequirement == "" {
			continue
		}
		count++
		dependencies = append(dependencies, strings.TrimSuffix(dep.OriginalRequirement, " 0"))
	}

	out.Print(locale.Tl("bundle_title", "[NOTICE]{{.V0}} Bundle[/RESET] includes {{.V1}} packages", *ingredient.Ingredient.Name, strconv.Itoa(count)))
	last := len(dependencies) - 1
	for i, dep := range dependencies {
		if i == last {
			out.Print(locale.Tl("bundle_package_name", "  └─ {{.V0}}", dep))
			continue
		}
		out.Print(locale.Tl("bundle_package_name", "  ├─ {{.V0}}", dep))
	}

	out.Print("")
	out.Print(locale.Tl("packages_auto_msg", "Packages are automatically added to your runtime."))
}
