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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPlatformsParams creates a new PlatformsParams object
// with the default values initialized.
func NewPlatformsParams() *PlatformsParams {

	return &PlatformsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPlatformsParamsWithTimeout creates a new PlatformsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPlatformsParamsWithTimeout(timeout time.Duration) *PlatformsParams {

	return &PlatformsParams{

		timeout: timeout,
	}
}

// NewPlatformsParamsWithContext creates a new PlatformsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPlatformsParamsWithContext(ctx context.Context) *PlatformsParams {

	return &PlatformsParams{

		Context: ctx,
	}
}

// NewPlatformsParamsWithHTTPClient creates a new PlatformsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPlatformsParamsWithHTTPClient(client *http.Client) *PlatformsParams {

	return &PlatformsParams{
		HTTPClient: client,
	}
}

/*PlatformsParams contains all the parameters to send to the API endpoint
for the platforms operation typically these are written to a http.Request
*/
type PlatformsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the platforms params
func (o *PlatformsParams) WithTimeout(timeout time.Duration) *PlatformsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the platforms params
func (o *PlatformsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the platforms params
func (o *PlatformsParams) WithContext(ctx context.Context) *PlatformsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the platforms params
func (o *PlatformsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the platforms params
func (o *PlatformsParams) WithHTTPClient(client *http.Client) *PlatformsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the platforms params
func (o *PlatformsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PlatformsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}