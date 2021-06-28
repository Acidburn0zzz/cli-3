// Code generated by go-swagger; DO NOT EDIT.

package github

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

// NewGetRepoRequirementsParams creates a new GetRepoRequirementsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRepoRequirementsParams() *GetRepoRequirementsParams {
	return &GetRepoRequirementsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepoRequirementsParamsWithTimeout creates a new GetRepoRequirementsParams object
// with the ability to set a timeout on a request.
func NewGetRepoRequirementsParamsWithTimeout(timeout time.Duration) *GetRepoRequirementsParams {
	return &GetRepoRequirementsParams{
		timeout: timeout,
	}
}

// NewGetRepoRequirementsParamsWithContext creates a new GetRepoRequirementsParams object
// with the ability to set a context for a request.
func NewGetRepoRequirementsParamsWithContext(ctx context.Context) *GetRepoRequirementsParams {
	return &GetRepoRequirementsParams{
		Context: ctx,
	}
}

// NewGetRepoRequirementsParamsWithHTTPClient creates a new GetRepoRequirementsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRepoRequirementsParamsWithHTTPClient(client *http.Client) *GetRepoRequirementsParams {
	return &GetRepoRequirementsParams{
		HTTPClient: client,
	}
}

/* GetRepoRequirementsParams contains all the parameters to send to the API endpoint
   for the get repo requirements operation.

   Typically these are written to a http.Request.
*/
type GetRepoRequirementsParams struct {

	// RepoURL.
	RepoURL string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get repo requirements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepoRequirementsParams) WithDefaults() *GetRepoRequirementsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get repo requirements params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRepoRequirementsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get repo requirements params
func (o *GetRepoRequirementsParams) WithTimeout(timeout time.Duration) *GetRepoRequirementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repo requirements params
func (o *GetRepoRequirementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repo requirements params
func (o *GetRepoRequirementsParams) WithContext(ctx context.Context) *GetRepoRequirementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repo requirements params
func (o *GetRepoRequirementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repo requirements params
func (o *GetRepoRequirementsParams) WithHTTPClient(client *http.Client) *GetRepoRequirementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repo requirements params
func (o *GetRepoRequirementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoURL adds the repoURL to the get repo requirements params
func (o *GetRepoRequirementsParams) WithRepoURL(repoURL string) *GetRepoRequirementsParams {
	o.SetRepoURL(repoURL)
	return o
}

// SetRepoURL adds the repoUrl to the get repo requirements params
func (o *GetRepoRequirementsParams) SetRepoURL(repoURL string) {
	o.RepoURL = repoURL
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoRequirementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repoUrl
	if err := r.SetPathParam("repoUrl", o.RepoURL); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
