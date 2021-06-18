// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IngredientVersionPagedList Ingredient Version Paged List
//
// A paginated list of ingredient versions
//
// swagger:model ingredientVersionPagedList
type IngredientVersionPagedList struct {

	// A page of ingredient versions
	// Required: true
	IngredientVersions []*IngredientVersion `json:"ingredient_versions"`

	// links
	// Required: true
	Links *PagingLinks `json:"links"`

	// paging
	// Required: true
	Paging *Paging `json:"paging"`
}

// Validate validates this ingredient version paged list
func (m *IngredientVersionPagedList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIngredientVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaging(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IngredientVersionPagedList) validateIngredientVersions(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_versions", "body", m.IngredientVersions); err != nil {
		return err
	}

	for i := 0; i < len(m.IngredientVersions); i++ {
		if swag.IsZero(m.IngredientVersions[i]) { // not required
			continue
		}

		if m.IngredientVersions[i] != nil {
			if err := m.IngredientVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingredient_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IngredientVersionPagedList) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *IngredientVersionPagedList) validatePaging(formats strfmt.Registry) error {

	if err := validate.Required("paging", "body", m.Paging); err != nil {
		return err
	}

	if m.Paging != nil {
		if err := m.Paging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ingredient version paged list based on the context it is used
func (m *IngredientVersionPagedList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIngredientVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePaging(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IngredientVersionPagedList) contextValidateIngredientVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IngredientVersions); i++ {

		if m.IngredientVersions[i] != nil {
			if err := m.IngredientVersions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ingredient_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IngredientVersionPagedList) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *IngredientVersionPagedList) contextValidatePaging(ctx context.Context, formats strfmt.Registry) error {

	if m.Paging != nil {
		if err := m.Paging.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("paging")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IngredientVersionPagedList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IngredientVersionPagedList) UnmarshalBinary(b []byte) error {
	var res IngredientVersionPagedList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
