package activate

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/internal/prompt"
	"github.com/ActiveState/cli/pkg/platform/api"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_client/organizations"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_client/projects"
	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/projectfile"
)

var exit = os.Exit

// NewExecute the new command.
func NewExecute(cmd *cobra.Command, args []string) {
	logging.Debug("Execute")

	name, fail := prompter.Input(locale.T("state_new_prompt_name"), "", prompt.InputRequired)
	if fail != nil {
		failures.Handle(fail, locale.T("error_state_new_aborted"))
		exit(1)
	}

	if !authentication.Get().Authenticated() && flag.Lookup("test.v") == nil {
		print.Error(locale.T("error_state_new_no_auth"))
		exit(1)
	}

	// If the user is not yet authenticated into the ActiveState Platform, it is a
	// simple prompt. Otherwise, fetch the list of organizations the user belongs
	// to and present the list to the user for a selection.
	owner, fail := promptForOwner()
	if fail != nil {
		failures.Handle(fail, locale.T("error_state_new_aborted"))
		exit(1)
	}

	// Create the project on the platform
	if fail = createPlatformProject(name, owner); fail != nil {
		failures.Handle(fail, locale.T("error_state_new_project_add"))
		exit(1)
	}

	path, fail := fetchPath(name)
	if fail != nil {
		failures.Handle(fail, locale.T("error_state_new_aborted"))
		exit(1)
	}

	// Create the project directory
	if fail = createProjectDir(path); fail != nil {
		failures.Handle(fail, locale.T("error_state_new_aborted"))
		exit(1)
	}

	projectURL := fmt.Sprintf("https://%s/%s/%s", constants.PlatformURL, owner, name)

	// Create the project locally on disk.
	if _, fail = projectfile.Create(projectURL, path); fail != nil {
		failures.Handle(fail, locale.T("error_state_new_aborted"))
		exit(1)
	}

	print.Line(locale.T("state_new_created", map[string]interface{}{"Dir": path}))
}

func promptForOwner() (string, *failures.Failure) {
	params := organizations.NewListOrganizationsParams()
	memberOnly := true
	params.SetMemberOnly(&memberOnly)
	orgs, err := authentication.Client().Organizations.ListOrganizations(params, authentication.ClientAuth())
	if err != nil {
		return "", api.FailUnknown.New("error_state_new_fetch_organizations")
	}
	owners := []string{}
	for _, org := range orgs.Payload {
		owners = append(owners, org.Name)
	}
	if len(owners) > 1 {
		return prompter.Select(locale.T("state_new_prompt_owner"), owners, "")
	}
	return owners[0], nil // auto-select only option
}

func fetchPath(projName string) (string, *failures.Failure) {
	cwd, _ := os.Getwd()
	files, _ := ioutil.ReadDir(cwd)

	if len(files) == 0 {
		// Current working directory is devoid of files. Use it as the path for
		// the new project.
		return cwd, nil
	}

	// Current working directory has files in it. Use a subdirectory with the
	// project name as the path for the new project.
	path := filepath.Join(cwd, projName)
	if _, err := os.Stat(path); err == nil {
		return "", failures.FailIO.New("error_state_new_exists")
	}

	return path, nil
}

func createPlatformProject(name, owner string) *failures.Failure {
	addParams := projects.NewAddProjectParams()
	addParams.SetOrganizationName(owner)
	addParams.SetProject(&mono_models.Project{Name: name})
	_, err := authentication.Client().Projects.AddProject(addParams, authentication.ClientAuth())
	if err != nil {
		return api.FailUnknown.Wrap(err)
	}
	return nil
}

func createProjectDir(path string) *failures.Failure {
	if _, err := os.Stat(path); err == nil {
		// Directory already exists
		files, _ := ioutil.ReadDir(path)
		if len(files) == 0 {
			return nil
		}
		return failures.FailIO.New("error_state_new_exists")
	}
	if err := os.MkdirAll(path, 0755); err != nil {
		return failures.FailIO.New("error_state_new_mkdir")
	}
	return nil
}
