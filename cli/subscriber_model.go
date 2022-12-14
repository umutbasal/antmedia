// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for Subscriber

// register flags to command
func registerModelSubscriberFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSubscriberConnected(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSubscriberSubscriberID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSubscriberType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSubscriberConnected(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	connectedDescription := `is subscriber connected`

	var connectedFlagName string
	if cmdPrefix == "" {
		connectedFlagName = "connected"
	} else {
		connectedFlagName = fmt.Sprintf("%v.connected", cmdPrefix)
	}

	var connectedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(connectedFlagName, connectedFlagDefault, connectedDescription)

	return nil
}

func registerSubscriberSubscriberID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	subscriberIdDescription := `the subscriber id of the subscriber`

	var subscriberIdFlagName string
	if cmdPrefix == "" {
		subscriberIdFlagName = "subscriberId"
	} else {
		subscriberIdFlagName = fmt.Sprintf("%v.subscriberId", cmdPrefix)
	}

	var subscriberIdFlagDefault string

	_ = cmd.PersistentFlags().String(subscriberIdFlagName, subscriberIdFlagDefault, subscriberIdDescription)

	return nil
}

func registerSubscriberType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := ` type of subscriber (play or publish)`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSubscriberFlags(depth int, m *models.Subscriber, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, connectedAdded := retrieveSubscriberConnectedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectedAdded

	err, subscriberIdAdded := retrieveSubscriberSubscriberIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || subscriberIdAdded

	err, typeAdded := retrieveSubscriberTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveSubscriberConnectedFlags(depth int, m *models.Subscriber, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectedFlagName := fmt.Sprintf("%v.connected", cmdPrefix)
	if cmd.Flags().Changed(connectedFlagName) {

		var connectedFlagName string
		if cmdPrefix == "" {
			connectedFlagName = "connected"
		} else {
			connectedFlagName = fmt.Sprintf("%v.connected", cmdPrefix)
		}

		connectedFlagValue, err := cmd.Flags().GetBool(connectedFlagName)
		if err != nil {
			return err, false
		}
		m.Connected = connectedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSubscriberSubscriberIDFlags(depth int, m *models.Subscriber, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	subscriberIdFlagName := fmt.Sprintf("%v.subscriberId", cmdPrefix)
	if cmd.Flags().Changed(subscriberIdFlagName) {

		var subscriberIdFlagName string
		if cmdPrefix == "" {
			subscriberIdFlagName = "subscriberId"
		} else {
			subscriberIdFlagName = fmt.Sprintf("%v.subscriberId", cmdPrefix)
		}

		subscriberIdFlagValue, err := cmd.Flags().GetString(subscriberIdFlagName)
		if err != nil {
			return err, false
		}
		m.SubscriberID = subscriberIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSubscriberTypeFlags(depth int, m *models.Subscriber, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
