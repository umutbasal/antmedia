// Code generated by go-swagger;

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ChangeLogSettingsReader is a Reader for the ChangeLogSettings structure.
type ChangeLogSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeLogSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeLogSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChangeLogSettingsOK creates a ChangeLogSettingsOK with default headers values
func NewChangeLogSettingsOK() *ChangeLogSettingsOK {
	return &ChangeLogSettingsOK{}
}

/* ChangeLogSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type ChangeLogSettingsOK struct {
	Payload string
}

func (o *ChangeLogSettingsOK) Error() string {
	return fmt.Sprintf("[GET /changeLogLevel/{level}][%d] changeLogSettingsOK  %+v", 200, o.Payload)
}
func (o *ChangeLogSettingsOK) GetPayload() string {
	return o.Payload
}

func (o *ChangeLogSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
