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

// IStreamCapableConnection i stream capable connection
//
// swagger:model IStreamCapableConnection
type IStreamCapableConnection struct {

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

	// streams map
	StreamsMap map[string]IClientStream `json:"streamsMap,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// written bytes
	WrittenBytes int64 `json:"writtenBytes,omitempty"`

	// written messages
	WrittenMessages int64 `json:"writtenMessages,omitempty"`
}

// Validate validates this i stream capable connection
func (m *IStreamCapableConnection) Validate(formats strfmt.Registry) error {
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

	if err := m.validateStreamsMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IStreamCapableConnection) validateAttributeNames(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("attributeNames", "body", m.AttributeNames); err != nil {
		return err
	}

	return nil
}

func (m *IStreamCapableConnection) validateClient(formats strfmt.Registry) error {
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

var iStreamCapableConnectionTypeEncodingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AMF0","AMF3","WEBSOCKET","SOCKETIO","RTP","SRTP","RAW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iStreamCapableConnectionTypeEncodingPropEnum = append(iStreamCapableConnectionTypeEncodingPropEnum, v)
	}
}

const (

	// IStreamCapableConnectionEncodingAMF0 captures enum value "AMF0"
	IStreamCapableConnectionEncodingAMF0 string = "AMF0"

	// IStreamCapableConnectionEncodingAMF3 captures enum value "AMF3"
	IStreamCapableConnectionEncodingAMF3 string = "AMF3"

	// IStreamCapableConnectionEncodingWEBSOCKET captures enum value "WEBSOCKET"
	IStreamCapableConnectionEncodingWEBSOCKET string = "WEBSOCKET"

	// IStreamCapableConnectionEncodingSOCKETIO captures enum value "SOCKETIO"
	IStreamCapableConnectionEncodingSOCKETIO string = "SOCKETIO"

	// IStreamCapableConnectionEncodingRTP captures enum value "RTP"
	IStreamCapableConnectionEncodingRTP string = "RTP"

	// IStreamCapableConnectionEncodingSRTP captures enum value "SRTP"
	IStreamCapableConnectionEncodingSRTP string = "SRTP"

	// IStreamCapableConnectionEncodingRAW captures enum value "RAW"
	IStreamCapableConnectionEncodingRAW string = "RAW"
)

// prop value enum
func (m *IStreamCapableConnection) validateEncodingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, iStreamCapableConnectionTypeEncodingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IStreamCapableConnection) validateEncoding(formats strfmt.Registry) error {
	if swag.IsZero(m.Encoding) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncodingEnum("encoding", "body", m.Encoding); err != nil {
		return err
	}

	return nil
}

func (m *IStreamCapableConnection) validateScope(formats strfmt.Registry) error {
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

func (m *IStreamCapableConnection) validateStreamsMap(formats strfmt.Registry) error {
	if swag.IsZero(m.StreamsMap) { // not required
		return nil
	}

	for k := range m.StreamsMap {

		if err := validate.Required("streamsMap"+"."+k, "body", m.StreamsMap[k]); err != nil {
			return err
		}
		if val, ok := m.StreamsMap[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("streamsMap" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("streamsMap" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this i stream capable connection based on the context it is used
func (m *IStreamCapableConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStreamsMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IStreamCapableConnection) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IStreamCapableConnection) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IStreamCapableConnection) contextValidateStreamsMap(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.StreamsMap {

		if val, ok := m.StreamsMap[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IStreamCapableConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IStreamCapableConnection) UnmarshalBinary(b []byte) error {
	var res IStreamCapableConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
