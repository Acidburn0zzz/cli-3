package alternative

import (
	"github.com/ActiveState/cli/pkg/platform/runtime2/build"
	"github.com/ActiveState/cli/pkg/platform/runtime2/rterrs"
)

//var _ setup.Setuper = &Setup{}
//var _ setup.ArtifactSetuper = &ArtifactSetup{}

type Setup struct {
}

type ArtifactSetup struct {
}

func NewSetup() *Setup {
	return &Setup{}
}

func NewArtifactSetup(artifactID build.ArtifactID) *ArtifactSetup {
	return &ArtifactSetup{}
}

func (s *Setup) PostInstall() error {
	return rterrs.NotImplemented
}

func (as *ArtifactSetup) NeedsSetup() bool {
	return true
}

func (as *ArtifactSetup) Move(from string) error {
	return rterrs.NotImplemented
}

func (as *ArtifactSetup) MetaDataCollection() error {
	return rterrs.NotImplemented
}

func (as *ArtifactSetup) Relocate() error {
	return rterrs.NotImplemented
}
