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
)

// NewGetHARepoParams creates a new GetHARepoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHARepoParams() *GetHARepoParams {
	return &GetHARepoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHARepoParamsWithTimeout creates a new GetHARepoParams object
// with the ability to set a timeout on a request.
func NewGetHARepoParamsWithTimeout(timeout time.Duration) *GetHARepoParams {
	return &GetHARepoParams{
		timeout: timeout,
	}
}

// NewGetHARepoParamsWithContext creates a new GetHARepoParams object
// with the ability to set a context for a request.
func NewGetHARepoParamsWithContext(ctx context.Context) *GetHARepoParams {
	return &GetHARepoParams{
		Context: ctx,
	}
}

// NewGetHARepoParamsWithHTTPClient creates a new GetHARepoParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHARepoParamsWithHTTPClient(client *http.Client) *GetHARepoParams {
	return &GetHARepoParams{
		HTTPClient: client,
	}
}

/* GetHARepoParams contains all the parameters to send to the API endpoint
   for the get h a repo operation.

   Typically these are written to a http.Request.
*/
type GetHARepoParams struct {

	/* Label.

	   HARepo Label
	*/
	Label string

	/* OrganizationName.

	   organization associated with repo
	*/
	OrganizationName string

	/* ProjectName.

	   project associated with repo
	*/
	ProjectName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get h a repo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHARepoParams) WithDefaults() *GetHARepoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get h a repo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHARepoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get h a repo params
func (o *GetHARepoParams) WithTimeout(timeout time.Duration) *GetHARepoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get h a repo params
func (o *GetHARepoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get h a repo params
func (o *GetHARepoParams) WithContext(ctx context.Context) *GetHARepoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get h a repo params
func (o *GetHARepoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get h a repo params
func (o *GetHARepoParams) WithHTTPClient(client *http.Client) *GetHARepoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get h a repo params
func (o *GetHARepoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLabel adds the label to the get h a repo params
func (o *GetHARepoParams) WithLabel(label string) *GetHARepoParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the get h a repo params
func (o *GetHARepoParams) SetLabel(label string) {
	o.Label = label
}

// WithOrganizationName adds the organizationName to the get h a repo params
func (o *GetHARepoParams) WithOrganizationName(organizationName string) *GetHARepoParams {
	o.SetOrganizationName(organizationName)
	return o
}

// SetOrganizationName adds the organizationName to the get h a repo params
func (o *GetHARepoParams) SetOrganizationName(organizationName string) {
	o.OrganizationName = organizationName
}

// WithProjectName adds the projectName to the get h a repo params
func (o *GetHARepoParams) WithProjectName(projectName string) *GetHARepoParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the get h a repo params
func (o *GetHARepoParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WriteToRequest writes these params to a swagger request
func (o *GetHARepoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param label
	if err := r.SetPathParam("label", o.Label); err != nil {
		return err
	}

	// path param organizationName
	if err := r.SetPathParam("organizationName", o.OrganizationName); err != nil {
		return err
	}

	// path param projectName
	if err := r.SetPathParam("projectName", o.ProjectName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}