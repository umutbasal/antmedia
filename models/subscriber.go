// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Subscriber The time based token subscriber class
//
// swagger:model Subscriber
type Subscriber struct {

	// is subscriber connected
	Connected bool `json:"connected,omitempty"`

	// the subscriber id of the subscriber
	SubscriberID string `json:"subscriberId,omitempty"`

	//  type of subscriber (play or publish)
	Type string `json:"type,omitempty"`
}

// Validate validates this subscriber
func (m *Subscriber) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this subscriber based on context it is used
func (m *Subscriber) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Subscriber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subscriber) UnmarshalBinary(b []byte) error {
	var res Subscriber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}