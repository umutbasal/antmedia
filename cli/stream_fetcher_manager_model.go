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

// Schema cli for StreamFetcherManager

// register flags to command
func registerModelStreamFetcherManagerFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStreamFetcherManagerDatastore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStreamFetcherManagerRestartStreamAutomatically(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStreamFetcherManagerStreamCheckerCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStreamFetcherManagerStreamCheckerInterval(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStreamFetcherManagerStreamFetcherList(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStreamFetcherManagerDatastore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var datastoreFlagName string
	if cmdPrefix == "" {
		datastoreFlagName = "datastore"
	} else {
		datastoreFlagName = fmt.Sprintf("%v.datastore", cmdPrefix)
	}

	if err := registerModelDataStoreFlags(depth+1, datastoreFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerStreamFetcherManagerRestartStreamAutomatically(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	restartStreamAutomaticallyDescription := ``

	var restartStreamAutomaticallyFlagName string
	if cmdPrefix == "" {
		restartStreamAutomaticallyFlagName = "restartStreamAutomatically"
	} else {
		restartStreamAutomaticallyFlagName = fmt.Sprintf("%v.restartStreamAutomatically", cmdPrefix)
	}

	var restartStreamAutomaticallyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(restartStreamAutomaticallyFlagName, restartStreamAutomaticallyFlagDefault, restartStreamAutomaticallyDescription)

	return nil
}

func registerStreamFetcherManagerStreamCheckerCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	streamCheckerCountDescription := ``

	var streamCheckerCountFlagName string
	if cmdPrefix == "" {
		streamCheckerCountFlagName = "streamCheckerCount"
	} else {
		streamCheckerCountFlagName = fmt.Sprintf("%v.streamCheckerCount", cmdPrefix)
	}

	var streamCheckerCountFlagDefault int32

	_ = cmd.PersistentFlags().Int32(streamCheckerCountFlagName, streamCheckerCountFlagDefault, streamCheckerCountDescription)

	return nil
}

func registerStreamFetcherManagerStreamCheckerInterval(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	streamCheckerIntervalDescription := ``

	var streamCheckerIntervalFlagName string
	if cmdPrefix == "" {
		streamCheckerIntervalFlagName = "streamCheckerInterval"
	} else {
		streamCheckerIntervalFlagName = fmt.Sprintf("%v.streamCheckerInterval", cmdPrefix)
	}

	var streamCheckerIntervalFlagDefault int32

	_ = cmd.PersistentFlags().Int32(streamCheckerIntervalFlagName, streamCheckerIntervalFlagDefault, streamCheckerIntervalDescription)

	return nil
}

func registerStreamFetcherManagerStreamFetcherList(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: streamFetcherList []*StreamFetcher array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStreamFetcherManagerFlags(depth int, m *models.StreamFetcherManager, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, datastoreAdded := retrieveStreamFetcherManagerDatastoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || datastoreAdded

	err, restartStreamAutomaticallyAdded := retrieveStreamFetcherManagerRestartStreamAutomaticallyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || restartStreamAutomaticallyAdded

	err, streamCheckerCountAdded := retrieveStreamFetcherManagerStreamCheckerCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamCheckerCountAdded

	err, streamCheckerIntervalAdded := retrieveStreamFetcherManagerStreamCheckerIntervalFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamCheckerIntervalAdded

	err, streamFetcherListAdded := retrieveStreamFetcherManagerStreamFetcherListFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamFetcherListAdded

	return nil, retAdded
}

func retrieveStreamFetcherManagerDatastoreFlags(depth int, m *models.StreamFetcherManager, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	datastoreFlagName := fmt.Sprintf("%v.datastore", cmdPrefix)
	if cmd.Flags().Changed(datastoreFlagName) {
		// info: complex object datastore DataStore is retrieved outside this Changed() block
	}
	datastoreFlagValue := m.Datastore
	if swag.IsZero(datastoreFlagValue) {
		datastoreFlagValue = &models.DataStore{}
	}

	err, datastoreAdded := retrieveModelDataStoreFlags(depth+1, datastoreFlagValue, datastoreFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || datastoreAdded
	if datastoreAdded {
		m.Datastore = datastoreFlagValue
	}

	return nil, retAdded
}

func retrieveStreamFetcherManagerRestartStreamAutomaticallyFlags(depth int, m *models.StreamFetcherManager, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	restartStreamAutomaticallyFlagName := fmt.Sprintf("%v.restartStreamAutomatically", cmdPrefix)
	if cmd.Flags().Changed(restartStreamAutomaticallyFlagName) {

		var restartStreamAutomaticallyFlagName string
		if cmdPrefix == "" {
			restartStreamAutomaticallyFlagName = "restartStreamAutomatically"
		} else {
			restartStreamAutomaticallyFlagName = fmt.Sprintf("%v.restartStreamAutomatically", cmdPrefix)
		}

		restartStreamAutomaticallyFlagValue, err := cmd.Flags().GetBool(restartStreamAutomaticallyFlagName)
		if err != nil {
			return err, false
		}
		m.RestartStreamAutomatically = restartStreamAutomaticallyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStreamFetcherManagerStreamCheckerCountFlags(depth int, m *models.StreamFetcherManager, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamCheckerCountFlagName := fmt.Sprintf("%v.streamCheckerCount", cmdPrefix)
	if cmd.Flags().Changed(streamCheckerCountFlagName) {

		var streamCheckerCountFlagName string
		if cmdPrefix == "" {
			streamCheckerCountFlagName = "streamCheckerCount"
		} else {
			streamCheckerCountFlagName = fmt.Sprintf("%v.streamCheckerCount", cmdPrefix)
		}

		streamCheckerCountFlagValue, err := cmd.Flags().GetInt32(streamCheckerCountFlagName)
		if err != nil {
			return err, false
		}
		m.StreamCheckerCount = streamCheckerCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStreamFetcherManagerStreamCheckerIntervalFlags(depth int, m *models.StreamFetcherManager, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamCheckerIntervalFlagName := fmt.Sprintf("%v.streamCheckerInterval", cmdPrefix)
	if cmd.Flags().Changed(streamCheckerIntervalFlagName) {

		var streamCheckerIntervalFlagName string
		if cmdPrefix == "" {
			streamCheckerIntervalFlagName = "streamCheckerInterval"
		} else {
			streamCheckerIntervalFlagName = fmt.Sprintf("%v.streamCheckerInterval", cmdPrefix)
		}

		streamCheckerIntervalFlagValue, err := cmd.Flags().GetInt32(streamCheckerIntervalFlagName)
		if err != nil {
			return err, false
		}
		m.StreamCheckerInterval = streamCheckerIntervalFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStreamFetcherManagerStreamFetcherListFlags(depth int, m *models.StreamFetcherManager, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamFetcherListFlagName := fmt.Sprintf("%v.streamFetcherList", cmdPrefix)
	if cmd.Flags().Changed(streamFetcherListFlagName) {
		// warning: streamFetcherList array type []*StreamFetcher is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
