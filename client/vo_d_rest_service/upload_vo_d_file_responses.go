// Code generated by go-swagger; DO NOT EDIT.

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

// UploadVoDFileReader is a Reader for the UploadVoDFile structure.
type UploadVoDFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadVoDFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadVoDFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadVoDFileOK creates a UploadVoDFileOK with default headers values
func NewUploadVoDFileOK() *UploadVoDFileOK {
	return &UploadVoDFileOK{}
}

/* UploadVoDFileOK describes a response with status code 200, with default header values.

successful operation
*/
type UploadVoDFileOK struct {
	Payload *models.Result
}

func (o *UploadVoDFileOK) Error() string {
	return fmt.Sprintf("[POST /v2/vods/create][%d] uploadVoDFileOK  %+v", 200, o.Payload)
}
func (o *UploadVoDFileOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *UploadVoDFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
