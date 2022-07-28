// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebRTCSendStats Aggregation of WebRTC Low Level Send Stats
//
// swagger:model WebRTCSendStats
type WebRTCSendStats struct {

	// Audio send stats
	AudioSendStats *WebRTCAudioSendStats `json:"audioSendStats,omitempty"`

	// Video send stats
	VideoSendStats *WebRTCVideoSendStats `json:"videoSendStats,omitempty"`
}

// Validate validates this web r t c send stats
func (m *WebRTCSendStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudioSendStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVideoSendStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebRTCSendStats) validateAudioSendStats(formats strfmt.Registry) error {
	if swag.IsZero(m.AudioSendStats) { // not required
		return nil
	}

	if m.AudioSendStats != nil {
		if err := m.AudioSendStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audioSendStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audioSendStats")
			}
			return err
		}
	}

	return nil
}

func (m *WebRTCSendStats) validateVideoSendStats(formats strfmt.Registry) error {
	if swag.IsZero(m.VideoSendStats) { // not required
		return nil
	}

	if m.VideoSendStats != nil {
		if err := m.VideoSendStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("videoSendStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("videoSendStats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this web r t c send stats based on the context it is used
func (m *WebRTCSendStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAudioSendStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVideoSendStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebRTCSendStats) contextValidateAudioSendStats(ctx context.Context, formats strfmt.Registry) error {

	if m.AudioSendStats != nil {
		if err := m.AudioSendStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("audioSendStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("audioSendStats")
			}
			return err
		}
	}

	return nil
}

func (m *WebRTCSendStats) contextValidateVideoSendStats(ctx context.Context, formats strfmt.Registry) error {

	if m.VideoSendStats != nil {
		if err := m.VideoSendStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("videoSendStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("videoSendStats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebRTCSendStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebRTCSendStats) UnmarshalBinary(b []byte) error {
	var res WebRTCSendStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
