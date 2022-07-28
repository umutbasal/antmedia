// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebRTCVideoSendStats web r t c video send stats
//
// swagger:model WebRTCVideoSendStats
type WebRTCVideoSendStats struct {

	// time ms
	TimeMs int64 `json:"timeMs,omitempty"`

	// video bytes sent
	VideoBytesSent int64 `json:"videoBytesSent,omitempty"`

	// video bytes sent per second
	VideoBytesSentPerSecond int64 `json:"videoBytesSentPerSecond,omitempty"`

	// video fir count
	VideoFirCount int64 `json:"videoFirCount,omitempty"`

	// video frames encoded
	VideoFramesEncoded int64 `json:"videoFramesEncoded,omitempty"`

	// video frames encoded per second
	VideoFramesEncodedPerSecond int64 `json:"videoFramesEncodedPerSecond,omitempty"`

	// video nack count
	VideoNackCount int64 `json:"videoNackCount,omitempty"`

	// video packets sent
	VideoPacketsSent int64 `json:"videoPacketsSent,omitempty"`

	// video packets sent per second
	VideoPacketsSentPerSecond int64 `json:"videoPacketsSentPerSecond,omitempty"`

	// video pli count
	VideoPliCount int64 `json:"videoPliCount,omitempty"`
}

// Validate validates this web r t c video send stats
func (m *WebRTCVideoSendStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this web r t c video send stats based on context it is used
func (m *WebRTCVideoSendStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebRTCVideoSendStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebRTCVideoSendStats) UnmarshalBinary(b []byte) error {
	var res WebRTCVideoSendStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
