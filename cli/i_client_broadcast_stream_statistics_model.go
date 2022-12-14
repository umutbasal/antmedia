// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for IClientBroadcastStreamStatistics

// register flags to command
func registerModelIClientBroadcastStreamStatisticsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIClientBroadcastStreamStatisticsActiveSubscribers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIClientBroadcastStreamStatisticsBytesReceived(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIClientBroadcastStreamStatisticsCreationTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIClientBroadcastStreamStatisticsCurrentTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIClientBroadcastStreamStatisticsMaxSubscribers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIClientBroadcastStreamStatisticsPublishedName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIClientBroadcastStreamStatisticsSaveFilename(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIClientBroadcastStreamStatisticsTotalSubscribers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIClientBroadcastStreamStatisticsActiveSubscribers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	activeSubscribersDescription := ``

	var activeSubscribersFlagName string
	if cmdPrefix == "" {
		activeSubscribersFlagName = "activeSubscribers"
	} else {
		activeSubscribersFlagName = fmt.Sprintf("%v.activeSubscribers", cmdPrefix)
	}

	var activeSubscribersFlagDefault int32

	_ = cmd.PersistentFlags().Int32(activeSubscribersFlagName, activeSubscribersFlagDefault, activeSubscribersDescription)

	return nil
}

func registerIClientBroadcastStreamStatisticsBytesReceived(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	bytesReceivedDescription := ``

	var bytesReceivedFlagName string
	if cmdPrefix == "" {
		bytesReceivedFlagName = "bytesReceived"
	} else {
		bytesReceivedFlagName = fmt.Sprintf("%v.bytesReceived", cmdPrefix)
	}

	var bytesReceivedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(bytesReceivedFlagName, bytesReceivedFlagDefault, bytesReceivedDescription)

	return nil
}

func registerIClientBroadcastStreamStatisticsCreationTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	creationTimeDescription := ``

	var creationTimeFlagName string
	if cmdPrefix == "" {
		creationTimeFlagName = "creationTime"
	} else {
		creationTimeFlagName = fmt.Sprintf("%v.creationTime", cmdPrefix)
	}

	var creationTimeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(creationTimeFlagName, creationTimeFlagDefault, creationTimeDescription)

	return nil
}

func registerIClientBroadcastStreamStatisticsCurrentTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	currentTimestampDescription := ``

	var currentTimestampFlagName string
	if cmdPrefix == "" {
		currentTimestampFlagName = "currentTimestamp"
	} else {
		currentTimestampFlagName = fmt.Sprintf("%v.currentTimestamp", cmdPrefix)
	}

	var currentTimestampFlagDefault int32

	_ = cmd.PersistentFlags().Int32(currentTimestampFlagName, currentTimestampFlagDefault, currentTimestampDescription)

	return nil
}

func registerIClientBroadcastStreamStatisticsMaxSubscribers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxSubscribersDescription := ``

	var maxSubscribersFlagName string
	if cmdPrefix == "" {
		maxSubscribersFlagName = "maxSubscribers"
	} else {
		maxSubscribersFlagName = fmt.Sprintf("%v.maxSubscribers", cmdPrefix)
	}

	var maxSubscribersFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxSubscribersFlagName, maxSubscribersFlagDefault, maxSubscribersDescription)

	return nil
}

func registerIClientBroadcastStreamStatisticsPublishedName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	publishedNameDescription := ``

	var publishedNameFlagName string
	if cmdPrefix == "" {
		publishedNameFlagName = "publishedName"
	} else {
		publishedNameFlagName = fmt.Sprintf("%v.publishedName", cmdPrefix)
	}

	var publishedNameFlagDefault string

	_ = cmd.PersistentFlags().String(publishedNameFlagName, publishedNameFlagDefault, publishedNameDescription)

	return nil
}

func registerIClientBroadcastStreamStatisticsSaveFilename(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	saveFilenameDescription := ``

	var saveFilenameFlagName string
	if cmdPrefix == "" {
		saveFilenameFlagName = "saveFilename"
	} else {
		saveFilenameFlagName = fmt.Sprintf("%v.saveFilename", cmdPrefix)
	}

	var saveFilenameFlagDefault string

	_ = cmd.PersistentFlags().String(saveFilenameFlagName, saveFilenameFlagDefault, saveFilenameDescription)

	return nil
}

func registerIClientBroadcastStreamStatisticsTotalSubscribers(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalSubscribersDescription := ``

	var totalSubscribersFlagName string
	if cmdPrefix == "" {
		totalSubscribersFlagName = "totalSubscribers"
	} else {
		totalSubscribersFlagName = fmt.Sprintf("%v.totalSubscribers", cmdPrefix)
	}

	var totalSubscribersFlagDefault int32

	_ = cmd.PersistentFlags().Int32(totalSubscribersFlagName, totalSubscribersFlagDefault, totalSubscribersDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIClientBroadcastStreamStatisticsFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, activeSubscribersAdded := retrieveIClientBroadcastStreamStatisticsActiveSubscribersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || activeSubscribersAdded

	err, bytesReceivedAdded := retrieveIClientBroadcastStreamStatisticsBytesReceivedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bytesReceivedAdded

	err, creationTimeAdded := retrieveIClientBroadcastStreamStatisticsCreationTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || creationTimeAdded

	err, currentTimestampAdded := retrieveIClientBroadcastStreamStatisticsCurrentTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currentTimestampAdded

	err, maxSubscribersAdded := retrieveIClientBroadcastStreamStatisticsMaxSubscribersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxSubscribersAdded

	err, publishedNameAdded := retrieveIClientBroadcastStreamStatisticsPublishedNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || publishedNameAdded

	err, saveFilenameAdded := retrieveIClientBroadcastStreamStatisticsSaveFilenameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || saveFilenameAdded

	err, totalSubscribersAdded := retrieveIClientBroadcastStreamStatisticsTotalSubscribersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalSubscribersAdded

	return nil, retAdded
}

func retrieveIClientBroadcastStreamStatisticsActiveSubscribersFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	activeSubscribersFlagName := fmt.Sprintf("%v.activeSubscribers", cmdPrefix)
	if cmd.Flags().Changed(activeSubscribersFlagName) {

		var activeSubscribersFlagName string
		if cmdPrefix == "" {
			activeSubscribersFlagName = "activeSubscribers"
		} else {
			activeSubscribersFlagName = fmt.Sprintf("%v.activeSubscribers", cmdPrefix)
		}

		activeSubscribersFlagValue, err := cmd.Flags().GetInt32(activeSubscribersFlagName)
		if err != nil {
			return err, false
		}
		m.ActiveSubscribers = activeSubscribersFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIClientBroadcastStreamStatisticsBytesReceivedFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	bytesReceivedFlagName := fmt.Sprintf("%v.bytesReceived", cmdPrefix)
	if cmd.Flags().Changed(bytesReceivedFlagName) {

		var bytesReceivedFlagName string
		if cmdPrefix == "" {
			bytesReceivedFlagName = "bytesReceived"
		} else {
			bytesReceivedFlagName = fmt.Sprintf("%v.bytesReceived", cmdPrefix)
		}

		bytesReceivedFlagValue, err := cmd.Flags().GetInt64(bytesReceivedFlagName)
		if err != nil {
			return err, false
		}
		m.BytesReceived = bytesReceivedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIClientBroadcastStreamStatisticsCreationTimeFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	creationTimeFlagName := fmt.Sprintf("%v.creationTime", cmdPrefix)
	if cmd.Flags().Changed(creationTimeFlagName) {

		var creationTimeFlagName string
		if cmdPrefix == "" {
			creationTimeFlagName = "creationTime"
		} else {
			creationTimeFlagName = fmt.Sprintf("%v.creationTime", cmdPrefix)
		}

		creationTimeFlagValue, err := cmd.Flags().GetInt64(creationTimeFlagName)
		if err != nil {
			return err, false
		}
		m.CreationTime = creationTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIClientBroadcastStreamStatisticsCurrentTimestampFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	currentTimestampFlagName := fmt.Sprintf("%v.currentTimestamp", cmdPrefix)
	if cmd.Flags().Changed(currentTimestampFlagName) {

		var currentTimestampFlagName string
		if cmdPrefix == "" {
			currentTimestampFlagName = "currentTimestamp"
		} else {
			currentTimestampFlagName = fmt.Sprintf("%v.currentTimestamp", cmdPrefix)
		}

		currentTimestampFlagValue, err := cmd.Flags().GetInt32(currentTimestampFlagName)
		if err != nil {
			return err, false
		}
		m.CurrentTimestamp = currentTimestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIClientBroadcastStreamStatisticsMaxSubscribersFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxSubscribersFlagName := fmt.Sprintf("%v.maxSubscribers", cmdPrefix)
	if cmd.Flags().Changed(maxSubscribersFlagName) {

		var maxSubscribersFlagName string
		if cmdPrefix == "" {
			maxSubscribersFlagName = "maxSubscribers"
		} else {
			maxSubscribersFlagName = fmt.Sprintf("%v.maxSubscribers", cmdPrefix)
		}

		maxSubscribersFlagValue, err := cmd.Flags().GetInt32(maxSubscribersFlagName)
		if err != nil {
			return err, false
		}
		m.MaxSubscribers = maxSubscribersFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIClientBroadcastStreamStatisticsPublishedNameFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	publishedNameFlagName := fmt.Sprintf("%v.publishedName", cmdPrefix)
	if cmd.Flags().Changed(publishedNameFlagName) {

		var publishedNameFlagName string
		if cmdPrefix == "" {
			publishedNameFlagName = "publishedName"
		} else {
			publishedNameFlagName = fmt.Sprintf("%v.publishedName", cmdPrefix)
		}

		publishedNameFlagValue, err := cmd.Flags().GetString(publishedNameFlagName)
		if err != nil {
			return err, false
		}
		m.PublishedName = publishedNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIClientBroadcastStreamStatisticsSaveFilenameFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	saveFilenameFlagName := fmt.Sprintf("%v.saveFilename", cmdPrefix)
	if cmd.Flags().Changed(saveFilenameFlagName) {

		var saveFilenameFlagName string
		if cmdPrefix == "" {
			saveFilenameFlagName = "saveFilename"
		} else {
			saveFilenameFlagName = fmt.Sprintf("%v.saveFilename", cmdPrefix)
		}

		saveFilenameFlagValue, err := cmd.Flags().GetString(saveFilenameFlagName)
		if err != nil {
			return err, false
		}
		m.SaveFilename = saveFilenameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIClientBroadcastStreamStatisticsTotalSubscribersFlags(depth int, m *models.IClientBroadcastStreamStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalSubscribersFlagName := fmt.Sprintf("%v.totalSubscribers", cmdPrefix)
	if cmd.Flags().Changed(totalSubscribersFlagName) {

		var totalSubscribersFlagName string
		if cmdPrefix == "" {
			totalSubscribersFlagName = "totalSubscribers"
		} else {
			totalSubscribersFlagName = fmt.Sprintf("%v.totalSubscribers", cmdPrefix)
		}

		totalSubscribersFlagValue, err := cmd.Flags().GetInt32(totalSubscribersFlagName)
		if err != nil {
			return err, false
		}
		m.TotalSubscribers = totalSubscribersFlagValue

		retAdded = true
	}

	return nil, retAdded
}
