// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for WebRTCViewerInfo

// register flags to command
func registerModelWebRTCViewerInfoFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWebRTCViewerInfoEdgeAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCViewerInfoStreamID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebRTCViewerInfoViewerID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebRTCViewerInfoEdgeAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	edgeAddressDescription := `IP address of the edge to which viewer is connected`

	var edgeAddressFlagName string
	if cmdPrefix == "" {
		edgeAddressFlagName = "edgeAddress"
	} else {
		edgeAddressFlagName = fmt.Sprintf("%v.edgeAddress", cmdPrefix)
	}

	var edgeAddressFlagDefault string

	_ = cmd.PersistentFlags().String(edgeAddressFlagName, edgeAddressFlagDefault, edgeAddressDescription)

	return nil
}

func registerWebRTCViewerInfoStreamID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	streamIdDescription := `stream id that viewer views`

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

func registerWebRTCViewerInfoViewerID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	viewerIdDescription := `the id of the viewer`

	var viewerIdFlagName string
	if cmdPrefix == "" {
		viewerIdFlagName = "viewerId"
	} else {
		viewerIdFlagName = fmt.Sprintf("%v.viewerId", cmdPrefix)
	}

	var viewerIdFlagDefault string

	_ = cmd.PersistentFlags().String(viewerIdFlagName, viewerIdFlagDefault, viewerIdDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelWebRTCViewerInfoFlags(depth int, m *models.WebRTCViewerInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, edgeAddressAdded := retrieveWebRTCViewerInfoEdgeAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || edgeAddressAdded

	err, streamIdAdded := retrieveWebRTCViewerInfoStreamIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamIdAdded

	err, viewerIdAdded := retrieveWebRTCViewerInfoViewerIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || viewerIdAdded

	return nil, retAdded
}

func retrieveWebRTCViewerInfoEdgeAddressFlags(depth int, m *models.WebRTCViewerInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	edgeAddressFlagName := fmt.Sprintf("%v.edgeAddress", cmdPrefix)
	if cmd.Flags().Changed(edgeAddressFlagName) {

		var edgeAddressFlagName string
		if cmdPrefix == "" {
			edgeAddressFlagName = "edgeAddress"
		} else {
			edgeAddressFlagName = fmt.Sprintf("%v.edgeAddress", cmdPrefix)
		}

		edgeAddressFlagValue, err := cmd.Flags().GetString(edgeAddressFlagName)
		if err != nil {
			return err, false
		}
		m.EdgeAddress = edgeAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebRTCViewerInfoStreamIDFlags(depth int, m *models.WebRTCViewerInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveWebRTCViewerInfoViewerIDFlags(depth int, m *models.WebRTCViewerInfo, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	viewerIdFlagName := fmt.Sprintf("%v.viewerId", cmdPrefix)
	if cmd.Flags().Changed(viewerIdFlagName) {

		var viewerIdFlagName string
		if cmdPrefix == "" {
			viewerIdFlagName = "viewerId"
		} else {
			viewerIdFlagName = fmt.Sprintf("%v.viewerId", cmdPrefix)
		}

		viewerIdFlagValue, err := cmd.Flags().GetString(viewerIdFlagName)
		if err != nil {
			return err, false
		}
		m.ViewerID = viewerIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}
