// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IAudioStreamCodec i audio stream codec
//
// swagger:model IAudioStreamCodec
type IAudioStreamCodec struct {

	// decoder configuration
	DecoderConfiguration *IoBuffer `json:"decoderConfiguration,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this i audio stream codec
func (m *IAudioStreamCodec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDecoderConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IAudioStreamCodec) validateDecoderConfiguration(formats strfmt.Registry) error {
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

// ContextValidate validate this i audio stream codec based on the context it is used
func (m *IAudioStreamCodec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDecoderConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IAudioStreamCodec) contextValidateDecoderConfiguration(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *IAudioStreamCodec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IAudioStreamCodec) UnmarshalBinary(b []byte) error {
	var res IAudioStreamCodec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
