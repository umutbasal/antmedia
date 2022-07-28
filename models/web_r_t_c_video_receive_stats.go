// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebRTCVideoReceiveStats web r t c video receive stats
//
// swagger:model WebRTCVideoReceiveStats
type WebRTCVideoReceiveStats struct {

	// video bytes received
	VideoBytesReceived int64 `json:"videoBytesReceived,omitempty"`

	// video bytes received per second
	VideoBytesReceivedPerSecond int64 `json:"videoBytesReceivedPerSecond,omitempty"`

	// video fir count
	VideoFirCount int64 `json:"videoFirCount,omitempty"`

	// video fraction lost
	VideoFractionLost float64 `json:"videoFractionLost,omitempty"`

	// video frame received
	VideoFrameReceived int64 `json:"videoFrameReceived,omitempty"`

	// video frame received per second
	VideoFrameReceivedPerSecond int64 `json:"videoFrameReceivedPerSecond,omitempty"`

	// video nack count
	VideoNackCount int64 `json:"videoNackCount,omitempty"`

	// video packets lost
	VideoPacketsLost int32 `json:"videoPacketsLost,omitempty"`

	// video packets received
	VideoPacketsReceived int64 `json:"videoPacketsReceived,omitempty"`

	// video packets received per second
	VideoPacketsReceivedPerSecond int64 `json:"videoPacketsReceivedPerSecond,omitempty"`

	// video pli count
	VideoPliCount int64 `json:"videoPliCount,omitempty"`
}

// Validate validates this web r t c video receive stats
func (m *WebRTCVideoReceiveStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this web r t c video receive stats based on context it is used
func (m *WebRTCVideoReceiveStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebRTCVideoReceiveStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebRTCVideoReceiveStats) UnmarshalBinary(b []byte) error {
	var res WebRTCVideoReceiveStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
