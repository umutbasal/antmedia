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

// GetTokenV2Reader is a Reader for the GetTokenV2 structure.
type GetTokenV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTokenV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTokenV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTokenV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTokenV2OK creates a GetTokenV2OK with default headers values
func NewGetTokenV2OK() *GetTokenV2OK {
	return &GetTokenV2OK{}
}

/* GetTokenV2OK describes a response with status code 200, with default header values.

Returns token
*/
type GetTokenV2OK struct {
	Payload *models.Token
}

func (o *GetTokenV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}/token][%d] getTokenV2OK  %+v", 200, o.Payload)
}
func (o *GetTokenV2OK) GetPayload() *models.Token {
	return o.Payload
}

func (o *GetTokenV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTokenV2BadRequest creates a GetTokenV2BadRequest with default headers values
func NewGetTokenV2BadRequest() *GetTokenV2BadRequest {
	return &GetTokenV2BadRequest{}
}

/* GetTokenV2BadRequest describes a response with status code 400, with default header values.

When there is an error in creating token
*/
type GetTokenV2BadRequest struct {
	Payload *models.Result
}

func (o *GetTokenV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}/token][%d] getTokenV2BadRequest  %+v", 400, o.Payload)
}
func (o *GetTokenV2BadRequest) GetPayload() *models.Result {
	return o.Payload
}

func (o *GetTokenV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
