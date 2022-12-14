// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Notify notify
//
// swagger:model Notify
type Notify struct {

	// action
	Action string `json:"action,omitempty"`

	// call
	Call *IServiceCall `json:"call,omitempty"`

	// connection params
	ConnectionParams map[string]interface{} `json:"connectionParams,omitempty"`

	// data
	Data *IoBuffer `json:"data,omitempty"`

	// data type
	// Format: byte
	DataType strfmt.Base64 `json:"dataType,omitempty"`

	// header
	Header *Header `json:"header,omitempty"`

	// object
	Object interface{} `json:"object,omitempty"`

	// source
	Source IEventListener `json:"source,omitempty"`

	// source type
	// Format: byte
	SourceType strfmt.Base64 `json:"sourceType,omitempty"`

	// timestamp
	Timestamp int32 `json:"timestamp,omitempty"`

	// transaction Id
	TransactionID int32 `json:"transactionId,omitempty"`

	// type
	// Enum: [SYSTEM STATUS SERVICE_CALL SHARED_OBJECT STREAM_ACTION STREAM_CONTROL STREAM_DATA CLIENT CLIENT_INVOKE CLIENT_NOTIFY SERVER]
	Type string `json:"type,omitempty"`
}

// Validate validates this notify
func (m *Notify) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notify) validateCall(formats strfmt.Registry) error {
	if swag.IsZero(m.Call) { // not required
		return nil
	}

	if m.Call != nil {
		if err := m.Call.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("call")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("call")
			}
			return err
		}
	}

	return nil
}

func (m *Notify) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *Notify) validateHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.Header) { // not required
		return nil
	}

	if m.Header != nil {
		if err := m.Header.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

var notifyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SYSTEM","STATUS","SERVICE_CALL","SHARED_OBJECT","STREAM_ACTION","STREAM_CONTROL","STREAM_DATA","CLIENT","CLIENT_INVOKE","CLIENT_NOTIFY","SERVER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notifyTypeTypePropEnum = append(notifyTypeTypePropEnum, v)
	}
}

const (

	// NotifyTypeSYSTEM captures enum value "SYSTEM"
	NotifyTypeSYSTEM string = "SYSTEM"

	// NotifyTypeSTATUS captures enum value "STATUS"
	NotifyTypeSTATUS string = "STATUS"

	// NotifyTypeSERVICECALL captures enum value "SERVICE_CALL"
	NotifyTypeSERVICECALL string = "SERVICE_CALL"

	// NotifyTypeSHAREDOBJECT captures enum value "SHARED_OBJECT"
	NotifyTypeSHAREDOBJECT string = "SHARED_OBJECT"

	// NotifyTypeSTREAMACTION captures enum value "STREAM_ACTION"
	NotifyTypeSTREAMACTION string = "STREAM_ACTION"

	// NotifyTypeSTREAMCONTROL captures enum value "STREAM_CONTROL"
	NotifyTypeSTREAMCONTROL string = "STREAM_CONTROL"

	// NotifyTypeSTREAMDATA captures enum value "STREAM_DATA"
	NotifyTypeSTREAMDATA string = "STREAM_DATA"

	// NotifyTypeCLIENT captures enum value "CLIENT"
	NotifyTypeCLIENT string = "CLIENT"

	// NotifyTypeCLIENTINVOKE captures enum value "CLIENT_INVOKE"
	NotifyTypeCLIENTINVOKE string = "CLIENT_INVOKE"

	// NotifyTypeCLIENTNOTIFY captures enum value "CLIENT_NOTIFY"
	NotifyTypeCLIENTNOTIFY string = "CLIENT_NOTIFY"

	// NotifyTypeSERVER captures enum value "SERVER"
	NotifyTypeSERVER string = "SERVER"
)

// prop value enum
func (m *Notify) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, notifyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Notify) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this notify based on the context it is used
func (m *Notify) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notify) contextValidateCall(ctx context.Context, formats strfmt.Registry) error {

	if m.Call != nil {
		if err := m.Call.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("call")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("call")
			}
			return err
		}
	}

	return nil
}

func (m *Notify) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {
		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *Notify) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.Header != nil {
		if err := m.Header.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notify) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notify) UnmarshalBinary(b []byte) error {
	var res Notify
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
