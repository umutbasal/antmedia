// Code generated by go-swagger; DO NOT EDIT.

package management_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"antmedia/models"
)

// DeleteApplicationReader is a Reader for the DeleteApplication structure.
type DeleteApplicationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteApplicationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApplicationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteApplicationOK creates a DeleteApplicationOK with default headers values
func NewDeleteApplicationOK() *DeleteApplicationOK {
	return &DeleteApplicationOK{}
}

/* DeleteApplicationOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteApplicationOK struct {
	Payload *models.Result
}

func (o *DeleteApplicationOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/applications/{appName}][%d] deleteApplicationOK  %+v", 200, o.Payload)
}
func (o *DeleteApplicationOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *DeleteApplicationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}