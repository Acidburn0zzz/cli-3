package runtime

import (
	"github.com/ActiveState/cli/pkg/platform/runtime2/rterrs"
	"github.com/ActiveState/cli/pkg/project"
)

type EnvProvider interface {
	Environ() (map[string]string, error)
}

type Runtime struct {
	ep EnvProvider
}

// New is the constructor function for alternative runtimes
func New(proj *project.Project) (*Runtime, error) {
	return nil, rterrs.NotImplemented
}

func (r *Runtime) Environ() (map[string]string, error) {
	return r.ep.Environ()
}
