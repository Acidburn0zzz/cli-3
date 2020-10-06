// Code generated by go-swagger; DO NOT EDIT.

package buildlogstream_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildFailed Build Failed
//
// A message indicating that a build failed.
//
// swagger:model BuildFailed
type BuildFailed struct {

	// All of the errors from the failed build. Note that these errors may not be suitable for presenting to users and should simply be logged for further investigation.
	// Required: true
	Errors []string `json:"errors"`

	// If true this failed build can be retried and it may succeed. If false, retrying this failed build will not change the outcome. This field's value is only valid when type is build_failed.
	// Required: true
	IsRetryable bool `json:"is_retryable"`

	// An S3 URI containing the log for this build.
	// Format: uri
	LogURI strfmt.URI `json:"log_uri,omitempty"`

	// A user-facing message describing the build results.
	Message string `json:"message,omitempty"`
}

// Validate validates this build failed
func (m *BuildFailed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsRetryable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildFailed) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	return nil
}

func (m *BuildFailed) validateIsRetryable(formats strfmt.Registry) error {

	if err := validate.Required("is_retryable", "body", bool(m.IsRetryable)); err != nil {
		return err
	}

	return nil
}

func (m *BuildFailed) validateLogURI(formats strfmt.Registry) error {

	if swag.IsZero(m.LogURI) { // not required
		return nil
	}

	if err := validate.FormatOf("log_uri", "body", "uri", m.LogURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildFailed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildFailed) UnmarshalBinary(b []byte) error {
	var res BuildFailed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
