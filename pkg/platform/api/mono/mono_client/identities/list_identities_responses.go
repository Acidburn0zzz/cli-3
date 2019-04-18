// Code generated by go-swagger; DO NOT EDIT.

package identities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	mono_models "github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// ListIdentitiesReader is a Reader for the ListIdentities structure.
type ListIdentitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIdentitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListIdentitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIdentitiesOK creates a ListIdentitiesOK with default headers values
func NewListIdentitiesOK() *ListIdentitiesOK {
	return &ListIdentitiesOK{}
}

/*ListIdentitiesOK handles this case with default header values.

Success
*/
type ListIdentitiesOK struct {
	Payload []*mono_models.Identity
}

func (o *ListIdentitiesOK) Error() string {
	return fmt.Sprintf("[GET /identities][%d] listIdentitiesOK  %+v", 200, o.Payload)
}

func (o *ListIdentitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}