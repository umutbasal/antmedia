// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for Result

// register flags to command
func registerModelResultFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerResultDataID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResultErrorID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResultMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerResultSuccess(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerResultDataID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dataIdDescription := `the id of the record if operation is about adding a record`

	var dataIdFlagName string
	if cmdPrefix == "" {
		dataIdFlagName = "dataId"
	} else {
		dataIdFlagName = fmt.Sprintf("%v.dataId", cmdPrefix)
	}

	var dataIdFlagDefault string

	_ = cmd.PersistentFlags().String(dataIdFlagName, dataIdFlagDefault, dataIdDescription)

	return nil
}

func registerResultErrorID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorIdDescription := `the id of errror of the operation result`

	var errorIdFlagName string
	if cmdPrefix == "" {
		errorIdFlagName = "errorId"
	} else {
		errorIdFlagName = fmt.Sprintf("%v.errorId", cmdPrefix)
	}

	var errorIdFlagDefault int32

	_ = cmd.PersistentFlags().Int32(errorIdFlagName, errorIdFlagDefault, errorIdDescription)

	return nil
}

func registerResultMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	messageDescription := `the message of the operation result`

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

func registerResultSuccess(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	successDescription := `the result of the operation`

	var successFlagName string
	if cmdPrefix == "" {
		successFlagName = "success"
	} else {
		successFlagName = fmt.Sprintf("%v.success", cmdPrefix)
	}

	var successFlagDefault bool

	_ = cmd.PersistentFlags().Bool(successFlagName, successFlagDefault, successDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelResultFlags(depth int, m *models.Result, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dataIdAdded := retrieveResultDataIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataIdAdded

	err, errorIdAdded := retrieveResultErrorIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorIdAdded

	err, messageAdded := retrieveResultMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || messageAdded

	err, successAdded := retrieveResultSuccessFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || successAdded

	return nil, retAdded
}

func retrieveResultDataIDFlags(depth int, m *models.Result, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataIdFlagName := fmt.Sprintf("%v.dataId", cmdPrefix)
	if cmd.Flags().Changed(dataIdFlagName) {

		var dataIdFlagName string
		if cmdPrefix == "" {
			dataIdFlagName = "dataId"
		} else {
			dataIdFlagName = fmt.Sprintf("%v.dataId", cmdPrefix)
		}

		dataIdFlagValue, err := cmd.Flags().GetString(dataIdFlagName)
		if err != nil {
			return err, false
		}
		m.DataID = dataIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResultErrorIDFlags(depth int, m *models.Result, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorIdFlagName := fmt.Sprintf("%v.errorId", cmdPrefix)
	if cmd.Flags().Changed(errorIdFlagName) {

		var errorIdFlagName string
		if cmdPrefix == "" {
			errorIdFlagName = "errorId"
		} else {
			errorIdFlagName = fmt.Sprintf("%v.errorId", cmdPrefix)
		}

		errorIdFlagValue, err := cmd.Flags().GetInt32(errorIdFlagName)
		if err != nil {
			return err, false
		}
		m.ErrorID = errorIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveResultMessageFlags(depth int, m *models.Result, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveResultSuccessFlags(depth int, m *models.Result, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	successFlagName := fmt.Sprintf("%v.success", cmdPrefix)
	if cmd.Flags().Changed(successFlagName) {

		var successFlagName string
		if cmdPrefix == "" {
			successFlagName = "success"
		} else {
			successFlagName = fmt.Sprintf("%v.success", cmdPrefix)
		}

		successFlagValue, err := cmd.Flags().GetBool(successFlagName)
		if err != nil {
			return err, false
		}
		m.Success = successFlagValue

		retAdded = true
	}

	return nil, retAdded
}
