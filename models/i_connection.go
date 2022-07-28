// Code generated by go-swagger; DO NOT EDIT.

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

// IConnection i connection
//
// swagger:model IConnection
type IConnection struct {

	// attribute names
	// Unique: true
	AttributeNames []string `json:"attributeNames"`

	// attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// basic scopes
	BasicScopes IteratorIBasicScope `json:"basicScopes,omitempty"`

	// client
	Client *IClient `json:"client,omitempty"`

	// client bytes read
	ClientBytesRead int64 `json:"clientBytesRead,omitempty"`

	// connect params
	ConnectParams map[string]interface{} `json:"connectParams,omitempty"`

	// connected
	Connected bool `json:"connected,omitempty"`

	// dropped messages
	DroppedMessages int64 `json:"droppedMessages,omitempty"`

	// encoding
	// Enum: [AMF0 AMF3 WEBSOCKET SOCKETIO RTP SRTP RAW]
	Encoding string `json:"encoding,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// last ping time
	LastPingTime int32 `json:"lastPingTime,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// pending messages
	PendingMessages int64 `json:"pendingMessages,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// read bytes
	ReadBytes int64 `json:"readBytes,omitempty"`

	// read messages
	ReadMessages int64 `json:"readMessages,omitempty"`

	// remote address
	RemoteAddress string `json:"remoteAddress,omitempty"`

	// remote addresses
	RemoteAddresses []string `json:"remoteAddresses"`

	// remote port
	RemotePort int32 `json:"remotePort,omitempty"`

	// scope
	Scope *IScope `json:"scope,omitempty"`

	// session Id
	SessionID string `json:"sessionId,omitempty"`

	// stream Id
	StreamID Number `json:"streamId,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// written bytes
	WrittenBytes int64 `json:"writtenBytes,omitempty"`

	// written messages
	WrittenMessages int64 `json:"writtenMessages,omitempty"`
}

// Validate validates this i connection
func (m *IConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncoding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IConnection) validateAttributeNames(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("attributeNames", "body", m.AttributeNames); err != nil {
		return err
	}

	return nil
}

func (m *IConnection) validateClient(formats strfmt.Registry) error {
	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

var iConnectionTypeEncodingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AMF0","AMF3","WEBSOCKET","SOCKETIO","RTP","SRTP","RAW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iConnectionTypeEncodingPropEnum = append(iConnectionTypeEncodingPropEnum, v)
	}
}

const (

	// IConnectionEncodingAMF0 captures enum value "AMF0"
	IConnectionEncodingAMF0 string = "AMF0"

	// IConnectionEncodingAMF3 captures enum value "AMF3"
	IConnectionEncodingAMF3 string = "AMF3"

	// IConnectionEncodingWEBSOCKET captures enum value "WEBSOCKET"
	IConnectionEncodingWEBSOCKET string = "WEBSOCKET"

	// IConnectionEncodingSOCKETIO captures enum value "SOCKETIO"
	IConnectionEncodingSOCKETIO string = "SOCKETIO"

	// IConnectionEncodingRTP captures enum value "RTP"
	IConnectionEncodingRTP string = "RTP"

	// IConnectionEncodingSRTP captures enum value "SRTP"
	IConnectionEncodingSRTP string = "SRTP"

	// IConnectionEncodingRAW captures enum value "RAW"
	IConnectionEncodingRAW string = "RAW"
)

// prop value enum
func (m *IConnection) validateEncodingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, iConnectionTypeEncodingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IConnection) validateEncoding(formats strfmt.Registry) error {
	if swag.IsZero(m.Encoding) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncodingEnum("encoding", "body", m.Encoding); err != nil {
		return err
	}

	return nil
}

func (m *IConnection) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this i connection based on the context it is used
func (m *IConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IConnection) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *IConnection) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if m.Scope != nil {
		if err := m.Scope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IConnection) UnmarshalBinary(b []byte) error {
	var res IConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
