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

// ClientConnectionManager client connection manager
//
// swagger:model ClientConnectionManager
type ClientConnectionManager struct {

	// scheme registry
	SchemeRegistry *SchemeRegistry `json:"schemeRegistry,omitempty"`
}

// Validate validates this client connection manager
func (m *ClientConnectionManager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchemeRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientConnectionManager) validateSchemeRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.SchemeRegistry) { // not required
		return nil
	}

	if m.SchemeRegistry != nil {
		if err := m.SchemeRegistry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schemeRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schemeRegistry")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this client connection manager based on the context it is used
func (m *ClientConnectionManager) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchemeRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientConnectionManager) contextValidateSchemeRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.SchemeRegistry != nil {
		if err := m.SchemeRegistry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schemeRegistry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schemeRegistry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientConnectionManager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientConnectionManager) UnmarshalBinary(b []byte) error {
	var res ClientConnectionManager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
