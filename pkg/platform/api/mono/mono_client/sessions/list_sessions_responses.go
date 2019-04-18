// Code generated by go-swagger; DO NOT EDIT.

package sessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// ListSessionsReader is a Reader for the ListSessions structure.
type ListSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListSessionsOK creates a ListSessionsOK with default headers values
func NewListSessionsOK() *ListSessionsOK {
	return &ListSessionsOK{}
}

/*ListSessionsOK handles this case with default header values.

Success
*/
type ListSessionsOK struct {
	Payload []*mono_models.Session
}

func (o *ListSessionsOK) Error() string {
	return fmt.Sprintf("[GET /sessions][%d] listSessionsOK  %+v", 200, o.Payload)
}

func (o *ListSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}