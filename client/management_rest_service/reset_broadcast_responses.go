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

// ResetBroadcastReader is a Reader for the ResetBroadcast structure.
type ResetBroadcastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetBroadcastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetBroadcastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewResetBroadcastOK creates a ResetBroadcastOK with default headers values
func NewResetBroadcastOK() *ResetBroadcastOK {
	return &ResetBroadcastOK{}
}

/* ResetBroadcastOK describes a response with status code 200, with default header values.

successful operation
*/
type ResetBroadcastOK struct {
	Payload *models.Result
}

func (o *ResetBroadcastOK) Error() string {
	return fmt.Sprintf("[POST /v2/applications/{appname}/reset][%d] resetBroadcastOK  %+v", 200, o.Payload)
}
func (o *ResetBroadcastOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *ResetBroadcastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
