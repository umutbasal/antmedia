// Code generated by go-swagger;

package management_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/umutbasal/antmedia/models"
)

// IsFirstLoginReader is a Reader for the IsFirstLogin structure.
type IsFirstLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsFirstLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsFirstLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIsFirstLoginOK creates a IsFirstLoginOK with default headers values
func NewIsFirstLoginOK() *IsFirstLoginOK {
	return &IsFirstLoginOK{}
}

/* IsFirstLoginOK describes a response with status code 200, with default header values.

successful operation
*/
type IsFirstLoginOK struct {
	Payload *models.Result
}

func (o *IsFirstLoginOK) Error() string {
	return fmt.Sprintf("[GET /v2/first-login-status][%d] isFirstLoginOK  %+v", 200, o.Payload)
}
func (o *IsFirstLoginOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *IsFirstLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
