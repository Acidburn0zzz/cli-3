// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetOrganizationInvitationsReader is a Reader for the GetOrganizationInvitations structure.
type GetOrganizationInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetOrganizationInvitationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationInvitationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationInvitationsOK creates a GetOrganizationInvitationsOK with default headers values
func NewGetOrganizationInvitationsOK() *GetOrganizationInvitationsOK {
	return &GetOrganizationInvitationsOK{}
}

/* GetOrganizationInvitationsOK describes a response with status code 200, with default header values.

Success
*/
type GetOrganizationInvitationsOK struct {
	Payload []*mono_models.Invitation
}

func (o *GetOrganizationInvitationsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/invitations][%d] getOrganizationInvitationsOK  %+v", 200, o.Payload)
}
func (o *GetOrganizationInvitationsOK) GetPayload() []*mono_models.Invitation {
	return o.Payload
}

func (o *GetOrganizationInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInvitationsForbidden creates a GetOrganizationInvitationsForbidden with default headers values
func NewGetOrganizationInvitationsForbidden() *GetOrganizationInvitationsForbidden {
	return &GetOrganizationInvitationsForbidden{}
}

/* GetOrganizationInvitationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationInvitationsForbidden struct {
	Payload *mono_models.Message
}

func (o *GetOrganizationInvitationsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/invitations][%d] getOrganizationInvitationsForbidden  %+v", 403, o.Payload)
}
func (o *GetOrganizationInvitationsForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetOrganizationInvitationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationInvitationsInternalServerError creates a GetOrganizationInvitationsInternalServerError with default headers values
func NewGetOrganizationInvitationsInternalServerError() *GetOrganizationInvitationsInternalServerError {
	return &GetOrganizationInvitationsInternalServerError{}
}

/* GetOrganizationInvitationsInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetOrganizationInvitationsInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetOrganizationInvitationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/invitations][%d] getOrganizationInvitationsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetOrganizationInvitationsInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetOrganizationInvitationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
