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

// IClusterStore i cluster store
//
// swagger:model IClusterStore
type IClusterStore struct {

	// all settings
	AllSettings []*AppSettings `json:"allSettings"`

	// node count
	NodeCount int64 `json:"nodeCount,omitempty"`
}

// Validate validates this i cluster store
func (m *IClusterStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IClusterStore) validateAllSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.AllSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.AllSettings); i++ {
		if swag.IsZero(m.AllSettings[i]) { // not required
			continue
		}

		if m.AllSettings[i] != nil {
			if err := m.AllSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this i cluster store based on the context it is used
func (m *IClusterStore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IClusterStore) contextValidateAllSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllSettings); i++ {

		if m.AllSettings[i] != nil {
			if err := m.AllSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IClusterStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IClusterStore) UnmarshalBinary(b []byte) error {
	var res IClusterStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
