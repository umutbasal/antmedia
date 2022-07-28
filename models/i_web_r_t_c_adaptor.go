// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IWebRTCAdaptor i web r t c adaptor
//
// swagger:model IWebRTCAdaptor
type IWebRTCAdaptor struct {

	// number of live streams
	NumberOfLiveStreams int32 `json:"numberOfLiveStreams,omitempty"`

	// number of total viewers
	NumberOfTotalViewers int32 `json:"numberOfTotalViewers,omitempty"`

	// streams
	// Unique: true
	Streams []string `json:"streams"`
}

// Validate validates this i web r t c adaptor
func (m *IWebRTCAdaptor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStreams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IWebRTCAdaptor) validateStreams(formats strfmt.Registry) error {
	if swag.IsZero(m.Streams) { // not required
		return nil
	}

	if err := validate.UniqueItems("streams", "body", m.Streams); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this i web r t c adaptor based on context it is used
func (m *IWebRTCAdaptor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IWebRTCAdaptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IWebRTCAdaptor) UnmarshalBinary(b []byte) error {
	var res IWebRTCAdaptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
