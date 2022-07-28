// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for PacketTime

// register flags to command
func registerModelPacketTimeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerPacketTimePacketTimeMs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerPacketTimeSystemTimeMs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerPacketTimePacketTimeMs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	packetTimeMsDescription := ``

	var packetTimeMsFlagName string
	if cmdPrefix == "" {
		packetTimeMsFlagName = "packetTimeMs"
	} else {
		packetTimeMsFlagName = fmt.Sprintf("%v.packetTimeMs", cmdPrefix)
	}

	var packetTimeMsFlagDefault int64

	_ = cmd.PersistentFlags().Int64(packetTimeMsFlagName, packetTimeMsFlagDefault, packetTimeMsDescription)

	return nil
}

func registerPacketTimeSystemTimeMs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	systemTimeMsDescription := ``

	var systemTimeMsFlagName string
	if cmdPrefix == "" {
		systemTimeMsFlagName = "systemTimeMs"
	} else {
		systemTimeMsFlagName = fmt.Sprintf("%v.systemTimeMs", cmdPrefix)
	}

	var systemTimeMsFlagDefault int64

	_ = cmd.PersistentFlags().Int64(systemTimeMsFlagName, systemTimeMsFlagDefault, systemTimeMsDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelPacketTimeFlags(depth int, m *models.PacketTime, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, packetTimeMsAdded := retrievePacketTimePacketTimeMsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || packetTimeMsAdded

	err, systemTimeMsAdded := retrievePacketTimeSystemTimeMsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || systemTimeMsAdded

	return nil, retAdded
}

func retrievePacketTimePacketTimeMsFlags(depth int, m *models.PacketTime, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	packetTimeMsFlagName := fmt.Sprintf("%v.packetTimeMs", cmdPrefix)
	if cmd.Flags().Changed(packetTimeMsFlagName) {

		var packetTimeMsFlagName string
		if cmdPrefix == "" {
			packetTimeMsFlagName = "packetTimeMs"
		} else {
			packetTimeMsFlagName = fmt.Sprintf("%v.packetTimeMs", cmdPrefix)
		}

		packetTimeMsFlagValue, err := cmd.Flags().GetInt64(packetTimeMsFlagName)
		if err != nil {
			return err, false
		}
		m.PacketTimeMs = packetTimeMsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrievePacketTimeSystemTimeMsFlags(depth int, m *models.PacketTime, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	systemTimeMsFlagName := fmt.Sprintf("%v.systemTimeMs", cmdPrefix)
	if cmd.Flags().Changed(systemTimeMsFlagName) {

		var systemTimeMsFlagName string
		if cmdPrefix == "" {
			systemTimeMsFlagName = "systemTimeMs"
		} else {
			systemTimeMsFlagName = fmt.Sprintf("%v.systemTimeMs", cmdPrefix)
		}

		systemTimeMsFlagValue, err := cmd.Flags().GetInt64(systemTimeMsFlagName)
		if err != nil {
			return err, false
		}
		m.SystemTimeMs = systemTimeMsFlagValue

		retAdded = true
	}

	return nil, retAdded
}
