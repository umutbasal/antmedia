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

// GetTotalBroadcastNumberV2Reader is a Reader for the GetTotalBroadcastNumberV2 structure.
type GetTotalBroadcastNumberV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTotalBroadcastNumberV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTotalBroadcastNumberV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTotalBroadcastNumberV2OK creates a GetTotalBroadcastNumberV2OK with default headers values
func NewGetTotalBroadcastNumberV2OK() *GetTotalBroadcastNumberV2OK {
	return &GetTotalBroadcastNumberV2OK{}
}

/* GetTotalBroadcastNumberV2OK describes a response with status code 200, with default header values.

successful operation
*/
type GetTotalBroadcastNumberV2OK struct {
	Payload *models.SimpleStat
}

func (o *GetTotalBroadcastNumberV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/count/{search}][%d] getTotalBroadcastNumberV2OK  %+v", 200, o.Payload)
}
func (o *GetTotalBroadcastNumberV2OK) GetPayload() *models.SimpleStat {
	return o.Payload
}

func (o *GetTotalBroadcastNumberV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleStat)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
