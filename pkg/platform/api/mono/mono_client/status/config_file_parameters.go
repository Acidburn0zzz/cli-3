// Code generated by go-swagger; DO NOT EDIT.

package status

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

// NewConfigFileParams creates a new ConfigFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfigFileParams() *ConfigFileParams {
	return &ConfigFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfigFileParamsWithTimeout creates a new ConfigFileParams object
// with the ability to set a timeout on a request.
func NewConfigFileParamsWithTimeout(timeout time.Duration) *ConfigFileParams {
	return &ConfigFileParams{
		timeout: timeout,
	}
}

// NewConfigFileParamsWithContext creates a new ConfigFileParams object
// with the ability to set a context for a request.
func NewConfigFileParamsWithContext(ctx context.Context) *ConfigFileParams {
	return &ConfigFileParams{
		Context: ctx,
	}
}

// NewConfigFileParamsWithHTTPClient creates a new ConfigFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfigFileParamsWithHTTPClient(client *http.Client) *ConfigFileParams {
	return &ConfigFileParams{
		HTTPClient: client,
	}
}

/* ConfigFileParams contains all the parameters to send to the API endpoint
   for the config file operation.

   Typically these are written to a http.Request.
*/
type ConfigFileParams struct {

	/* IdentityID.

	   Identity UUID (optional)

	   Format: uuid
	*/
	IdentityID *strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the config file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigFileParams) WithDefaults() *ConfigFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the config file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the config file params
func (o *ConfigFileParams) WithTimeout(timeout time.Duration) *ConfigFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the config file params
func (o *ConfigFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the config file params
func (o *ConfigFileParams) WithContext(ctx context.Context) *ConfigFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the config file params
func (o *ConfigFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the config file params
func (o *ConfigFileParams) WithHTTPClient(client *http.Client) *ConfigFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the config file params
func (o *ConfigFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentityID adds the identityID to the config file params
func (o *ConfigFileParams) WithIdentityID(identityID *strfmt.UUID) *ConfigFileParams {
	o.SetIdentityID(identityID)
	return o
}

// SetIdentityID adds the identityId to the config file params
func (o *ConfigFileParams) SetIdentityID(identityID *strfmt.UUID) {
	o.IdentityID = identityID
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IdentityID != nil {

		// query param identityID
		var qrIdentityID strfmt.UUID

		if o.IdentityID != nil {
			qrIdentityID = *o.IdentityID
		}
		qIdentityID := qrIdentityID.String()
		if qIdentityID != "" {

			if err := r.SetQueryParam("identityID", qIdentityID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
