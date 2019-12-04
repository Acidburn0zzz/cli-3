package show

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bndr/gotabulate"
	"github.com/spf13/cobra"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/constraints"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/internal/updater"
	"github.com/ActiveState/cli/pkg/cmdlets/commands"
	prj "github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
)

// Command is the show command's definition.
var Command = &commands.Command{
	Name:        "show",
	Description: "show_project",
	Run:         Execute,

	Arguments: []*commands.Argument{
		&commands.Argument{
			Name:        "remote",
			Description: "arg_state_show_remote_description",
			Variable:    &Args.Remote,
		},
	},
}

// Args holds the arg values passed through the command line.
var Args struct {
	Remote string
}

var Flags struct {
	Output *string
}

type platform struct {
	Name         string `json:"name"`
	Os           string `json:"os,omitempty"`
	Version      string `json:"version,omitempty"`
	Architecture string `json:"architecture,omitempty"`
}

type language struct {
	Name    string `json:"name"`
	Version string `json:"version,omitemtpy"`
}

type script struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type event struct {
	Name string `json:"name"`
}

type projectData struct {
	Name         string     `json:"name"`
	Organization string     `json:"organization"`
	Platforms    []platform `json:"platforms,omitempty"`
	Languages    []language `json:"languages,omitempty"`
	Scripts      []script   `json:"scripts,omitempty"`
	Events       []event    `json:"events,omitempty"`
}

// Execute the show command.
func Execute(cmd *cobra.Command, args []string) {
	logging.Debug("Execute")

	var project *prj.Project
	if Args.Remote == "" {
		project = prj.Get()
	} else {
		path := Args.Remote
		projectFilePath := filepath.Join(Args.Remote, constants.ConfigFileName)
		if _, err := os.Stat(path); err != nil {
			print.Error(locale.T("err_state_show_path_does_not_exist"))
			return
		} else if _, err := os.Stat(projectFilePath); err != nil {
			print.Error(locale.T("err_state_show_no_config"))
			return
		}
		projectFile, err := projectfile.Parse(projectFilePath)
		if err != nil {
			logging.Errorf("Unable to parse activestate.yaml: %s", err)
			print.Error(locale.T("err_state_show_project_parse"))
			return
		}
		var fail *failures.Failure
		project, fail = prj.New(projectFile)
		if fail != nil {
			failures.Handle(fail.ToError(), fail.Message)
			return
		}
	}

	updater.PrintUpdateMessage()

	output := commands.Output(strings.ToLower(*Flags.Output))
	switch output {
	case commands.JSON, commands.EditorV0:
		fail := printProjectJSON(project)
		if fail != nil {
			failures.Handle(fail, locale.T("err_state_show_print_json"))
		}
	default:
		printProject(project)
	}
}

func printProjectJSON(project *prj.Project) *failures.Failure {
	resultData, err := json.Marshal(newProject(project))
	if err != nil {
		return failures.FailOS.Wrap(err)
	}

	print.Line(string(resultData))
	return nil
}

func newProject(proj *prj.Project) projectData {
	r := projectData{
		Name:         proj.Name(),
		Organization: proj.Owner(),
	}
	source := proj.Source()

	for _, plat := range source.Platforms {
		r.Platforms = append(r.Platforms, platform{
			Name:         plat.Name,
			Os:           plat.Os,
			Version:      plat.Version,
			Architecture: plat.Architecture,
		})
	}

	for _, lang := range source.Languages {
		r.Languages = append(r.Languages, language{
			Name:    lang.Name,
			Version: lang.Version,
		})
	}

	for _, s := range source.Scripts {
		if !constraints.IsConstrained(s.Constraints) {
			r.Scripts = append(r.Scripts, script{
				Name:        s.Name,
				Description: s.Description,
			})
		}
	}

	for _, e := range source.Events {
		if !constraints.IsConstrained(e.Constraints) {
			r.Events = append(r.Events, event{Name: e.Name})
		}
	}

	return r
}

func printProject(project *prj.Project) {
	print.BoldInline("%s: ", locale.T("print_state_show_name"))
	print.Line("%s", project.Name())

	print.BoldInline("%s: ", locale.T("print_state_show_organization"))
	print.Line("%s", project.Owner())

	print.Line("")

	printPlatforms(project.Source())
	printLanguages(project.Source())
	printScripts(project.Source())
	printEvents(project.Source())
}

func printPlatforms(project *projectfile.Project) {
	if len(project.Platforms) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, platform := range project.Platforms {
		constrained := "*"
		if !constraints.PlatformMatches(platform) {
			constrained = ""
		}
		v := fmt.Sprintf("%s%s %s %s (%s)", constrained, platform.Os, platform.Version, platform.Architecture, platform.Name)
		rows = append(rows, []interface{}{v})
	}

	print.BoldInline("%s:", locale.T("print_state_show_platforms"))
	printTable(rows)
}

func printEvents(project *projectfile.Project) {
	if len(project.Events) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, event := range project.Events {
		if !constraints.IsConstrained(event.Constraints) {
			rows = append(rows, []interface{}{event.Name})
		}
	}

	print.BoldInline("%s:", locale.T("print_state_show_events"))
	printTable(rows)
}

func printScripts(project *projectfile.Project) {
	if len(project.Scripts) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, script := range project.Scripts {
		if !constraints.IsConstrained(script.Constraints) {
			rows = append(rows, []interface{}{script.Name, script.Description})
		}
	}

	print.BoldInline("%s:", locale.T("print_state_show_scripts"))
	printTable(rows)
}

func printLanguages(project *projectfile.Project) {
	if len(project.Languages) == 0 {
		return
	}

	rows := [][]interface{}{}
	for _, language := range project.Languages {
		if !constraints.IsConstrained(language.Constraints) {
			rows = append(rows, []interface{}{language.Name, language.Version})
		}
	}

	print.BoldInline("%s:", locale.T("print_state_show_languages"))
	printTable(rows)
}

func printTable(rows [][]interface{}) {
	t := gotabulate.Create(rows)

	// gotabulate tries to make the first row the headers, so use some empty header instead
	// this is also the reason why we're using BoldInLine, since the header line will act as the newline
	t.SetHeaders([]string{""})

	t.SetHideLines([]string{"betweenLine", "top", "aboveTitle", "belowheader", "LineTop", "LineBottom", "bottomLine"}) // Don't print whitespace lines
	t.SetAlign("left")
	print.Line(t.Render("plain"))
}
