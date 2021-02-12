package build

import (
	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// ArtifactID represents an artifact ID
type ArtifactID string

type Artifact struct {
	ArtifactID   ArtifactID
	Name         string
	Dependencies []ArtifactID
	// ...
}

func MakeArtifactsFromRecipe(recipe *inventory_models.Recipe) map[ArtifactID]Artifact {
	panic("implement me")
}

type Changes struct {
	Added   []ArtifactID
	Updated []ArtifactID
	Removed []ArtifactID
}
