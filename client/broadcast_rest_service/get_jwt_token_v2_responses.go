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

// GetJwtTokenV2Reader is a Reader for the GetJwtTokenV2 structure.
type GetJwtTokenV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJwtTokenV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJwtTokenV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJwtTokenV2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJwtTokenV2OK creates a GetJwtTokenV2OK with default headers values
func NewGetJwtTokenV2OK() *GetJwtTokenV2OK {
	return &GetJwtTokenV2OK{}
}

/* GetJwtTokenV2OK describes a response with status code 200, with default header values.

Returns token
*/
type GetJwtTokenV2OK struct {
	Payload *models.Token
}

func (o *GetJwtTokenV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}/jwt-token][%d] getJwtTokenV2OK  %+v", 200, o.Payload)
}
func (o *GetJwtTokenV2OK) GetPayload() *models.Token {
	return o.Payload
}

func (o *GetJwtTokenV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJwtTokenV2BadRequest creates a GetJwtTokenV2BadRequest with default headers values
func NewGetJwtTokenV2BadRequest() *GetJwtTokenV2BadRequest {
	return &GetJwtTokenV2BadRequest{}
}

/* GetJwtTokenV2BadRequest describes a response with status code 400, with default header values.

When there is an error in creating token
*/
type GetJwtTokenV2BadRequest struct {
	Payload *models.Result
}

func (o *GetJwtTokenV2BadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}/jwt-token][%d] getJwtTokenV2BadRequest  %+v", 400, o.Payload)
}
func (o *GetJwtTokenV2BadRequest) GetPayload() *models.Result {
	return o.Payload
}

func (o *GetJwtTokenV2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}