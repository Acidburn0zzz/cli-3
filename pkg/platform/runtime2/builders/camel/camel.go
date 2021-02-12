package camel

import (
	"github.com/ActiveState/cli/pkg/platform/runtime2/rterrs"
	"github.com/ActiveState/cli/pkg/project"
)

//var _ runtime.EnvProvider = &Camel{}

// Camel is the specialization of a runtime for Camel builds
type Camel struct{}

// New is the constructor function for Camel runtimes
func New(proj *project.Project) (*Camel, error) {
	return nil, rterrs.NotImplemented
}

func (a *Camel) Environ() (map[string]string, error) {
	return nil, rterrs.NotImplemented
}
