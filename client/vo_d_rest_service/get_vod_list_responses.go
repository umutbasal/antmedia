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

// GetVodListReader is a Reader for the GetVodList structure.
type GetVodListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVodListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVodListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVodListOK creates a GetVodListOK with default headers values
func NewGetVodListOK() *GetVodListOK {
	return &GetVodListOK{}
}

/* GetVodListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetVodListOK struct {
	Payload []*models.VoD
}

func (o *GetVodListOK) Error() string {
	return fmt.Sprintf("[GET /v2/vods/list/{offset}/{size}][%d] getVodListOK  %+v", 200, o.Payload)
}
func (o *GetVodListOK) GetPayload() []*models.VoD {
	return o.Payload
}

func (o *GetVodListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
