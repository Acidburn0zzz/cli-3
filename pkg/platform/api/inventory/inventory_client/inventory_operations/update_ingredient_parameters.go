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

// NewUpdateIngredientParams creates a new UpdateIngredientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateIngredientParams() *UpdateIngredientParams {
	return &UpdateIngredientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateIngredientParamsWithTimeout creates a new UpdateIngredientParams object
// with the ability to set a timeout on a request.
func NewUpdateIngredientParamsWithTimeout(timeout time.Duration) *UpdateIngredientParams {
	return &UpdateIngredientParams{
		timeout: timeout,
	}
}

// NewUpdateIngredientParamsWithContext creates a new UpdateIngredientParams object
// with the ability to set a context for a request.
func NewUpdateIngredientParamsWithContext(ctx context.Context) *UpdateIngredientParams {
	return &UpdateIngredientParams{
		Context: ctx,
	}
}

// NewUpdateIngredientParamsWithHTTPClient creates a new UpdateIngredientParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateIngredientParamsWithHTTPClient(client *http.Client) *UpdateIngredientParams {
	return &UpdateIngredientParams{
		HTTPClient: client,
	}
}

/* UpdateIngredientParams contains all the parameters to send to the API endpoint
   for the update ingredient operation.

   Typically these are written to a http.Request.
*/
type UpdateIngredientParams struct {

	// Ingredient.
	Ingredient *inventory_models.IngredientUpdate

	// IngredientID.
	//
	// Format: uuid
	IngredientID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update ingredient params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIngredientParams) WithDefaults() *UpdateIngredientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update ingredient params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateIngredientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update ingredient params
func (o *UpdateIngredientParams) WithTimeout(timeout time.Duration) *UpdateIngredientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ingredient params
func (o *UpdateIngredientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ingredient params
func (o *UpdateIngredientParams) WithContext(ctx context.Context) *UpdateIngredientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ingredient params
func (o *UpdateIngredientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ingredient params
func (o *UpdateIngredientParams) WithHTTPClient(client *http.Client) *UpdateIngredientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ingredient params
func (o *UpdateIngredientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredient adds the ingredient to the update ingredient params
func (o *UpdateIngredientParams) WithIngredient(ingredient *inventory_models.IngredientUpdate) *UpdateIngredientParams {
	o.SetIngredient(ingredient)
	return o
}

// SetIngredient adds the ingredient to the update ingredient params
func (o *UpdateIngredientParams) SetIngredient(ingredient *inventory_models.IngredientUpdate) {
	o.Ingredient = ingredient
}

// WithIngredientID adds the ingredientID to the update ingredient params
func (o *UpdateIngredientParams) WithIngredientID(ingredientID strfmt.UUID) *UpdateIngredientParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the update ingredient params
func (o *UpdateIngredientParams) SetIngredientID(ingredientID strfmt.UUID) {
	o.IngredientID = ingredientID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateIngredientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Ingredient != nil {
		if err := r.SetBodyParam(o.Ingredient); err != nil {
			return err
		}
	}

	// path param ingredient_id
	if err := r.SetPathParam("ingredient_id", o.IngredientID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
