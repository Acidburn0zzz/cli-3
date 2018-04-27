// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// OrganizationEditableAddOns organization editable add ons
// swagger:model organizationEditableAddOns
type OrganizationEditableAddOns map[string]AddOn

// Validate validates this organization editable add ons
func (m OrganizationEditableAddOns) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("", "body", OrganizationEditableAddOns(m)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}