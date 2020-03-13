// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewLoginWithGithubParams creates a new LoginWithGithubParams object
// with the default values initialized.
func NewLoginWithGithubParams() *LoginWithGithubParams {
	var ()
	return &LoginWithGithubParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoginWithGithubParamsWithTimeout creates a new LoginWithGithubParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoginWithGithubParamsWithTimeout(timeout time.Duration) *LoginWithGithubParams {
	var ()
	return &LoginWithGithubParams{

		timeout: timeout,
	}
}

// NewLoginWithGithubParamsWithContext creates a new LoginWithGithubParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoginWithGithubParamsWithContext(ctx context.Context) *LoginWithGithubParams {
	var ()
	return &LoginWithGithubParams{

		Context: ctx,
	}
}

// NewLoginWithGithubParamsWithHTTPClient creates a new LoginWithGithubParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoginWithGithubParamsWithHTTPClient(client *http.Client) *LoginWithGithubParams {
	var ()
	return &LoginWithGithubParams{
		HTTPClient: client,
	}
}

/*LoginWithGithubParams contains all the parameters to send to the API endpoint
for the login with github operation typically these are written to a http.Request
*/
type LoginWithGithubParams struct {

	/*Code*/
	Code *string
	/*Error*/
	Error *string
	/*InviteCode*/
	InviteCode *string
	/*NextRoute*/
	NextRoute *string
	/*State*/
	State *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the login with github params
func (o *LoginWithGithubParams) WithTimeout(timeout time.Duration) *LoginWithGithubParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login with github params
func (o *LoginWithGithubParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login with github params
func (o *LoginWithGithubParams) WithContext(ctx context.Context) *LoginWithGithubParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login with github params
func (o *LoginWithGithubParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login with github params
func (o *LoginWithGithubParams) WithHTTPClient(client *http.Client) *LoginWithGithubParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login with github params
func (o *LoginWithGithubParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCode adds the code to the login with github params
func (o *LoginWithGithubParams) WithCode(code *string) *LoginWithGithubParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the login with github params
func (o *LoginWithGithubParams) SetCode(code *string) {
	o.Code = code
}

// WithError adds the error to the login with github params
func (o *LoginWithGithubParams) WithError(error *string) *LoginWithGithubParams {
	o.SetError(error)
	return o
}

// SetError adds the error to the login with github params
func (o *LoginWithGithubParams) SetError(error *string) {
	o.Error = error
}

// WithInviteCode adds the inviteCode to the login with github params
func (o *LoginWithGithubParams) WithInviteCode(inviteCode *string) *LoginWithGithubParams {
	o.SetInviteCode(inviteCode)
	return o
}

// SetInviteCode adds the inviteCode to the login with github params
func (o *LoginWithGithubParams) SetInviteCode(inviteCode *string) {
	o.InviteCode = inviteCode
}

// WithNextRoute adds the nextRoute to the login with github params
func (o *LoginWithGithubParams) WithNextRoute(nextRoute *string) *LoginWithGithubParams {
	o.SetNextRoute(nextRoute)
	return o
}

// SetNextRoute adds the nextRoute to the login with github params
func (o *LoginWithGithubParams) SetNextRoute(nextRoute *string) {
	o.NextRoute = nextRoute
}

// WithState adds the state to the login with github params
func (o *LoginWithGithubParams) WithState(state *string) *LoginWithGithubParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the login with github params
func (o *LoginWithGithubParams) SetState(state *string) {
	o.State = state
}

// WriteToRequest writes these params to a swagger request
func (o *LoginWithGithubParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Code != nil {

		// query param code
		var qrCode string
		if o.Code != nil {
			qrCode = *o.Code
		}
		qCode := qrCode
		if qCode != "" {
			if err := r.SetQueryParam("code", qCode); err != nil {
				return err
			}
		}

	}

	if o.Error != nil {

		// query param error
		var qrError string
		if o.Error != nil {
			qrError = *o.Error
		}
		qError := qrError
		if qError != "" {
			if err := r.SetQueryParam("error", qError); err != nil {
				return err
			}
		}

	}

	if o.InviteCode != nil {

		// query param inviteCode
		var qrInviteCode string
		if o.InviteCode != nil {
			qrInviteCode = *o.InviteCode
		}
		qInviteCode := qrInviteCode
		if qInviteCode != "" {
			if err := r.SetQueryParam("inviteCode", qInviteCode); err != nil {
				return err
			}
		}

	}

	if o.NextRoute != nil {

		// query param nextRoute
		var qrNextRoute string
		if o.NextRoute != nil {
			qrNextRoute = *o.NextRoute
		}
		qNextRoute := qrNextRoute
		if qNextRoute != "" {
			if err := r.SetQueryParam("nextRoute", qNextRoute); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
