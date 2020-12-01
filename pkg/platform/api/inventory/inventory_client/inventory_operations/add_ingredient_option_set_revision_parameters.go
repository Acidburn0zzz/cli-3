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

// NewAddIngredientOptionSetRevisionParams creates a new AddIngredientOptionSetRevisionParams object
// with the default values initialized.
func NewAddIngredientOptionSetRevisionParams() *AddIngredientOptionSetRevisionParams {
	var ()
	return &AddIngredientOptionSetRevisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddIngredientOptionSetRevisionParamsWithTimeout creates a new AddIngredientOptionSetRevisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddIngredientOptionSetRevisionParamsWithTimeout(timeout time.Duration) *AddIngredientOptionSetRevisionParams {
	var ()
	return &AddIngredientOptionSetRevisionParams{

		timeout: timeout,
	}
}

// NewAddIngredientOptionSetRevisionParamsWithContext creates a new AddIngredientOptionSetRevisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddIngredientOptionSetRevisionParamsWithContext(ctx context.Context) *AddIngredientOptionSetRevisionParams {
	var ()
	return &AddIngredientOptionSetRevisionParams{

		Context: ctx,
	}
}

// NewAddIngredientOptionSetRevisionParamsWithHTTPClient creates a new AddIngredientOptionSetRevisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddIngredientOptionSetRevisionParamsWithHTTPClient(client *http.Client) *AddIngredientOptionSetRevisionParams {
	var ()
	return &AddIngredientOptionSetRevisionParams{
		HTTPClient: client,
	}
}

/*AddIngredientOptionSetRevisionParams contains all the parameters to send to the API endpoint
for the add ingredient option set revision operation typically these are written to a http.Request
*/
type AddIngredientOptionSetRevisionParams struct {

	/*IngredientOptionSetID*/
	IngredientOptionSetID strfmt.UUID
	/*IngredientOptionSetRevision*/
	IngredientOptionSetRevision *inventory_models.IngredientOptionSetRevisionCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) WithTimeout(timeout time.Duration) *AddIngredientOptionSetRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) WithContext(ctx context.Context) *AddIngredientOptionSetRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) WithHTTPClient(client *http.Client) *AddIngredientOptionSetRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredientOptionSetID adds the ingredientOptionSetID to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) WithIngredientOptionSetID(ingredientOptionSetID strfmt.UUID) *AddIngredientOptionSetRevisionParams {
	o.SetIngredientOptionSetID(ingredientOptionSetID)
	return o
}

// SetIngredientOptionSetID adds the ingredientOptionSetId to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) SetIngredientOptionSetID(ingredientOptionSetID strfmt.UUID) {
	o.IngredientOptionSetID = ingredientOptionSetID
}

// WithIngredientOptionSetRevision adds the ingredientOptionSetRevision to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) WithIngredientOptionSetRevision(ingredientOptionSetRevision *inventory_models.IngredientOptionSetRevisionCore) *AddIngredientOptionSetRevisionParams {
	o.SetIngredientOptionSetRevision(ingredientOptionSetRevision)
	return o
}

// SetIngredientOptionSetRevision adds the ingredientOptionSetRevision to the add ingredient option set revision params
func (o *AddIngredientOptionSetRevisionParams) SetIngredientOptionSetRevision(ingredientOptionSetRevision *inventory_models.IngredientOptionSetRevisionCore) {
	o.IngredientOptionSetRevision = ingredientOptionSetRevision
}

// WriteToRequest writes these params to a swagger request
func (o *AddIngredientOptionSetRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ingredient_option_set_id
	if err := r.SetPathParam("ingredient_option_set_id", o.IngredientOptionSetID.String()); err != nil {
		return err
	}

	if o.IngredientOptionSetRevision != nil {
		if err := r.SetBodyParam(o.IngredientOptionSetRevision); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
