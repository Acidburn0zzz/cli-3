// Code generated by go-swagger; DO NOT EDIT.

package tiers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTiersParams creates a new GetTiersParams object
// with the default values initialized.
func NewGetTiersParams() *GetTiersParams {

	return &GetTiersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTiersParamsWithTimeout creates a new GetTiersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTiersParamsWithTimeout(timeout time.Duration) *GetTiersParams {

	return &GetTiersParams{

		timeout: timeout,
	}
}

// NewGetTiersParamsWithContext creates a new GetTiersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTiersParamsWithContext(ctx context.Context) *GetTiersParams {

	return &GetTiersParams{

		Context: ctx,
	}
}

// NewGetTiersParamsWithHTTPClient creates a new GetTiersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTiersParamsWithHTTPClient(client *http.Client) *GetTiersParams {

	return &GetTiersParams{
		HTTPClient: client,
	}
}

/*GetTiersParams contains all the parameters to send to the API endpoint
for the get tiers operation typically these are written to a http.Request
*/
type GetTiersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tiers params
func (o *GetTiersParams) WithTimeout(timeout time.Duration) *GetTiersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tiers params
func (o *GetTiersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tiers params
func (o *GetTiersParams) WithContext(ctx context.Context) *GetTiersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tiers params
func (o *GetTiersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tiers params
func (o *GetTiersParams) WithHTTPClient(client *http.Client) *GetTiersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tiers params
func (o *GetTiersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetTiersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
