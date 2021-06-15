// Code generated by go-swagger; DO NOT EDIT.

package ingredients

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

// NewGetIngredientVersionParams creates a new GetIngredientVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIngredientVersionParams() *GetIngredientVersionParams {
	return &GetIngredientVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIngredientVersionParamsWithTimeout creates a new GetIngredientVersionParams object
// with the ability to set a timeout on a request.
func NewGetIngredientVersionParamsWithTimeout(timeout time.Duration) *GetIngredientVersionParams {
	return &GetIngredientVersionParams{
		timeout: timeout,
	}
}

// NewGetIngredientVersionParamsWithContext creates a new GetIngredientVersionParams object
// with the ability to set a context for a request.
func NewGetIngredientVersionParamsWithContext(ctx context.Context) *GetIngredientVersionParams {
	return &GetIngredientVersionParams{
		Context: ctx,
	}
}

// NewGetIngredientVersionParamsWithHTTPClient creates a new GetIngredientVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIngredientVersionParamsWithHTTPClient(client *http.Client) *GetIngredientVersionParams {
	return &GetIngredientVersionParams{
		HTTPClient: client,
	}
}

/* GetIngredientVersionParams contains all the parameters to send to the API endpoint
   for the get ingredient version operation.

   Typically these are written to a http.Request.
*/
type GetIngredientVersionParams struct {

	// IngredientID.
	//
	// Format: uuid
	IngredientID strfmt.UUID

	// Version.
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ingredient version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIngredientVersionParams) WithDefaults() *GetIngredientVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ingredient version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIngredientVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ingredient version params
func (o *GetIngredientVersionParams) WithTimeout(timeout time.Duration) *GetIngredientVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ingredient version params
func (o *GetIngredientVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ingredient version params
func (o *GetIngredientVersionParams) WithContext(ctx context.Context) *GetIngredientVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ingredient version params
func (o *GetIngredientVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ingredient version params
func (o *GetIngredientVersionParams) WithHTTPClient(client *http.Client) *GetIngredientVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ingredient version params
func (o *GetIngredientVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredientID adds the ingredientID to the get ingredient version params
func (o *GetIngredientVersionParams) WithIngredientID(ingredientID strfmt.UUID) *GetIngredientVersionParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the get ingredient version params
func (o *GetIngredientVersionParams) SetIngredientID(ingredientID strfmt.UUID) {
	o.IngredientID = ingredientID
}

// WithVersion adds the version to the get ingredient version params
func (o *GetIngredientVersionParams) WithVersion(version string) *GetIngredientVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get ingredient version params
func (o *GetIngredientVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetIngredientVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ingredientID
	if err := r.SetPathParam("ingredientID", o.IngredientID.String()); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
