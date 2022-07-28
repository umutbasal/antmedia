// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebRTCViewerInfo Stores the info for a WebRTC viewer
//
// swagger:model WebRTCViewerInfo
type WebRTCViewerInfo struct {

	// IP address of the edge to which viewer is connected
	EdgeAddress string `json:"edgeAddress,omitempty"`

	// stream id that viewer views
	StreamID string `json:"streamId,omitempty"`

	// the id of the viewer
	ViewerID string `json:"viewerId,omitempty"`
}

// Validate validates this web r t c viewer info
func (m *WebRTCViewerInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this web r t c viewer info based on context it is used
func (m *WebRTCViewerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebRTCViewerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebRTCViewerInfo) UnmarshalBinary(b []byte) error {
	var res WebRTCViewerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}