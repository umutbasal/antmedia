// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for IClusterNotifier

// register flags to command
func registerModelIClusterNotifierFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIClusterNotifierClusterStore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIClusterNotifierClusterStore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var clusterStoreFlagName string
	if cmdPrefix == "" {
		clusterStoreFlagName = "clusterStore"
	} else {
		clusterStoreFlagName = fmt.Sprintf("%v.clusterStore", cmdPrefix)
	}

	if err := registerModelIClusterStoreFlags(depth+1, clusterStoreFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIClusterNotifierFlags(depth int, m *models.IClusterNotifier, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, clusterStoreAdded := retrieveIClusterNotifierClusterStoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterStoreAdded

	return nil, retAdded
}

func retrieveIClusterNotifierClusterStoreFlags(depth int, m *models.IClusterNotifier, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clusterStoreFlagName := fmt.Sprintf("%v.clusterStore", cmdPrefix)
	if cmd.Flags().Changed(clusterStoreFlagName) {
		// info: complex object clusterStore IClusterStore is retrieved outside this Changed() block
	}
	clusterStoreFlagValue := m.ClusterStore
	if swag.IsZero(clusterStoreFlagValue) {
		clusterStoreFlagValue = &models.IClusterStore{}
	}

	err, clusterStoreAdded := retrieveModelIClusterStoreFlags(depth+1, clusterStoreFlagValue, clusterStoreFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterStoreAdded
	if clusterStoreAdded {
		m.ClusterStore = clusterStoreFlagValue
	}

	return nil, retAdded
}
