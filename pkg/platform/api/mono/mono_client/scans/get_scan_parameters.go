// Code generated by go-swagger; DO NOT EDIT.

package scans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetScanParams creates a new GetScanParams object
// with the default values initialized.
func NewGetScanParams() *GetScanParams {
	var ()
	return &GetScanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScanParamsWithTimeout creates a new GetScanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScanParamsWithTimeout(timeout time.Duration) *GetScanParams {
	var ()
	return &GetScanParams{

		timeout: timeout,
	}
}

// NewGetScanParamsWithContext creates a new GetScanParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScanParamsWithContext(ctx context.Context) *GetScanParams {
	var ()
	return &GetScanParams{

		Context: ctx,
	}
}

// NewGetScanParamsWithHTTPClient creates a new GetScanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScanParamsWithHTTPClient(client *http.Client) *GetScanParams {
	var ()
	return &GetScanParams{
		HTTPClient: client,
	}
}

/*GetScanParams contains all the parameters to send to the API endpoint
for the get scan operation typically these are written to a http.Request
*/
type GetScanParams struct {

	/*ScanID
	  Unique ID of scan

	*/
	ScanID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scan params
func (o *GetScanParams) WithTimeout(timeout time.Duration) *GetScanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scan params
func (o *GetScanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scan params
func (o *GetScanParams) WithContext(ctx context.Context) *GetScanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scan params
func (o *GetScanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scan params
func (o *GetScanParams) WithHTTPClient(client *http.Client) *GetScanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scan params
func (o *GetScanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScanID adds the scanID to the get scan params
func (o *GetScanParams) WithScanID(scanID strfmt.UUID) *GetScanParams {
	o.SetScanID(scanID)
	return o
}

// SetScanID adds the scanId to the get scan params
func (o *GetScanParams) SetScanID(scanID strfmt.UUID) {
	o.ScanID = scanID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scanID
	if err := r.SetPathParam("scanID", o.ScanID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
