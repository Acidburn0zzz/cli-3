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

// NewAddImageRevisionParams creates a new AddImageRevisionParams object
// with the default values initialized.
func NewAddImageRevisionParams() *AddImageRevisionParams {
	var ()
	return &AddImageRevisionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddImageRevisionParamsWithTimeout creates a new AddImageRevisionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddImageRevisionParamsWithTimeout(timeout time.Duration) *AddImageRevisionParams {
	var ()
	return &AddImageRevisionParams{

		timeout: timeout,
	}
}

// NewAddImageRevisionParamsWithContext creates a new AddImageRevisionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddImageRevisionParamsWithContext(ctx context.Context) *AddImageRevisionParams {
	var ()
	return &AddImageRevisionParams{

		Context: ctx,
	}
}

// NewAddImageRevisionParamsWithHTTPClient creates a new AddImageRevisionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddImageRevisionParamsWithHTTPClient(client *http.Client) *AddImageRevisionParams {
	var ()
	return &AddImageRevisionParams{
		HTTPClient: client,
	}
}

/*AddImageRevisionParams contains all the parameters to send to the API endpoint
for the add image revision operation typically these are written to a http.Request
*/
type AddImageRevisionParams struct {

	/*ImageID*/
	ImageID strfmt.UUID
	/*ImageRevision*/
	ImageRevision *inventory_models.ImageRevisionCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add image revision params
func (o *AddImageRevisionParams) WithTimeout(timeout time.Duration) *AddImageRevisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add image revision params
func (o *AddImageRevisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add image revision params
func (o *AddImageRevisionParams) WithContext(ctx context.Context) *AddImageRevisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add image revision params
func (o *AddImageRevisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add image revision params
func (o *AddImageRevisionParams) WithHTTPClient(client *http.Client) *AddImageRevisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add image revision params
func (o *AddImageRevisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageID adds the imageID to the add image revision params
func (o *AddImageRevisionParams) WithImageID(imageID strfmt.UUID) *AddImageRevisionParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the add image revision params
func (o *AddImageRevisionParams) SetImageID(imageID strfmt.UUID) {
	o.ImageID = imageID
}

// WithImageRevision adds the imageRevision to the add image revision params
func (o *AddImageRevisionParams) WithImageRevision(imageRevision *inventory_models.ImageRevisionCore) *AddImageRevisionParams {
	o.SetImageRevision(imageRevision)
	return o
}

// SetImageRevision adds the imageRevision to the add image revision params
func (o *AddImageRevisionParams) SetImageRevision(imageRevision *inventory_models.ImageRevisionCore) {
	o.ImageRevision = imageRevision
}

// WriteToRequest writes these params to a swagger request
func (o *AddImageRevisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param image_id
	if err := r.SetPathParam("image_id", o.ImageID.String()); err != nil {
		return err
	}

	if o.ImageRevision != nil {
		if err := r.SetBodyParam(o.ImageRevision); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
