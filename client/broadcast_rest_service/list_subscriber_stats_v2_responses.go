// Code generated by go-swagger;

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

// ListSubscriberStatsV2Reader is a Reader for the ListSubscriberStatsV2 structure.
type ListSubscriberStatsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSubscriberStatsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSubscriberStatsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListSubscriberStatsV2OK creates a ListSubscriberStatsV2OK with default headers values
func NewListSubscriberStatsV2OK() *ListSubscriberStatsV2OK {
	return &ListSubscriberStatsV2OK{}
}

/* ListSubscriberStatsV2OK describes a response with status code 200, with default header values.

successful operation
*/
type ListSubscriberStatsV2OK struct {
	Payload []*models.SubscriberStats
}

func (o *ListSubscriberStatsV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}/subscriber-stats/list/{offset}/{size}][%d] listSubscriberStatsV2OK  %+v", 200, o.Payload)
}
func (o *ListSubscriberStatsV2OK) GetPayload() []*models.SubscriberStats {
	return o.Payload
}

func (o *ListSubscriberStatsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
