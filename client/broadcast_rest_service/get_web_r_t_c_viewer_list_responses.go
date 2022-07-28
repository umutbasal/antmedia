// Code generated by go-swagger;

package broadcast_rest_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"antmedia/models"
)

// GetWebRTCViewerListReader is a Reader for the GetWebRTCViewerList structure.
type GetWebRTCViewerListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebRTCViewerListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebRTCViewerListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWebRTCViewerListOK creates a GetWebRTCViewerListOK with default headers values
func NewGetWebRTCViewerListOK() *GetWebRTCViewerListOK {
	return &GetWebRTCViewerListOK{}
}

/* GetWebRTCViewerListOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWebRTCViewerListOK struct {
	Payload []*models.WebRTCViewerInfo
}

func (o *GetWebRTCViewerListOK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/webrtc-viewers/list/{offset}/{size}][%d] getWebRTCViewerListOK  %+v", 200, o.Payload)
}
func (o *GetWebRTCViewerListOK) GetPayload() []*models.WebRTCViewerInfo {
	return o.Payload
}

func (o *GetWebRTCViewerListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
