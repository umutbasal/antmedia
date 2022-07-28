// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for ClientBroadcastStream

// register flags to command
func registerModelClientBroadcastStreamFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerClientBroadcastStreamAbsoluteStartTimeMs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamActiveSubscribers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamAutomaticRecording(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamBroadcastStreamPublishName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamBytesReceived(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamClientBufferDuration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamCodecInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamConnection(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamCreationTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamCurrentTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamMaxSubscribers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamMetaData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamParameters(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamProvider(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamPublishedName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamRecording(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamSaveFilename(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamScope(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamState(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamStatistics(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamStreamID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamStreamListeners(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClientBroadcastStreamTotalSubscribers(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerClientBroadcastStreamAbsoluteStartTimeMs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	absoluteStartTimeMsDescription := ``

	var absoluteStartTimeMsFlagName string
	if cmdPrefix == "" {
		absoluteStartTimeMsFlagName = "absoluteStartTimeMs"
	} else {
		absoluteStartTimeMsFlagName = fmt.Sprintf("%v.absoluteStartTimeMs", cmdPrefix)
	}

	var absoluteStartTimeMsFlagDefault int64

	_ = cmd.PersistentFlags().Int64(absoluteStartTimeMsFlagName, absoluteStartTimeMsFlagDefault, absoluteStartTimeMsDescription)

	return nil
}

func registerClientBroadcastStreamActiveSubscribers(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerClientBroadcastStreamAutomaticRecording(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	automaticRecordingDescription := ``

	var automaticRecordingFlagName string
	if cmdPrefix == "" {
		automaticRecordingFlagName = "automaticRecording"
	} else {
		automaticRecordingFlagName = fmt.Sprintf("%v.automaticRecording", cmdPrefix)
	}

	var automaticRecordingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(automaticRecordingFlagName, automaticRecordingFlagDefault, automaticRecordingDescription)

	return nil
}

func registerClientBroadcastStreamBroadcastStreamPublishName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	broadcastStreamPublishNameDescription := ``

	var broadcastStreamPublishNameFlagName string
	if cmdPrefix == "" {
		broadcastStreamPublishNameFlagName = "broadcastStreamPublishName"
	} else {
		broadcastStreamPublishNameFlagName = fmt.Sprintf("%v.broadcastStreamPublishName", cmdPrefix)
	}

	var broadcastStreamPublishNameFlagDefault string

	_ = cmd.PersistentFlags().String(broadcastStreamPublishNameFlagName, broadcastStreamPublishNameFlagDefault, broadcastStreamPublishNameDescription)

	return nil
}

func registerClientBroadcastStreamBytesReceived(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerClientBroadcastStreamClientBufferDuration(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	clientBufferDurationDescription := ``

	var clientBufferDurationFlagName string
	if cmdPrefix == "" {
		clientBufferDurationFlagName = "clientBufferDuration"
	} else {
		clientBufferDurationFlagName = fmt.Sprintf("%v.clientBufferDuration", cmdPrefix)
	}

	var clientBufferDurationFlagDefault int32

	_ = cmd.PersistentFlags().Int32(clientBufferDurationFlagName, clientBufferDurationFlagDefault, clientBufferDurationDescription)

	return nil
}

func registerClientBroadcastStreamCodecInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var codecInfoFlagName string
	if cmdPrefix == "" {
		codecInfoFlagName = "codecInfo"
	} else {
		codecInfoFlagName = fmt.Sprintf("%v.codecInfo", cmdPrefix)
	}

	if err := registerModelIStreamCodecInfoFlags(depth+1, codecInfoFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClientBroadcastStreamConnection(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var connectionFlagName string
	if cmdPrefix == "" {
		connectionFlagName = "connection"
	} else {
		connectionFlagName = fmt.Sprintf("%v.connection", cmdPrefix)
	}

	if err := registerModelIStreamCapableConnectionFlags(depth+1, connectionFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClientBroadcastStreamCreationTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerClientBroadcastStreamCurrentTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerClientBroadcastStreamMaxSubscribers(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerClientBroadcastStreamMetaData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var metaDataFlagName string
	if cmdPrefix == "" {
		metaDataFlagName = "metaData"
	} else {
		metaDataFlagName = fmt.Sprintf("%v.metaData", cmdPrefix)
	}

	if err := registerModelNotifyFlags(depth+1, metaDataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClientBroadcastStreamName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerClientBroadcastStreamParameters(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: parameters map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerClientBroadcastStreamProvider(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: provider IProvider map type is not supported by go-swagger cli yet

	return nil
}

func registerClientBroadcastStreamPublishedName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerClientBroadcastStreamRecording(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	recordingDescription := ``

	var recordingFlagName string
	if cmdPrefix == "" {
		recordingFlagName = "recording"
	} else {
		recordingFlagName = fmt.Sprintf("%v.recording", cmdPrefix)
	}

	var recordingFlagDefault bool

	_ = cmd.PersistentFlags().Bool(recordingFlagName, recordingFlagDefault, recordingDescription)

	return nil
}

func registerClientBroadcastStreamSaveFilename(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerClientBroadcastStreamScope(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var scopeFlagName string
	if cmdPrefix == "" {
		scopeFlagName = "scope"
	} else {
		scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
	}

	if err := registerModelIScopeFlags(depth+1, scopeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClientBroadcastStreamState(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	stateDescription := `Enum: ["INIT","UNINIT","OPEN","CLOSED","STARTED","STOPPED","PLAYING","PAUSED","RESUMED","END","SEEK"]. `

	var stateFlagName string
	if cmdPrefix == "" {
		stateFlagName = "state"
	} else {
		stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
	}

	var stateFlagDefault string

	_ = cmd.PersistentFlags().String(stateFlagName, stateFlagDefault, stateDescription)

	if err := cmd.RegisterFlagCompletionFunc(stateFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["INIT","UNINIT","OPEN","CLOSED","STARTED","STOPPED","PLAYING","PAUSED","RESUMED","END","SEEK"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerClientBroadcastStreamStatistics(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var statisticsFlagName string
	if cmdPrefix == "" {
		statisticsFlagName = "statistics"
	} else {
		statisticsFlagName = fmt.Sprintf("%v.statistics", cmdPrefix)
	}

	if err := registerModelIClientBroadcastStreamStatisticsFlags(depth+1, statisticsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClientBroadcastStreamStreamID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: streamId Number map type is not supported by go-swagger cli yet

	return nil
}

func registerClientBroadcastStreamStreamListeners(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: streamListeners []IStreamListener array type is not supported by go-swagger cli yet

	return nil
}

func registerClientBroadcastStreamTotalSubscribers(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelClientBroadcastStreamFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, absoluteStartTimeMsAdded := retrieveClientBroadcastStreamAbsoluteStartTimeMsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || absoluteStartTimeMsAdded

	err, activeSubscribersAdded := retrieveClientBroadcastStreamActiveSubscribersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || activeSubscribersAdded

	err, automaticRecordingAdded := retrieveClientBroadcastStreamAutomaticRecordingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || automaticRecordingAdded

	err, broadcastStreamPublishNameAdded := retrieveClientBroadcastStreamBroadcastStreamPublishNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || broadcastStreamPublishNameAdded

	err, bytesReceivedAdded := retrieveClientBroadcastStreamBytesReceivedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || bytesReceivedAdded

	err, clientBufferDurationAdded := retrieveClientBroadcastStreamClientBufferDurationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clientBufferDurationAdded

	err, codecInfoAdded := retrieveClientBroadcastStreamCodecInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || codecInfoAdded

	err, connectionAdded := retrieveClientBroadcastStreamConnectionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionAdded

	err, creationTimeAdded := retrieveClientBroadcastStreamCreationTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || creationTimeAdded

	err, currentTimestampAdded := retrieveClientBroadcastStreamCurrentTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || currentTimestampAdded

	err, maxSubscribersAdded := retrieveClientBroadcastStreamMaxSubscribersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxSubscribersAdded

	err, metaDataAdded := retrieveClientBroadcastStreamMetaDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaDataAdded

	err, nameAdded := retrieveClientBroadcastStreamNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, parametersAdded := retrieveClientBroadcastStreamParametersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parametersAdded

	err, providerAdded := retrieveClientBroadcastStreamProviderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || providerAdded

	err, publishedNameAdded := retrieveClientBroadcastStreamPublishedNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || publishedNameAdded

	err, recordingAdded := retrieveClientBroadcastStreamRecordingFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || recordingAdded

	err, saveFilenameAdded := retrieveClientBroadcastStreamSaveFilenameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || saveFilenameAdded

	err, scopeAdded := retrieveClientBroadcastStreamScopeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeAdded

	err, stateAdded := retrieveClientBroadcastStreamStateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stateAdded

	err, statisticsAdded := retrieveClientBroadcastStreamStatisticsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statisticsAdded

	err, streamIdAdded := retrieveClientBroadcastStreamStreamIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamIdAdded

	err, streamListenersAdded := retrieveClientBroadcastStreamStreamListenersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamListenersAdded

	err, totalSubscribersAdded := retrieveClientBroadcastStreamTotalSubscribersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalSubscribersAdded

	return nil, retAdded
}

func retrieveClientBroadcastStreamAbsoluteStartTimeMsFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	absoluteStartTimeMsFlagName := fmt.Sprintf("%v.absoluteStartTimeMs", cmdPrefix)
	if cmd.Flags().Changed(absoluteStartTimeMsFlagName) {

		var absoluteStartTimeMsFlagName string
		if cmdPrefix == "" {
			absoluteStartTimeMsFlagName = "absoluteStartTimeMs"
		} else {
			absoluteStartTimeMsFlagName = fmt.Sprintf("%v.absoluteStartTimeMs", cmdPrefix)
		}

		absoluteStartTimeMsFlagValue, err := cmd.Flags().GetInt64(absoluteStartTimeMsFlagName)
		if err != nil {
			return err, false
		}
		m.AbsoluteStartTimeMs = absoluteStartTimeMsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamActiveSubscribersFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClientBroadcastStreamAutomaticRecordingFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	automaticRecordingFlagName := fmt.Sprintf("%v.automaticRecording", cmdPrefix)
	if cmd.Flags().Changed(automaticRecordingFlagName) {

		var automaticRecordingFlagName string
		if cmdPrefix == "" {
			automaticRecordingFlagName = "automaticRecording"
		} else {
			automaticRecordingFlagName = fmt.Sprintf("%v.automaticRecording", cmdPrefix)
		}

		automaticRecordingFlagValue, err := cmd.Flags().GetBool(automaticRecordingFlagName)
		if err != nil {
			return err, false
		}
		m.AutomaticRecording = automaticRecordingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamBroadcastStreamPublishNameFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	broadcastStreamPublishNameFlagName := fmt.Sprintf("%v.broadcastStreamPublishName", cmdPrefix)
	if cmd.Flags().Changed(broadcastStreamPublishNameFlagName) {

		var broadcastStreamPublishNameFlagName string
		if cmdPrefix == "" {
			broadcastStreamPublishNameFlagName = "broadcastStreamPublishName"
		} else {
			broadcastStreamPublishNameFlagName = fmt.Sprintf("%v.broadcastStreamPublishName", cmdPrefix)
		}

		broadcastStreamPublishNameFlagValue, err := cmd.Flags().GetString(broadcastStreamPublishNameFlagName)
		if err != nil {
			return err, false
		}
		m.BroadcastStreamPublishName = broadcastStreamPublishNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamBytesReceivedFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClientBroadcastStreamClientBufferDurationFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clientBufferDurationFlagName := fmt.Sprintf("%v.clientBufferDuration", cmdPrefix)
	if cmd.Flags().Changed(clientBufferDurationFlagName) {

		var clientBufferDurationFlagName string
		if cmdPrefix == "" {
			clientBufferDurationFlagName = "clientBufferDuration"
		} else {
			clientBufferDurationFlagName = fmt.Sprintf("%v.clientBufferDuration", cmdPrefix)
		}

		clientBufferDurationFlagValue, err := cmd.Flags().GetInt32(clientBufferDurationFlagName)
		if err != nil {
			return err, false
		}
		m.ClientBufferDuration = clientBufferDurationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamCodecInfoFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	codecInfoFlagName := fmt.Sprintf("%v.codecInfo", cmdPrefix)
	if cmd.Flags().Changed(codecInfoFlagName) {
		// info: complex object codecInfo IStreamCodecInfo is retrieved outside this Changed() block
	}
	codecInfoFlagValue := m.CodecInfo
	if swag.IsZero(codecInfoFlagValue) {
		codecInfoFlagValue = &models.IStreamCodecInfo{}
	}

	err, codecInfoAdded := retrieveModelIStreamCodecInfoFlags(depth+1, codecInfoFlagValue, codecInfoFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || codecInfoAdded
	if codecInfoAdded {
		m.CodecInfo = codecInfoFlagValue
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamConnectionFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectionFlagName := fmt.Sprintf("%v.connection", cmdPrefix)
	if cmd.Flags().Changed(connectionFlagName) {
		// info: complex object connection IStreamCapableConnection is retrieved outside this Changed() block
	}
	connectionFlagValue := m.Connection
	if swag.IsZero(connectionFlagValue) {
		connectionFlagValue = &models.IStreamCapableConnection{}
	}

	err, connectionAdded := retrieveModelIStreamCapableConnectionFlags(depth+1, connectionFlagValue, connectionFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionAdded
	if connectionAdded {
		m.Connection = connectionFlagValue
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamCreationTimeFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClientBroadcastStreamCurrentTimestampFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClientBroadcastStreamMaxSubscribersFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClientBroadcastStreamMetaDataFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metaDataFlagName := fmt.Sprintf("%v.metaData", cmdPrefix)
	if cmd.Flags().Changed(metaDataFlagName) {
		// info: complex object metaData Notify is retrieved outside this Changed() block
	}
	metaDataFlagValue := m.MetaData
	if swag.IsZero(metaDataFlagValue) {
		metaDataFlagValue = &models.Notify{}
	}

	err, metaDataAdded := retrieveModelNotifyFlags(depth+1, metaDataFlagValue, metaDataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metaDataAdded
	if metaDataAdded {
		m.MetaData = metaDataFlagValue
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamNameFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamParametersFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parametersFlagName := fmt.Sprintf("%v.parameters", cmdPrefix)
	if cmd.Flags().Changed(parametersFlagName) {
		// warning: parameters map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamProviderFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	providerFlagName := fmt.Sprintf("%v.provider", cmdPrefix)
	if cmd.Flags().Changed(providerFlagName) {
		// warning: provider map type IProvider is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamPublishedNameFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClientBroadcastStreamRecordingFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	recordingFlagName := fmt.Sprintf("%v.recording", cmdPrefix)
	if cmd.Flags().Changed(recordingFlagName) {

		var recordingFlagName string
		if cmdPrefix == "" {
			recordingFlagName = "recording"
		} else {
			recordingFlagName = fmt.Sprintf("%v.recording", cmdPrefix)
		}

		recordingFlagValue, err := cmd.Flags().GetBool(recordingFlagName)
		if err != nil {
			return err, false
		}
		m.Recording = recordingFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamSaveFilenameFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClientBroadcastStreamScopeFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scopeFlagName := fmt.Sprintf("%v.scope", cmdPrefix)
	if cmd.Flags().Changed(scopeFlagName) {
		// info: complex object scope IScope is retrieved outside this Changed() block
	}
	scopeFlagValue := m.Scope
	if swag.IsZero(scopeFlagValue) {
		scopeFlagValue = &models.IScope{}
	}

	err, scopeAdded := retrieveModelIScopeFlags(depth+1, scopeFlagValue, scopeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeAdded
	if scopeAdded {
		m.Scope = scopeFlagValue
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamStateFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stateFlagName := fmt.Sprintf("%v.state", cmdPrefix)
	if cmd.Flags().Changed(stateFlagName) {

		var stateFlagName string
		if cmdPrefix == "" {
			stateFlagName = "state"
		} else {
			stateFlagName = fmt.Sprintf("%v.state", cmdPrefix)
		}

		stateFlagValue, err := cmd.Flags().GetString(stateFlagName)
		if err != nil {
			return err, false
		}
		m.State = stateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamStatisticsFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statisticsFlagName := fmt.Sprintf("%v.statistics", cmdPrefix)
	if cmd.Flags().Changed(statisticsFlagName) {
		// info: complex object statistics IClientBroadcastStreamStatistics is retrieved outside this Changed() block
	}
	statisticsFlagValue := m.Statistics
	if swag.IsZero(statisticsFlagValue) {
		statisticsFlagValue = &models.IClientBroadcastStreamStatistics{}
	}

	err, statisticsAdded := retrieveModelIClientBroadcastStreamStatisticsFlags(depth+1, statisticsFlagValue, statisticsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statisticsAdded
	if statisticsAdded {
		m.Statistics = statisticsFlagValue
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamStreamIDFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamIdFlagName := fmt.Sprintf("%v.streamId", cmdPrefix)
	if cmd.Flags().Changed(streamIdFlagName) {
		// warning: streamId map type Number is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamStreamListenersFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamListenersFlagName := fmt.Sprintf("%v.streamListeners", cmdPrefix)
	if cmd.Flags().Changed(streamListenersFlagName) {
		// warning: streamListeners array type []IStreamListener is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveClientBroadcastStreamTotalSubscribersFlags(depth int, m *models.ClientBroadcastStream, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
