// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for IClusterStore

// register flags to command
func registerModelIClusterStoreFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIClusterStoreAllSettings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIClusterStoreNodeCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIClusterStoreAllSettings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: allSettings []*AppSettings array type is not supported by go-swagger cli yet

	return nil
}

func registerIClusterStoreNodeCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodeCountDescription := ``

	var nodeCountFlagName string
	if cmdPrefix == "" {
		nodeCountFlagName = "nodeCount"
	} else {
		nodeCountFlagName = fmt.Sprintf("%v.nodeCount", cmdPrefix)
	}

	var nodeCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(nodeCountFlagName, nodeCountFlagDefault, nodeCountDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIClusterStoreFlags(depth int, m *models.IClusterStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, allSettingsAdded := retrieveIClusterStoreAllSettingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || allSettingsAdded

	err, nodeCountAdded := retrieveIClusterStoreNodeCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodeCountAdded

	return nil, retAdded
}

func retrieveIClusterStoreAllSettingsFlags(depth int, m *models.IClusterStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	allSettingsFlagName := fmt.Sprintf("%v.allSettings", cmdPrefix)
	if cmd.Flags().Changed(allSettingsFlagName) {
		// warning: allSettings array type []*AppSettings is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIClusterStoreNodeCountFlags(depth int, m *models.IClusterStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nodeCountFlagName := fmt.Sprintf("%v.nodeCount", cmdPrefix)
	if cmd.Flags().Changed(nodeCountFlagName) {

		var nodeCountFlagName string
		if cmdPrefix == "" {
			nodeCountFlagName = "nodeCount"
		} else {
			nodeCountFlagName = fmt.Sprintf("%v.nodeCount", cmdPrefix)
		}

		nodeCountFlagValue, err := cmd.Flags().GetInt64(nodeCountFlagName)
		if err != nil {
			return err, false
		}
		m.NodeCount = nodeCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}