package runtime

import (
	"testing"

	"github.com/ActiveState/cli/pkg/platform/apiclient"
	"github.com/ActiveState/cli/pkg/platform/runtime2/build"
	"github.com/ActiveState/cli/pkg/project"
)

// TestSetup
func TestSetup(t *testing.T) {
	var mockProject *project.Project
	var mockMessageHandler build.MessageHandler
	mockClient := apiclient.NewMock()
	// specify behavior of mock client here.
	// ...

	{
		s := NewSetup(mockProject, mockMessageHandler, mockClient)

		s.InstallRuntime()
		// TODO: check error

		// TODO: check messageHandler calls

		r, _ := s.InstalledRuntime()
		// TODO: check runtime works
		r.Environ()
	}

	// or...

	{
		s := NewSetup(mockProject, mockMessageHandler, mockClient)
		r, err := s.InstallUninstalled(s.InstalledRuntime())
		t.Log(err)
		t.Log(r.Environ())
	}
}
