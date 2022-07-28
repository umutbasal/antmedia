// Code generated by go-swagger;

package broadcast_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/umutbasal/antmedia/models"
)

// StopMoveReader is a Reader for the StopMove structure.
type StopMoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopMoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopMoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStopMoveOK creates a StopMoveOK with default headers values
func NewStopMoveOK() *StopMoveOK {
	return &StopMoveOK{}
}

/* StopMoveOK describes a response with status code 200, with default header values.

successful operation
*/
type StopMoveOK struct {
	Payload *models.Result
}

func (o *StopMoveOK) Error() string {
	return fmt.Sprintf("[POST /v2/broadcasts/{id}/ip-camera/stop-move][%d] stopMoveOK  %+v", 200, o.Payload)
}
func (o *StopMoveOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *StopMoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
