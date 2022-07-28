// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for FrameData

// register flags to command
func registerModelFrameDataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerFrameDataFrame(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerFrameDataFrame(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var frameFlagName string
	if cmdPrefix == "" {
		frameFlagName = "frame"
	} else {
		frameFlagName = fmt.Sprintf("%v.frame", cmdPrefix)
	}

	if err := registerModelIoBufferFlags(depth+1, frameFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelFrameDataFlags(depth int, m *models.FrameData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, frameAdded := retrieveFrameDataFrameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || frameAdded

	return nil, retAdded
}

func retrieveFrameDataFrameFlags(depth int, m *models.FrameData, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	frameFlagName := fmt.Sprintf("%v.frame", cmdPrefix)
	if cmd.Flags().Changed(frameFlagName) {
		// info: complex object frame IoBuffer is retrieved outside this Changed() block
	}
	frameFlagValue := m.Frame
	if swag.IsZero(frameFlagValue) {
		frameFlagValue = &models.IoBuffer{}
	}

	err, frameAdded := retrieveModelIoBufferFlags(depth+1, frameFlagValue, frameFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || frameAdded
	if frameAdded {
		m.Frame = frameFlagValue
	}

	return nil, retAdded
}
