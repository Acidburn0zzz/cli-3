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

// NewAddCPUExtensionRevisionParams creates a new AddCPUExtensionRevisionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddCPUExtensionRevisionParams() *AddCPUExtensionRevisionParams {
	return &AddCPUExtensionRevisionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddCPUExtensionRevisionParamsWithTimeout creates a new AddCPUExtensionRevisionParams object
// with the ability to set a timeout on a request.
func NewAddCPUExtensionRevisionParamsWithTimeout(timeout time.Duration) *AddCPUExtensionRevisionParams {
	return &AddCPUExtensionRevisionParams{
		timeout: timeout,
	}
}

// NewAddCPUExtensionRevisionParamsWithContext creates a new AddCPUExtensionRevisionParams object
// with the ability to set a context for a request.
func NewAddCPUExtensionRevisionParamsWithContext(ctx context.Context) *AddCPUExtensionRevisionParams {
	return &AddCPUExtensionRevisionParams{
		Context: ctx,
	}
}

// NewAddCPUExtensionRevisionParamsWithHTTPClient creates a new AddCPUExtensionRevisionParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddCPUExtensionRevisionParamsWithHTTPClient(client *http.Client) *AddCPUExtensionRevisionParams {
	return &AddCPUExtensionRevisionParams{
		HTTPClient: client,
	}
}

/* AddCPUExtensionRevisionParams contains all the parameters to send to the API endpoint
   for the add Cpu extension revision operation.

   Typically these are written to a http.Request.
*/
type AddCPUExtensionRevisionParams struct {

	// CPUExtensionID.
	//
	// Format: uuid
	CPUExtensionID strfmt.UUID

	// CPUExtensionRevision.
	CPUExtensionRevision *inventory_models.RevisionedFeatureProvider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add Cpu extension revision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCPUExtensionRevisionParams) WithDefaults() *AddCPUExtensionRevisionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add Cpu extension revision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddCPUExtensionRevisionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) WithTimeout(timeout time.Duration) *AddCPUExtensionRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) WithContext(ctx context.Context) *AddCPUExtensionRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) WithHTTPClient(client *http.Client) *AddCPUExtensionRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCPUExtensionID adds the cPUExtensionID to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) WithCPUExtensionID(cPUExtensionID strfmt.UUID) *AddCPUExtensionRevisionParams {
	o.SetCPUExtensionID(cPUExtensionID)
	return o
}

// SetCPUExtensionID adds the cpuExtensionId to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) SetCPUExtensionID(cPUExtensionID strfmt.UUID) {
	o.CPUExtensionID = cPUExtensionID
}

// WithCPUExtensionRevision adds the cPUExtensionRevision to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) WithCPUExtensionRevision(cPUExtensionRevision *inventory_models.RevisionedFeatureProvider) *AddCPUExtensionRevisionParams {
	o.SetCPUExtensionRevision(cPUExtensionRevision)
	return o
}

// SetCPUExtensionRevision adds the cpuExtensionRevision to the add Cpu extension revision params
func (o *AddCPUExtensionRevisionParams) SetCPUExtensionRevision(cPUExtensionRevision *inventory_models.RevisionedFeatureProvider) {
	o.CPUExtensionRevision = cPUExtensionRevision
}

// WriteToRequest writes these params to a swagger request
func (o *AddCPUExtensionRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cpu_extension_id
	if err := r.SetPathParam("cpu_extension_id", o.CPUExtensionID.String()); err != nil {
		return err
	}
	if o.CPUExtensionRevision != nil {
		if err := r.SetBodyParam(o.CPUExtensionRevision); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
