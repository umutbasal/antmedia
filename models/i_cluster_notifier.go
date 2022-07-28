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

// IClusterNotifier i cluster notifier
//
// swagger:model IClusterNotifier
type IClusterNotifier struct {

	// cluster store
	ClusterStore *IClusterStore `json:"clusterStore,omitempty"`
}

// Validate validates this i cluster notifier
func (m *IClusterNotifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterStore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IClusterNotifier) validateClusterStore(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterStore) { // not required
		return nil
	}

	if m.ClusterStore != nil {
		if err := m.ClusterStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterStore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterStore")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this i cluster notifier based on the context it is used
func (m *IClusterNotifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterStore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IClusterNotifier) contextValidateClusterStore(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterStore != nil {
		if err := m.ClusterStore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterStore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterStore")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IClusterNotifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IClusterNotifier) UnmarshalBinary(b []byte) error {
	var res IClusterNotifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
