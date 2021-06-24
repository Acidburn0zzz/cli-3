// Code generated by go-swagger; DO NOT EDIT.

package languages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// ListLanguagesReader is a Reader for the ListLanguages structure.
type ListLanguagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLanguagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLanguagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListLanguagesOK creates a ListLanguagesOK with default headers values
func NewListLanguagesOK() *ListLanguagesOK {
	return &ListLanguagesOK{}
}

/* ListLanguagesOK describes a response with status code 200, with default header values.

Success
*/
type ListLanguagesOK struct {
	Payload []*mono_models.Language
}

func (o *ListLanguagesOK) Error() string {
	return fmt.Sprintf("[GET /languages][%d] listLanguagesOK  %+v", 200, o.Payload)
}
func (o *ListLanguagesOK) GetPayload() []*mono_models.Language {
	return o.Payload
}

func (o *ListLanguagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
