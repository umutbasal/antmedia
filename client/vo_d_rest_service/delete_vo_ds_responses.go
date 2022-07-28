// Code generated by go-swagger;

package vo_d_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"antmedia/models"
)

// DeleteVoDsReader is a Reader for the DeleteVoDs structure.
type DeleteVoDsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVoDsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVoDsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVoDsOK creates a DeleteVoDsOK with default headers values
func NewDeleteVoDsOK() *DeleteVoDsOK {
	return &DeleteVoDsOK{}
}

/* DeleteVoDsOK describes a response with status code 200, with default header values.

successful operation
*/
type DeleteVoDsOK struct {
	Payload *models.Result
}

func (o *DeleteVoDsOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/vods/bulk][%d] deleteVoDsOK  %+v", 200, o.Payload)
}
func (o *DeleteVoDsOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *DeleteVoDsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
