// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for WebRTCAudioReceiveStats

// register flags to command
func registerModelWebRTCAudioReceiveStatsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWebRTCAudioReceiveStatsAudioBytesReceived(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCAudioReceiveStatsAudioBytesReceivedPerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCAudioReceiveStatsAudioFractionLost(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCAudioReceiveStatsAudioJitter(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCAudioReceiveStatsAudioPacketsLost(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCAudioReceiveStatsAudioPacketsReceived(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCAudioReceiveStatsAudioPacketsReceivedPerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebRTCAudioReceiveStatsAudioBytesReceived(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	audioBytesReceivedDescription := ``

	var audioBytesReceivedFlagName string
	if cmdPrefix == "" {
		audioBytesReceivedFlagName = "audioBytesReceived"
	} else {
		audioBytesReceivedFlagName = fmt.Sprintf("%v.audioBytesReceived", cmdPrefix)
	}

	var audioBytesReceivedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(audioBytesReceivedFlagName, audioBytesReceivedFlagDefault, audioBytesReceivedDescription)

	return nil
}

func registerWebRTCAudioReceiveStatsAudioBytesReceivedPerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	audioBytesReceivedPerSecondDescription := ``

	var audioBytesReceivedPerSecondFlagName string
	if cmdPrefix == "" {
		audioBytesReceivedPerSecondFlagName = "audioBytesReceivedPerSecond"
	} else {
		audioBytesReceivedPerSecondFlagName = fmt.Sprintf("%v.audioBytesReceivedPerSecond", cmdPrefix)
	}

	var audioBytesReceivedPerSecondFlagDefault int64

	_ = cmd.PersistentFlags().Int64(audioBytesReceivedPerSecondFlagName, audioBytesReceivedPerSecondFlagDefault, audioBytesReceivedPerSecondDescription)

	return nil
}

func registerWebRTCAudioReceiveStatsAudioFractionLost(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	audioFractionLostDescription := ``

	var audioFractionLostFlagName string
	if cmdPrefix == "" {
		audioFractionLostFlagName = "audioFractionLost"
	} else {
		audioFractionLostFlagName = fmt.Sprintf("%v.audioFractionLost", cmdPrefix)
	}

	var audioFractionLostFlagDefault float64

	_ = cmd.PersistentFlags().Float64(audioFractionLostFlagName, audioFractionLostFlagDefault, audioFractionLostDescription)

	return nil
}

func registerWebRTCAudioReceiveStatsAudioJitter(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	audioJitterDescription := ``

	var audioJitterFlagName string
	if cmdPrefix == "" {
		audioJitterFlagName = "audioJitter"
	} else {
		audioJitterFlagName = fmt.Sprintf("%v.audioJitter", cmdPrefix)
	}

	var audioJitterFlagDefault float64

	_ = cmd.PersistentFlags().Float64(audioJitterFlagName, audioJitterFlagDefault, audioJitterDescription)

	return nil
}

func registerWebRTCAudioReceiveStatsAudioPacketsLost(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	audioPacketsLostDescription := ``

	var audioPacketsLostFlagName string
	if cmdPrefix == "" {
		audioPacketsLostFlagName = "audioPacketsLost"
	} else {
		audioPacketsLostFlagName = fmt.Sprintf("%v.audioPacketsLost", cmdPrefix)
	}

	var audioPacketsLostFlagDefault int32

	_ = cmd.PersistentFlags().Int32(audioPacketsLostFlagName, audioPacketsLostFlagDefault, audioPacketsLostDescription)

	return nil
}

func registerWebRTCAudioReceiveStatsAudioPacketsReceived(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	audioPacketsReceivedDescription := ``

	var audioPacketsReceivedFlagName string
	if cmdPrefix == "" {
		audioPacketsReceivedFlagName = "audioPacketsReceived"
	} else {
		audioPacketsReceivedFlagName = fmt.Sprintf("%v.audioPacketsReceived", cmdPrefix)
	}

	var audioPacketsReceivedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(audioPacketsReceivedFlagName, audioPacketsReceivedFlagDefault, audioPacketsReceivedDescription)

	return nil
}

func registerWebRTCAudioReceiveStatsAudioPacketsReceivedPerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	audioPacketsReceivedPerSecondDescription := ``

	var audioPacketsReceivedPerSecondFlagName string
	if cmdPrefix == "" {
		audioPacketsReceivedPerSecondFlagName = "audioPacketsReceivedPerSecond"
	} else {
		audioPacketsReceivedPerSecondFlagName = fmt.Sprintf("%v.audioPacketsReceivedPerSecond", cmdPrefix)
	}

	var audioPacketsReceivedPerSecondFlagDefault int64

	_ = cmd.PersistentFlags().Int64(audioPacketsReceivedPerSecondFlagName, audioPacketsReceivedPerSecondFlagDefault, audioPacketsReceivedPerSecondDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelWebRTCAudioReceiveStatsFlags(depth int, m *models.WebRTCAudioReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, audioBytesReceivedAdded := retrieveWebRTCAudioReceiveStatsAudioBytesReceivedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || audioBytesReceivedAdded

	err, audioBytesReceivedPerSecondAdded := retrieveWebRTCAudioReceiveStatsAudioBytesReceivedPerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || audioBytesReceivedPerSecondAdded

	err, audioFractionLostAdded := retrieveWebRTCAudioReceiveStatsAudioFractionLostFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || audioFractionLostAdded

	err, audioJitterAdded := retrieveWebRTCAudioReceiveStatsAudioJitterFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || audioJitterAdded

	err, audioPacketsLostAdded := retrieveWebRTCAudioReceiveStatsAudioPacketsLostFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || audioPacketsLostAdded

	err, audioPacketsReceivedAdded := retrieveWebRTCAudioReceiveStatsAudioPacketsReceivedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || audioPacketsReceivedAdded

	err, audioPacketsReceivedPerSecondAdded := retrieveWebRTCAudioReceiveStatsAudioPacketsReceivedPerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || audioPacketsReceivedPerSecondAdded

	return nil, retAdded
}

func retrieveWebRTCAudioReceiveStatsAudioBytesReceivedFlags(depth int, m *models.WebRTCAudioReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	audioBytesReceivedFlagName := fmt.Sprintf("%v.audioBytesReceived", cmdPrefix)
	if cmd.Flags().Changed(audioBytesReceivedFlagName) {

		var audioBytesReceivedFlagName string
		if cmdPrefix == "" {
			audioBytesReceivedFlagName = "audioBytesReceived"
		} else {
			audioBytesReceivedFlagName = fmt.Sprintf("%v.audioBytesReceived", cmdPrefix)
		}

		audioBytesReceivedFlagValue, err := cmd.Flags().GetInt64(audioBytesReceivedFlagName)
		if err != nil {
			return err, false
		}
		m.AudioBytesReceived = audioBytesReceivedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCAudioReceiveStatsAudioBytesReceivedPerSecondFlags(depth int, m *models.WebRTCAudioReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	audioBytesReceivedPerSecondFlagName := fmt.Sprintf("%v.audioBytesReceivedPerSecond", cmdPrefix)
	if cmd.Flags().Changed(audioBytesReceivedPerSecondFlagName) {

		var audioBytesReceivedPerSecondFlagName string
		if cmdPrefix == "" {
			audioBytesReceivedPerSecondFlagName = "audioBytesReceivedPerSecond"
		} else {
			audioBytesReceivedPerSecondFlagName = fmt.Sprintf("%v.audioBytesReceivedPerSecond", cmdPrefix)
		}

		audioBytesReceivedPerSecondFlagValue, err := cmd.Flags().GetInt64(audioBytesReceivedPerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.AudioBytesReceivedPerSecond = audioBytesReceivedPerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCAudioReceiveStatsAudioFractionLostFlags(depth int, m *models.WebRTCAudioReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	audioFractionLostFlagName := fmt.Sprintf("%v.audioFractionLost", cmdPrefix)
	if cmd.Flags().Changed(audioFractionLostFlagName) {

		var audioFractionLostFlagName string
		if cmdPrefix == "" {
			audioFractionLostFlagName = "audioFractionLost"
		} else {
			audioFractionLostFlagName = fmt.Sprintf("%v.audioFractionLost", cmdPrefix)
		}

		audioFractionLostFlagValue, err := cmd.Flags().GetFloat64(audioFractionLostFlagName)
		if err != nil {
			return err, false
		}
		m.AudioFractionLost = audioFractionLostFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCAudioReceiveStatsAudioJitterFlags(depth int, m *models.WebRTCAudioReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	audioJitterFlagName := fmt.Sprintf("%v.audioJitter", cmdPrefix)
	if cmd.Flags().Changed(audioJitterFlagName) {

		var audioJitterFlagName string
		if cmdPrefix == "" {
			audioJitterFlagName = "audioJitter"
		} else {
			audioJitterFlagName = fmt.Sprintf("%v.audioJitter", cmdPrefix)
		}

		audioJitterFlagValue, err := cmd.Flags().GetFloat64(audioJitterFlagName)
		if err != nil {
			return err, false
		}
		m.AudioJitter = audioJitterFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCAudioReceiveStatsAudioPacketsLostFlags(depth int, m *models.WebRTCAudioReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	audioPacketsLostFlagName := fmt.Sprintf("%v.audioPacketsLost", cmdPrefix)
	if cmd.Flags().Changed(audioPacketsLostFlagName) {

		var audioPacketsLostFlagName string
		if cmdPrefix == "" {
			audioPacketsLostFlagName = "audioPacketsLost"
		} else {
			audioPacketsLostFlagName = fmt.Sprintf("%v.audioPacketsLost", cmdPrefix)
		}

		audioPacketsLostFlagValue, err := cmd.Flags().GetInt32(audioPacketsLostFlagName)
		if err != nil {
			return err, false
		}
		m.AudioPacketsLost = audioPacketsLostFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCAudioReceiveStatsAudioPacketsReceivedFlags(depth int, m *models.WebRTCAudioReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	audioPacketsReceivedFlagName := fmt.Sprintf("%v.audioPacketsReceived", cmdPrefix)
	if cmd.Flags().Changed(audioPacketsReceivedFlagName) {

		var audioPacketsReceivedFlagName string
		if cmdPrefix == "" {
			audioPacketsReceivedFlagName = "audioPacketsReceived"
		} else {
			audioPacketsReceivedFlagName = fmt.Sprintf("%v.audioPacketsReceived", cmdPrefix)
		}

		audioPacketsReceivedFlagValue, err := cmd.Flags().GetInt64(audioPacketsReceivedFlagName)
		if err != nil {
			return err, false
		}
		m.AudioPacketsReceived = audioPacketsReceivedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCAudioReceiveStatsAudioPacketsReceivedPerSecondFlags(depth int, m *models.WebRTCAudioReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	audioPacketsReceivedPerSecondFlagName := fmt.Sprintf("%v.audioPacketsReceivedPerSecond", cmdPrefix)
	if cmd.Flags().Changed(audioPacketsReceivedPerSecondFlagName) {

		var audioPacketsReceivedPerSecondFlagName string
		if cmdPrefix == "" {
			audioPacketsReceivedPerSecondFlagName = "audioPacketsReceivedPerSecond"
		} else {
			audioPacketsReceivedPerSecondFlagName = fmt.Sprintf("%v.audioPacketsReceivedPerSecond", cmdPrefix)
		}

		audioPacketsReceivedPerSecondFlagValue, err := cmd.Flags().GetInt64(audioPacketsReceivedPerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.AudioPacketsReceivedPerSecond = audioPacketsReceivedPerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}
