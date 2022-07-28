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

// GetRTMPToWebRTCStatsReader is a Reader for the GetRTMPToWebRTCStats structure.
type GetRTMPToWebRTCStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRTMPToWebRTCStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRTMPToWebRTCStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRTMPToWebRTCStatsOK creates a GetRTMPToWebRTCStatsOK with default headers values
func NewGetRTMPToWebRTCStatsOK() *GetRTMPToWebRTCStatsOK {
	return &GetRTMPToWebRTCStatsOK{}
}

/* GetRTMPToWebRTCStatsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRTMPToWebRTCStatsOK struct {
	Payload *models.RTMPToWebRTCStats
}

func (o *GetRTMPToWebRTCStatsOK) Error() string {
	return fmt.Sprintf("[GET /v2/broadcasts/{id}/rtmp-to-webrtc-stats][%d] getRTMPToWebRTCStatsOK  %+v", 200, o.Payload)
}
func (o *GetRTMPToWebRTCStatsOK) GetPayload() *models.RTMPToWebRTCStats {
	return o.Payload
}

func (o *GetRTMPToWebRTCStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RTMPToWebRTCStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
