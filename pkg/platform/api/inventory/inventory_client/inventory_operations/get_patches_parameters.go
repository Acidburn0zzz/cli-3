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
	"github.com/go-openapi/swag"
)

// NewGetPatchesParams creates a new GetPatchesParams object
// with the default values initialized.
func NewGetPatchesParams() *GetPatchesParams {
	var (
		limitDefault = int64(50)
		pageDefault  = int64(1)
	)
	return &GetPatchesParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatchesParamsWithTimeout creates a new GetPatchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatchesParamsWithTimeout(timeout time.Duration) *GetPatchesParams {
	var (
		limitDefault = int64(50)
		pageDefault  = int64(1)
	)
	return &GetPatchesParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		timeout: timeout,
	}
}

// NewGetPatchesParamsWithContext creates a new GetPatchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatchesParamsWithContext(ctx context.Context) *GetPatchesParams {
	var (
		limitDefault = int64(50)
		pageDefault  = int64(1)
	)
	return &GetPatchesParams{
		Limit: &limitDefault,
		Page:  &pageDefault,

		Context: ctx,
	}
}

// NewGetPatchesParamsWithHTTPClient creates a new GetPatchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatchesParamsWithHTTPClient(client *http.Client) *GetPatchesParams {
	var (
		limitDefault = int64(50)
		pageDefault  = int64(1)
	)
	return &GetPatchesParams{
		Limit:      &limitDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*GetPatchesParams contains all the parameters to send to the API endpoint
for the get patches operation typically these are written to a http.Request
*/
type GetPatchesParams struct {

	/*Limit
	  The maximum number of items returned per page

	*/
	Limit *int64
	/*Page
	  The page number returned

	*/
	Page *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get patches params
func (o *GetPatchesParams) WithTimeout(timeout time.Duration) *GetPatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get patches params
func (o *GetPatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get patches params
func (o *GetPatchesParams) WithContext(ctx context.Context) *GetPatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get patches params
func (o *GetPatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get patches params
func (o *GetPatchesParams) WithHTTPClient(client *http.Client) *GetPatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get patches params
func (o *GetPatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get patches params
func (o *GetPatchesParams) WithLimit(limit *int64) *GetPatchesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get patches params
func (o *GetPatchesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get patches params
func (o *GetPatchesParams) WithPage(page *int64) *GetPatchesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get patches params
func (o *GetPatchesParams) SetPage(page *int64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
