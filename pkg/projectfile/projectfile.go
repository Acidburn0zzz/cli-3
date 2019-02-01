package projectfile

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	yaml "gopkg.in/yaml.v2"
)

var (
	// FailNoProject identifies a failure as being due to a missing project file
	FailNoProject = failures.Type("projectfile.fail.noproject")

	// FailParseProject identifies a failure as being due inability to parse file contents
	FailParseProject = failures.Type("projectfile.fail.parseproject")

	// FailValidate identifies a failure during validation
	FailValidate = failures.Type("projectfile.fail.validate")

	// FailInvalidVersion identifies a failure as being due to an invalid version format
	FailInvalidVersion = failures.Type("projectfile.fail.version")
)

// Version is used in cases where we only care about parsing the version field. In all other cases the version is parsed via
// the Project struct
type Version struct {
	Version string `yaml:"version"`
}

// Project covers the top level project structure of our yaml
type Project struct {
	Name         string      `yaml:"name"`
	Owner        string      `yaml:"owner"`
	Namespace    string      `yaml:"namespace"`
	Version      string      `yaml:"version"`
	Environments string      `yaml:"environments"`
	Platforms    []Platform  `yaml:"platforms"`
	Languages    []Language  `yaml:"languages"`
	Variables    []*Variable `yaml:"variables"`
	Events       []Event     `yaml:"events"`
	Scripts      []Script    `yaml:"scripts"`
	path         string      // "private"
}

// Platform covers the platform structure of our yaml
type Platform struct {
	Name         string `yaml:"name"`
	Os           string `yaml:"os"`
	Version      string `yaml:"version"`
	Architecture string `yaml:"architecture"`
	Libc         string `yaml:"libc"`
	Compiler     string `yaml:"compiler"`
}

// Build covers the build map, which can go under languages or packages
// Build can hold variable keys, so we cannot predict what they are, hence why it is a map
type Build map[string]string

// Language covers the language structure, which goes under Project
type Language struct {
	Name        string     `yaml:"name"`
	Version     string     `yaml:"version"`
	Constraints Constraint `yaml:"constraints"`
	Build       Build      `yaml:"build"`
	Packages    []Package  `yaml:"packages"`
}

// Constraint covers the constraint structure, which can go under almost any other struct
type Constraint struct {
	Platform    string `yaml:"platform"`
	Environment string `yaml:"environment"`
}

// Package covers the package structure, which goes under the language struct
type Package struct {
	Name        string     `yaml:"name"`
	Version     string     `yaml:"version"`
	Constraints Constraint `yaml:"constraints"`
	Build       Build      `yaml:"build"`
}

// Event covers the event structure, which goes under Project
type Event struct {
	Name        string     `yaml:"name"`
	Value       string     `yaml:"value"`
	Constraints Constraint `yaml:"constraints"`
}

// Script covers the script structure, which goes under Project
type Script struct {
	Name        string     `yaml:"name"`
	Value       string     `yaml:"value"`
	Standalone  bool       `yaml:"standalone"`
	Constraints Constraint `yaml:"constraints"`
}

var persistentProject *Project

// Parse the given filepath, which should be the full path to an activestate.yaml file
func Parse(filepath string) (*Project, error) {
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	project := Project{}
	err = yaml.Unmarshal([]byte(dat), &project)
	project.path = filepath

	if err != nil {
		return nil, FailNoProject.New(locale.T("err_project_parse", map[string]interface{}{"Error": err.Error()}))
	}

	return &project, project.Parse()
}

// Parse further processes the current file by parsing mixed values (something go-yaml doesnt handle)
func (p *Project) Parse() *failures.Failure {
	for _, variable := range p.Variables {
		fail := variable.Parse()
		if fail != nil {
			return fail
		}
	}

	return nil
}

// Path returns the project's activestate.yaml file path.
func (p *Project) Path() string {
	return p.path
}

// SetPath sets the path of the project file and should generally only be used by tests
func (p *Project) SetPath(path string) {
	p.path = path
}

// Save the project to its activestate.yaml file
func (p *Project) Save() *failures.Failure {
	dat, err := yaml.Marshal(p)
	if err != nil {
		return failures.FailMarshal.Wrap(err)
	}

	f, err := os.Create(p.Path())
	defer f.Close()
	if err != nil {
		return failures.FailIO.Wrap(err)
	}

	_, err = f.Write([]byte(dat))
	if err != nil {
		return failures.FailIO.Wrap(err)
	}

	return nil
}

// Returns the path to the project activestate.yaml
func getProjectFilePath() (string, *failures.Failure) {
	projectFilePath := os.Getenv(constants.ProjectEnvVarName)
	if projectFilePath != "" {
		return projectFilePath, nil
	}

	root, err := os.Getwd()
	if err != nil {
		logging.Warning("Could not get project root path: %v", err)
		return "", failures.FailOS.Wrap(err)
	}
	return fileutils.FindFileInPath(root, constants.ConfigFileName)
}

// Get returns the project configration in an unsafe manner (exits if errors occur)
func Get() *Project {
	project, err := GetSafe()
	if err != nil {
		failures.Handle(err, locale.T("err_project_file_unavailable"))
		os.Exit(1)
	}

	return project
}

// GetSafe returns the project configuration in a safe manner (returns error)
func GetSafe() (*Project, *failures.Failure) {
	if persistentProject != nil {
		return persistentProject, nil
	}

	// we do not want to use a path provided by state if we're running tests
	projectFilePath, failure := getProjectFilePath()
	if failure != nil {
		return nil, failure
	}

	_, err := ioutil.ReadFile(projectFilePath)
	if err != nil {
		logging.Warning("Cannot load config file: %v", err)
		return nil, FailNoProject.New(locale.T("err_no_projectfile"))
	}
	project, err := Parse(projectFilePath)
	if err != nil {
		return nil, FailParseProject.New(locale.T("err_parse_project"))
	}
	project.Persist()
	return project, nil
}

// ParseVersion parses the version field from the projectfile, and ONLY the version field. This is to ensure it doesn't
// trip over older activestate.yaml's with breaking changes
func ParseVersion() (string, *failures.Failure) {
	var projectFilePath string
	if persistentProject != nil {
		projectFilePath = persistentProject.Path()
	} else {
		var fail *failures.Failure
		projectFilePath, fail = getProjectFilePath()
		if fail != nil {
			// Not being able to find a project file is not a failure for the purposes of this function
			return "", nil
		}
	}

	if projectFilePath == "" {
		return "", nil
	}

	dat, err := ioutil.ReadFile(projectFilePath)
	if err != nil {
		return "", failures.FailIO.Wrap(err)
	}

	versionStruct := Version{}
	err = yaml.Unmarshal([]byte(dat), &versionStruct)
	if err != nil {
		return "", FailParseProject.Wrap(err)
	}

	if versionStruct.Version == "" {
		return "", nil
	}

	version := strings.TrimSpace(versionStruct.Version)
	match, fail := regexp.MatchString("^\\d+\\.\\d+\\.\\d+-\\d+$", version)
	if fail != nil || !match {
		return "", FailInvalidVersion.New(locale.T("err_invalid_version"))
	}

	return versionStruct.Version, nil
}

// Reset the current state, which unsets the persistent project
func Reset() {
	persistentProject = nil
	os.Unsetenv(constants.ProjectEnvVarName)
}

// Persist "activates" the given project and makes it such that subsequent calls
// to Get() return this project.
// Only one project can persist at a time.
func (p *Project) Persist() {
	persistentProject = p
	os.Setenv(constants.ProjectEnvVarName, p.Path())
}
