// Code generated by go-swagger; DO NOT EDIT.

package management_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PutV2UsersReader is a Reader for the PutV2Users structure.
type PutV2UsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutV2UsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutV2UsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutV2UsersOK creates a PutV2UsersOK with default headers values
func NewPutV2UsersOK() *PutV2UsersOK {
	return &PutV2UsersOK{}
}

/* PutV2UsersOK describes a response with status code 200, with default header values.

successful operation
*/
type PutV2UsersOK struct {
	Payload []interface{}
}

func (o *PutV2UsersOK) Error() string {
	return fmt.Sprintf("[PUT /v2/users][%d] putV2UsersOK  %+v", 200, o.Payload)
}
func (o *PutV2UsersOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *PutV2UsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}