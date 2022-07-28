// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorageClient storage client
//
// swagger:model StorageClient
type StorageClient struct {

	// access key
	AccessKey string `json:"accessKey,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// endpoint
	Endpoint string `json:"endpoint,omitempty"`

	// permission
	Permission string `json:"permission,omitempty"`

	// region
	Region string `json:"region,omitempty"`

	// secret key
	SecretKey string `json:"secretKey,omitempty"`

	// storage class
	StorageClass string `json:"storageClass,omitempty"`

	// storage name
	StorageName string `json:"storageName,omitempty"`
}

// Validate validates this storage client
func (m *StorageClient) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this storage client based on context it is used
func (m *StorageClient) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StorageClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageClient) UnmarshalBinary(b []byte) error {
	var res StorageClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}