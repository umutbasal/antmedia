// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for Environment

// register flags to command
func registerModelEnvironmentFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEnvironmentActiveProfiles(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEnvironmentDefaultProfiles(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEnvironmentActiveProfiles(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: activeProfiles []string array type is not supported by go-swagger cli yet

	return nil
}

func registerEnvironmentDefaultProfiles(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: defaultProfiles []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEnvironmentFlags(depth int, m *models.Environment, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, activeProfilesAdded := retrieveEnvironmentActiveProfilesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || activeProfilesAdded

	err, defaultProfilesAdded := retrieveEnvironmentDefaultProfilesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultProfilesAdded

	return nil, retAdded
}

func retrieveEnvironmentActiveProfilesFlags(depth int, m *models.Environment, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	activeProfilesFlagName := fmt.Sprintf("%v.activeProfiles", cmdPrefix)
	if cmd.Flags().Changed(activeProfilesFlagName) {
		// warning: activeProfiles array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveEnvironmentDefaultProfilesFlags(depth int, m *models.Environment, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultProfilesFlagName := fmt.Sprintf("%v.defaultProfiles", cmdPrefix)
	if cmd.Flags().Changed(defaultProfilesFlagName) {
		// warning: defaultProfiles array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}