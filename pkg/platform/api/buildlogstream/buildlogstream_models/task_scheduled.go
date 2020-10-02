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

// TaskScheduled Task Scheduled
//
// A message indicating that a task will run as part of a build.
//
// swagger:model taskScheduled
type TaskScheduled struct {

	// task
	// Required: true
	Task *string `json:"task"`
}

// Validate validates this task scheduled
func (m *TaskScheduled) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskScheduled) validateTask(formats strfmt.Registry) error {

	if err := validate.Required("task", "body", m.Task); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskScheduled) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskScheduled) UnmarshalBinary(b []byte) error {
	var res TaskScheduled
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
