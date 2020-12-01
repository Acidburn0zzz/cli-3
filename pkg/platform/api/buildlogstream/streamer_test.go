package buildlogstream

import (
	"fmt"
	"testing"

	"github.com/autarch/testify/assert"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

func intToUUID(number int) strfmt.UUID {
	return strfmt.UUID(fmt.Sprintf("00000000-0000-0000-0000-%012d", number))
}

func intMapToUUIDMap(in map[int][]int) map[strfmt.UUID][]strfmt.UUID {
	out := map[strfmt.UUID][]strfmt.UUID{}
	for k := range in {
		kk := intToUUID(k)
		out[kk] = []strfmt.UUID{}
		for _, vv := range in[k] {
			out[kk] = append(out[kk], intToUUID(vv))
		}
	}
	return out
}

func intsToArtifactMap(in []int) map[strfmt.UUID]ArtifactMapping {
	out := map[strfmt.UUID]ArtifactMapping{}
	for _, v := range in {
		out[intToUUID(v)] = ArtifactMapping{}
	}
	return out
}

type depGraph map[int][]int

func depGraphsToResolvedIngredients(dgs depGraph) []*inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItems {
	pn := "language"

	res := make([]*inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItems, 0, len(dgs))
	for d, dchildren := range dgs {
		uuid := intToUUID(d)
		deps := make([]*inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItemsDependenciesItems, 0, len(dchildren))
		for _, dc := range dchildren {
			duuid := intToUUID(dc)

			deps = append(deps, &inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItemsDependenciesItems{
				IngredientVersionID: &duuid,
			})
		}
		name := fmt.Sprintf("pkg%02d", d)
		resolved := &inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItems{
			Ingredient: &inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredient{
				V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1: inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1{
					V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1AllOf0: inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientAllOf1AllOf0{
						Name:             &name,
						PrimaryNamespace: &pn,
					},
				},
			},
			IngredientVersion: &inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersion{
				V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf0: inventory_models.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf0{
					IngredientVersionID: &uuid,
				},
			},
			Dependencies: deps,
		}
		res = append(res, resolved)
	}
	return res
}

func TestFetchDepTree(t *testing.T) {
	dg := depGraph{
		1: {11, 12, 19, 2, 900},
		2: {21, 29, 900},
		3: {31, 1, 39, 900},
	}
	ingredients := depGraphsToResolvedIngredients(dg)
	ingredientMap := intsToArtifactMap([]int{1, 2, 3, 11, 12, 900, 21, 31})

	depTree, recursive := fetchDepTree(ingredients, ingredientMap)

	expectedDirect := intMapToUUIDMap(map[int][]int{
		1: {11, 12, 2, 900},
		2: {21, 900},
		3: {31, 1, 900},
	})

	expectedRecursive := intMapToUUIDMap(map[int][]int{
		1: {11, 12, 2, 21, 900},
		2: {21, 900},
		3: {31, 1, 11, 12, 2, 21, 900},
	})

	assert.Equal(t, expectedDirect, depTree)
	assert.Equal(t, expectedRecursive, recursive)
}

func TestFetchRecursiveDepTree(t *testing.T) {
	dg := depGraph{
		1: {2},
		2: {1},
	}
	ingredients := depGraphsToResolvedIngredients(dg)
	ingredientMap := intsToArtifactMap([]int{1, 2})

	depTree, recursive := fetchDepTree(ingredients, ingredientMap)

	expectedDirect := intMapToUUIDMap(map[int][]int{
		1: {2},
		2: {1},
	})

	expectedRecursive := intMapToUUIDMap(map[int][]int{
		1: {2, 1},
		2: {1, 2},
	})

	assert.Equal(t, expectedDirect, depTree)
	assert.Equal(t, expectedRecursive, recursive)
}
