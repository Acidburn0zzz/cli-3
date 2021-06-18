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

// NewGetCPUArchitectureParams creates a new GetCPUArchitectureParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCPUArchitectureParams() *GetCPUArchitectureParams {
	return &GetCPUArchitectureParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCPUArchitectureParamsWithTimeout creates a new GetCPUArchitectureParams object
// with the ability to set a timeout on a request.
func NewGetCPUArchitectureParamsWithTimeout(timeout time.Duration) *GetCPUArchitectureParams {
	return &GetCPUArchitectureParams{
		timeout: timeout,
	}
}

// NewGetCPUArchitectureParamsWithContext creates a new GetCPUArchitectureParams object
// with the ability to set a context for a request.
func NewGetCPUArchitectureParamsWithContext(ctx context.Context) *GetCPUArchitectureParams {
	return &GetCPUArchitectureParams{
		Context: ctx,
	}
}

// NewGetCPUArchitectureParamsWithHTTPClient creates a new GetCPUArchitectureParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCPUArchitectureParamsWithHTTPClient(client *http.Client) *GetCPUArchitectureParams {
	return &GetCPUArchitectureParams{
		HTTPClient: client,
	}
}

/* GetCPUArchitectureParams contains all the parameters to send to the API endpoint
   for the get Cpu architecture operation.

   Typically these are written to a http.Request.
*/
type GetCPUArchitectureParams struct {

	/* AllowUnstable.

	   Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version
	*/
	AllowUnstable *bool

	// CPUArchitectureID.
	//
	// Format: uuid
	CPUArchitectureID strfmt.UUID

	/* StateAt.

	   Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	   Format: date-time
	*/
	StateAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Cpu architecture params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCPUArchitectureParams) WithDefaults() *GetCPUArchitectureParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Cpu architecture params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCPUArchitectureParams) SetDefaults() {
	var (
		allowUnstableDefault = bool(false)
	)

	val := GetCPUArchitectureParams{
		AllowUnstable: &allowUnstableDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Cpu architecture params
func (o *GetCPUArchitectureParams) WithTimeout(timeout time.Duration) *GetCPUArchitectureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Cpu architecture params
func (o *GetCPUArchitectureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Cpu architecture params
func (o *GetCPUArchitectureParams) WithContext(ctx context.Context) *GetCPUArchitectureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Cpu architecture params
func (o *GetCPUArchitectureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Cpu architecture params
func (o *GetCPUArchitectureParams) WithHTTPClient(client *http.Client) *GetCPUArchitectureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Cpu architecture params
func (o *GetCPUArchitectureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get Cpu architecture params
func (o *GetCPUArchitectureParams) WithAllowUnstable(allowUnstable *bool) *GetCPUArchitectureParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get Cpu architecture params
func (o *GetCPUArchitectureParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithCPUArchitectureID adds the cPUArchitectureID to the get Cpu architecture params
func (o *GetCPUArchitectureParams) WithCPUArchitectureID(cPUArchitectureID strfmt.UUID) *GetCPUArchitectureParams {
	o.SetCPUArchitectureID(cPUArchitectureID)
	return o
}

// SetCPUArchitectureID adds the cpuArchitectureId to the get Cpu architecture params
func (o *GetCPUArchitectureParams) SetCPUArchitectureID(cPUArchitectureID strfmt.UUID) {
	o.CPUArchitectureID = cPUArchitectureID
}

// WithStateAt adds the stateAt to the get Cpu architecture params
func (o *GetCPUArchitectureParams) WithStateAt(stateAt *strfmt.DateTime) *GetCPUArchitectureParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get Cpu architecture params
func (o *GetCPUArchitectureParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetCPUArchitectureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param cpu_architecture_id
	if err := r.SetPathParam("cpu_architecture_id", o.CPUArchitectureID.String()); err != nil {
		return err
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
