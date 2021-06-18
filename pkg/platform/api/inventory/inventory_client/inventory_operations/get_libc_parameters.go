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

// NewGetLibcParams creates a new GetLibcParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLibcParams() *GetLibcParams {
	return &GetLibcParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLibcParamsWithTimeout creates a new GetLibcParams object
// with the ability to set a timeout on a request.
func NewGetLibcParamsWithTimeout(timeout time.Duration) *GetLibcParams {
	return &GetLibcParams{
		timeout: timeout,
	}
}

// NewGetLibcParamsWithContext creates a new GetLibcParams object
// with the ability to set a context for a request.
func NewGetLibcParamsWithContext(ctx context.Context) *GetLibcParams {
	return &GetLibcParams{
		Context: ctx,
	}
}

// NewGetLibcParamsWithHTTPClient creates a new GetLibcParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLibcParamsWithHTTPClient(client *http.Client) *GetLibcParams {
	return &GetLibcParams{
		HTTPClient: client,
	}
}

/* GetLibcParams contains all the parameters to send to the API endpoint
   for the get libc operation.

   Typically these are written to a http.Request.
*/
type GetLibcParams struct {

	// LibcID.
	//
	// Format: uuid
	LibcID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get libc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLibcParams) WithDefaults() *GetLibcParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get libc params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLibcParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get libc params
func (o *GetLibcParams) WithTimeout(timeout time.Duration) *GetLibcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get libc params
func (o *GetLibcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get libc params
func (o *GetLibcParams) WithContext(ctx context.Context) *GetLibcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get libc params
func (o *GetLibcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get libc params
func (o *GetLibcParams) WithHTTPClient(client *http.Client) *GetLibcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get libc params
func (o *GetLibcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLibcID adds the libcID to the get libc params
func (o *GetLibcParams) WithLibcID(libcID strfmt.UUID) *GetLibcParams {
	o.SetLibcID(libcID)
	return o
}

// SetLibcID adds the libcId to the get libc params
func (o *GetLibcParams) SetLibcID(libcID strfmt.UUID) {
	o.LibcID = libcID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLibcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param libc_id
	if err := r.SetPathParam("libc_id", o.LibcID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
