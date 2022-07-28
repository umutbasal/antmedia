// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaglibDescriptor taglib descriptor
//
// swagger:model TaglibDescriptor
type TaglibDescriptor struct {

	// taglib location
	TaglibLocation string `json:"taglibLocation,omitempty"`

	// taglib URI
	TaglibURI string `json:"taglibURI,omitempty"`
}

// Validate validates this taglib descriptor
func (m *TaglibDescriptor) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this taglib descriptor based on context it is used
func (m *TaglibDescriptor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaglibDescriptor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaglibDescriptor) UnmarshalBinary(b []byte) error {
	var res TaglibDescriptor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
