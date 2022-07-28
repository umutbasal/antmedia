// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for WebRTCVideoSendStats

// register flags to command
func registerModelWebRTCVideoSendStatsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWebRTCVideoSendStatsTimeMs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoBytesSent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoBytesSentPerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoFirCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoFramesEncoded(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoFramesEncodedPerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoNackCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoPacketsSent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoPacketsSentPerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoSendStatsVideoPliCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebRTCVideoSendStatsTimeMs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timeMsDescription := ``

	var timeMsFlagName string
	if cmdPrefix == "" {
		timeMsFlagName = "timeMs"
	} else {
		timeMsFlagName = fmt.Sprintf("%v.timeMs", cmdPrefix)
	}

	var timeMsFlagDefault int64

	_ = cmd.PersistentFlags().Int64(timeMsFlagName, timeMsFlagDefault, timeMsDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoBytesSent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoBytesSentDescription := ``

	var videoBytesSentFlagName string
	if cmdPrefix == "" {
		videoBytesSentFlagName = "videoBytesSent"
	} else {
		videoBytesSentFlagName = fmt.Sprintf("%v.videoBytesSent", cmdPrefix)
	}

	var videoBytesSentFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoBytesSentFlagName, videoBytesSentFlagDefault, videoBytesSentDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoBytesSentPerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoBytesSentPerSecondDescription := ``

	var videoBytesSentPerSecondFlagName string
	if cmdPrefix == "" {
		videoBytesSentPerSecondFlagName = "videoBytesSentPerSecond"
	} else {
		videoBytesSentPerSecondFlagName = fmt.Sprintf("%v.videoBytesSentPerSecond", cmdPrefix)
	}

	var videoBytesSentPerSecondFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoBytesSentPerSecondFlagName, videoBytesSentPerSecondFlagDefault, videoBytesSentPerSecondDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoFirCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoFirCountDescription := ``

	var videoFirCountFlagName string
	if cmdPrefix == "" {
		videoFirCountFlagName = "videoFirCount"
	} else {
		videoFirCountFlagName = fmt.Sprintf("%v.videoFirCount", cmdPrefix)
	}

	var videoFirCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoFirCountFlagName, videoFirCountFlagDefault, videoFirCountDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoFramesEncoded(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoFramesEncodedDescription := ``

	var videoFramesEncodedFlagName string
	if cmdPrefix == "" {
		videoFramesEncodedFlagName = "videoFramesEncoded"
	} else {
		videoFramesEncodedFlagName = fmt.Sprintf("%v.videoFramesEncoded", cmdPrefix)
	}

	var videoFramesEncodedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoFramesEncodedFlagName, videoFramesEncodedFlagDefault, videoFramesEncodedDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoFramesEncodedPerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoFramesEncodedPerSecondDescription := ``

	var videoFramesEncodedPerSecondFlagName string
	if cmdPrefix == "" {
		videoFramesEncodedPerSecondFlagName = "videoFramesEncodedPerSecond"
	} else {
		videoFramesEncodedPerSecondFlagName = fmt.Sprintf("%v.videoFramesEncodedPerSecond", cmdPrefix)
	}

	var videoFramesEncodedPerSecondFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoFramesEncodedPerSecondFlagName, videoFramesEncodedPerSecondFlagDefault, videoFramesEncodedPerSecondDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoNackCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoNackCountDescription := ``

	var videoNackCountFlagName string
	if cmdPrefix == "" {
		videoNackCountFlagName = "videoNackCount"
	} else {
		videoNackCountFlagName = fmt.Sprintf("%v.videoNackCount", cmdPrefix)
	}

	var videoNackCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoNackCountFlagName, videoNackCountFlagDefault, videoNackCountDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoPacketsSent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoPacketsSentDescription := ``

	var videoPacketsSentFlagName string
	if cmdPrefix == "" {
		videoPacketsSentFlagName = "videoPacketsSent"
	} else {
		videoPacketsSentFlagName = fmt.Sprintf("%v.videoPacketsSent", cmdPrefix)
	}

	var videoPacketsSentFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoPacketsSentFlagName, videoPacketsSentFlagDefault, videoPacketsSentDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoPacketsSentPerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoPacketsSentPerSecondDescription := ``

	var videoPacketsSentPerSecondFlagName string
	if cmdPrefix == "" {
		videoPacketsSentPerSecondFlagName = "videoPacketsSentPerSecond"
	} else {
		videoPacketsSentPerSecondFlagName = fmt.Sprintf("%v.videoPacketsSentPerSecond", cmdPrefix)
	}

	var videoPacketsSentPerSecondFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoPacketsSentPerSecondFlagName, videoPacketsSentPerSecondFlagDefault, videoPacketsSentPerSecondDescription)

	return nil
}

func registerWebRTCVideoSendStatsVideoPliCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoPliCountDescription := ``

	var videoPliCountFlagName string
	if cmdPrefix == "" {
		videoPliCountFlagName = "videoPliCount"
	} else {
		videoPliCountFlagName = fmt.Sprintf("%v.videoPliCount", cmdPrefix)
	}

	var videoPliCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoPliCountFlagName, videoPliCountFlagDefault, videoPliCountDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelWebRTCVideoSendStatsFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, timeMsAdded := retrieveWebRTCVideoSendStatsTimeMsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timeMsAdded

	err, videoBytesSentAdded := retrieveWebRTCVideoSendStatsVideoBytesSentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoBytesSentAdded

	err, videoBytesSentPerSecondAdded := retrieveWebRTCVideoSendStatsVideoBytesSentPerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoBytesSentPerSecondAdded

	err, videoFirCountAdded := retrieveWebRTCVideoSendStatsVideoFirCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoFirCountAdded

	err, videoFramesEncodedAdded := retrieveWebRTCVideoSendStatsVideoFramesEncodedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoFramesEncodedAdded

	err, videoFramesEncodedPerSecondAdded := retrieveWebRTCVideoSendStatsVideoFramesEncodedPerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoFramesEncodedPerSecondAdded

	err, videoNackCountAdded := retrieveWebRTCVideoSendStatsVideoNackCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoNackCountAdded

	err, videoPacketsSentAdded := retrieveWebRTCVideoSendStatsVideoPacketsSentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoPacketsSentAdded

	err, videoPacketsSentPerSecondAdded := retrieveWebRTCVideoSendStatsVideoPacketsSentPerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoPacketsSentPerSecondAdded

	err, videoPliCountAdded := retrieveWebRTCVideoSendStatsVideoPliCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoPliCountAdded

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsTimeMsFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timeMsFlagName := fmt.Sprintf("%v.timeMs", cmdPrefix)
	if cmd.Flags().Changed(timeMsFlagName) {

		var timeMsFlagName string
		if cmdPrefix == "" {
			timeMsFlagName = "timeMs"
		} else {
			timeMsFlagName = fmt.Sprintf("%v.timeMs", cmdPrefix)
		}

		timeMsFlagValue, err := cmd.Flags().GetInt64(timeMsFlagName)
		if err != nil {
			return err, false
		}
		m.TimeMs = timeMsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoBytesSentFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoBytesSentFlagName := fmt.Sprintf("%v.videoBytesSent", cmdPrefix)
	if cmd.Flags().Changed(videoBytesSentFlagName) {

		var videoBytesSentFlagName string
		if cmdPrefix == "" {
			videoBytesSentFlagName = "videoBytesSent"
		} else {
			videoBytesSentFlagName = fmt.Sprintf("%v.videoBytesSent", cmdPrefix)
		}

		videoBytesSentFlagValue, err := cmd.Flags().GetInt64(videoBytesSentFlagName)
		if err != nil {
			return err, false
		}
		m.VideoBytesSent = videoBytesSentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoBytesSentPerSecondFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoBytesSentPerSecondFlagName := fmt.Sprintf("%v.videoBytesSentPerSecond", cmdPrefix)
	if cmd.Flags().Changed(videoBytesSentPerSecondFlagName) {

		var videoBytesSentPerSecondFlagName string
		if cmdPrefix == "" {
			videoBytesSentPerSecondFlagName = "videoBytesSentPerSecond"
		} else {
			videoBytesSentPerSecondFlagName = fmt.Sprintf("%v.videoBytesSentPerSecond", cmdPrefix)
		}

		videoBytesSentPerSecondFlagValue, err := cmd.Flags().GetInt64(videoBytesSentPerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.VideoBytesSentPerSecond = videoBytesSentPerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoFirCountFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoFirCountFlagName := fmt.Sprintf("%v.videoFirCount", cmdPrefix)
	if cmd.Flags().Changed(videoFirCountFlagName) {

		var videoFirCountFlagName string
		if cmdPrefix == "" {
			videoFirCountFlagName = "videoFirCount"
		} else {
			videoFirCountFlagName = fmt.Sprintf("%v.videoFirCount", cmdPrefix)
		}

		videoFirCountFlagValue, err := cmd.Flags().GetInt64(videoFirCountFlagName)
		if err != nil {
			return err, false
		}
		m.VideoFirCount = videoFirCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoFramesEncodedFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoFramesEncodedFlagName := fmt.Sprintf("%v.videoFramesEncoded", cmdPrefix)
	if cmd.Flags().Changed(videoFramesEncodedFlagName) {

		var videoFramesEncodedFlagName string
		if cmdPrefix == "" {
			videoFramesEncodedFlagName = "videoFramesEncoded"
		} else {
			videoFramesEncodedFlagName = fmt.Sprintf("%v.videoFramesEncoded", cmdPrefix)
		}

		videoFramesEncodedFlagValue, err := cmd.Flags().GetInt64(videoFramesEncodedFlagName)
		if err != nil {
			return err, false
		}
		m.VideoFramesEncoded = videoFramesEncodedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoFramesEncodedPerSecondFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoFramesEncodedPerSecondFlagName := fmt.Sprintf("%v.videoFramesEncodedPerSecond", cmdPrefix)
	if cmd.Flags().Changed(videoFramesEncodedPerSecondFlagName) {

		var videoFramesEncodedPerSecondFlagName string
		if cmdPrefix == "" {
			videoFramesEncodedPerSecondFlagName = "videoFramesEncodedPerSecond"
		} else {
			videoFramesEncodedPerSecondFlagName = fmt.Sprintf("%v.videoFramesEncodedPerSecond", cmdPrefix)
		}

		videoFramesEncodedPerSecondFlagValue, err := cmd.Flags().GetInt64(videoFramesEncodedPerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.VideoFramesEncodedPerSecond = videoFramesEncodedPerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoNackCountFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoNackCountFlagName := fmt.Sprintf("%v.videoNackCount", cmdPrefix)
	if cmd.Flags().Changed(videoNackCountFlagName) {

		var videoNackCountFlagName string
		if cmdPrefix == "" {
			videoNackCountFlagName = "videoNackCount"
		} else {
			videoNackCountFlagName = fmt.Sprintf("%v.videoNackCount", cmdPrefix)
		}

		videoNackCountFlagValue, err := cmd.Flags().GetInt64(videoNackCountFlagName)
		if err != nil {
			return err, false
		}
		m.VideoNackCount = videoNackCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoPacketsSentFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoPacketsSentFlagName := fmt.Sprintf("%v.videoPacketsSent", cmdPrefix)
	if cmd.Flags().Changed(videoPacketsSentFlagName) {

		var videoPacketsSentFlagName string
		if cmdPrefix == "" {
			videoPacketsSentFlagName = "videoPacketsSent"
		} else {
			videoPacketsSentFlagName = fmt.Sprintf("%v.videoPacketsSent", cmdPrefix)
		}

		videoPacketsSentFlagValue, err := cmd.Flags().GetInt64(videoPacketsSentFlagName)
		if err != nil {
			return err, false
		}
		m.VideoPacketsSent = videoPacketsSentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoPacketsSentPerSecondFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoPacketsSentPerSecondFlagName := fmt.Sprintf("%v.videoPacketsSentPerSecond", cmdPrefix)
	if cmd.Flags().Changed(videoPacketsSentPerSecondFlagName) {

		var videoPacketsSentPerSecondFlagName string
		if cmdPrefix == "" {
			videoPacketsSentPerSecondFlagName = "videoPacketsSentPerSecond"
		} else {
			videoPacketsSentPerSecondFlagName = fmt.Sprintf("%v.videoPacketsSentPerSecond", cmdPrefix)
		}

		videoPacketsSentPerSecondFlagValue, err := cmd.Flags().GetInt64(videoPacketsSentPerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.VideoPacketsSentPerSecond = videoPacketsSentPerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoSendStatsVideoPliCountFlags(depth int, m *models.WebRTCVideoSendStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoPliCountFlagName := fmt.Sprintf("%v.videoPliCount", cmdPrefix)
	if cmd.Flags().Changed(videoPliCountFlagName) {

		var videoPliCountFlagName string
		if cmdPrefix == "" {
			videoPliCountFlagName = "videoPliCount"
		} else {
			videoPliCountFlagName = fmt.Sprintf("%v.videoPliCount", cmdPrefix)
		}

		videoPliCountFlagValue, err := cmd.Flags().GetInt64(videoPliCountFlagName)
		if err != nil {
			return err, false
		}
		m.VideoPliCount = videoPliCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}