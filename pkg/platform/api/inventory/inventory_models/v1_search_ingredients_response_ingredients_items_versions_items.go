// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SearchIngredientsResponseIngredientsItemsVersionsItems Search ingredients response version
// swagger:model v1SearchIngredientsResponseIngredientsItemsVersionsItems
type V1SearchIngredientsResponseIngredientsItemsVersionsItems struct {

	// is indemnified
	// Required: true
	IsIndemnified *bool `json:"is_indemnified"`

	// link
	// Format: uri
	Link strfmt.URI `json:"link,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this v1 search ingredients response ingredients items versions items
func (m *V1SearchIngredientsResponseIngredientsItemsVersionsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIsIndemnified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SearchIngredientsResponseIngredientsItemsVersionsItems) validateIsIndemnified(formats strfmt.Registry) error {

	if err := validate.Required("is_indemnified", "body", m.IsIndemnified); err != nil {
		return err
	}

	return nil
}

func (m *V1SearchIngredientsResponseIngredientsItemsVersionsItems) validateLink(formats strfmt.Registry) error {

	if swag.IsZero(m.Link) { // not required
		return nil
	}

	if err := validate.FormatOf("link", "body", "uri", m.Link.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SearchIngredientsResponseIngredientsItemsVersionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SearchIngredientsResponseIngredientsItemsVersionsItems) UnmarshalBinary(b []byte) error {
	var res V1SearchIngredientsResponseIngredientsItemsVersionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
