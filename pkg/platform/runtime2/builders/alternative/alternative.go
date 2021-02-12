package alternative

import (
	"github.com/ActiveState/cli/pkg/platform/runtime2/rterrs"
	"github.com/ActiveState/cli/pkg/project"
)

//var _ runtime.EnvProvider = &Alternative{}

// Alternative is the specialization of a runtime for alternative builds
type Alternative struct{}

// New is the constructor function for alternative runtimes
func New(proj *project.Project) (*Alternative, error) {
	return nil, rterrs.NotImplemented
}

func (a *Alternative) Environ() (map[string]string, error) {
	return nil, rterrs.NotImplemented
}
