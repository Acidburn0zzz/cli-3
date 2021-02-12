// Package impl is the top-level package for all runtime implementations.
// Currently, the ActiveState Platform supports two build engines (Camel and
// Alternative), their implementations can be found in sub-packages relative to
// this directory.
package build

import (
	"github.com/ActiveState/cli/pkg/platform/api/buildlogstream"
	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// MessageHandler is the interface for callback functions that are called during
// runtime set-up when progress messages can be forwarded to the user
type MessageHandler interface {
	buildlogstream.MessageHandler

	// ChangeSummary summarizes the changes to the current project during the InstallRuntime() call.
	// This summary is printed as soon as possible, providing the State Tool user with an idea of the complexity of the requested build.
	// The arguments are for the changes introduced in the latest commit that this Setup is setting up.
	// TODO: Decide if we want to have a method to de-activate the change summary for activations where it does not make sense.
	ChangeSummary(artifacts map[ArtifactID]Artifact, requested Changes, changed Changes)
	ArtifactDownloadStarting(artifactName string)
	ArtifactDownloadCompleted(artifactName string, number, total int)
	ArtifactDownloadFailed(artifactName string, errorMsg string)
}

// Engine describes the build engine that was used to build the runtime
type Engine int

const (
	// UnknownEngine represents an engine unknown to the runtime.
	UnknownEngine Engine = iota

	// Camel is the legacy build engine, that builds Active{Python,Perl,Tcl}
	// distributions
	Camel

	// Alternative is the new alternative build orchestration framework
	Alternative

	// Hybrid wraps Camel.
	Hybrid
)

// Result is the unified response of a Build request
type Result struct {
	Engine Engine
	Recipe *inventory_models.Recipe
}

// Logger is an interface to communicate with the build log streamer
type Logger interface {
	Wait()
	Close()
	// BuiltArtifactChannel returns a channel of artifact IDs for artifacts that are ready to be downloaded.
	Built() chan ArtifactID
	Err() <-chan error
}
