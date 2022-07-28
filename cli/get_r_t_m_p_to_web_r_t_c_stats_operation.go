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

// makeOperationBroadcastRestServiceGetRTMPToWebRTCStatsCmd returns a cmd to handle operation getRTMPToWebRTCStats
func makeOperationBroadcastRestServiceGetRTMPToWebRTCStatsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getRTMPToWebRTCStats",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetRTMPToWebRTCStats,
	}

	if err := registerOperationBroadcastRestServiceGetRTMPToWebRTCStatsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetRTMPToWebRTCStats uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetRTMPToWebRTCStats(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetRTMPToWebRTCStatsParams()
	if err, _ := retrieveOperationBroadcastRestServiceGetRTMPToWebRTCStatsIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetRTMPToWebRTCStatsResult(appCli.BroadcastRestService.GetRTMPToWebRTCStats(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetRTMPToWebRTCStatsParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetRTMPToWebRTCStatsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetRTMPToWebRTCStatsIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetRTMPToWebRTCStatsIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. the id of the stream`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceGetRTMPToWebRTCStatsIDFlag(m *broadcast_rest_service.GetRTMPToWebRTCStatsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceGetRTMPToWebRTCStatsResult parses request result and return the string content
func parseOperationBroadcastRestServiceGetRTMPToWebRTCStatsResult(resp0 *broadcast_rest_service.GetRTMPToWebRTCStatsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.GetRTMPToWebRTCStatsOK)
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