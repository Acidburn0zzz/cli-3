package languages

import (
	"errors"
	"strings"

	"github.com/ActiveState/cli/internal/captain"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/project"
)

type Update struct {
	out     output.Outputer
	project *project.Project
}

type primeable interface {
	primer.Projecter
	primer.Outputer
}

func NewUpdate(prime primeable) *Update {
	return &Update{prime.Output(), prime.Project()}
}

type UpdateParams struct {
	Language captain.LanguageVersion
}

func (u *Update) Run(params *UpdateParams) error {
	lang := &model.Language{
		Name: params.Language.Name(), Version: params.Language.Version(),
	}

	if u.project == nil {
		return locale.NewInputError("err_no_project")
	}

	err := ensureLanguageProject(lang, u.project)
	if err != nil {
		return err
	}

	err = ensureLanguagePlatform(lang)
	if err != nil {
		return err
	}

	err = ensureVersion(lang)
	if err != nil {
		if lang.Version == "" {
			return locale.WrapInputError(err, "err_language_project", "Language: {{.V0}} is already installed, you can update it by running {{.V0}}@<version>", lang.Name)
		}
		return err
	}

	err = removeLanguage(u.project, lang.Name)
	if err != nil {
		return err
	}

	return addLanguage(u.project, lang)
}

func ensureLanguagePlatform(language *model.Language) error {
	platformLanguages, err := model.FetchLanguages()
	if err != nil {
		return err
	}

	for _, pl := range platformLanguages {
		if strings.ToLower(pl.Name) == strings.ToLower(language.Name) {
			return nil
		}
	}

	return errors.New(locale.Tr("err_update_not_found", language.Name))
}

func ensureLanguageProject(language *model.Language, project *project.Project) error {
	targetCommitID, err := model.LatestCommitID(project.Owner(), project.Name())
	if err != nil {
		return err
	}

	platformLanguage, err := model.FetchLanguageForCommit(*targetCommitID)
	if err != nil {
		return err
	}

	if platformLanguage.Name != language.Name {
		return locale.NewInputError("err_language_mismatch")
	}
	return nil
}

type fetchVersionsFunc func(name string) ([]string, error)

func ensureVersion(language *model.Language) error {
	return ensureVersionTestable(language, model.FetchLanguageVersions)
}

func ensureVersionTestable(language *model.Language, fetchVersions fetchVersionsFunc) error {
	if language.Version == "" {
		return locale.NewInputError("err_language_no_version", "No language version provided")
	}

	versions, err := fetchVersions(language.Name)
	if err != nil {
		return err
	}

	for _, ver := range versions {
		if language.Version == ver {
			return nil
		}
	}

	return locale.NewInputError("err_language_version_not_found", "", language.Version, language.Name)
}

func removeLanguage(project *project.Project, current string) error {
	targetCommitID, err := model.LatestCommitID(project.Owner(), project.Name())
	if err != nil {
		return err
	}

	platformLanguage, err := model.FetchLanguageForCommit(*targetCommitID)
	if err != nil {
		return err
	}

	err = model.CommitLanguage(project.Owner(), project.Name(), model.OperationRemoved, platformLanguage.Name, platformLanguage.Version)
	if err != nil {
		return err
	}

	return nil
}

func addLanguage(project *project.Project, lang *model.Language) error {
	err := model.CommitLanguage(project.Owner(), project.Name(), model.OperationAdded, lang.Name, lang.Version)
	if err != nil {
		return err
	}

	return nil
}
