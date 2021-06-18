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

// NewAddAuthorParams creates a new AddAuthorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAuthorParams() *AddAuthorParams {
	return &AddAuthorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAuthorParamsWithTimeout creates a new AddAuthorParams object
// with the ability to set a timeout on a request.
func NewAddAuthorParamsWithTimeout(timeout time.Duration) *AddAuthorParams {
	return &AddAuthorParams{
		timeout: timeout,
	}
}

// NewAddAuthorParamsWithContext creates a new AddAuthorParams object
// with the ability to set a context for a request.
func NewAddAuthorParamsWithContext(ctx context.Context) *AddAuthorParams {
	return &AddAuthorParams{
		Context: ctx,
	}
}

// NewAddAuthorParamsWithHTTPClient creates a new AddAuthorParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAuthorParamsWithHTTPClient(client *http.Client) *AddAuthorParams {
	return &AddAuthorParams{
		HTTPClient: client,
	}
}

/* AddAuthorParams contains all the parameters to send to the API endpoint
   for the add author operation.

   Typically these are written to a http.Request.
*/
type AddAuthorParams struct {

	// Author.
	Author *inventory_models.AuthorCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add author params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAuthorParams) WithDefaults() *AddAuthorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add author params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAuthorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add author params
func (o *AddAuthorParams) WithTimeout(timeout time.Duration) *AddAuthorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add author params
func (o *AddAuthorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add author params
func (o *AddAuthorParams) WithContext(ctx context.Context) *AddAuthorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add author params
func (o *AddAuthorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add author params
func (o *AddAuthorParams) WithHTTPClient(client *http.Client) *AddAuthorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add author params
func (o *AddAuthorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthor adds the author to the add author params
func (o *AddAuthorParams) WithAuthor(author *inventory_models.AuthorCore) *AddAuthorParams {
	o.SetAuthor(author)
	return o
}

// SetAuthor adds the author to the add author params
func (o *AddAuthorParams) SetAuthor(author *inventory_models.AuthorCore) {
	o.Author = author
}

// WriteToRequest writes these params to a swagger request
func (o *AddAuthorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Author != nil {
		if err := r.SetBodyParam(o.Author); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
