// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IGlobalScope i global scope
//
// swagger:model IGlobalScope
type IGlobalScope struct {

	// attribute names
	// Unique: true
	AttributeNames []string `json:"attributeNames"`

	// attributes
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// class loader
	ClassLoader *ClassLoader `json:"classLoader,omitempty"`

	// client connections
	// Unique: true
	ClientConnections []*IConnection `json:"clientConnections"`

	// clients
	// Unique: true
	Clients []*IClient `json:"clients"`

	// connections
	Connections [][]*IConnection `json:"connections"`

	// context
	Context *IContext `json:"context,omitempty"`

	// context path
	ContextPath string `json:"contextPath,omitempty"`

	// depth
	Depth int32 `json:"depth,omitempty"`

	// event listeners
	// Unique: true
	EventListeners []IEventListener `json:"eventListeners"`

	// handler
	Handler IScopeHandler `json:"handler,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parent
	Parent *IScope `json:"parent,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// scope names
	// Unique: true
	ScopeNames []string `json:"scopeNames"`

	// server
	Server *IServer `json:"server,omitempty"`

	// service handler names
	// Unique: true
	ServiceHandlerNames []string `json:"serviceHandlerNames"`

	// statistics
	Statistics *IScopeStatistics `json:"statistics,omitempty"`

	// store
	Store *IPersistenceStore `json:"store,omitempty"`

	// type
	// Enum: [UNDEFINED GLOBAL APPLICATION ROOM BROADCAST SHARED_OBJECT]
	Type string `json:"type,omitempty"`

	// valid
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this i global scope
func (m *IGlobalScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClassLoader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventListeners(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceHandlerNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStore(formats); err != nil {
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

func (m *IGlobalScope) validateAttributeNames(formats strfmt.Registry) error {
	if swag.IsZero(m.AttributeNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("attributeNames", "body", m.AttributeNames); err != nil {
		return err
	}

	return nil
}

func (m *IGlobalScope) validateClassLoader(formats strfmt.Registry) error {
	if swag.IsZero(m.ClassLoader) { // not required
		return nil
	}

	if m.ClassLoader != nil {
		if err := m.ClassLoader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classLoader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("classLoader")
			}
			return err
		}
	}

	return nil
}

func (m *IGlobalScope) validateClientConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientConnections) { // not required
		return nil
	}

	if err := validate.UniqueItems("clientConnections", "body", m.ClientConnections); err != nil {
		return err
	}

	for i := 0; i < len(m.ClientConnections); i++ {
		if swag.IsZero(m.ClientConnections[i]) { // not required
			continue
		}

		if m.ClientConnections[i] != nil {
			if err := m.ClientConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientConnections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clientConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IGlobalScope) validateClients(formats strfmt.Registry) error {
	if swag.IsZero(m.Clients) { // not required
		return nil
	}

	if err := validate.UniqueItems("clients", "body", m.Clients); err != nil {
		return err
	}

	for i := 0; i < len(m.Clients); i++ {
		if swag.IsZero(m.Clients[i]) { // not required
			continue
		}

		if m.Clients[i] != nil {
			if err := m.Clients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clients" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IGlobalScope) validateConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.Connections) { // not required
		return nil
	}

	for i := 0; i < len(m.Connections); i++ {

		if err := validate.UniqueItems("connections"+"."+strconv.Itoa(i), "body", m.Connections[i]); err != nil {
			return err
		}

		for ii := 0; ii < len(m.Connections[i]); ii++ {
			if swag.IsZero(m.Connections[i][ii]) { // not required
				continue
			}

			if m.Connections[i][ii] != nil {
				if err := m.Connections[i][ii].Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("connections" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("connections" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					}
					return err
				}
			}

		}

	}

	return nil
}

func (m *IGlobalScope) validateContext(formats strfmt.Registry) error {
	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if m.Context != nil {
		if err := m.Context.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

func (m *IGlobalScope) validateEventListeners(formats strfmt.Registry) error {
	if swag.IsZero(m.EventListeners) { // not required
		return nil
	}

	if err := validate.UniqueItems("eventListeners", "body", m.EventListeners); err != nil {
		return err
	}

	return nil
}

func (m *IGlobalScope) validateParent(formats strfmt.Registry) error {
	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *IGlobalScope) validateScopeNames(formats strfmt.Registry) error {
	if swag.IsZero(m.ScopeNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("scopeNames", "body", m.ScopeNames); err != nil {
		return err
	}

	return nil
}

func (m *IGlobalScope) validateServer(formats strfmt.Registry) error {
	if swag.IsZero(m.Server) { // not required
		return nil
	}

	if m.Server != nil {
		if err := m.Server.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *IGlobalScope) validateServiceHandlerNames(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceHandlerNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("serviceHandlerNames", "body", m.ServiceHandlerNames); err != nil {
		return err
	}

	return nil
}

func (m *IGlobalScope) validateStatistics(formats strfmt.Registry) error {
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

func (m *IGlobalScope) validateStore(formats strfmt.Registry) error {
	if swag.IsZero(m.Store) { // not required
		return nil
	}

	if m.Store != nil {
		if err := m.Store.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("store")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("store")
			}
			return err
		}
	}

	return nil
}

var iGlobalScopeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNDEFINED","GLOBAL","APPLICATION","ROOM","BROADCAST","SHARED_OBJECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iGlobalScopeTypeTypePropEnum = append(iGlobalScopeTypeTypePropEnum, v)
	}
}

const (

	// IGlobalScopeTypeUNDEFINED captures enum value "UNDEFINED"
	IGlobalScopeTypeUNDEFINED string = "UNDEFINED"

	// IGlobalScopeTypeGLOBAL captures enum value "GLOBAL"
	IGlobalScopeTypeGLOBAL string = "GLOBAL"

	// IGlobalScopeTypeAPPLICATION captures enum value "APPLICATION"
	IGlobalScopeTypeAPPLICATION string = "APPLICATION"

	// IGlobalScopeTypeROOM captures enum value "ROOM"
	IGlobalScopeTypeROOM string = "ROOM"

	// IGlobalScopeTypeBROADCAST captures enum value "BROADCAST"
	IGlobalScopeTypeBROADCAST string = "BROADCAST"

	// IGlobalScopeTypeSHAREDOBJECT captures enum value "SHARED_OBJECT"
	IGlobalScopeTypeSHAREDOBJECT string = "SHARED_OBJECT"
)

// prop value enum
func (m *IGlobalScope) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, iGlobalScopeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IGlobalScope) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this i global scope based on the context it is used
func (m *IGlobalScope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClassLoader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClientConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatistics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStore(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IGlobalScope) contextValidateClassLoader(ctx context.Context, formats strfmt.Registry) error {

	if m.ClassLoader != nil {
		if err := m.ClassLoader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classLoader")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("classLoader")
			}
			return err
		}
	}

	return nil
}

func (m *IGlobalScope) contextValidateClientConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClientConnections); i++ {

		if m.ClientConnections[i] != nil {
			if err := m.ClientConnections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientConnections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clientConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IGlobalScope) contextValidateClients(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Clients); i++ {

		if m.Clients[i] != nil {
			if err := m.Clients[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clients" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("clients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IGlobalScope) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connections); i++ {

		for ii := 0; ii < len(m.Connections[i]); ii++ {

			if m.Connections[i][ii] != nil {
				if err := m.Connections[i][ii].ContextValidate(ctx, formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("connections" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("connections" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					}
					return err
				}
			}

		}

	}

	return nil
}

func (m *IGlobalScope) contextValidateContext(ctx context.Context, formats strfmt.Registry) error {

	if m.Context != nil {
		if err := m.Context.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("context")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("context")
			}
			return err
		}
	}

	return nil
}

func (m *IGlobalScope) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {
		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *IGlobalScope) contextValidateServer(ctx context.Context, formats strfmt.Registry) error {

	if m.Server != nil {
		if err := m.Server.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("server")
			}
			return err
		}
	}

	return nil
}

func (m *IGlobalScope) contextValidateStatistics(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IGlobalScope) contextValidateStore(ctx context.Context, formats strfmt.Registry) error {

	if m.Store != nil {
		if err := m.Store.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("store")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("store")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IGlobalScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IGlobalScope) UnmarshalBinary(b []byte) error {
	var res IGlobalScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}