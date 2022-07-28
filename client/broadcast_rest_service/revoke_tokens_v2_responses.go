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

// RevokeTokensV2Reader is a Reader for the RevokeTokensV2 structure.
type RevokeTokensV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeTokensV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeTokensV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRevokeTokensV2OK creates a RevokeTokensV2OK with default headers values
func NewRevokeTokensV2OK() *RevokeTokensV2OK {
	return &RevokeTokensV2OK{}
}

/* RevokeTokensV2OK describes a response with status code 200, with default header values.

successful operation
*/
type RevokeTokensV2OK struct {
	Payload *models.Result
}

func (o *RevokeTokensV2OK) Error() string {
	return fmt.Sprintf("[DELETE /v2/broadcasts/{id}/tokens][%d] revokeTokensV2OK  %+v", 200, o.Payload)
}
func (o *RevokeTokensV2OK) GetPayload() *models.Result {
	return o.Payload
}

func (o *RevokeTokensV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
