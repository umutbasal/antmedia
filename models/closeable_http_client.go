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

// CloseableHTTPClient closeable Http client
//
// swagger:model CloseableHttpClient
type CloseableHTTPClient struct {

	// connection manager
	ConnectionManager *ClientConnectionManager `json:"connectionManager,omitempty"`

	// params
	Params HTTPParams `json:"params,omitempty"`
}

// Validate validates this closeable Http client
func (m *CloseableHTTPClient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionManager(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloseableHTTPClient) validateConnectionManager(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionManager) { // not required
		return nil
	}

	if m.ConnectionManager != nil {
		if err := m.ConnectionManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionManager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionManager")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this closeable Http client based on the context it is used
func (m *CloseableHTTPClient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectionManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloseableHTTPClient) contextValidateConnectionManager(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectionManager != nil {
		if err := m.ConnectionManager.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionManager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionManager")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloseableHTTPClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloseableHTTPClient) UnmarshalBinary(b []byte) error {
	var res CloseableHTTPClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}