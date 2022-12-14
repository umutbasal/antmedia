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

// GetCameraErrorV2Reader is a Reader for the GetCameraErrorV2 structure.
type GetCameraErrorV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCameraErrorV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCameraErrorV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCameraErrorV2OK creates a GetCameraErrorV2OK with default headers values
func NewGetCameraErrorV2OK() *GetCameraErrorV2OK {
	return &GetCameraErrorV2OK{}
}

/* GetCameraErrorV2OK describes a response with status code 200, with default header values.

successful operation
*/
type GetCameraErrorV2OK struct {
	Payload *models.Result
}

func (o *GetCameraErrorV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{streamId}/ip-camera-error][%d] getCameraErrorV2OK  %+v", 200, o.Payload)
}
func (o *GetCameraErrorV2OK) GetPayload() *models.Result {
	return o.Payload
}

func (o *GetCameraErrorV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
