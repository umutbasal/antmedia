// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/client/broadcast_rest_service"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBroadcastRestServiceStopPlayingCmd returns a cmd to handle operation stopPlaying
func makeOperationBroadcastRestServiceStopPlayingCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "stopPlaying",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceStopPlaying,
	}

	if err := registerOperationBroadcastRestServiceStopPlayingParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceStopPlaying uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceStopPlaying(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewStopPlayingParams()
	if err, _ := retrieveOperationBroadcastRestServiceStopPlayingWebrtcViewerIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceStopPlayingResult(appCli.BroadcastRestService.StopPlaying(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceStopPlayingParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceStopPlayingParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceStopPlayingWebrtcViewerIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceStopPlayingWebrtcViewerIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	webrtcViewerIdDescription := `Required. the id of the webrtc viewer.`

	var webrtcViewerIdFlagName string
	if cmdPrefix == "" {
		webrtcViewerIdFlagName = "webrtc-viewer-id"
	} else {
		webrtcViewerIdFlagName = fmt.Sprintf("%v.webrtc-viewer-id", cmdPrefix)
	}

	var webrtcViewerIdFlagDefault string

	_ = cmd.PersistentFlags().String(webrtcViewerIdFlagName, webrtcViewerIdFlagDefault, webrtcViewerIdDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceStopPlayingWebrtcViewerIDFlag(m *broadcast_rest_service.StopPlayingParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("webrtc-viewer-id") {

		var webrtcViewerIdFlagName string
		if cmdPrefix == "" {
			webrtcViewerIdFlagName = "webrtc-viewer-id"
		} else {
			webrtcViewerIdFlagName = fmt.Sprintf("%v.webrtc-viewer-id", cmdPrefix)
		}

		webrtcViewerIdFlagValue, err := cmd.Flags().GetString(webrtcViewerIdFlagName)
		if err != nil {
			return err, false
		}
		m.WebrtcViewerID = webrtcViewerIdFlagValue

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceStopPlayingResult parses request result and return the string content
func parseOperationBroadcastRestServiceStopPlayingResult(resp0 *broadcast_rest_service.StopPlayingOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.StopPlayingOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
