// Code generated by go-swagger; DO NOT EDIT.

package broadcast_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"antmedia/models"
)

// CreateBroadcastReader is a Reader for the CreateBroadcast structure.
type CreateBroadcastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBroadcastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateBroadcastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateBroadcastBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateBroadcastOK creates a CreateBroadcastOK with default headers values
func NewCreateBroadcastOK() *CreateBroadcastOK {
	return &CreateBroadcastOK{}
}

/* CreateBroadcastOK describes a response with status code 200, with default header values.

Returns the created stream
*/
type CreateBroadcastOK struct {
	Payload *models.Broadcast
}

func (o *CreateBroadcastOK) Error() string {
	return fmt.Sprintf("[POST /v2/broadcasts/create][%d] createBroadcastOK  %+v", 200, o.Payload)
}
func (o *CreateBroadcastOK) GetPayload() *models.Broadcast {
	return o.Payload
}

func (o *CreateBroadcastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Broadcast)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBroadcastBadRequest creates a CreateBroadcastBadRequest with default headers values
func NewCreateBroadcastBadRequest() *CreateBroadcastBadRequest {
	return &CreateBroadcastBadRequest{}
}

/* CreateBroadcastBadRequest describes a response with status code 400, with default header values.

If stream id is already used in the data store, it returns error
*/
type CreateBroadcastBadRequest struct {
	Payload *models.Result
}

func (o *CreateBroadcastBadRequest) Error() string {
	return fmt.Sprintf("[POST /v2/broadcasts/create][%d] createBroadcastBadRequest  %+v", 400, o.Payload)
}
func (o *CreateBroadcastBadRequest) GetPayload() *models.Result {
	return o.Payload
}

func (o *CreateBroadcastBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}