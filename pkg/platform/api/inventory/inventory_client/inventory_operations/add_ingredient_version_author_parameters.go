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

// NewAddIngredientVersionAuthorParams creates a new AddIngredientVersionAuthorParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddIngredientVersionAuthorParams() *AddIngredientVersionAuthorParams {
	return &AddIngredientVersionAuthorParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddIngredientVersionAuthorParamsWithTimeout creates a new AddIngredientVersionAuthorParams object
// with the ability to set a timeout on a request.
func NewAddIngredientVersionAuthorParamsWithTimeout(timeout time.Duration) *AddIngredientVersionAuthorParams {
	return &AddIngredientVersionAuthorParams{
		timeout: timeout,
	}
}

// NewAddIngredientVersionAuthorParamsWithContext creates a new AddIngredientVersionAuthorParams object
// with the ability to set a context for a request.
func NewAddIngredientVersionAuthorParamsWithContext(ctx context.Context) *AddIngredientVersionAuthorParams {
	return &AddIngredientVersionAuthorParams{
		Context: ctx,
	}
}

// NewAddIngredientVersionAuthorParamsWithHTTPClient creates a new AddIngredientVersionAuthorParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddIngredientVersionAuthorParamsWithHTTPClient(client *http.Client) *AddIngredientVersionAuthorParams {
	return &AddIngredientVersionAuthorParams{
		HTTPClient: client,
	}
}

/* AddIngredientVersionAuthorParams contains all the parameters to send to the API endpoint
   for the add ingredient version author operation.

   Typically these are written to a http.Request.
*/
type AddIngredientVersionAuthorParams struct {

	// AuthorID.
	AuthorID *inventory_models.AddIngredientVersionAuthorParamsBody

	// IngredientID.
	//
	// Format: uuid
	IngredientID strfmt.UUID

	// IngredientVersionID.
	//
	// Format: uuid
	IngredientVersionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add ingredient version author params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddIngredientVersionAuthorParams) WithDefaults() *AddIngredientVersionAuthorParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add ingredient version author params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddIngredientVersionAuthorParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) WithTimeout(timeout time.Duration) *AddIngredientVersionAuthorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) WithContext(ctx context.Context) *AddIngredientVersionAuthorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) WithHTTPClient(client *http.Client) *AddIngredientVersionAuthorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorID adds the authorID to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) WithAuthorID(authorID *inventory_models.AddIngredientVersionAuthorParamsBody) *AddIngredientVersionAuthorParams {
	o.SetAuthorID(authorID)
	return o
}

// SetAuthorID adds the authorId to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) SetAuthorID(authorID *inventory_models.AddIngredientVersionAuthorParamsBody) {
	o.AuthorID = authorID
}

// WithIngredientID adds the ingredientID to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) WithIngredientID(ingredientID strfmt.UUID) *AddIngredientVersionAuthorParams {
	o.SetIngredientID(ingredientID)
	return o
}

// SetIngredientID adds the ingredientId to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) SetIngredientID(ingredientID strfmt.UUID) {
	o.IngredientID = ingredientID
}

// WithIngredientVersionID adds the ingredientVersionID to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) WithIngredientVersionID(ingredientVersionID strfmt.UUID) *AddIngredientVersionAuthorParams {
	o.SetIngredientVersionID(ingredientVersionID)
	return o
}

// SetIngredientVersionID adds the ingredientVersionId to the add ingredient version author params
func (o *AddIngredientVersionAuthorParams) SetIngredientVersionID(ingredientVersionID strfmt.UUID) {
	o.IngredientVersionID = ingredientVersionID
}

// WriteToRequest writes these params to a swagger request
func (o *AddIngredientVersionAuthorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.AuthorID != nil {
		if err := r.SetBodyParam(o.AuthorID); err != nil {
			return err
		}
	}

	// path param ingredient_id
	if err := r.SetPathParam("ingredient_id", o.IngredientID.String()); err != nil {
		return err
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
