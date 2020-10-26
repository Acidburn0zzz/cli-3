// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems Build Flag Condition Sub Schema
//
// A build flag name and value that must be included in a recipe for the containing entity to apply. If nothing in the recipe matches this condition, the containing entity is disable/cannot be used.
// swagger:model v1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems
type V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems struct {

	// Whether the build flag value must match this condition's value (eq) in order to satisfy the condition or must not match (ne)
	// Required: true
	// Enum: [eq ne]
	Comparator *string `json:"comparator"`

	// The name of the build flag conditioned upon
	// Required: true
	Name *string `json:"name"`

	// The value of the build flag conditioned upon
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this v1 recipe response recipes items resolved ingredients items ingredient version all of3 all of0 ingredient options items condition sets items build flag conditions items
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItemsTypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItemsTypeComparatorPropEnum = append(v1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItemsTypeComparatorPropEnum, v)
	}
}

const (

	// V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItemsComparatorEq captures enum value "eq"
	V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItemsComparatorEq string = "eq"

	// V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItemsComparatorNe captures enum value "ne"
	V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItemsComparatorNe string = "ne"
)

// prop value enum
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItemsTypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems) UnmarshalBinary(b []byte) error {
	var res V1RecipeResponseRecipesItemsResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}