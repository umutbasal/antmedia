// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for SubscriberStats

// register flags to command
func registerModelSubscriberStatsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSubscriberStatsAvgAudioBitrate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSubscriberStatsAvgVideoBitrate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSubscriberStatsConnectionEvents(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSubscriberStatsStreamID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSubscriberStatsSubscriberID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSubscriberStatsAvgAudioBitrate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	avgAudioBitrateDescription := `average audio bitrate for a subscriber`

	var avgAudioBitrateFlagName string
	if cmdPrefix == "" {
		avgAudioBitrateFlagName = "avgAudioBitrate"
	} else {
		avgAudioBitrateFlagName = fmt.Sprintf("%v.avgAudioBitrate", cmdPrefix)
	}

	var avgAudioBitrateFlagDefault int64

	_ = cmd.PersistentFlags().Int64(avgAudioBitrateFlagName, avgAudioBitrateFlagDefault, avgAudioBitrateDescription)

	return nil
}

func registerSubscriberStatsAvgVideoBitrate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	avgVideoBitrateDescription := `average video bitrate for a subscriber`

	var avgVideoBitrateFlagName string
	if cmdPrefix == "" {
		avgVideoBitrateFlagName = "avgVideoBitrate"
	} else {
		avgVideoBitrateFlagName = fmt.Sprintf("%v.avgVideoBitrate", cmdPrefix)
	}

	var avgVideoBitrateFlagDefault int64

	_ = cmd.PersistentFlags().Int64(avgVideoBitrateFlagName, avgVideoBitrateFlagDefault, avgVideoBitrateDescription)

	return nil
}

func registerSubscriberStatsConnectionEvents(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: connectionEvents []*ConnectionEvent array type is not supported by go-swagger cli yet

	return nil
}

func registerSubscriberStatsStreamID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	streamIdDescription := `the stream id of the token`

	var streamIdFlagName string
	if cmdPrefix == "" {
		streamIdFlagName = "streamId"
	} else {
		streamIdFlagName = fmt.Sprintf("%v.streamId", cmdPrefix)
	}

	var streamIdFlagDefault string

	_ = cmd.PersistentFlags().String(streamIdFlagName, streamIdFlagDefault, streamIdDescription)

	return nil
}

func registerSubscriberStatsSubscriberID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	subscriberIdDescription := `the subscriber id of the subscriber`

	var subscriberIdFlagName string
	if cmdPrefix == "" {
		subscriberIdFlagName = "subscriberId"
	} else {
		subscriberIdFlagName = fmt.Sprintf("%v.subscriberId", cmdPrefix)
	}

	var subscriberIdFlagDefault string

	_ = cmd.PersistentFlags().String(subscriberIdFlagName, subscriberIdFlagDefault, subscriberIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSubscriberStatsFlags(depth int, m *models.SubscriberStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, avgAudioBitrateAdded := retrieveSubscriberStatsAvgAudioBitrateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || avgAudioBitrateAdded

	err, avgVideoBitrateAdded := retrieveSubscriberStatsAvgVideoBitrateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || avgVideoBitrateAdded

	err, connectionEventsAdded := retrieveSubscriberStatsConnectionEventsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionEventsAdded

	err, streamIdAdded := retrieveSubscriberStatsStreamIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamIdAdded

	err, subscriberIdAdded := retrieveSubscriberStatsSubscriberIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || subscriberIdAdded

	return nil, retAdded
}

func retrieveSubscriberStatsAvgAudioBitrateFlags(depth int, m *models.SubscriberStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	avgAudioBitrateFlagName := fmt.Sprintf("%v.avgAudioBitrate", cmdPrefix)
	if cmd.Flags().Changed(avgAudioBitrateFlagName) {

		var avgAudioBitrateFlagName string
		if cmdPrefix == "" {
			avgAudioBitrateFlagName = "avgAudioBitrate"
		} else {
			avgAudioBitrateFlagName = fmt.Sprintf("%v.avgAudioBitrate", cmdPrefix)
		}

		avgAudioBitrateFlagValue, err := cmd.Flags().GetInt64(avgAudioBitrateFlagName)
		if err != nil {
			return err, false
		}
		m.AvgAudioBitrate = avgAudioBitrateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSubscriberStatsAvgVideoBitrateFlags(depth int, m *models.SubscriberStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	avgVideoBitrateFlagName := fmt.Sprintf("%v.avgVideoBitrate", cmdPrefix)
	if cmd.Flags().Changed(avgVideoBitrateFlagName) {

		var avgVideoBitrateFlagName string
		if cmdPrefix == "" {
			avgVideoBitrateFlagName = "avgVideoBitrate"
		} else {
			avgVideoBitrateFlagName = fmt.Sprintf("%v.avgVideoBitrate", cmdPrefix)
		}

		avgVideoBitrateFlagValue, err := cmd.Flags().GetInt64(avgVideoBitrateFlagName)
		if err != nil {
			return err, false
		}
		m.AvgVideoBitrate = avgVideoBitrateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSubscriberStatsConnectionEventsFlags(depth int, m *models.SubscriberStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectionEventsFlagName := fmt.Sprintf("%v.connectionEvents", cmdPrefix)
	if cmd.Flags().Changed(connectionEventsFlagName) {
		// warning: connectionEvents array type []*ConnectionEvent is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveSubscriberStatsStreamIDFlags(depth int, m *models.SubscriberStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamIdFlagName := fmt.Sprintf("%v.streamId", cmdPrefix)
	if cmd.Flags().Changed(streamIdFlagName) {

		var streamIdFlagName string
		if cmdPrefix == "" {
			streamIdFlagName = "streamId"
		} else {
			streamIdFlagName = fmt.Sprintf("%v.streamId", cmdPrefix)
		}

		streamIdFlagValue, err := cmd.Flags().GetString(streamIdFlagName)
		if err != nil {
			return err, false
		}
		m.StreamID = streamIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSubscriberStatsSubscriberIDFlags(depth int, m *models.SubscriberStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	subscriberIdFlagName := fmt.Sprintf("%v.subscriberId", cmdPrefix)
	if cmd.Flags().Changed(subscriberIdFlagName) {

		var subscriberIdFlagName string
		if cmdPrefix == "" {
			subscriberIdFlagName = "subscriberId"
		} else {
			subscriberIdFlagName = fmt.Sprintf("%v.subscriberId", cmdPrefix)
		}

		subscriberIdFlagValue, err := cmd.Flags().GetString(subscriberIdFlagName)
		if err != nil {
			return err, false
		}
		m.SubscriberID = subscriberIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
