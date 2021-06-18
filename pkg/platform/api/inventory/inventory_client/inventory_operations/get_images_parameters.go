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

// NewGetImagesParams creates a new GetImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetImagesParams() *GetImagesParams {
	return &GetImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetImagesParamsWithTimeout creates a new GetImagesParams object
// with the ability to set a timeout on a request.
func NewGetImagesParamsWithTimeout(timeout time.Duration) *GetImagesParams {
	return &GetImagesParams{
		timeout: timeout,
	}
}

// NewGetImagesParamsWithContext creates a new GetImagesParams object
// with the ability to set a context for a request.
func NewGetImagesParamsWithContext(ctx context.Context) *GetImagesParams {
	return &GetImagesParams{
		Context: ctx,
	}
}

// NewGetImagesParamsWithHTTPClient creates a new GetImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetImagesParamsWithHTTPClient(client *http.Client) *GetImagesParams {
	return &GetImagesParams{
		HTTPClient: client,
	}
}

/* GetImagesParams contains all the parameters to send to the API endpoint
   for the get images operation.

   Typically these are written to a http.Request.
*/
type GetImagesParams struct {

	/* AllowUnstable.

	   Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version
	*/
	AllowUnstable *bool

	/* Limit.

	   The maximum number of items returned per page

	   Default: 50
	*/
	Limit *int64

	/* Page.

	   The page number returned

	   Default: 1
	*/
	Page *int64

	/* StateAt.

	   Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	   Format: date-time
	*/
	StateAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImagesParams) WithDefaults() *GetImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetImagesParams) SetDefaults() {
	var (
		allowUnstableDefault = bool(false)

		limitDefault = int64(50)

		pageDefault = int64(1)
	)

	val := GetImagesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get images params
func (o *GetImagesParams) WithTimeout(timeout time.Duration) *GetImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get images params
func (o *GetImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get images params
func (o *GetImagesParams) WithContext(ctx context.Context) *GetImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get images params
func (o *GetImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get images params
func (o *GetImagesParams) WithHTTPClient(client *http.Client) *GetImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get images params
func (o *GetImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get images params
func (o *GetImagesParams) WithAllowUnstable(allowUnstable *bool) *GetImagesParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get images params
func (o *GetImagesParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithLimit adds the limit to the get images params
func (o *GetImagesParams) WithLimit(limit *int64) *GetImagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get images params
func (o *GetImagesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get images params
func (o *GetImagesParams) WithPage(page *int64) *GetImagesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get images params
func (o *GetImagesParams) SetPage(page *int64) {
	o.Page = page
}

// WithStateAt adds the stateAt to the get images params
func (o *GetImagesParams) WithStateAt(stateAt *strfmt.DateTime) *GetImagesParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get images params
func (o *GetImagesParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowUnstable != nil {

		// query param allow_unstable
		var qrAllowUnstable bool

		if o.AllowUnstable != nil {
			qrAllowUnstable = *o.AllowUnstable
		}
		qAllowUnstable := swag.FormatBool(qrAllowUnstable)
		if qAllowUnstable != "" {

			if err := r.SetQueryParam("allow_unstable", qAllowUnstable); err != nil {
				return err
			}
		}
	}

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

	if o.StateAt != nil {

		// query param state_at
		var qrStateAt strfmt.DateTime

		if o.StateAt != nil {
			qrStateAt = *o.StateAt
		}
		qStateAt := qrStateAt.String()
		if qStateAt != "" {

			if err := r.SetQueryParam("state_at", qStateAt); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
