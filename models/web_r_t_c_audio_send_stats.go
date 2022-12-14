// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebRTCAudioSendStats web r t c audio send stats
//
// swagger:model WebRTCAudioSendStats
type WebRTCAudioSendStats struct {

	// audio bytes sent
	AudioBytesSent int64 `json:"audioBytesSent,omitempty"`

	// audio bytes sent per second
	AudioBytesSentPerSecond int64 `json:"audioBytesSentPerSecond,omitempty"`

	// audio packets per second
	AudioPacketsPerSecond int64 `json:"audioPacketsPerSecond,omitempty"`

	// audio packets sent
	AudioPacketsSent int64 `json:"audioPacketsSent,omitempty"`

	// time ms
	TimeMs int64 `json:"timeMs,omitempty"`
}

// Validate validates this web r t c audio send stats
func (m *WebRTCAudioSendStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this web r t c audio send stats based on context it is used
func (m *WebRTCAudioSendStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebRTCAudioSendStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebRTCAudioSendStats) UnmarshalBinary(b []byte) error {
	var res WebRTCAudioSendStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
