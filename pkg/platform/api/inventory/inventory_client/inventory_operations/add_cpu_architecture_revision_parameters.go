// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// NewAddCPUArchitectureRevisionParams creates a new AddCPUArchitectureRevisionParams object
// with the default values initialized.
func NewAddCPUArchitectureRevisionParams() *AddCPUArchitectureRevisionParams {
	var ()
	return &AddCPUArchitectureRevisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCPUArchitectureRevisionParamsWithTimeout creates a new AddCPUArchitectureRevisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCPUArchitectureRevisionParamsWithTimeout(timeout time.Duration) *AddCPUArchitectureRevisionParams {
	var ()
	return &AddCPUArchitectureRevisionParams{

		timeout: timeout,
	}
}

// NewAddCPUArchitectureRevisionParamsWithContext creates a new AddCPUArchitectureRevisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCPUArchitectureRevisionParamsWithContext(ctx context.Context) *AddCPUArchitectureRevisionParams {
	var ()
	return &AddCPUArchitectureRevisionParams{

		Context: ctx,
	}
}

// NewAddCPUArchitectureRevisionParamsWithHTTPClient creates a new AddCPUArchitectureRevisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCPUArchitectureRevisionParamsWithHTTPClient(client *http.Client) *AddCPUArchitectureRevisionParams {
	var ()
	return &AddCPUArchitectureRevisionParams{
		HTTPClient: client,
	}
}

/*AddCPUArchitectureRevisionParams contains all the parameters to send to the API endpoint
for the add Cpu architecture revision operation typically these are written to a http.Request
*/
type AddCPUArchitectureRevisionParams struct {

	/*CPUArchitectureID*/
	CPUArchitectureID strfmt.UUID
	/*CPUArchitectureRevision*/
	CPUArchitectureRevision *inventory_models.RevisionedFeatureProvider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) WithTimeout(timeout time.Duration) *AddCPUArchitectureRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) WithContext(ctx context.Context) *AddCPUArchitectureRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) WithHTTPClient(client *http.Client) *AddCPUArchitectureRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCPUArchitectureID adds the cPUArchitectureID to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) WithCPUArchitectureID(cPUArchitectureID strfmt.UUID) *AddCPUArchitectureRevisionParams {
	o.SetCPUArchitectureID(cPUArchitectureID)
	return o
}

// SetCPUArchitectureID adds the cpuArchitectureId to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) SetCPUArchitectureID(cPUArchitectureID strfmt.UUID) {
	o.CPUArchitectureID = cPUArchitectureID
}

// WithCPUArchitectureRevision adds the cPUArchitectureRevision to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) WithCPUArchitectureRevision(cPUArchitectureRevision *inventory_models.RevisionedFeatureProvider) *AddCPUArchitectureRevisionParams {
	o.SetCPUArchitectureRevision(cPUArchitectureRevision)
	return o
}

// SetCPUArchitectureRevision adds the cpuArchitectureRevision to the add Cpu architecture revision params
func (o *AddCPUArchitectureRevisionParams) SetCPUArchitectureRevision(cPUArchitectureRevision *inventory_models.RevisionedFeatureProvider) {
	o.CPUArchitectureRevision = cPUArchitectureRevision
}

// WriteToRequest writes these params to a swagger request
func (o *AddCPUArchitectureRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cpu_architecture_id
	if err := r.SetPathParam("cpu_architecture_id", o.CPUArchitectureID.String()); err != nil {
		return err
	}

	if o.CPUArchitectureRevision != nil {
		if err := r.SetBodyParam(o.CPUArchitectureRevision); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
