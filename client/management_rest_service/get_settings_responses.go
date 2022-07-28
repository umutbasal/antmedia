// Code generated by go-swagger;

package management_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"antmedia/models"
)

// GetSettingsReader is a Reader for the GetSettings structure.
type GetSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSettingsOK creates a GetSettingsOK with default headers values
func NewGetSettingsOK() *GetSettingsOK {
	return &GetSettingsOK{}
}

/* GetSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSettingsOK struct {
	Payload *models.Result
}

func (o *GetSettingsOK) Error() string {
	return fmt.Sprintf("[GET /v2/applications/settings/{appname}][%d] getSettingsOK  %+v", 200, o.Payload)
}
func (o *GetSettingsOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *GetSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
