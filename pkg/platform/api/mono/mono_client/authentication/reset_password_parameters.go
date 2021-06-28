// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// NewResetPasswordParams creates a new ResetPasswordParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetPasswordParams() *ResetPasswordParams {
	return &ResetPasswordParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetPasswordParamsWithTimeout creates a new ResetPasswordParams object
// with the ability to set a timeout on a request.
func NewResetPasswordParamsWithTimeout(timeout time.Duration) *ResetPasswordParams {
	return &ResetPasswordParams{
		timeout: timeout,
	}
}

// NewResetPasswordParamsWithContext creates a new ResetPasswordParams object
// with the ability to set a context for a request.
func NewResetPasswordParamsWithContext(ctx context.Context) *ResetPasswordParams {
	return &ResetPasswordParams{
		Context: ctx,
	}
}

// NewResetPasswordParamsWithHTTPClient creates a new ResetPasswordParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetPasswordParamsWithHTTPClient(client *http.Client) *ResetPasswordParams {
	return &ResetPasswordParams{
		HTTPClient: client,
	}
}

/* ResetPasswordParams contains all the parameters to send to the API endpoint
   for the reset password operation.

   Typically these are written to a http.Request.
*/
type ResetPasswordParams struct {

	/* ResetRequest.

	   Reset Request
	*/
	ResetRequest *mono_models.PasswordReset

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetPasswordParams) WithDefaults() *ResetPasswordParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset password params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetPasswordParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset password params
func (o *ResetPasswordParams) WithTimeout(timeout time.Duration) *ResetPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset password params
func (o *ResetPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset password params
func (o *ResetPasswordParams) WithContext(ctx context.Context) *ResetPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset password params
func (o *ResetPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset password params
func (o *ResetPasswordParams) WithHTTPClient(client *http.Client) *ResetPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset password params
func (o *ResetPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResetRequest adds the resetRequest to the reset password params
func (o *ResetPasswordParams) WithResetRequest(resetRequest *mono_models.PasswordReset) *ResetPasswordParams {
	o.SetResetRequest(resetRequest)
	return o
}

// SetResetRequest adds the resetRequest to the reset password params
func (o *ResetPasswordParams) SetResetRequest(resetRequest *mono_models.PasswordReset) {
	o.ResetRequest = resetRequest
}

// WriteToRequest writes these params to a swagger request
func (o *ResetPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ResetRequest != nil {
		if err := r.SetBodyParam(o.ResetRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
