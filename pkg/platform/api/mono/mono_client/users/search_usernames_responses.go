// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// SearchUsernamesReader is a Reader for the SearchUsernames structure.
type SearchUsernamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUsernamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUsernamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewSearchUsernamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchUsernamesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchUsernamesOK creates a SearchUsernamesOK with default headers values
func NewSearchUsernamesOK() *SearchUsernamesOK {
	return &SearchUsernamesOK{}
}

/* SearchUsernamesOK describes a response with status code 200, with default header values.

Search for users matching the given search string
*/
type SearchUsernamesOK struct {
	Payload []*mono_models.User
}

func (o *SearchUsernamesOK) Error() string {
	return fmt.Sprintf("[POST /users/search_usernames][%d] searchUsernamesOK  %+v", 200, o.Payload)
}
func (o *SearchUsernamesOK) GetPayload() []*mono_models.User {
	return o.Payload
}

func (o *SearchUsernamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsernamesForbidden creates a SearchUsernamesForbidden with default headers values
func NewSearchUsernamesForbidden() *SearchUsernamesForbidden {
	return &SearchUsernamesForbidden{}
}

/* SearchUsernamesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchUsernamesForbidden struct {
	Payload *mono_models.Message
}

func (o *SearchUsernamesForbidden) Error() string {
	return fmt.Sprintf("[POST /users/search_usernames][%d] searchUsernamesForbidden  %+v", 403, o.Payload)
}
func (o *SearchUsernamesForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SearchUsernamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUsernamesInternalServerError creates a SearchUsernamesInternalServerError with default headers values
func NewSearchUsernamesInternalServerError() *SearchUsernamesInternalServerError {
	return &SearchUsernamesInternalServerError{}
}

/* SearchUsernamesInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SearchUsernamesInternalServerError struct {
	Payload *mono_models.Message
}

func (o *SearchUsernamesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/search_usernames][%d] searchUsernamesInternalServerError  %+v", 500, o.Payload)
}
func (o *SearchUsernamesInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SearchUsernamesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SearchUsernamesBody search usernames body
swagger:model SearchUsernamesBody
*/
type SearchUsernamesBody struct {

	// The search query
	Query string `json:"query,omitempty"`
}

// Validate validates this search usernames body
func (o *SearchUsernamesBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search usernames body based on context it is used
func (o *SearchUsernamesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchUsernamesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchUsernamesBody) UnmarshalBinary(b []byte) error {
	var res SearchUsernamesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
