// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetLoginJwtToken(params *GetLoginJwtTokenParams, opts ...ClientOption) error

	GetLogout(params *GetLogoutParams, opts ...ClientOption) (*GetLogoutNoContent, error)

	GetRenew(params *GetRenewParams, opts ...ClientOption) (*GetRenewOK, error)

	PostLogin(params *PostLoginParams, opts ...ClientOption) (*PostLoginOK, error)

	AddToken(params *AddTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddTokenOK, error)

	ChangePassword(params *ChangePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangePasswordOK, error)

	DeleteToken(params *DeleteTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTokenOK, error)

	DisableTOTP(params *DisableTOTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableTOTPOK, error)

	EnableTOTP(params *EnableTOTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableTOTPOK, error)

	ListTokens(params *ListTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTokensOK, error)

	LoginAs(params *LoginAsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LoginAsOK, error)

	LoginWithGithub(params *LoginWithGithubParams, opts ...ClientOption) error

	NewTOTP(params *NewTOTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NewTOTPOK, error)

	RequestReset(params *RequestResetParams, opts ...ClientOption) (*RequestResetOK, error)

	ResetPassword(params *ResetPasswordParams, opts ...ClientOption) (*ResetPasswordOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetLoginJwtToken logins with a valid j w t and redirect to a platform URL

  Login with a valid JWT and redirect to a platform URL
*/
func (a *Client) GetLoginJwtToken(params *GetLoginJwtTokenParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoginJwtTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLoginJwtToken",
		Method:             "GET",
		PathPattern:        "/login/jwt/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLoginJwtTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  GetLogout renews a valid j w t

  Log out of the current session
*/
func (a *Client) GetLogout(params *GetLogoutParams, opts ...ClientOption) (*GetLogoutNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogoutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLogout",
		Method:             "GET",
		PathPattern:        "/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLogoutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLogoutNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLogout: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetRenew renews a valid j w t

  Renew your current JWT to forestall expiration
*/
func (a *Client) GetRenew(params *GetRenewParams, opts ...ClientOption) (*GetRenewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRenewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetRenew",
		Method:             "GET",
		PathPattern:        "/renew",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRenewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetRenewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetRenew: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostLogin trades your credentials for a j w t

  Supply either username/password OR token
*/
func (a *Client) PostLogin(params *PostLoginParams, opts ...ClientOption) (*PostLoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostLogin",
		Method:             "POST",
		PathPattern:        "/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostLoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLogin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddToken generates an API token for current user

  Produces an API token for use with automated API clients
*/
func (a *Client) AddToken(params *AddTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addToken",
		Method:             "POST",
		PathPattern:        "/apikeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ChangePassword changes the current password

  Prompts for current password which is used to change it to something new.
*/
func (a *Client) ChangePassword(params *ChangePasswordParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ChangePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangePasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "changePassword",
		Method:             "POST",
		PathPattern:        "/change-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangePasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changePassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteToken deletes an API token

  Deletes the specified API Token
*/
func (a *Client) DeleteToken(params *DeleteTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteToken",
		Method:             "DELETE",
		PathPattern:        "/apikeys/{tokenID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DisableTOTP disables t o t p

  Disable TOTP authentication
*/
func (a *Client) DisableTOTP(params *DisableTOTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableTOTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableTOTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disableTOTP",
		Method:             "DELETE",
		PathPattern:        "/totp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisableTOTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableTOTPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disableTOTP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EnableTOTP enables t o t p

  Enable TOTP authentication by performing initial code validation
*/
func (a *Client) EnableTOTP(params *EnableTOTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableTOTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableTOTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enableTOTP",
		Method:             "POST",
		PathPattern:        "/totp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EnableTOTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableTOTPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enableTOTP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListTokens lists current user s API tokens

  List of all active API Tokens for current user
*/
func (a *Client) ListTokens(params *ListTokensParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTokensOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTokensParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTokens",
		Method:             "GET",
		PathPattern:        "/apikeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTokensReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListTokensOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTokens: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LoginAs logins as given user requires you to be a superuser
*/
func (a *Client) LoginAs(params *LoginAsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LoginAsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginAsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "loginAs",
		Method:             "POST",
		PathPattern:        "/login/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LoginAsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginAsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for loginAs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LoginWithGithub callbacks endpoint for github auth
*/
func (a *Client) LoginWithGithub(params *LoginWithGithubParams, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginWithGithubParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "loginWithGithub",
		Method:             "GET",
		PathPattern:        "/githubLogin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LoginWithGithubReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  NewTOTP sets up a new t o t p key

  Establish the private key for two-factor authentication
*/
func (a *Client) NewTOTP(params *NewTOTPParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NewTOTPOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNewTOTPParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "newTOTP",
		Method:             "GET",
		PathPattern:        "/totp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NewTOTPReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*NewTOTPOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for newTOTP: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RequestReset requests a password recovery email

  Sends a link which can be used to reset a forgotten password.
*/
func (a *Client) RequestReset(params *RequestResetParams, opts ...ClientOption) (*RequestResetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestResetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "requestReset",
		Method:             "POST",
		PathPattern:        "/request-reset/{email}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RequestResetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RequestResetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for requestReset: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ResetPassword resets a forgotten password

  Sends a link which can be used to reset a forgotten password.
*/
func (a *Client) ResetPassword(params *ResetPasswordParams, opts ...ClientOption) (*ResetPasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResetPasswordParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "resetPassword",
		Method:             "POST",
		PathPattern:        "/reset-password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ResetPasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResetPasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for resetPassword: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
