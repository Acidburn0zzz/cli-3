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

// AuthorCore Author Core
//
// The core fields of a author
//
// swagger:model authorCore
type AuthorCore struct {

	// The author's email address
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// The author's name
	// Required: true
	Name *string `json:"name"`

	// The author's websites
	// Required: true
	// Unique: true
	Websites []strfmt.URI `json:"websites"`
}

// Validate validates this author core
func (m *AuthorCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsites(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorCore) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuthorCore) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *AuthorCore) validateWebsites(formats strfmt.Registry) error {

	if err := validate.Required("websites", "body", m.Websites); err != nil {
		return err
	}

	if err := validate.UniqueItems("websites", "body", m.Websites); err != nil {
		return err
	}

	for i := 0; i < len(m.Websites); i++ {

		if err := validate.FormatOf("websites"+"."+strconv.Itoa(i), "body", "uri", m.Websites[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this author core based on context it is used
func (m *AuthorCore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AuthorCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorCore) UnmarshalBinary(b []byte) error {
	var res AuthorCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
