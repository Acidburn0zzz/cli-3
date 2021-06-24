// Code generated by go-swagger; DO NOT EDIT.

package github

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

// NewGetLoginsAndReposParams creates a new GetLoginsAndReposParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoginsAndReposParams() *GetLoginsAndReposParams {
	return &GetLoginsAndReposParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginsAndReposParamsWithTimeout creates a new GetLoginsAndReposParams object
// with the ability to set a timeout on a request.
func NewGetLoginsAndReposParamsWithTimeout(timeout time.Duration) *GetLoginsAndReposParams {
	return &GetLoginsAndReposParams{
		timeout: timeout,
	}
}

// NewGetLoginsAndReposParamsWithContext creates a new GetLoginsAndReposParams object
// with the ability to set a context for a request.
func NewGetLoginsAndReposParamsWithContext(ctx context.Context) *GetLoginsAndReposParams {
	return &GetLoginsAndReposParams{
		Context: ctx,
	}
}

// NewGetLoginsAndReposParamsWithHTTPClient creates a new GetLoginsAndReposParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoginsAndReposParamsWithHTTPClient(client *http.Client) *GetLoginsAndReposParams {
	return &GetLoginsAndReposParams{
		HTTPClient: client,
	}
}

/* GetLoginsAndReposParams contains all the parameters to send to the API endpoint
   for the get logins and repos operation.

   Typically these are written to a http.Request.
*/
type GetLoginsAndReposParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get logins and repos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoginsAndReposParams) WithDefaults() *GetLoginsAndReposParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get logins and repos params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoginsAndReposParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get logins and repos params
func (o *GetLoginsAndReposParams) WithTimeout(timeout time.Duration) *GetLoginsAndReposParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get logins and repos params
func (o *GetLoginsAndReposParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get logins and repos params
func (o *GetLoginsAndReposParams) WithContext(ctx context.Context) *GetLoginsAndReposParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get logins and repos params
func (o *GetLoginsAndReposParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get logins and repos params
func (o *GetLoginsAndReposParams) WithHTTPClient(client *http.Client) *GetLoginsAndReposParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get logins and repos params
func (o *GetLoginsAndReposParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginsAndReposParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
