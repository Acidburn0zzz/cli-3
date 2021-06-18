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

// NewAddKernelGPUArchitectureParams creates a new AddKernelGPUArchitectureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddKernelGPUArchitectureParams() *AddKernelGPUArchitectureParams {
	return &AddKernelGPUArchitectureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddKernelGPUArchitectureParamsWithTimeout creates a new AddKernelGPUArchitectureParams object
// with the ability to set a timeout on a request.
func NewAddKernelGPUArchitectureParamsWithTimeout(timeout time.Duration) *AddKernelGPUArchitectureParams {
	return &AddKernelGPUArchitectureParams{
		timeout: timeout,
	}
}

// NewAddKernelGPUArchitectureParamsWithContext creates a new AddKernelGPUArchitectureParams object
// with the ability to set a context for a request.
func NewAddKernelGPUArchitectureParamsWithContext(ctx context.Context) *AddKernelGPUArchitectureParams {
	return &AddKernelGPUArchitectureParams{
		Context: ctx,
	}
}

// NewAddKernelGPUArchitectureParamsWithHTTPClient creates a new AddKernelGPUArchitectureParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddKernelGPUArchitectureParamsWithHTTPClient(client *http.Client) *AddKernelGPUArchitectureParams {
	return &AddKernelGPUArchitectureParams{
		HTTPClient: client,
	}
}

/* AddKernelGPUArchitectureParams contains all the parameters to send to the API endpoint
   for the add kernel g p u architecture operation.

   Typically these are written to a http.Request.
*/
type AddKernelGPUArchitectureParams struct {

	// GpuArchitectureID.
	GpuArchitectureID *inventory_models.AddKernelGPUArchitectureParamsBody

	// KernelID.
	//
	// Format: uuid
	KernelID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add kernel g p u architecture params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddKernelGPUArchitectureParams) WithDefaults() *AddKernelGPUArchitectureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add kernel g p u architecture params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddKernelGPUArchitectureParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) WithTimeout(timeout time.Duration) *AddKernelGPUArchitectureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) WithContext(ctx context.Context) *AddKernelGPUArchitectureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) WithHTTPClient(client *http.Client) *AddKernelGPUArchitectureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGpuArchitectureID adds the gpuArchitectureID to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) WithGpuArchitectureID(gpuArchitectureID *inventory_models.AddKernelGPUArchitectureParamsBody) *AddKernelGPUArchitectureParams {
	o.SetGpuArchitectureID(gpuArchitectureID)
	return o
}

// SetGpuArchitectureID adds the gpuArchitectureId to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) SetGpuArchitectureID(gpuArchitectureID *inventory_models.AddKernelGPUArchitectureParamsBody) {
	o.GpuArchitectureID = gpuArchitectureID
}

// WithKernelID adds the kernelID to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) WithKernelID(kernelID strfmt.UUID) *AddKernelGPUArchitectureParams {
	o.SetKernelID(kernelID)
	return o
}

// SetKernelID adds the kernelId to the add kernel g p u architecture params
func (o *AddKernelGPUArchitectureParams) SetKernelID(kernelID strfmt.UUID) {
	o.KernelID = kernelID
}

// WriteToRequest writes these params to a swagger request
func (o *AddKernelGPUArchitectureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.GpuArchitectureID != nil {
		if err := r.SetBodyParam(o.GpuArchitectureID); err != nil {
			return err
		}
	}

	// path param kernel_id
	if err := r.SetPathParam("kernel_id", o.KernelID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
