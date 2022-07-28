// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Endpoint The endpoint class, such as Facebook, Twitter or custom RTMP endpoints
//
// swagger:model Endpoint
type Endpoint struct {

	// the endpoint service id, this field holds the id of the endpoint
	EndpointServiceID string `json:"endpointServiceId,omitempty"`

	// the RTMP URL of the endpoint
	RtmpURL string `json:"rtmpUrl,omitempty"`

	// Status of the RTMP muxer, possible values are started, finished, failed, broadcasting, {@link IAntMediaStreamHandler#BROADCAST_STATUS_*}
	Status string `json:"status,omitempty"`

	// the service name like facebook, periscope, youtube or generic
	Type string `json:"type,omitempty"`
}

// Validate validates this endpoint
func (m *Endpoint) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this endpoint based on context it is used
func (m *Endpoint) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Endpoint) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Endpoint) UnmarshalBinary(b []byte) error {
	var res Endpoint
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
