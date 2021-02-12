package runtime

import (
	"errors"
	"sync"

	"github.com/ActiveState/cli/pkg/platform/api/buildlogstream"
	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
	"github.com/ActiveState/cli/pkg/platform/runtime2/build"
	"github.com/ActiveState/cli/pkg/platform/runtime2/builders/alternative"
	"github.com/ActiveState/cli/pkg/platform/runtime2/rterrs"
	"github.com/ActiveState/cli/pkg/project"
)

// ErrNotInstalled is returned when the runtime is not locally installed yet.
// See the `setup.Setup` on how to set up a runtime installation.
var ErrNotInstalled = errors.New("Runtime not installed yet")

// IsNotInstalledError is a convenience function to checks if an error is NotInstalledError
func IsNotInstalledError(err error) bool {
	return errors.Is(err, ErrNotInstalled)
}

type ClientProvider interface {
	Solve() (*inventory_models.Order, error)
	Build(*inventory_models.Order) (*build.Result, error)
	BuildLog(msgHandler buildlogstream.MessageHandler, recipe *inventory_models.Recipe) (build.Logger, error)
}

// Setuper is the interface for an implementation of runtime setup functions
// These need to be specialized for each BuildEngine type
type Setuper interface {
	PostInstall() error
}

// ArtifactSetuper is the interface for an implementation of artifact setup functions
// These need to be specialized for each BuildEngine type
type ArtifactSetuper interface {
	NeedsSetup() bool
	Move(tmpInstallDir string) error
	MetaDataCollection() error
	Relocate() error
}

// maximum number of parallel artifact installations
const maxConcurrency = 3

// Setup provides methods to setup a fully-function runtime that *only* requires interactions with the local file system.
type Setup struct {
	client     ClientProvider
	msgHandler build.MessageHandler
}

// New returns a new Setup instance that can install a Runtime locally on the machine.
func NewSetup(project *project.Project, msgHandler build.MessageHandler, api ClientProvider) *Setup {
	panic("implement me")
}

// InstalledRuntime returns a locally installed Runtime instance.
//
// If the runtime cannot be initialized a NotInstalledError is returned.
func (s *Setup) InstalledRuntime() (*Runtime, error) {
	// check if complete installation can be found locally or:
	//   return ErrNotInstalled
	// next: try to load from local installation
	return nil, rterrs.NotImplemented
}

// InstallRuntime installs the runtime locally, such that it can be retrieved with the InstalledRuntime function afterwards.
func (s *Setup) InstallRuntime() error {
	// Get order for commit
	order, err := s.client.Solve()
	if err != nil {
		return err
	}

	// Request build
	buildResult, err := s.client.Build(order)
	if err != nil {
		return err
	}

	// Create the setup implementation based on the build engine (alternative or camel)
	var setupImpl Setuper
	setupImpl = s.selectSetupImplementation(buildResult.Engine)

	// Compute and handle the change summary
	artifacts := build.MakeArtifactsFromRecipe(buildResult.Recipe)
	requestedArtifacts, changedArtifacts := s.changeSummaryArgs(buildResult)
	s.msgHandler.ChangeSummary(artifacts, requestedArtifacts, changedArtifacts)

	// Access the build log to receive build updates.
	// Note: This may not actually connect to the build log if the build has already finished.
	buildLog, err := s.client.BuildLog(s.msgHandler, buildResult.Recipe)
	if err != nil {
		return err
	}
	defer buildLog.Wait()

	// wait for artifacts to be built and set them up in parallel with maximum concurrency
	ready := buildLog.Built()
	var wg sync.WaitGroup
	for i := 0; i < maxConcurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for a := range ready {
				// setup
				s.setupArtifact(buildResult.Engine, a)
			}
		}()
	}
	wg.Wait()

	err = <-buildLog.Err()
	if err != nil {
		return err
	}

	setupImpl.PostInstall()

	return rterrs.NotImplemented
}

func (s *Setup) InstallUninstalled(r *Runtime, err error) (*Runtime, error) {
	if err == nil || !IsNotInstalledError(err) {
		return r, err
	}
	if err := s.InstallRuntime(); err != nil {
		return nil, err
	}
	return s.InstalledRuntime()
}

// setupArtifact sets up artifact
// The artifact is downloaded, unpacked and then processed by the artifact setup implementation
func (s *Setup) setupArtifact(buildEngine build.Engine, a build.ArtifactID) {
	as := s.selectArtifactSetupImplementation(buildEngine, a)
	if !as.NeedsSetup() {
		return
	}

	tarball := s.downloadArtifactTarball(a)

	unpackedDir := s.unpackTarball(tarball)

	as.Move(unpackedDir)
	as.MetaDataCollection()
	as.Relocate()

	panic("implement error handling")
}

func (s *Setup) changeSummaryArgs(buildResult *build.Result) (requested build.Changes, changed build.Changes) {
	panic("implement me")
}

// downloadArtifactTarball retrieves the tarball for an artifactID
// Note: the tarball may also be retrieved from a local cache directory if that is available.
func (s *Setup) downloadArtifactTarball(artifactID build.ArtifactID) string {
	s.msgHandler.ArtifactDownloadStarting("artifactName")
	panic("implement me")
}

func (s *Setup) unpackTarball(tarballPath string) string {
	panic("implement me")
}

func (s *Setup) selectSetupImplementation(buildEngine build.Engine) Setuper {
	if buildEngine == build.Alternative {
		return alternative.NewSetup()
	}
	panic("implement me")
}

func (s *Setup) selectArtifactSetupImplementation(buildEngine build.Engine, a build.ArtifactID) ArtifactSetuper {
	if buildEngine == build.Alternative {
		return alternative.NewArtifactSetup(a)
	}
	panic("implement me")
}
