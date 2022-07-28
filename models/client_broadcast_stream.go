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

// ClientBroadcastStream client broadcast stream
//
// swagger:model ClientBroadcastStream
type ClientBroadcastStream struct {

	// absolute start time ms
	AbsoluteStartTimeMs int64 `json:"absoluteStartTimeMs,omitempty"`

	// active subscribers
	ActiveSubscribers int32 `json:"activeSubscribers,omitempty"`

	// automatic recording
	AutomaticRecording bool `json:"automaticRecording,omitempty"`

	// broadcast stream publish name
	BroadcastStreamPublishName string `json:"broadcastStreamPublishName,omitempty"`

	// bytes received
	BytesReceived int64 `json:"bytesReceived,omitempty"`

	// client buffer duration
	ClientBufferDuration int32 `json:"clientBufferDuration,omitempty"`

	// codec info
	CodecInfo *IStreamCodecInfo `json:"codecInfo,omitempty"`

	// connection
	Connection *IStreamCapableConnection `json:"connection,omitempty"`

	// creation time
	CreationTime int64 `json:"creationTime,omitempty"`

	// current timestamp
	CurrentTimestamp int32 `json:"currentTimestamp,omitempty"`

	// max subscribers
	MaxSubscribers int32 `json:"maxSubscribers,omitempty"`

	// meta data
	MetaData *Notify `json:"metaData,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parameters
	Parameters map[string]string `json:"parameters,omitempty"`

	// provider
	Provider IProvider `json:"provider,omitempty"`

	// published name
	PublishedName string `json:"publishedName,omitempty"`

	// recording
	Recording bool `json:"recording,omitempty"`

	// save filename
	SaveFilename string `json:"saveFilename,omitempty"`

	// scope
	Scope *IScope `json:"scope,omitempty"`

	// state
	// Enum: [INIT UNINIT OPEN CLOSED STARTED STOPPED PLAYING PAUSED RESUMED END SEEK]
	State string `json:"state,omitempty"`

	// statistics
	Statistics *IClientBroadcastStreamStatistics `json:"statistics,omitempty"`

	// stream Id
	StreamID Number `json:"streamId,omitempty"`

	// stream listeners
	StreamListeners []IStreamListener `json:"streamListeners"`

	// total subscribers
	TotalSubscribers int32 `json:"totalSubscribers,omitempty"`
}

// Validate validates this client broadcast stream
func (m *ClientBroadcastStream) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodecInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientBroadcastStream) validateCodecInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CodecInfo) { // not required
		return nil
	}

	if m.CodecInfo != nil {
		if err := m.CodecInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codecInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codecInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ClientBroadcastStream) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *ClientBroadcastStream) validateMetaData(formats strfmt.Registry) error {
	if swag.IsZero(m.MetaData) { // not required
		return nil
	}

	if m.MetaData != nil {
		if err := m.MetaData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metaData")
			}
			return err
		}
	}

	return nil
}

func (m *ClientBroadcastStream) validateScope(formats strfmt.Registry) error {
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

var clientBroadcastStreamTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INIT","UNINIT","OPEN","CLOSED","STARTED","STOPPED","PLAYING","PAUSED","RESUMED","END","SEEK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clientBroadcastStreamTypeStatePropEnum = append(clientBroadcastStreamTypeStatePropEnum, v)
	}
}

const (

	// ClientBroadcastStreamStateINIT captures enum value "INIT"
	ClientBroadcastStreamStateINIT string = "INIT"

	// ClientBroadcastStreamStateUNINIT captures enum value "UNINIT"
	ClientBroadcastStreamStateUNINIT string = "UNINIT"

	// ClientBroadcastStreamStateOPEN captures enum value "OPEN"
	ClientBroadcastStreamStateOPEN string = "OPEN"

	// ClientBroadcastStreamStateCLOSED captures enum value "CLOSED"
	ClientBroadcastStreamStateCLOSED string = "CLOSED"

	// ClientBroadcastStreamStateSTARTED captures enum value "STARTED"
	ClientBroadcastStreamStateSTARTED string = "STARTED"

	// ClientBroadcastStreamStateSTOPPED captures enum value "STOPPED"
	ClientBroadcastStreamStateSTOPPED string = "STOPPED"

	// ClientBroadcastStreamStatePLAYING captures enum value "PLAYING"
	ClientBroadcastStreamStatePLAYING string = "PLAYING"

	// ClientBroadcastStreamStatePAUSED captures enum value "PAUSED"
	ClientBroadcastStreamStatePAUSED string = "PAUSED"

	// ClientBroadcastStreamStateRESUMED captures enum value "RESUMED"
	ClientBroadcastStreamStateRESUMED string = "RESUMED"

	// ClientBroadcastStreamStateEND captures enum value "END"
	ClientBroadcastStreamStateEND string = "END"

	// ClientBroadcastStreamStateSEEK captures enum value "SEEK"
	ClientBroadcastStreamStateSEEK string = "SEEK"
)

// prop value enum
func (m *ClientBroadcastStream) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clientBroadcastStreamTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClientBroadcastStream) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *ClientBroadcastStream) validateStatistics(formats strfmt.Registry) error {
	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	if m.Statistics != nil {
		if err := m.Statistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this client broadcast stream based on the context it is used
func (m *ClientBroadcastStream) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCodecInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetaData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatistics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClientBroadcastStream) contextValidateCodecInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CodecInfo != nil {
		if err := m.CodecInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("codecInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("codecInfo")
			}
			return err
		}
	}

	return nil
}

func (m *ClientBroadcastStream) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {
		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *ClientBroadcastStream) contextValidateMetaData(ctx context.Context, formats strfmt.Registry) error {

	if m.MetaData != nil {
		if err := m.MetaData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metaData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metaData")
			}
			return err
		}
	}

	return nil
}

func (m *ClientBroadcastStream) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClientBroadcastStream) contextValidateStatistics(ctx context.Context, formats strfmt.Registry) error {

	if m.Statistics != nil {
		if err := m.Statistics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClientBroadcastStream) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientBroadcastStream) UnmarshalBinary(b []byte) error {
	var res ClientBroadcastStream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
