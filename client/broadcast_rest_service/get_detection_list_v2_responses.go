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

// GetDetectionListV2Reader is a Reader for the GetDetectionListV2 structure.
type GetDetectionListV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDetectionListV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDetectionListV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDetectionListV2OK creates a GetDetectionListV2OK with default headers values
func NewGetDetectionListV2OK() *GetDetectionListV2OK {
	return &GetDetectionListV2OK{}
}

/* GetDetectionListV2OK describes a response with status code 200, with default header values.

successful operation
*/
type GetDetectionListV2OK struct {
	Payload []*models.TensorFlowObject
}

func (o *GetDetectionListV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}/detections/{offset}/{size}][%d] getDetectionListV2OK  %+v", 200, o.Payload)
}
func (o *GetDetectionListV2OK) GetPayload() []*models.TensorFlowObject {
	return o.Payload
}

func (o *GetDetectionListV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
