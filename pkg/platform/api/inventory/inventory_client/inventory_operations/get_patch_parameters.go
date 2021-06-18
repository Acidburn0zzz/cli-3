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
)

// NewGetPatchParams creates a new GetPatchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPatchParams() *GetPatchParams {
	return &GetPatchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatchParamsWithTimeout creates a new GetPatchParams object
// with the ability to set a timeout on a request.
func NewGetPatchParamsWithTimeout(timeout time.Duration) *GetPatchParams {
	return &GetPatchParams{
		timeout: timeout,
	}
}

// NewGetPatchParamsWithContext creates a new GetPatchParams object
// with the ability to set a context for a request.
func NewGetPatchParamsWithContext(ctx context.Context) *GetPatchParams {
	return &GetPatchParams{
		Context: ctx,
	}
}

// NewGetPatchParamsWithHTTPClient creates a new GetPatchParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPatchParamsWithHTTPClient(client *http.Client) *GetPatchParams {
	return &GetPatchParams{
		HTTPClient: client,
	}
}

/* GetPatchParams contains all the parameters to send to the API endpoint
   for the get patch operation.

   Typically these are written to a http.Request.
*/
type GetPatchParams struct {

	// PatchID.
	//
	// Format: uuid
	PatchID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPatchParams) WithDefaults() *GetPatchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get patch params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPatchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get patch params
func (o *GetPatchParams) WithTimeout(timeout time.Duration) *GetPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get patch params
func (o *GetPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get patch params
func (o *GetPatchParams) WithContext(ctx context.Context) *GetPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get patch params
func (o *GetPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get patch params
func (o *GetPatchParams) WithHTTPClient(client *http.Client) *GetPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get patch params
func (o *GetPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatchID adds the patchID to the get patch params
func (o *GetPatchParams) WithPatchID(patchID strfmt.UUID) *GetPatchParams {
	o.SetPatchID(patchID)
	return o
}

// SetPatchID adds the patchId to the get patch params
func (o *GetPatchParams) SetPatchID(patchID strfmt.UUID) {
	o.PatchID = patchID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param patch_id
	if err := r.SetPathParam("patch_id", o.PatchID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
