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

// ListSubscriberV2Reader is a Reader for the ListSubscriberV2 structure.
type ListSubscriberV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSubscriberV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSubscriberV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSubscriberV2OK creates a ListSubscriberV2OK with default headers values
func NewListSubscriberV2OK() *ListSubscriberV2OK {
	return &ListSubscriberV2OK{}
}

/* ListSubscriberV2OK describes a response with status code 200, with default header values.

successful operation
*/
type ListSubscriberV2OK struct {
	Payload []*models.Subscriber
}

func (o *ListSubscriberV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}/subscribers/list/{offset}/{size}][%d] listSubscriberV2OK  %+v", 200, o.Payload)
}
func (o *ListSubscriberV2OK) GetPayload() []*models.Subscriber {
	return o.Payload
}

func (o *ListSubscriberV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}