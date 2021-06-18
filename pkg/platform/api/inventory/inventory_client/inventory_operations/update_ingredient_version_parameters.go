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

// NewUpdateIngredientVersionParams creates a new UpdateIngredientVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIngredientVersionParams() *UpdateIngredientVersionParams {
	return &UpdateIngredientVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIngredientVersionParamsWithTimeout creates a new UpdateIngredientVersionParams object
// with the ability to set a timeout on a request.
func NewUpdateIngredientVersionParamsWithTimeout(timeout time.Duration) *UpdateIngredientVersionParams {
	return &UpdateIngredientVersionParams{
		timeout: timeout,
	}
}

// NewUpdateIngredientVersionParamsWithContext creates a new UpdateIngredientVersionParams object
// with the ability to set a context for a request.
func NewUpdateIngredientVersionParamsWithContext(ctx context.Context) *UpdateIngredientVersionParams {
	return &UpdateIngredientVersionParams{
		Context: ctx,
	}
}

// NewUpdateIngredientVersionParamsWithHTTPClient creates a new UpdateIngredientVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIngredientVersionParamsWithHTTPClient(client *http.Client) *UpdateIngredientVersionParams {
	return &UpdateIngredientVersionParams{
		HTTPClient: client,
	}
}

/* UpdateIngredientVersionParams contains all the parameters to send to the API endpoint
   for the update ingredient version operation.

   Typically these are written to a http.Request.
*/
type UpdateIngredientVersionParams struct {

	// IngredientID.
	//
	// Format: uuid
	IngredientID strfmt.UUID

	// IngredientVersion.
	IngredientVersion *inventory_models.IngredientVersionUpdate

	// IngredientVersionID.
	//
	// Format: uuid
	IngredientVersionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update ingredient version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIngredientVersionParams) WithDefaults() *UpdateIngredientVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update ingredient version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIngredientVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update ingredient version params
func (o *UpdateIngredientVersionParams) WithTimeout(timeout time.Duration) *UpdateIngredientVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ingredient version params
func (o *UpdateIngredientVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ingredient version params
func (o *UpdateIngredientVersionParams) WithContext(ctx context.Context) *UpdateIngredientVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ingredient version params
func (o *UpdateIngredientVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ingredient version params
func (o *UpdateIngredientVersionParams) WithHTTPClient(client *http.Client) *UpdateIngredientVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ingredient version params
func (o *UpdateIngredientVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredientID adds the ingredientID to the update ingredient version params
func (o *UpdateIngredientVersionParams) WithIngredientID(ingredientID strfmt.UUID) *UpdateIngredientVersionParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the update ingredient version params
func (o *UpdateIngredientVersionParams) SetIngredientID(ingredientID strfmt.UUID) {
	o.IngredientID = ingredientID
}

// WithIngredientVersion adds the ingredientVersion to the update ingredient version params
func (o *UpdateIngredientVersionParams) WithIngredientVersion(ingredientVersion *inventory_models.IngredientVersionUpdate) *UpdateIngredientVersionParams {
	o.SetIngredientVersion(ingredientVersion)
	return o
}

// SetIngredientVersion adds the ingredientVersion to the update ingredient version params
func (o *UpdateIngredientVersionParams) SetIngredientVersion(ingredientVersion *inventory_models.IngredientVersionUpdate) {
	o.IngredientVersion = ingredientVersion
}

// WithIngredientVersionID adds the ingredientVersionID to the update ingredient version params
func (o *UpdateIngredientVersionParams) WithIngredientVersionID(ingredientVersionID strfmt.UUID) *UpdateIngredientVersionParams {
	o.SetIngredientVersionID(ingredientVersionID)
	return o
}

// SetIngredientVersionID adds the ingredientVersionId to the update ingredient version params
func (o *UpdateIngredientVersionParams) SetIngredientVersionID(ingredientVersionID strfmt.UUID) {
	o.IngredientVersionID = ingredientVersionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIngredientVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ingredient_id
	if err := r.SetPathParam("ingredient_id", o.IngredientID.String()); err != nil {
		return err
	}
	if o.IngredientVersion != nil {
		if err := r.SetBodyParam(o.IngredientVersion); err != nil {
			return err
		}
	}

	// path param ingredient_version_id
	if err := r.SetPathParam("ingredient_version_id", o.IngredientVersionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
