// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for DataStore

// register flags to command
func registerModelDataStoreFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerDataStoreActiveBroadcastCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataStoreAvailable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataStoreBroadcastCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataStoreExternalStreamsList(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataStoreTotalBroadcastNumber(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataStoreTotalVodNumber(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataStoreTotalWebRTCViewersCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerDataStoreWriteStatsToDatastore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerDataStoreActiveBroadcastCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	activeBroadcastCountDescription := ``

	var activeBroadcastCountFlagName string
	if cmdPrefix == "" {
		activeBroadcastCountFlagName = "activeBroadcastCount"
	} else {
		activeBroadcastCountFlagName = fmt.Sprintf("%v.activeBroadcastCount", cmdPrefix)
	}

	var activeBroadcastCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(activeBroadcastCountFlagName, activeBroadcastCountFlagDefault, activeBroadcastCountDescription)

	return nil
}

func registerDataStoreAvailable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	availableDescription := ``

	var availableFlagName string
	if cmdPrefix == "" {
		availableFlagName = "available"
	} else {
		availableFlagName = fmt.Sprintf("%v.available", cmdPrefix)
	}

	var availableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(availableFlagName, availableFlagDefault, availableDescription)

	return nil
}

func registerDataStoreBroadcastCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	broadcastCountDescription := ``

	var broadcastCountFlagName string
	if cmdPrefix == "" {
		broadcastCountFlagName = "broadcastCount"
	} else {
		broadcastCountFlagName = fmt.Sprintf("%v.broadcastCount", cmdPrefix)
	}

	var broadcastCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(broadcastCountFlagName, broadcastCountFlagDefault, broadcastCountDescription)

	return nil
}

func registerDataStoreExternalStreamsList(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: externalStreamsList []*Broadcast array type is not supported by go-swagger cli yet

	return nil
}

func registerDataStoreTotalBroadcastNumber(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalBroadcastNumberDescription := ``

	var totalBroadcastNumberFlagName string
	if cmdPrefix == "" {
		totalBroadcastNumberFlagName = "totalBroadcastNumber"
	} else {
		totalBroadcastNumberFlagName = fmt.Sprintf("%v.totalBroadcastNumber", cmdPrefix)
	}

	var totalBroadcastNumberFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalBroadcastNumberFlagName, totalBroadcastNumberFlagDefault, totalBroadcastNumberDescription)

	return nil
}

func registerDataStoreTotalVodNumber(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalVodNumberDescription := ``

	var totalVodNumberFlagName string
	if cmdPrefix == "" {
		totalVodNumberFlagName = "totalVodNumber"
	} else {
		totalVodNumberFlagName = fmt.Sprintf("%v.totalVodNumber", cmdPrefix)
	}

	var totalVodNumberFlagDefault int64

	_ = cmd.PersistentFlags().Int64(totalVodNumberFlagName, totalVodNumberFlagDefault, totalVodNumberDescription)

	return nil
}

func registerDataStoreTotalWebRTCViewersCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalWebRTCViewersCountDescription := ``

	var totalWebRTCViewersCountFlagName string
	if cmdPrefix == "" {
		totalWebRTCViewersCountFlagName = "totalWebRTCViewersCount"
	} else {
		totalWebRTCViewersCountFlagName = fmt.Sprintf("%v.totalWebRTCViewersCount", cmdPrefix)
	}

	var totalWebRTCViewersCountFlagDefault int32

	_ = cmd.PersistentFlags().Int32(totalWebRTCViewersCountFlagName, totalWebRTCViewersCountFlagDefault, totalWebRTCViewersCountDescription)

	return nil
}

func registerDataStoreWriteStatsToDatastore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	writeStatsToDatastoreDescription := ``

	var writeStatsToDatastoreFlagName string
	if cmdPrefix == "" {
		writeStatsToDatastoreFlagName = "writeStatsToDatastore"
	} else {
		writeStatsToDatastoreFlagName = fmt.Sprintf("%v.writeStatsToDatastore", cmdPrefix)
	}

	var writeStatsToDatastoreFlagDefault bool

	_ = cmd.PersistentFlags().Bool(writeStatsToDatastoreFlagName, writeStatsToDatastoreFlagDefault, writeStatsToDatastoreDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelDataStoreFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, activeBroadcastCountAdded := retrieveDataStoreActiveBroadcastCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || activeBroadcastCountAdded

	err, availableAdded := retrieveDataStoreAvailableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || availableAdded

	err, broadcastCountAdded := retrieveDataStoreBroadcastCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || broadcastCountAdded

	err, externalStreamsListAdded := retrieveDataStoreExternalStreamsListFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || externalStreamsListAdded

	err, totalBroadcastNumberAdded := retrieveDataStoreTotalBroadcastNumberFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalBroadcastNumberAdded

	err, totalVodNumberAdded := retrieveDataStoreTotalVodNumberFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalVodNumberAdded

	err, totalWebRTCViewersCountAdded := retrieveDataStoreTotalWebRTCViewersCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalWebRTCViewersCountAdded

	err, writeStatsToDatastoreAdded := retrieveDataStoreWriteStatsToDatastoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || writeStatsToDatastoreAdded

	return nil, retAdded
}

func retrieveDataStoreActiveBroadcastCountFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	activeBroadcastCountFlagName := fmt.Sprintf("%v.activeBroadcastCount", cmdPrefix)
	if cmd.Flags().Changed(activeBroadcastCountFlagName) {

		var activeBroadcastCountFlagName string
		if cmdPrefix == "" {
			activeBroadcastCountFlagName = "activeBroadcastCount"
		} else {
			activeBroadcastCountFlagName = fmt.Sprintf("%v.activeBroadcastCount", cmdPrefix)
		}

		activeBroadcastCountFlagValue, err := cmd.Flags().GetInt64(activeBroadcastCountFlagName)
		if err != nil {
			return err, false
		}
		m.ActiveBroadcastCount = activeBroadcastCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataStoreAvailableFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	availableFlagName := fmt.Sprintf("%v.available", cmdPrefix)
	if cmd.Flags().Changed(availableFlagName) {

		var availableFlagName string
		if cmdPrefix == "" {
			availableFlagName = "available"
		} else {
			availableFlagName = fmt.Sprintf("%v.available", cmdPrefix)
		}

		availableFlagValue, err := cmd.Flags().GetBool(availableFlagName)
		if err != nil {
			return err, false
		}
		m.Available = availableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataStoreBroadcastCountFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	broadcastCountFlagName := fmt.Sprintf("%v.broadcastCount", cmdPrefix)
	if cmd.Flags().Changed(broadcastCountFlagName) {

		var broadcastCountFlagName string
		if cmdPrefix == "" {
			broadcastCountFlagName = "broadcastCount"
		} else {
			broadcastCountFlagName = fmt.Sprintf("%v.broadcastCount", cmdPrefix)
		}

		broadcastCountFlagValue, err := cmd.Flags().GetInt64(broadcastCountFlagName)
		if err != nil {
			return err, false
		}
		m.BroadcastCount = broadcastCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataStoreExternalStreamsListFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	externalStreamsListFlagName := fmt.Sprintf("%v.externalStreamsList", cmdPrefix)
	if cmd.Flags().Changed(externalStreamsListFlagName) {
		// warning: externalStreamsList array type []*Broadcast is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveDataStoreTotalBroadcastNumberFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalBroadcastNumberFlagName := fmt.Sprintf("%v.totalBroadcastNumber", cmdPrefix)
	if cmd.Flags().Changed(totalBroadcastNumberFlagName) {

		var totalBroadcastNumberFlagName string
		if cmdPrefix == "" {
			totalBroadcastNumberFlagName = "totalBroadcastNumber"
		} else {
			totalBroadcastNumberFlagName = fmt.Sprintf("%v.totalBroadcastNumber", cmdPrefix)
		}

		totalBroadcastNumberFlagValue, err := cmd.Flags().GetInt64(totalBroadcastNumberFlagName)
		if err != nil {
			return err, false
		}
		m.TotalBroadcastNumber = totalBroadcastNumberFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataStoreTotalVodNumberFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalVodNumberFlagName := fmt.Sprintf("%v.totalVodNumber", cmdPrefix)
	if cmd.Flags().Changed(totalVodNumberFlagName) {

		var totalVodNumberFlagName string
		if cmdPrefix == "" {
			totalVodNumberFlagName = "totalVodNumber"
		} else {
			totalVodNumberFlagName = fmt.Sprintf("%v.totalVodNumber", cmdPrefix)
		}

		totalVodNumberFlagValue, err := cmd.Flags().GetInt64(totalVodNumberFlagName)
		if err != nil {
			return err, false
		}
		m.TotalVodNumber = totalVodNumberFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataStoreTotalWebRTCViewersCountFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalWebRTCViewersCountFlagName := fmt.Sprintf("%v.totalWebRTCViewersCount", cmdPrefix)
	if cmd.Flags().Changed(totalWebRTCViewersCountFlagName) {

		var totalWebRTCViewersCountFlagName string
		if cmdPrefix == "" {
			totalWebRTCViewersCountFlagName = "totalWebRTCViewersCount"
		} else {
			totalWebRTCViewersCountFlagName = fmt.Sprintf("%v.totalWebRTCViewersCount", cmdPrefix)
		}

		totalWebRTCViewersCountFlagValue, err := cmd.Flags().GetInt32(totalWebRTCViewersCountFlagName)
		if err != nil {
			return err, false
		}
		m.TotalWebRTCViewersCount = totalWebRTCViewersCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveDataStoreWriteStatsToDatastoreFlags(depth int, m *models.DataStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	writeStatsToDatastoreFlagName := fmt.Sprintf("%v.writeStatsToDatastore", cmdPrefix)
	if cmd.Flags().Changed(writeStatsToDatastoreFlagName) {

		var writeStatsToDatastoreFlagName string
		if cmdPrefix == "" {
			writeStatsToDatastoreFlagName = "writeStatsToDatastore"
		} else {
			writeStatsToDatastoreFlagName = fmt.Sprintf("%v.writeStatsToDatastore", cmdPrefix)
		}

		writeStatsToDatastoreFlagValue, err := cmd.Flags().GetBool(writeStatsToDatastoreFlagName)
		if err != nil {
			return err, false
		}
		m.WriteStatsToDatastore = writeStatsToDatastoreFlagValue

		retAdded = true
	}

	return nil, retAdded
}
