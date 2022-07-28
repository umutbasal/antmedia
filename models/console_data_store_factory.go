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

// ConsoleDataStoreFactory console data store factory
//
// swagger:model ConsoleDataStoreFactory
type ConsoleDataStoreFactory struct {

	// app name
	AppName string `json:"appName,omitempty"`

	// data store
	DataStore *AbstractConsoleDataStore `json:"dataStore,omitempty"`

	// db host
	DbHost string `json:"dbHost,omitempty"`

	// db name
	DbName string `json:"dbName,omitempty"`

	// db password
	DbPassword string `json:"dbPassword,omitempty"`

	// db type
	DbType string `json:"dbType,omitempty"`

	// db user
	DbUser string `json:"dbUser,omitempty"`
}

// Validate validates this console data store factory
func (m *ConsoleDataStoreFactory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataStore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsoleDataStoreFactory) validateDataStore(formats strfmt.Registry) error {
	if swag.IsZero(m.DataStore) { // not required
		return nil
	}

	if m.DataStore != nil {
		if err := m.DataStore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataStore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataStore")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this console data store factory based on the context it is used
func (m *ConsoleDataStoreFactory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataStore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsoleDataStoreFactory) contextValidateDataStore(ctx context.Context, formats strfmt.Registry) error {

	if m.DataStore != nil {
		if err := m.DataStore.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataStore")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dataStore")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsoleDataStoreFactory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsoleDataStoreFactory) UnmarshalBinary(b []byte) error {
	var res ConsoleDataStoreFactory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
