// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationContext application context
//
// swagger:model ApplicationContext
type ApplicationContext struct {

	// application name
	ApplicationName string `json:"applicationName,omitempty"`

	// autowire capable bean factory
	AutowireCapableBeanFactory AutowireCapableBeanFactory `json:"autowireCapableBeanFactory,omitempty"`

	// bean definition count
	BeanDefinitionCount int32 `json:"beanDefinitionCount,omitempty"`

	// bean definition names
	BeanDefinitionNames []string `json:"beanDefinitionNames"`

	// class loader
	ClassLoader *ClassLoader `json:"classLoader,omitempty"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// environment
	Environment *Environment `json:"environment,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// parent
	Parent *ApplicationContext `json:"parent,omitempty"`

	// parent bean factory
	ParentBeanFactory BeanFactory `json:"parentBeanFactory,omitempty"`

	// startup date
	StartupDate int64 `json:"startupDate,omitempty"`
}

// Validate validates this application context
func (m *ApplicationContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassLoader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationContext) validateClassLoader(formats strfmt.Registry) error {
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

func (m *ApplicationContext) validateEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.Environment) { // not required
		return nil
	}

	if m.Environment != nil {
		if err := m.Environment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationContext) validateParent(formats strfmt.Registry) error {
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

// ContextValidate validate this application context based on the context it is used
func (m *ApplicationContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClassLoader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEnvironment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationContext) contextValidateClassLoader(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ApplicationContext) contextValidateEnvironment(ctx context.Context, formats strfmt.Registry) error {

	if m.Environment != nil {
		if err := m.Environment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("environment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("environment")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationContext) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ApplicationContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationContext) UnmarshalBinary(b []byte) error {
	var res ApplicationContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
