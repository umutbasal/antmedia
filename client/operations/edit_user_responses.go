// Code generated by go-swagger;

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/umutbasal/antmedia/models"
)

// EditUserReader is a Reader for the EditUser structure.
type EditUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEditUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEditUserOK creates a EditUserOK with default headers values
func NewEditUserOK() *EditUserOK {
	return &EditUserOK{}
}

/* EditUserOK describes a response with status code 200, with default header values.

successful operation
*/
type EditUserOK struct {
	Payload *models.Result
}

func (o *EditUserOK) Error() string {
	return fmt.Sprintf("[POST /editUser][%d] editUserOK  %+v", 200, o.Payload)
}
func (o *EditUserOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *EditUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
