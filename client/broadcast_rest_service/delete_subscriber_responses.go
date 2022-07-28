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

// DeleteSubscriberReader is a Reader for the DeleteSubscriber structure.
type DeleteSubscriberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubscriberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSubscriberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteSubscriberOK creates a DeleteSubscriberOK with default headers values
func NewDeleteSubscriberOK() *DeleteSubscriberOK {
	return &DeleteSubscriberOK{}
}

/* DeleteSubscriberOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteSubscriberOK struct {
	Payload *models.Result
}

func (o *DeleteSubscriberOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/broadcasts/{id}/subscribers/{sid}][%d] deleteSubscriberOK  %+v", 200, o.Payload)
}
func (o *DeleteSubscriberOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *DeleteSubscriberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
