// Code generated by go-swagger; DO NOT EDIT.

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

// GetJVMMemoryInfoReader is a Reader for the GetJVMMemoryInfo structure.
type GetJVMMemoryInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJVMMemoryInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJVMMemoryInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJVMMemoryInfoOK creates a GetJVMMemoryInfoOK with default headers values
func NewGetJVMMemoryInfoOK() *GetJVMMemoryInfoOK {
	return &GetJVMMemoryInfoOK{}
}

/* GetJVMMemoryInfoOK describes a response with status code 200, with default header values.

successful operation
*/
type GetJVMMemoryInfoOK struct {
	Payload *models.Result
}

func (o *GetJVMMemoryInfoOK) Error() string {
	return fmt.Sprintf("[GET /v2/jvm-memory-status][%d] getJVmMemoryInfoOK  %+v", 200, o.Payload)
}
func (o *GetJVMMemoryInfoOK) GetPayload() *models.Result {
	return o.Payload
}

func (o *GetJVMMemoryInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Result)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
