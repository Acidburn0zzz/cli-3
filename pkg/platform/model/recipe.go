package model

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/sysinfo"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/retryhttp"
	"github.com/ActiveState/cli/pkg/platform/api"
	"github.com/ActiveState/cli/pkg/platform/api/inventory"
	iop "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_client/inventory_operations"
	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
)

// Fail types for this package
var (
	FailOrderRecipes   = failures.Type("model.fail.orderrecipes", api.FailUnknown)
	FailRecipeNotFound = failures.Type("model.fail.recipe.notfound", failures.FailNonFatal)

	FailUnsupportedPlatform = failures.Type("model.fail.unsupportedplatform")
	FailNoRecipes           = failures.Type("model.fail.norecipes", api.FailNotFound)
)

// HostPlatform stores a reference to current platform
var HostPlatform string

// Recipe aliases recipe model
type Recipe = inventory_models.V1RecipeResponseRecipesItems

// OrderAnnotations are sent with every order for analytical purposes described here:
// https://docs.google.com/document/d/1nXeNCRWX-4ULtk20t3C7kTZDCSBJchJJHhzz-lQuVsU/edit#heading=h.o93wm4bt5ul9
type OrderAnnotations struct {
	CommitID     string `json:"commit_id"`
	Project      string `json:"project"`
	Organization string `json:"organization"`
}

func init() {
	HostPlatform = sysinfo.OS().String()
}

// FetchRawRecipeForCommit returns a recipe from a project based off a commitID
func FetchRawRecipeForCommit(commitID strfmt.UUID, owner, project string) (string, *failures.Failure) {
	return fetchRawRecipe(commitID, owner, project, nil)
}

// FetchRawRecipeForCommitAndPlatform returns a recipe from a project based off a commitID and platform
func FetchRawRecipeForCommitAndPlatform(commitID strfmt.UUID, owner, project string, platform string) (string, *failures.Failure) {
	return fetchRawRecipe(commitID, owner, project, &platform)
}

// FetchRecipeIDForCommit returns a recipe ID for a project based on the given commitID and the current platform
func FetchRecipeIDForCommit(commitID strfmt.UUID, owner, project, orgID string, private bool) (*strfmt.UUID, *failures.Failure) {
	return fetchRecipeID(commitID, owner, project, orgID, private, &HostPlatform)
}

// FetchRecipeIDForCommitAndPlatform returns a recipe ID for a project based on the given commitID and platform string
func FetchRecipeIDForCommitAndPlatform(commitID strfmt.UUID, owner, project, orgID string, private bool, hostPlatform string) (*strfmt.UUID, *failures.Failure) {
	return fetchRecipeID(commitID, owner, project, orgID, private, &hostPlatform)
}

func FetchRecipeByID(recipeID strfmt.UUID) (*inventory_models.Recipe, error) {
	params := iop.NewGetSolutionRecipeParams()
	params.RecipeID = recipeID

	client, _ := inventory.Init()
	recipe, err := client.GetSolutionRecipe(params, authentication.ClientAuth())
	if err != nil {
		return nil, errs.Wrap(err, "Unknown error while retrieving full order, error: %v, recipe: %s", err, string(recipeID))
	}

	return recipe.Payload.Recipe, nil
}

func ResolveRecipe(commitID strfmt.UUID, owner, projectName string, project *mono_models.Project) (*inventory_models.Recipe, error) {
	if commitID == "" {
		return nil, locale.NewError("err_no_commit", "Missing Commit ID")
	}

	private := false
	orgID := strfmt.UUID(constants.ValidZeroUUID)

	if project != nil {
		orgID = project.OrganizationID
		private = project.Private
	}

	recipeID, fail := FetchRecipeIDForCommitAndPlatform(commitID, owner, projectName, orgID.String(), private, HostPlatform)
	if fail != nil {
		return nil, fail.ToError()
	}

	if recipeID == nil {
		return nil, errs.New("recipeID is nil")
	}

	recipe, err := FetchRecipeByID(*recipeID)
	if err != nil {
		return nil, errs.Wrap(err, "Could not fetch recipe")
	}

	if recipe.RecipeID == nil {
		return nil, errs.New("Resulting recipe does not have a recipeID")
	}

	return recipe, nil
}

func fetchRawRecipe(commitID strfmt.UUID, owner, project string, hostPlatform *string) (string, *failures.Failure) {
	_, transport := inventory.Init()

	var err error
	params := iop.NewResolveRecipesParams()
	params.SetHTTPClient(retryhttp.DefaultClient.StandardClient())
	params.SetTimeout(time.Second * 60)
	params.Order, err = commitToOrder(commitID, owner, project)
	if err != nil {
		return "", FailOrderRecipes.Wrap(err)
	}

	var fail *failures.Failure
	if hostPlatform != nil {
		params.Order.Platforms, fail = filterPlatformIDs(*hostPlatform, runtime.GOARCH, params.Order.Platforms)
		if fail != nil {
			return "", fail
		}
	}

	recipe, err := inventory.ResolveRecipes(transport, params, authentication.ClientAuth())
	if err != nil {
		if err == context.DeadlineExceeded {
			return "", FailOrderRecipes.New("request_timed_out")
		}

		orderBody, err2 := json.Marshal(params.Order)
		if err2 != nil {
			orderBody = []byte(fmt.Sprintf("Could not marshal order, error: %v", err2))
		}
		switch rrErr := err.(type) {
		case *iop.ResolveRecipesDefault:
			msg := *rrErr.Payload.Message
			logging.Error("Could not resolve order, error: %s, order: %s", msg, string(orderBody))
			return "", FailOrderRecipes.New("err_solve_order", msg)
		case *iop.ResolveRecipesBadRequest:
			msg := *rrErr.Payload.Message
			logging.Error("Bad request while resolving order, error: %s, order: %s", msg, string(orderBody))
			return "", FailOrderRecipes.New("err_order_bad_request", commitID.String(), msg)
		default:
			logging.Error("Unknown error while resolving order, error: %v, order: %s", err, string(orderBody))
			return "", FailOrderRecipes.Wrap(err, "err_order_unknown")
		}
	}

	return recipe, nil
}

func commitToOrder(commitID strfmt.UUID, owner, project string) (*inventory_models.Order, error) {
	monoOrder, err := FetchOrderFromCommit(commitID)
	if err != nil {
		return nil, FailOrderRecipes.Wrap(err, locale.T("err_order_recipe")).ToError()
	}

	orderData, err := monoOrder.MarshalBinary()
	if err != nil {
		return nil, failures.FailMarshal.New(locale.T("err_order_marshal")).ToError()
	}

	order := &inventory_models.Order{}
	err = order.UnmarshalBinary(orderData)
	if err != nil {
		return nil, failures.FailMarshal.New(locale.T("err_order_marshal")).ToError()
	}

	order.Annotations = OrderAnnotations{
		CommitID:     commitID.String(),
		Project:      project,
		Organization: owner,
	}

	return order, nil
}

func fetchRecipeID(commitID strfmt.UUID, owner, project, orgID string, private bool, hostPlatform *string) (*strfmt.UUID, *failures.Failure) {
	var err error
	params := iop.NewSolveOrderParams()
	params.SetHTTPClient(retryhttp.DefaultClient.StandardClient())
	params.SetTimeout(time.Second * 60)
	params.Order, err = commitToOrder(commitID, owner, project)
	if err != nil {
		return nil, FailOrderRecipes.Wrap(err)
	}

	client, _ := inventory.Init()

	response, err := client.SolveOrder(params, authentication.ClientAuth())
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, FailOrderRecipes.New("request_timed_out")
		}

		orderBody, _ := json.Marshal(params.Order)
		switch rrErr := err.(type) {
		case *iop.SolveOrderDefault:
			msg := *rrErr.Payload.Message
			logging.Error("Could not solve order, error: %s, order: %s", msg, string(orderBody))
			return nil, FailOrderRecipes.New("err_solve_order", msg)
		case *iop.SolveOrderBadRequest:
			msg := *rrErr.Payload.Message
			logging.Error("Bad request while resolving order, error: %s, order: %s", msg, string(orderBody))
			return nil, FailOrderRecipes.New("err_order_bad_request", commitID.String(), msg)
		default:
			logging.Error("Unknown error while resolving order, error: %v, order: %s", err, string(orderBody))
			return nil, FailOrderRecipes.Wrap(err, "err_order_unknown")
		}
	}

	platformIDs, fail := filterPlatformIDs(*hostPlatform, runtime.GOARCH, params.Order.Platforms)
	if fail != nil {
		return nil, FailOrderRecipes.Wrap(fail)
	}
	if len(platformIDs) == 0 {
		return nil, FailOrderRecipes.New("err_recipe_no_platform")
	} else if len(platformIDs) > 1 {
		logging.Debug("Received multiple platform IDs.  Picking the first one.")
	}
	platformID := platformIDs[0].String()

	if _, ok := response.Payload[platformID]; !ok {
		return nil, FailOrderRecipes.New("err_recipe_not_found")
	}
	return response.Payload[platformID].RecipeID, nil
}

func IngredientVersionMap(recipe *inventory_models.Recipe) map[strfmt.UUID]*inventory_models.ResolvedIngredient {
	ingredientVersionMap := map[strfmt.UUID]*inventory_models.ResolvedIngredient{}

	for _, re := range recipe.ResolvedIngredients {
		if re.Ingredient.PrimaryNamespace != nil && (*re.Ingredient.PrimaryNamespace == "builder" || *re.Ingredient.PrimaryNamespace == "builder-lib" || *re.Ingredient.PrimaryNamespace == string(NamespaceLanguage)) {
			continue
		}

		if re.IngredientVersion == nil || re.IngredientVersion.IngredientVersionID == nil {
			continue
		}
		ingredientVersionMap[*re.IngredientVersion.IngredientVersionID] = re
	}
	return ingredientVersionMap
}

func ArtifactMap(recipe *inventory_models.Recipe) map[strfmt.UUID]*inventory_models.ResolvedIngredient {
	artifactMap := map[strfmt.UUID]*inventory_models.ResolvedIngredient{}

	for _, re := range recipe.ResolvedIngredients {
		if re.Ingredient.PrimaryNamespace != nil && (*re.Ingredient.PrimaryNamespace == "builder" || *re.Ingredient.PrimaryNamespace == "builder-lib" || *re.Ingredient.PrimaryNamespace == string(NamespaceLanguage)) {
			continue
		}

		artifactMap[re.ArtifactID] = re
	}
	return artifactMap
}

func ArtifactDescription(artifactID strfmt.UUID, artifactMap map[strfmt.UUID]*inventory_models.ResolvedIngredient) string {
	v, ok := artifactMap[artifactID]
	if !ok || v.Ingredient.Name == nil {
		return locale.Tl("unknown_artifact_description", "Artifact {{.V0}}", artifactID.String())
	}

	version := ""
	if v.IngredientVersion != nil {
		version = "@" + *v.IngredientVersion.Version
	}

	return *v.Ingredient.Name + version
}

func ParseDepTree(ingredients []*inventory_models.ResolvedIngredient, ingredientVersionMap map[strfmt.UUID]*inventory_models.ResolvedIngredient) (directdeptree map[strfmt.UUID][]strfmt.UUID, recursive map[strfmt.UUID][]strfmt.UUID) {
	directdeptree = map[strfmt.UUID][]strfmt.UUID{}
	for _, ingredient := range ingredients {
		if ingredient.IngredientVersion == nil || ingredient.IngredientVersion.IngredientVersionID == nil {
			continue
		}

		id := ingredient.IngredientVersion.IngredientVersionID
		// skip ingredients that are not mapped to artifacts
		if _, ok := ingredientVersionMap[*id]; !ok {
			continue
		}
		// Construct directdeptree entry
		if _, ok := directdeptree[*id]; !ok {
			directdeptree[*id] = []strfmt.UUID{}
		}

		// Add direct dependencies
		for _, dep := range ingredient.Dependencies {
			if dep.IngredientVersionID == nil {
				continue
			}
			// skip ingredients that are not mapped to artifacts
			if _, ok := ingredientVersionMap[*dep.IngredientVersionID]; !ok {
				continue
			}
			directdeptree[*id] = append(directdeptree[*id], *dep.IngredientVersionID)
		}
	}

	// Now resolve ALL dependencies, not just the direct ones
	deptree := map[strfmt.UUID][]strfmt.UUID{}
	for ingredientID := range directdeptree {
		deps := []strfmt.UUID{}
		deptree[ingredientID] = recursiveDeps(deps, directdeptree, ingredientID, map[strfmt.UUID]struct{}{})
	}

	return directdeptree, deptree
}

func recursiveDeps(deps []strfmt.UUID, directdeptree map[strfmt.UUID][]strfmt.UUID, id strfmt.UUID, skip map[strfmt.UUID]struct{}) []strfmt.UUID {
	if _, ok := directdeptree[id]; !ok {
		return deps
	}

	for _, dep := range directdeptree[id] {
		if _, ok := skip[dep]; ok {
			continue
		}
		skip[dep] = struct{}{}

		deps = append(deps, dep)
		deps = recursiveDeps(deps, directdeptree, dep, skip)
	}

	return deps
}
