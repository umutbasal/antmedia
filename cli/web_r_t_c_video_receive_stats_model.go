// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for WebRTCVideoReceiveStats

// register flags to command
func registerModelWebRTCVideoReceiveStatsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWebRTCVideoReceiveStatsVideoBytesReceived(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoBytesReceivedPerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoFirCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoFractionLost(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoFrameReceived(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoFrameReceivedPerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoNackCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoPacketsLost(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoPacketsReceived(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoPacketsReceivedPerSecond(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCVideoReceiveStatsVideoPliCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebRTCVideoReceiveStatsVideoBytesReceived(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoBytesReceivedDescription := ``

	var videoBytesReceivedFlagName string
	if cmdPrefix == "" {
		videoBytesReceivedFlagName = "videoBytesReceived"
	} else {
		videoBytesReceivedFlagName = fmt.Sprintf("%v.videoBytesReceived", cmdPrefix)
	}

	var videoBytesReceivedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoBytesReceivedFlagName, videoBytesReceivedFlagDefault, videoBytesReceivedDescription)

	return nil
}

func registerWebRTCVideoReceiveStatsVideoBytesReceivedPerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoBytesReceivedPerSecondDescription := ``

	var videoBytesReceivedPerSecondFlagName string
	if cmdPrefix == "" {
		videoBytesReceivedPerSecondFlagName = "videoBytesReceivedPerSecond"
	} else {
		videoBytesReceivedPerSecondFlagName = fmt.Sprintf("%v.videoBytesReceivedPerSecond", cmdPrefix)
	}

	var videoBytesReceivedPerSecondFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoBytesReceivedPerSecondFlagName, videoBytesReceivedPerSecondFlagDefault, videoBytesReceivedPerSecondDescription)

	return nil
}

func registerWebRTCVideoReceiveStatsVideoFirCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerWebRTCVideoReceiveStatsVideoFractionLost(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoFractionLostDescription := ``

	var videoFractionLostFlagName string
	if cmdPrefix == "" {
		videoFractionLostFlagName = "videoFractionLost"
	} else {
		videoFractionLostFlagName = fmt.Sprintf("%v.videoFractionLost", cmdPrefix)
	}

	var videoFractionLostFlagDefault float64

	_ = cmd.PersistentFlags().Float64(videoFractionLostFlagName, videoFractionLostFlagDefault, videoFractionLostDescription)

	return nil
}

func registerWebRTCVideoReceiveStatsVideoFrameReceived(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoFrameReceivedDescription := ``

	var videoFrameReceivedFlagName string
	if cmdPrefix == "" {
		videoFrameReceivedFlagName = "videoFrameReceived"
	} else {
		videoFrameReceivedFlagName = fmt.Sprintf("%v.videoFrameReceived", cmdPrefix)
	}

	var videoFrameReceivedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoFrameReceivedFlagName, videoFrameReceivedFlagDefault, videoFrameReceivedDescription)

	return nil
}

func registerWebRTCVideoReceiveStatsVideoFrameReceivedPerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoFrameReceivedPerSecondDescription := ``

	var videoFrameReceivedPerSecondFlagName string
	if cmdPrefix == "" {
		videoFrameReceivedPerSecondFlagName = "videoFrameReceivedPerSecond"
	} else {
		videoFrameReceivedPerSecondFlagName = fmt.Sprintf("%v.videoFrameReceivedPerSecond", cmdPrefix)
	}

	var videoFrameReceivedPerSecondFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoFrameReceivedPerSecondFlagName, videoFrameReceivedPerSecondFlagDefault, videoFrameReceivedPerSecondDescription)

	return nil
}

func registerWebRTCVideoReceiveStatsVideoNackCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerWebRTCVideoReceiveStatsVideoPacketsLost(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoPacketsLostDescription := ``

	var videoPacketsLostFlagName string
	if cmdPrefix == "" {
		videoPacketsLostFlagName = "videoPacketsLost"
	} else {
		videoPacketsLostFlagName = fmt.Sprintf("%v.videoPacketsLost", cmdPrefix)
	}

	var videoPacketsLostFlagDefault int32

	_ = cmd.PersistentFlags().Int32(videoPacketsLostFlagName, videoPacketsLostFlagDefault, videoPacketsLostDescription)

	return nil
}

func registerWebRTCVideoReceiveStatsVideoPacketsReceived(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoPacketsReceivedDescription := ``

	var videoPacketsReceivedFlagName string
	if cmdPrefix == "" {
		videoPacketsReceivedFlagName = "videoPacketsReceived"
	} else {
		videoPacketsReceivedFlagName = fmt.Sprintf("%v.videoPacketsReceived", cmdPrefix)
	}

	var videoPacketsReceivedFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoPacketsReceivedFlagName, videoPacketsReceivedFlagDefault, videoPacketsReceivedDescription)

	return nil
}

func registerWebRTCVideoReceiveStatsVideoPacketsReceivedPerSecond(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	videoPacketsReceivedPerSecondDescription := ``

	var videoPacketsReceivedPerSecondFlagName string
	if cmdPrefix == "" {
		videoPacketsReceivedPerSecondFlagName = "videoPacketsReceivedPerSecond"
	} else {
		videoPacketsReceivedPerSecondFlagName = fmt.Sprintf("%v.videoPacketsReceivedPerSecond", cmdPrefix)
	}

	var videoPacketsReceivedPerSecondFlagDefault int64

	_ = cmd.PersistentFlags().Int64(videoPacketsReceivedPerSecondFlagName, videoPacketsReceivedPerSecondFlagDefault, videoPacketsReceivedPerSecondDescription)

	return nil
}

func registerWebRTCVideoReceiveStatsVideoPliCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelWebRTCVideoReceiveStatsFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, videoBytesReceivedAdded := retrieveWebRTCVideoReceiveStatsVideoBytesReceivedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoBytesReceivedAdded

	err, videoBytesReceivedPerSecondAdded := retrieveWebRTCVideoReceiveStatsVideoBytesReceivedPerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoBytesReceivedPerSecondAdded

	err, videoFirCountAdded := retrieveWebRTCVideoReceiveStatsVideoFirCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoFirCountAdded

	err, videoFractionLostAdded := retrieveWebRTCVideoReceiveStatsVideoFractionLostFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoFractionLostAdded

	err, videoFrameReceivedAdded := retrieveWebRTCVideoReceiveStatsVideoFrameReceivedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoFrameReceivedAdded

	err, videoFrameReceivedPerSecondAdded := retrieveWebRTCVideoReceiveStatsVideoFrameReceivedPerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoFrameReceivedPerSecondAdded

	err, videoNackCountAdded := retrieveWebRTCVideoReceiveStatsVideoNackCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoNackCountAdded

	err, videoPacketsLostAdded := retrieveWebRTCVideoReceiveStatsVideoPacketsLostFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoPacketsLostAdded

	err, videoPacketsReceivedAdded := retrieveWebRTCVideoReceiveStatsVideoPacketsReceivedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoPacketsReceivedAdded

	err, videoPacketsReceivedPerSecondAdded := retrieveWebRTCVideoReceiveStatsVideoPacketsReceivedPerSecondFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoPacketsReceivedPerSecondAdded

	err, videoPliCountAdded := retrieveWebRTCVideoReceiveStatsVideoPliCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || videoPliCountAdded

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoBytesReceivedFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoBytesReceivedFlagName := fmt.Sprintf("%v.videoBytesReceived", cmdPrefix)
	if cmd.Flags().Changed(videoBytesReceivedFlagName) {

		var videoBytesReceivedFlagName string
		if cmdPrefix == "" {
			videoBytesReceivedFlagName = "videoBytesReceived"
		} else {
			videoBytesReceivedFlagName = fmt.Sprintf("%v.videoBytesReceived", cmdPrefix)
		}

		videoBytesReceivedFlagValue, err := cmd.Flags().GetInt64(videoBytesReceivedFlagName)
		if err != nil {
			return err, false
		}
		m.VideoBytesReceived = videoBytesReceivedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoBytesReceivedPerSecondFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoBytesReceivedPerSecondFlagName := fmt.Sprintf("%v.videoBytesReceivedPerSecond", cmdPrefix)
	if cmd.Flags().Changed(videoBytesReceivedPerSecondFlagName) {

		var videoBytesReceivedPerSecondFlagName string
		if cmdPrefix == "" {
			videoBytesReceivedPerSecondFlagName = "videoBytesReceivedPerSecond"
		} else {
			videoBytesReceivedPerSecondFlagName = fmt.Sprintf("%v.videoBytesReceivedPerSecond", cmdPrefix)
		}

		videoBytesReceivedPerSecondFlagValue, err := cmd.Flags().GetInt64(videoBytesReceivedPerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.VideoBytesReceivedPerSecond = videoBytesReceivedPerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoFirCountFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveWebRTCVideoReceiveStatsVideoFractionLostFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoFractionLostFlagName := fmt.Sprintf("%v.videoFractionLost", cmdPrefix)
	if cmd.Flags().Changed(videoFractionLostFlagName) {

		var videoFractionLostFlagName string
		if cmdPrefix == "" {
			videoFractionLostFlagName = "videoFractionLost"
		} else {
			videoFractionLostFlagName = fmt.Sprintf("%v.videoFractionLost", cmdPrefix)
		}

		videoFractionLostFlagValue, err := cmd.Flags().GetFloat64(videoFractionLostFlagName)
		if err != nil {
			return err, false
		}
		m.VideoFractionLost = videoFractionLostFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoFrameReceivedFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoFrameReceivedFlagName := fmt.Sprintf("%v.videoFrameReceived", cmdPrefix)
	if cmd.Flags().Changed(videoFrameReceivedFlagName) {

		var videoFrameReceivedFlagName string
		if cmdPrefix == "" {
			videoFrameReceivedFlagName = "videoFrameReceived"
		} else {
			videoFrameReceivedFlagName = fmt.Sprintf("%v.videoFrameReceived", cmdPrefix)
		}

		videoFrameReceivedFlagValue, err := cmd.Flags().GetInt64(videoFrameReceivedFlagName)
		if err != nil {
			return err, false
		}
		m.VideoFrameReceived = videoFrameReceivedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoFrameReceivedPerSecondFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoFrameReceivedPerSecondFlagName := fmt.Sprintf("%v.videoFrameReceivedPerSecond", cmdPrefix)
	if cmd.Flags().Changed(videoFrameReceivedPerSecondFlagName) {

		var videoFrameReceivedPerSecondFlagName string
		if cmdPrefix == "" {
			videoFrameReceivedPerSecondFlagName = "videoFrameReceivedPerSecond"
		} else {
			videoFrameReceivedPerSecondFlagName = fmt.Sprintf("%v.videoFrameReceivedPerSecond", cmdPrefix)
		}

		videoFrameReceivedPerSecondFlagValue, err := cmd.Flags().GetInt64(videoFrameReceivedPerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.VideoFrameReceivedPerSecond = videoFrameReceivedPerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoNackCountFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveWebRTCVideoReceiveStatsVideoPacketsLostFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoPacketsLostFlagName := fmt.Sprintf("%v.videoPacketsLost", cmdPrefix)
	if cmd.Flags().Changed(videoPacketsLostFlagName) {

		var videoPacketsLostFlagName string
		if cmdPrefix == "" {
			videoPacketsLostFlagName = "videoPacketsLost"
		} else {
			videoPacketsLostFlagName = fmt.Sprintf("%v.videoPacketsLost", cmdPrefix)
		}

		videoPacketsLostFlagValue, err := cmd.Flags().GetInt32(videoPacketsLostFlagName)
		if err != nil {
			return err, false
		}
		m.VideoPacketsLost = videoPacketsLostFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoPacketsReceivedFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoPacketsReceivedFlagName := fmt.Sprintf("%v.videoPacketsReceived", cmdPrefix)
	if cmd.Flags().Changed(videoPacketsReceivedFlagName) {

		var videoPacketsReceivedFlagName string
		if cmdPrefix == "" {
			videoPacketsReceivedFlagName = "videoPacketsReceived"
		} else {
			videoPacketsReceivedFlagName = fmt.Sprintf("%v.videoPacketsReceived", cmdPrefix)
		}

		videoPacketsReceivedFlagValue, err := cmd.Flags().GetInt64(videoPacketsReceivedFlagName)
		if err != nil {
			return err, false
		}
		m.VideoPacketsReceived = videoPacketsReceivedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoPacketsReceivedPerSecondFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	videoPacketsReceivedPerSecondFlagName := fmt.Sprintf("%v.videoPacketsReceivedPerSecond", cmdPrefix)
	if cmd.Flags().Changed(videoPacketsReceivedPerSecondFlagName) {

		var videoPacketsReceivedPerSecondFlagName string
		if cmdPrefix == "" {
			videoPacketsReceivedPerSecondFlagName = "videoPacketsReceivedPerSecond"
		} else {
			videoPacketsReceivedPerSecondFlagName = fmt.Sprintf("%v.videoPacketsReceivedPerSecond", cmdPrefix)
		}

		videoPacketsReceivedPerSecondFlagValue, err := cmd.Flags().GetInt64(videoPacketsReceivedPerSecondFlagName)
		if err != nil {
			return err, false
		}
		m.VideoPacketsReceivedPerSecond = videoPacketsReceivedPerSecondFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCVideoReceiveStatsVideoPliCountFlags(depth int, m *models.WebRTCVideoReceiveStats, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
