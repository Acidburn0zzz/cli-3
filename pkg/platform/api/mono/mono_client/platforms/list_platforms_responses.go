// Code generated by go-swagger; DO NOT EDIT.

package platforms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// ListPlatformsReader is a Reader for the ListPlatforms structure.
type ListPlatformsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPlatformsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPlatformsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListPlatformsOK creates a ListPlatformsOK with default headers values
func NewListPlatformsOK() *ListPlatformsOK {
	return &ListPlatformsOK{}
}

/* ListPlatformsOK describes a response with status code 200, with default header values.

Success
*/
type ListPlatformsOK struct {
	Payload []*mono_models.Platform
}

func (o *ListPlatformsOK) Error() string {
	return fmt.Sprintf("[GET /platforms][%d] listPlatformsOK  %+v", 200, o.Payload)
}
func (o *ListPlatformsOK) GetPayload() []*mono_models.Platform {
	return o.Payload
}

func (o *ListPlatformsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
