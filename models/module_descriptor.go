// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModuleDescriptor module descriptor
//
// swagger:model ModuleDescriptor
type ModuleDescriptor struct {

	// automatic
	Automatic bool `json:"automatic,omitempty"`

	// open
	Open bool `json:"open,omitempty"`
}

// Validate validates this module descriptor
func (m *ModuleDescriptor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this module descriptor based on context it is used
func (m *ModuleDescriptor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModuleDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModuleDescriptor) UnmarshalBinary(b []byte) error {
	var res ModuleDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
