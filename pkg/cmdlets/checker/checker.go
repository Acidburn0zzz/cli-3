package checker

import (
	"errors"
	"net"
	"strconv"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/svcmanager"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/project"
	"golang.org/x/net/context"
)

// RunCommitsBehindNotifier checks for the commits behind count based on the
// provided project and displays the results to the user in a helpful manner.
func RunCommitsBehindNotifier(p *project.Project,out output.Outputer) {
	latestCommitID, err := model.BranchCommitID(p.Owner(), p.Name(), p.BranchName())
	if err != nil {
		logging.Error("Can not get branch info for %s/%s", p.Owner(), p.Name())
		return
	}

	if latestCommitID == nil {
		logging.Debug("Latest commit is nil")
		return
	}

	count, err := model.CommitsBehind(*latestCommitID, p.CommitUUID())
	if err != nil {
		if errors.Is(err, model.ErrCommitCountUnknowable) {
			out.Notice(output.Heading(locale.Tr("runtime_update_notice_unknown_count")))
			out.Notice(locale.Tr("runtime_update_help", p.Owner(), p.Name()))
			return
		}

		logging.Warning(locale.T("err_could_not_get_commit_behind_count"))
		return
	}
	if count > 0 {
		ct := strconv.Itoa(count)
		out.Notice(output.Heading(locale.Tr("runtime_update_notice_known_count", ct)))
		out.Notice(locale.Tr("runtime_update_help", p.Owner(), p.Name()))
	}
}

func RunUpdateNotifier(cfg *config.Instance, out output.Outputer) {
	ctx, _ := context.WithTimeout(context.Background(), svcmanager.MinimalTimeout)
	svc, err := model.NewSvcModel(ctx, cfg)
	if err != nil {
		logging.Error("Could not init svc model when running update notifier, error: %v", errs.JoinMessage(err))
		return
	}
	up, err := svc.CheckUpdate()
	if err != nil {
		var timeoutErr net.Error
		if errors.As(err, &timeoutErr) && timeoutErr.Timeout() {
			logging.Debug("CheckUpdate timed out")
			return
		}
		logging.Error("Could not check for update when running update notifier, error: %v", errs.JoinMessage(err))
		return
	}
	if up == nil {
		return
	}
	out.Notice(output.Heading(locale.Tr("update_available_header")))
	out.Notice(locale.Tr("update_available", constants.VersionNumber, up.Version))
}
