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

// GetConferenceRoomListReader is a Reader for the GetConferenceRoomList structure.
type GetConferenceRoomListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConferenceRoomListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConferenceRoomListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConferenceRoomListOK creates a GetConferenceRoomListOK with default headers values
func NewGetConferenceRoomListOK() *GetConferenceRoomListOK {
	return &GetConferenceRoomListOK{}
}

/* GetConferenceRoomListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConferenceRoomListOK struct {
	Payload []*models.ConferenceRoom
}

func (o *GetConferenceRoomListOK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/conference-rooms/list/{offset}/{size}][%d] getConferenceRoomListOK  %+v", 200, o.Payload)
}
func (o *GetConferenceRoomListOK) GetPayload() []*models.ConferenceRoom {
	return o.Payload
}

func (o *GetConferenceRoomListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}