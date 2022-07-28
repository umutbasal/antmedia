// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for Exception

// register flags to command
func registerModelExceptionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerExceptionCause(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExceptionLocalizedMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExceptionMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExceptionStackTrace(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExceptionSuppressed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerExceptionCause(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var causeFlagName string
	if cmdPrefix == "" {
		causeFlagName = "cause"
	} else {
		causeFlagName = fmt.Sprintf("%v.cause", cmdPrefix)
	}

	if err := registerModelThrowableFlags(depth+1, causeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerExceptionLocalizedMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	localizedMessageDescription := ``

	var localizedMessageFlagName string
	if cmdPrefix == "" {
		localizedMessageFlagName = "localizedMessage"
	} else {
		localizedMessageFlagName = fmt.Sprintf("%v.localizedMessage", cmdPrefix)
	}

	var localizedMessageFlagDefault string

	_ = cmd.PersistentFlags().String(localizedMessageFlagName, localizedMessageFlagDefault, localizedMessageDescription)

	return nil
}

func registerExceptionMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := ``

	var messageFlagName string
	if cmdPrefix == "" {
		messageFlagName = "message"
	} else {
		messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
	}

	var messageFlagDefault string

	_ = cmd.PersistentFlags().String(messageFlagName, messageFlagDefault, messageDescription)

	return nil
}

func registerExceptionStackTrace(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: stackTrace []*StackTraceElement array type is not supported by go-swagger cli yet

	return nil
}

func registerExceptionSuppressed(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: suppressed []*Throwable array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelExceptionFlags(depth int, m *models.Exception, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, causeAdded := retrieveExceptionCauseFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || causeAdded

	err, localizedMessageAdded := retrieveExceptionLocalizedMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || localizedMessageAdded

	err, messageAdded := retrieveExceptionMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, stackTraceAdded := retrieveExceptionStackTraceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stackTraceAdded

	err, suppressedAdded := retrieveExceptionSuppressedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || suppressedAdded

	return nil, retAdded
}

func retrieveExceptionCauseFlags(depth int, m *models.Exception, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	causeFlagName := fmt.Sprintf("%v.cause", cmdPrefix)
	if cmd.Flags().Changed(causeFlagName) {
		// info: complex object cause Throwable is retrieved outside this Changed() block
	}
	causeFlagValue := m.Cause
	if swag.IsZero(causeFlagValue) {
		causeFlagValue = &models.Throwable{}
	}

	err, causeAdded := retrieveModelThrowableFlags(depth+1, causeFlagValue, causeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || causeAdded
	if causeAdded {
		m.Cause = causeFlagValue
	}

	return nil, retAdded
}

func retrieveExceptionLocalizedMessageFlags(depth int, m *models.Exception, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	localizedMessageFlagName := fmt.Sprintf("%v.localizedMessage", cmdPrefix)
	if cmd.Flags().Changed(localizedMessageFlagName) {

		var localizedMessageFlagName string
		if cmdPrefix == "" {
			localizedMessageFlagName = "localizedMessage"
		} else {
			localizedMessageFlagName = fmt.Sprintf("%v.localizedMessage", cmdPrefix)
		}

		localizedMessageFlagValue, err := cmd.Flags().GetString(localizedMessageFlagName)
		if err != nil {
			return err, false
		}
		m.LocalizedMessage = localizedMessageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExceptionMessageFlags(depth int, m *models.Exception, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	messageFlagName := fmt.Sprintf("%v.message", cmdPrefix)
	if cmd.Flags().Changed(messageFlagName) {

		var messageFlagName string
		if cmdPrefix == "" {
			messageFlagName = "message"
		} else {
			messageFlagName = fmt.Sprintf("%v.message", cmdPrefix)
		}

		messageFlagValue, err := cmd.Flags().GetString(messageFlagName)
		if err != nil {
			return err, false
		}
		m.Message = messageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExceptionStackTraceFlags(depth int, m *models.Exception, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stackTraceFlagName := fmt.Sprintf("%v.stackTrace", cmdPrefix)
	if cmd.Flags().Changed(stackTraceFlagName) {
		// warning: stackTrace array type []*StackTraceElement is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveExceptionSuppressedFlags(depth int, m *models.Exception, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	suppressedFlagName := fmt.Sprintf("%v.suppressed", cmdPrefix)
	if cmd.Flags().Changed(suppressedFlagName) {
		// warning: suppressed array type []*Throwable is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
