// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GitRepoRequirements git repo requirements
//
// swagger:model GitRepoRequirements
type GitRepoRequirements struct {

	// language
	Language string `json:"language,omitempty"`

	// requirements
	Requirements string `json:"requirements,omitempty"`
}

// Validate validates this git repo requirements
func (m *GitRepoRequirements) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this git repo requirements based on context it is used
func (m *GitRepoRequirements) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GitRepoRequirements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GitRepoRequirements) UnmarshalBinary(b []byte) error {
	var res GitRepoRequirements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
