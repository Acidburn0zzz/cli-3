// Code generated by go-swagger; DO NOT EDIT.

package s3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// SignS3URIReader is a Reader for the SignS3URI structure.
type SignS3URIReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SignS3URIReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSignS3URIOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSignS3URIBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSignS3URIForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSignS3URIInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSignS3URIOK creates a SignS3URIOK with default headers values
func NewSignS3URIOK() *SignS3URIOK {
	return &SignS3URIOK{}
}

/* SignS3URIOK describes a response with status code 200, with default header values.

Success
*/
type SignS3URIOK struct {
	Payload *mono_models.SignedURI
}

func (o *SignS3URIOK) Error() string {
	return fmt.Sprintf("[GET /s3/sign/{URI}][%d] signS3UriOK  %+v", 200, o.Payload)
}
func (o *SignS3URIOK) GetPayload() *mono_models.SignedURI {
	return o.Payload
}

func (o *SignS3URIOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.SignedURI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignS3URIBadRequest creates a SignS3URIBadRequest with default headers values
func NewSignS3URIBadRequest() *SignS3URIBadRequest {
	return &SignS3URIBadRequest{}
}

/* SignS3URIBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SignS3URIBadRequest struct {
	Payload *mono_models.Message
}

func (o *SignS3URIBadRequest) Error() string {
	return fmt.Sprintf("[GET /s3/sign/{URI}][%d] signS3UriBadRequest  %+v", 400, o.Payload)
}
func (o *SignS3URIBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SignS3URIBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignS3URIForbidden creates a SignS3URIForbidden with default headers values
func NewSignS3URIForbidden() *SignS3URIForbidden {
	return &SignS3URIForbidden{}
}

/* SignS3URIForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SignS3URIForbidden struct {
	Payload *mono_models.Message
}

func (o *SignS3URIForbidden) Error() string {
	return fmt.Sprintf("[GET /s3/sign/{URI}][%d] signS3UriForbidden  %+v", 403, o.Payload)
}
func (o *SignS3URIForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SignS3URIForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSignS3URIInternalServerError creates a SignS3URIInternalServerError with default headers values
func NewSignS3URIInternalServerError() *SignS3URIInternalServerError {
	return &SignS3URIInternalServerError{}
}

/* SignS3URIInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type SignS3URIInternalServerError struct {
	Payload *mono_models.Message
}

func (o *SignS3URIInternalServerError) Error() string {
	return fmt.Sprintf("[GET /s3/sign/{URI}][%d] signS3UriInternalServerError  %+v", 500, o.Payload)
}
func (o *SignS3URIInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *SignS3URIInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
