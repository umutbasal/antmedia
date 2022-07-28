// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for ClientConnectionManager

// register flags to command
func registerModelClientConnectionManagerFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerClientConnectionManagerSchemeRegistry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerClientConnectionManagerSchemeRegistry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var schemeRegistryFlagName string
	if cmdPrefix == "" {
		schemeRegistryFlagName = "schemeRegistry"
	} else {
		schemeRegistryFlagName = fmt.Sprintf("%v.schemeRegistry", cmdPrefix)
	}

	if err := registerModelSchemeRegistryFlags(depth+1, schemeRegistryFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelClientConnectionManagerFlags(depth int, m *models.ClientConnectionManager, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, schemeRegistryAdded := retrieveClientConnectionManagerSchemeRegistryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || schemeRegistryAdded

	return nil, retAdded
}

func retrieveClientConnectionManagerSchemeRegistryFlags(depth int, m *models.ClientConnectionManager, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	schemeRegistryFlagName := fmt.Sprintf("%v.schemeRegistry", cmdPrefix)
	if cmd.Flags().Changed(schemeRegistryFlagName) {
		// info: complex object schemeRegistry SchemeRegistry is retrieved outside this Changed() block
	}
	schemeRegistryFlagValue := m.SchemeRegistry
	if swag.IsZero(schemeRegistryFlagValue) {
		schemeRegistryFlagValue = &models.SchemeRegistry{}
	}

	err, schemeRegistryAdded := retrieveModelSchemeRegistryFlags(depth+1, schemeRegistryFlagValue, schemeRegistryFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || schemeRegistryAdded
	if schemeRegistryAdded {
		m.SchemeRegistry = schemeRegistryFlagValue
	}

	return nil, retAdded
}