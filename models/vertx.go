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

// Vertx vertx
//
// swagger:model Vertx
type Vertx struct {

	// clustered
	Clustered bool `json:"clustered,omitempty"`

	// metrics enabled
	MetricsEnabled bool `json:"metricsEnabled,omitempty"`

	// native transport enabled
	NativeTransportEnabled bool `json:"nativeTransportEnabled,omitempty"`

	// or create context
	OrCreateContext *Context `json:"orCreateContext,omitempty"`
}

// Validate validates this vertx
func (m *Vertx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrCreateContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vertx) validateOrCreateContext(formats strfmt.Registry) error {
	if swag.IsZero(m.OrCreateContext) { // not required
		return nil
	}

	if m.OrCreateContext != nil {
		if err := m.OrCreateContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orCreateContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orCreateContext")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vertx based on the context it is used
func (m *Vertx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrCreateContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vertx) contextValidateOrCreateContext(ctx context.Context, formats strfmt.Registry) error {

	if m.OrCreateContext != nil {
		if err := m.OrCreateContext.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orCreateContext")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orCreateContext")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vertx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vertx) UnmarshalBinary(b []byte) error {
	var res Vertx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
