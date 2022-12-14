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

// AbstractConsoleDataStore abstract console data store
//
// swagger:model AbstractConsoleDataStore
type AbstractConsoleDataStore struct {

	// available
	Available bool `json:"available,omitempty"`

	// invalid login count map
	InvalidLoginCountMap map[string]int32 `json:"invalidLoginCountMap,omitempty"`

	// is blocked map
	IsBlockedMap map[string]bool `json:"isBlockedMap,omitempty"`

	// number of user records
	NumberOfUserRecords int32 `json:"numberOfUserRecords,omitempty"`

	// user list
	UserList []*User `json:"userList"`
}

// Validate validates this abstract console data store
func (m *AbstractConsoleDataStore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AbstractConsoleDataStore) validateUserList(formats strfmt.Registry) error {
	if swag.IsZero(m.UserList) { // not required
		return nil
	}

	for i := 0; i < len(m.UserList); i++ {
		if swag.IsZero(m.UserList[i]) { // not required
			continue
		}

		if m.UserList[i] != nil {
			if err := m.UserList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this abstract console data store based on the context it is used
func (m *AbstractConsoleDataStore) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AbstractConsoleDataStore) contextValidateUserList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserList); i++ {

		if m.UserList[i] != nil {
			if err := m.UserList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AbstractConsoleDataStore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AbstractConsoleDataStore) UnmarshalBinary(b []byte) error {
	var res AbstractConsoleDataStore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
