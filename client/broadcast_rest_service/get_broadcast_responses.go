// Code generated by go-swagger;

package broadcast_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetBroadcastReader is a Reader for the GetBroadcast structure.
type GetBroadcastReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBroadcastReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBroadcastOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBroadcastNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBroadcastOK creates a GetBroadcastOK with default headers values
func NewGetBroadcastOK() *GetBroadcastOK {
	return &GetBroadcastOK{}
}

/* GetBroadcastOK describes a response with status code 200, with default header values.

Return the broadcast object
*/
type GetBroadcastOK struct {
}

func (o *GetBroadcastOK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}][%d] getBroadcastOK ", 200)
}

func (o *GetBroadcastOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBroadcastNotFound creates a GetBroadcastNotFound with default headers values
func NewGetBroadcastNotFound() *GetBroadcastNotFound {
	return &GetBroadcastNotFound{}
}

/* GetBroadcastNotFound describes a response with status code 404, with default header values.

Broadcast object not found
*/
type GetBroadcastNotFound struct {
}

func (o *GetBroadcastNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}][%d] getBroadcastNotFound ", 404)
}

func (o *GetBroadcastNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
