// Code generated by go-swagger;

package management_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/umutbasal/antmedia/models"
)

// GetLogFileReader is a Reader for the GetLogFile structure.
type GetLogFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogFileOK creates a GetLogFileOK with default headers values
func NewGetLogFileOK() *GetLogFileOK {
	return &GetLogFileOK{}
}

/* GetLogFileOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLogFileOK struct {
	Payload *models.Result
}

func (o *GetLogFileOK) Error() string {
	return fmt.Sprintf("[GET /v2/log-file/{offsetSize}/{charSize}][%d] getLogFileOK  %+v", 200, o.Payload)
}
func (o *GetLogFileOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *GetLogFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
