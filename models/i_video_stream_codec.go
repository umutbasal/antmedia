// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IVideoStreamCodec i video stream codec
//
// swagger:model IVideoStreamCodec
type IVideoStreamCodec struct {

	// decoder configuration
	DecoderConfiguration *IoBuffer `json:"decoderConfiguration,omitempty"`

	// keyframe
	Keyframe *IoBuffer `json:"keyframe,omitempty"`

	// keyframes
	Keyframes []*FrameData `json:"keyframes"`

	// name
	Name string `json:"name,omitempty"`

	// num interframes
	NumInterframes int32 `json:"numInterframes,omitempty"`
}

// Validate validates this i video stream codec
func (m *IVideoStreamCodec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecoderConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyframe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyframes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IVideoStreamCodec) validateDecoderConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.DecoderConfiguration) { // not required
		return nil
	}

	if m.DecoderConfiguration != nil {
		if err := m.DecoderConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decoderConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("decoderConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *IVideoStreamCodec) validateKeyframe(formats strfmt.Registry) error {
	if swag.IsZero(m.Keyframe) { // not required
		return nil
	}

	if m.Keyframe != nil {
		if err := m.Keyframe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyframe")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyframe")
			}
			return err
		}
	}

	return nil
}

func (m *IVideoStreamCodec) validateKeyframes(formats strfmt.Registry) error {
	if swag.IsZero(m.Keyframes) { // not required
		return nil
	}

	for i := 0; i < len(m.Keyframes); i++ {
		if swag.IsZero(m.Keyframes[i]) { // not required
			continue
		}

		if m.Keyframes[i] != nil {
			if err := m.Keyframes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keyframes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("keyframes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this i video stream codec based on the context it is used
func (m *IVideoStreamCodec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDecoderConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyframe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyframes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IVideoStreamCodec) contextValidateDecoderConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.DecoderConfiguration != nil {
		if err := m.DecoderConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("decoderConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("decoderConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *IVideoStreamCodec) contextValidateKeyframe(ctx context.Context, formats strfmt.Registry) error {

	if m.Keyframe != nil {
		if err := m.Keyframe.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyframe")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyframe")
			}
			return err
		}
	}

	return nil
}

func (m *IVideoStreamCodec) contextValidateKeyframes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Keyframes); i++ {

		if m.Keyframes[i] != nil {
			if err := m.Keyframes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keyframes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("keyframes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IVideoStreamCodec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IVideoStreamCodec) UnmarshalBinary(b []byte) error {
	var res IVideoStreamCodec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
