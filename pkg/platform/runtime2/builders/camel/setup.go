package camel

import (
	"github.com/ActiveState/cli/pkg/platform/runtime2/build"
	"github.com/ActiveState/cli/pkg/platform/runtime2/setup"
)

var _ setup.Setuper = &Setup{}
var _ setup.ArtifactSetuper = &ArtifactSetup{}

type Setup struct {
}

func NewSetup() *Setup {
	return &Setup{}
}

func (s *Setup) PostInstall() error {
	panic("implement me")
}

type ArtifactSetup struct {
}

func NewArtifactSetup(artifactID build.ArtifactID) *ArtifactSetup {
	return &ArtifactSetup{}
}

func (as *ArtifactSetup) NeedsSetup() bool {
	panic("implement me")
}

func (as *ArtifactSetup) Move(from string) error {
	panic("implement me")
}

func (as *ArtifactSetup) MetaDataCollection() error {
	panic("implement me")
}

func (as *ArtifactSetup) Relocate() error {
	panic("implement me")
}
