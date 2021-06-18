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

// NewAddPatchParams creates a new AddPatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddPatchParams() *AddPatchParams {
	return &AddPatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddPatchParamsWithTimeout creates a new AddPatchParams object
// with the ability to set a timeout on a request.
func NewAddPatchParamsWithTimeout(timeout time.Duration) *AddPatchParams {
	return &AddPatchParams{
		timeout: timeout,
	}
}

// NewAddPatchParamsWithContext creates a new AddPatchParams object
// with the ability to set a context for a request.
func NewAddPatchParamsWithContext(ctx context.Context) *AddPatchParams {
	return &AddPatchParams{
		Context: ctx,
	}
}

// NewAddPatchParamsWithHTTPClient creates a new AddPatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddPatchParamsWithHTTPClient(client *http.Client) *AddPatchParams {
	return &AddPatchParams{
		HTTPClient: client,
	}
}

/* AddPatchParams contains all the parameters to send to the API endpoint
   for the add patch operation.

   Typically these are written to a http.Request.
*/
type AddPatchParams struct {

	// Patch.
	Patch *inventory_models.PatchCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPatchParams) WithDefaults() *AddPatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add patch params
func (o *AddPatchParams) WithTimeout(timeout time.Duration) *AddPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add patch params
func (o *AddPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add patch params
func (o *AddPatchParams) WithContext(ctx context.Context) *AddPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add patch params
func (o *AddPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add patch params
func (o *AddPatchParams) WithHTTPClient(client *http.Client) *AddPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add patch params
func (o *AddPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatch adds the patch to the add patch params
func (o *AddPatchParams) WithPatch(patch *inventory_models.PatchCore) *AddPatchParams {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the add patch params
func (o *AddPatchParams) SetPatch(patch *inventory_models.PatchCore) {
	o.Patch = patch
}

// WriteToRequest writes these params to a swagger request
func (o *AddPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
