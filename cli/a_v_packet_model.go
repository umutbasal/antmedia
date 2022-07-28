// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for AVPacket

// register flags to command
func registerModelAVPacketFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAVPacketNull(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAVPacketPointer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAVPacketNull(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerAVPacketPointer(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelAVPacketFlags(depth int, m *models.AVPacket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, nullAdded := retrieveAVPacketNullFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nullAdded

	err, pointerAdded := retrieveAVPacketPointerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pointerAdded

	return nil, retAdded
}

func retrieveAVPacketNullFlags(depth int, m *models.AVPacket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAVPacketPointerFlags(depth int, m *models.AVPacket, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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