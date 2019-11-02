// Code generated by go-swagger; DO NOT EDIT.

package headchef_operations

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

// NewGetBuildStatusParams creates a new GetBuildStatusParams object
// with the default values initialized.
func NewGetBuildStatusParams() *GetBuildStatusParams {
	var ()
	return &GetBuildStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBuildStatusParamsWithTimeout creates a new GetBuildStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBuildStatusParamsWithTimeout(timeout time.Duration) *GetBuildStatusParams {
	var ()
	return &GetBuildStatusParams{

		timeout: timeout,
	}
}

// NewGetBuildStatusParamsWithContext creates a new GetBuildStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBuildStatusParamsWithContext(ctx context.Context) *GetBuildStatusParams {
	var ()
	return &GetBuildStatusParams{

		Context: ctx,
	}
}

// NewGetBuildStatusParamsWithHTTPClient creates a new GetBuildStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBuildStatusParamsWithHTTPClient(client *http.Client) *GetBuildStatusParams {
	var ()
	return &GetBuildStatusParams{
		HTTPClient: client,
	}
}

/*GetBuildStatusParams contains all the parameters to send to the API endpoint
for the get build status operation typically these are written to a http.Request
*/
type GetBuildStatusParams struct {

	/*BuildRequestID*/
	BuildRequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get build status params
func (o *GetBuildStatusParams) WithTimeout(timeout time.Duration) *GetBuildStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get build status params
func (o *GetBuildStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get build status params
func (o *GetBuildStatusParams) WithContext(ctx context.Context) *GetBuildStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get build status params
func (o *GetBuildStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get build status params
func (o *GetBuildStatusParams) WithHTTPClient(client *http.Client) *GetBuildStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get build status params
func (o *GetBuildStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildRequestID adds the buildRequestID to the get build status params
func (o *GetBuildStatusParams) WithBuildRequestID(buildRequestID strfmt.UUID) *GetBuildStatusParams {
	o.SetBuildRequestID(buildRequestID)
	return o
}

// SetBuildRequestID adds the buildRequestId to the get build status params
func (o *GetBuildStatusParams) SetBuildRequestID(buildRequestID strfmt.UUID) {
	o.BuildRequestID = buildRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBuildStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param build_request_id
	if err := r.SetPathParam("build_request_id", o.BuildRequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}