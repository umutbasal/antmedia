// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for AVRational

// register flags to command
func registerModelAVRationalFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAVRationalNull(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAVRationalPointer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAVRationalNull(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nullDescription := ``

	var nullFlagName string
	if cmdPrefix == "" {
		nullFlagName = "null"
	} else {
		nullFlagName = fmt.Sprintf("%v.null", cmdPrefix)
	}

	var nullFlagDefault bool

	_ = cmd.PersistentFlags().Bool(nullFlagName, nullFlagDefault, nullDescription)

	return nil
}

func registerAVRationalPointer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var pointerFlagName string
	if cmdPrefix == "" {
		pointerFlagName = "pointer"
	} else {
		pointerFlagName = fmt.Sprintf("%v.pointer", cmdPrefix)
	}

	if err := registerModelPointerFlags(depth+1, pointerFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAVRationalFlags(depth int, m *models.AVRational, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nullAdded := retrieveAVRationalNullFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nullAdded

	err, pointerAdded := retrieveAVRationalPointerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pointerAdded

	return nil, retAdded
}

func retrieveAVRationalNullFlags(depth int, m *models.AVRational, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nullFlagName := fmt.Sprintf("%v.null", cmdPrefix)
	if cmd.Flags().Changed(nullFlagName) {

		var nullFlagName string
		if cmdPrefix == "" {
			nullFlagName = "null"
		} else {
			nullFlagName = fmt.Sprintf("%v.null", cmdPrefix)
		}

		nullFlagValue, err := cmd.Flags().GetBool(nullFlagName)
		if err != nil {
			return err, false
		}
		m.Null = nullFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAVRationalPointerFlags(depth int, m *models.AVRational, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pointerFlagName := fmt.Sprintf("%v.pointer", cmdPrefix)
	if cmd.Flags().Changed(pointerFlagName) {
		// info: complex object pointer Pointer is retrieved outside this Changed() block
	}
	pointerFlagValue := m.Pointer
	if swag.IsZero(pointerFlagValue) {
		pointerFlagValue = &models.Pointer{}
	}

	err, pointerAdded := retrieveModelPointerFlags(depth+1, pointerFlagValue, pointerFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pointerAdded
	if pointerAdded {
		m.Pointer = pointerFlagValue
	}

	return nil, retAdded
}
