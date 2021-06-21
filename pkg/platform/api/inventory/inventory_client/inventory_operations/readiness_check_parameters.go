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

// NewReadinessCheckParams creates a new ReadinessCheckParams object
// with the default values initialized.
func NewReadinessCheckParams() *ReadinessCheckParams {

	return &ReadinessCheckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadinessCheckParamsWithTimeout creates a new ReadinessCheckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadinessCheckParamsWithTimeout(timeout time.Duration) *ReadinessCheckParams {

	return &ReadinessCheckParams{

		timeout: timeout,
	}
}

// NewReadinessCheckParamsWithContext creates a new ReadinessCheckParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadinessCheckParamsWithContext(ctx context.Context) *ReadinessCheckParams {

	return &ReadinessCheckParams{

		Context: ctx,
	}
}

// NewReadinessCheckParamsWithHTTPClient creates a new ReadinessCheckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadinessCheckParamsWithHTTPClient(client *http.Client) *ReadinessCheckParams {

	return &ReadinessCheckParams{
		HTTPClient: client,
	}
}

/*ReadinessCheckParams contains all the parameters to send to the API endpoint
for the readiness check operation typically these are written to a http.Request
*/
type ReadinessCheckParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the readiness check params
func (o *ReadinessCheckParams) WithTimeout(timeout time.Duration) *ReadinessCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the readiness check params
func (o *ReadinessCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the readiness check params
func (o *ReadinessCheckParams) WithContext(ctx context.Context) *ReadinessCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the readiness check params
func (o *ReadinessCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the readiness check params
func (o *ReadinessCheckParams) WithHTTPClient(client *http.Client) *ReadinessCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the readiness check params
func (o *ReadinessCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ReadinessCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
