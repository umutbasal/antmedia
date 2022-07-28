// Code generated by go-swagger; DO NOT EDIT.

package broadcast_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRoomInfoReader is a Reader for the GetRoomInfo structure.
type GetRoomInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoomInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoomInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoomInfoOK creates a GetRoomInfoOK with default headers values
func NewGetRoomInfoOK() *GetRoomInfoOK {
	return &GetRoomInfoOK{}
}

/* GetRoomInfoOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoomInfoOK struct {
	Payload []string
}

func (o *GetRoomInfoOK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/conference-rooms/{room_id}/room-info][%d] getRoomInfoOK  %+v", 200, o.Payload)
}
func (o *GetRoomInfoOK) GetPayload() []string {
	return o.Payload
}

func (o *GetRoomInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
