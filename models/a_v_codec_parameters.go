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

// AVCodecParameters a v codec parameters
//
// swagger:model AVCodecParameters
type AVCodecParameters struct {

	// null
	Null bool `json:"null,omitempty"`

	// pointer
	Pointer *Pointer `json:"pointer,omitempty"`
}

// Validate validates this a v codec parameters
func (m *AVCodecParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePointer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AVCodecParameters) validatePointer(formats strfmt.Registry) error {
	if swag.IsZero(m.Pointer) { // not required
		return nil
	}

	if m.Pointer != nil {
		if err := m.Pointer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pointer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pointer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a v codec parameters based on the context it is used
func (m *AVCodecParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePointer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AVCodecParameters) contextValidatePointer(ctx context.Context, formats strfmt.Registry) error {

	if m.Pointer != nil {
		if err := m.Pointer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pointer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pointer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AVCodecParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AVCodecParameters) UnmarshalBinary(b []byte) error {
	var res AVCodecParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
