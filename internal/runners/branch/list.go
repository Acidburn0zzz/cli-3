package branch

import (
	"fmt"
	"sort"

	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/project"
)

type primeable interface {
	primer.Outputer
	primer.Projecter
}

type List struct {
	out     output.Outputer
	project *project.Project
}

func NewList(prime primeable) *List {
	return &List{
		out:     prime.Output(),
		project: prime.Project(),
	}
}

func (l *List) Run() error {
	logging.Debug("ExecuteList")

	project, err := model.FetchProjectByName(l.project.Owner(), l.project.Name())
	if err != nil {
		return locale.WrapError(err, "err_fetch_project", l.project.Namespace().String())
	}

	localBranch := l.project.BranchName()

	var branches []string
	for _, branch := range project.Branches {
		branchName := branch.Label
		if branchName == localBranch {
			branchName = fmt.Sprintf("[ACTIONABLE]%s[/RESET]", branchName)
		}
		branches = append(branches, branchName)
	}
	sort.Slice(branches, func(i, j int) bool {
		return branches[i] < branches[j]
	})

	l.out.Print(branches)

	return nil
}
