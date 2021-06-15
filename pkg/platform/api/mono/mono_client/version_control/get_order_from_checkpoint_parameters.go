// Code generated by go-swagger; DO NOT EDIT.

package version_control

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

// NewGetOrderFromCheckpointParams creates a new GetOrderFromCheckpointParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrderFromCheckpointParams() *GetOrderFromCheckpointParams {
	return &GetOrderFromCheckpointParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrderFromCheckpointParamsWithTimeout creates a new GetOrderFromCheckpointParams object
// with the ability to set a timeout on a request.
func NewGetOrderFromCheckpointParamsWithTimeout(timeout time.Duration) *GetOrderFromCheckpointParams {
	return &GetOrderFromCheckpointParams{
		timeout: timeout,
	}
}

// NewGetOrderFromCheckpointParamsWithContext creates a new GetOrderFromCheckpointParams object
// with the ability to set a context for a request.
func NewGetOrderFromCheckpointParamsWithContext(ctx context.Context) *GetOrderFromCheckpointParams {
	return &GetOrderFromCheckpointParams{
		Context: ctx,
	}
}

// NewGetOrderFromCheckpointParamsWithHTTPClient creates a new GetOrderFromCheckpointParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrderFromCheckpointParamsWithHTTPClient(client *http.Client) *GetOrderFromCheckpointParams {
	return &GetOrderFromCheckpointParams{
		HTTPClient: client,
	}
}

/* GetOrderFromCheckpointParams contains all the parameters to send to the API endpoint
   for the get order from checkpoint operation.

   Typically these are written to a http.Request.
*/
type GetOrderFromCheckpointParams struct {

	// OrderInfo.
	OrderInfo *mono_models.OrderInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get order from checkpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrderFromCheckpointParams) WithDefaults() *GetOrderFromCheckpointParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get order from checkpoint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrderFromCheckpointParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get order from checkpoint params
func (o *GetOrderFromCheckpointParams) WithTimeout(timeout time.Duration) *GetOrderFromCheckpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get order from checkpoint params
func (o *GetOrderFromCheckpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get order from checkpoint params
func (o *GetOrderFromCheckpointParams) WithContext(ctx context.Context) *GetOrderFromCheckpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get order from checkpoint params
func (o *GetOrderFromCheckpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get order from checkpoint params
func (o *GetOrderFromCheckpointParams) WithHTTPClient(client *http.Client) *GetOrderFromCheckpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get order from checkpoint params
func (o *GetOrderFromCheckpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderInfo adds the orderInfo to the get order from checkpoint params
func (o *GetOrderFromCheckpointParams) WithOrderInfo(orderInfo *mono_models.OrderInfo) *GetOrderFromCheckpointParams {
	o.SetOrderInfo(orderInfo)
	return o
}

// SetOrderInfo adds the orderInfo to the get order from checkpoint params
func (o *GetOrderFromCheckpointParams) SetOrderInfo(orderInfo *mono_models.OrderInfo) {
	o.OrderInfo = orderInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrderFromCheckpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.OrderInfo != nil {
		if err := r.SetBodyParam(o.OrderInfo); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
