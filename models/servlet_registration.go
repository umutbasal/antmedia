// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServletRegistration servlet registration
//
// swagger:model ServletRegistration
type ServletRegistration struct {

	// class name
	ClassName string `json:"className,omitempty"`

	// init parameters
	InitParameters map[string]string `json:"initParameters,omitempty"`

	// mappings
	Mappings []string `json:"mappings"`

	// name
	Name string `json:"name,omitempty"`

	// run as role
	RunAsRole string `json:"runAsRole,omitempty"`
}

// Validate validates this servlet registration
func (m *ServletRegistration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this servlet registration based on context it is used
func (m *ServletRegistration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServletRegistration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServletRegistration) UnmarshalBinary(b []byte) error {
	var res ServletRegistration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}