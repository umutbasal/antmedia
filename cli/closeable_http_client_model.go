// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for CloseableHTTPClient

// register flags to command
func registerModelCloseableHTTPClientFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCloseableHTTPClientConnectionManager(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCloseableHTTPClientParams(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCloseableHTTPClientConnectionManager(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var connectionManagerFlagName string
	if cmdPrefix == "" {
		connectionManagerFlagName = "connectionManager"
	} else {
		connectionManagerFlagName = fmt.Sprintf("%v.connectionManager", cmdPrefix)
	}

	if err := registerModelClientConnectionManagerFlags(depth+1, connectionManagerFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerCloseableHTTPClientParams(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: params HTTPParams map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelCloseableHTTPClientFlags(depth int, m *models.CloseableHTTPClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, connectionManagerAdded := retrieveCloseableHTTPClientConnectionManagerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionManagerAdded

	err, paramsAdded := retrieveCloseableHTTPClientParamsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || paramsAdded

	return nil, retAdded
}

func retrieveCloseableHTTPClientConnectionManagerFlags(depth int, m *models.CloseableHTTPClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectionManagerFlagName := fmt.Sprintf("%v.connectionManager", cmdPrefix)
	if cmd.Flags().Changed(connectionManagerFlagName) {
		// info: complex object connectionManager ClientConnectionManager is retrieved outside this Changed() block
	}
	connectionManagerFlagValue := m.ConnectionManager
	if swag.IsZero(connectionManagerFlagValue) {
		connectionManagerFlagValue = &models.ClientConnectionManager{}
	}

	err, connectionManagerAdded := retrieveModelClientConnectionManagerFlags(depth+1, connectionManagerFlagValue, connectionManagerFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionManagerAdded
	if connectionManagerAdded {
		m.ConnectionManager = connectionManagerFlagValue
	}

	return nil, retAdded
}

func retrieveCloseableHTTPClientParamsFlags(depth int, m *models.CloseableHTTPClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	paramsFlagName := fmt.Sprintf("%v.params", cmdPrefix)
	if cmd.Flags().Changed(paramsFlagName) {
		// warning: params map type HTTPParams is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
