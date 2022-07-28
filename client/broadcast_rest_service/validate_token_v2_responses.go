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

// ValidateTokenV2Reader is a Reader for the ValidateTokenV2 structure.
type ValidateTokenV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateTokenV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateTokenV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateTokenV2OK creates a ValidateTokenV2OK with default headers values
func NewValidateTokenV2OK() *ValidateTokenV2OK {
	return &ValidateTokenV2OK{}
}

/* ValidateTokenV2OK describes a response with status code 200, with default header values.

successful operation
*/
type ValidateTokenV2OK struct {
	Payload *models.Result
}

func (o *ValidateTokenV2OK) Error() string {
	return fmt.Sprintf("[POST /v2/broadcasts/validate-token][%d] validateTokenV2OK  %+v", 200, o.Payload)
}
func (o *ValidateTokenV2OK) GetPayload() *models.Result {
	return o.Payload
}

func (o *ValidateTokenV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}