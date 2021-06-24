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

// NewAddTagParams creates a new AddTagParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddTagParams() *AddTagParams {
	return &AddTagParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddTagParamsWithTimeout creates a new AddTagParams object
// with the ability to set a timeout on a request.
func NewAddTagParamsWithTimeout(timeout time.Duration) *AddTagParams {
	return &AddTagParams{
		timeout: timeout,
	}
}

// NewAddTagParamsWithContext creates a new AddTagParams object
// with the ability to set a context for a request.
func NewAddTagParamsWithContext(ctx context.Context) *AddTagParams {
	return &AddTagParams{
		Context: ctx,
	}
}

// NewAddTagParamsWithHTTPClient creates a new AddTagParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddTagParamsWithHTTPClient(client *http.Client) *AddTagParams {
	return &AddTagParams{
		HTTPClient: client,
	}
}

/* AddTagParams contains all the parameters to send to the API endpoint
   for the add tag operation.

   Typically these are written to a http.Request.
*/
type AddTagParams struct {

	/* OrganizationName.

	   organization associated with tag
	*/
	OrganizationName string

	/* ProjectName.

	   project associated with tag
	*/
	ProjectName string

	// Tag.
	Tag *mono_models.TagEditable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTagParams) WithDefaults() *AddTagParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add tag params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddTagParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add tag params
func (o *AddTagParams) WithTimeout(timeout time.Duration) *AddTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add tag params
func (o *AddTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add tag params
func (o *AddTagParams) WithContext(ctx context.Context) *AddTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add tag params
func (o *AddTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add tag params
func (o *AddTagParams) WithHTTPClient(client *http.Client) *AddTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add tag params
func (o *AddTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationName adds the organizationName to the add tag params
func (o *AddTagParams) WithOrganizationName(organizationName string) *AddTagParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the add tag params
func (o *AddTagParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithProjectName adds the projectName to the add tag params
func (o *AddTagParams) WithProjectName(projectName string) *AddTagParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the add tag params
func (o *AddTagParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithTag adds the tag to the add tag params
func (o *AddTagParams) WithTag(tag *mono_models.TagEditable) *AddTagParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the add tag params
func (o *AddTagParams) SetTag(tag *mono_models.TagEditable) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *AddTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}
	if o.Tag != nil {
		if err := r.SetBodyParam(o.Tag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
