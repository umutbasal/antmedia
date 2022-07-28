// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"antmedia/models"
)

// GetNodeCountReader is a Reader for the GetNodeCount structure.
type GetNodeCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNodeCountOK creates a GetNodeCountOK with default headers values
func NewGetNodeCountOK() *GetNodeCountOK {
	return &GetNodeCountOK{}
}

/* GetNodeCountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetNodeCountOK struct {
	Payload *models.SimpleStat
}

func (o *GetNodeCountOK) Error() string {
	return fmt.Sprintf("[GET /v2/cluster/node-count][%d] getNodeCountOK  %+v", 200, o.Payload)
}
func (o *GetNodeCountOK) GetPayload() *models.SimpleStat {
	return o.Payload
}

func (o *GetNodeCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SimpleStat)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}