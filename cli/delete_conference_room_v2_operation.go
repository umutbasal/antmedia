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

// makeOperationBroadcastRestServiceDeleteConferenceRoomV2Cmd returns a cmd to handle operation deleteConferenceRoomV2
func makeOperationBroadcastRestServiceDeleteConferenceRoomV2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteConferenceRoomV2",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceDeleteConferenceRoomV2,
	}

	if err := registerOperationBroadcastRestServiceDeleteConferenceRoomV2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceDeleteConferenceRoomV2 uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceDeleteConferenceRoomV2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewDeleteConferenceRoomV2Params()
	if err, _ := retrieveOperationBroadcastRestServiceDeleteConferenceRoomV2RoomIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceDeleteConferenceRoomV2Result(appCli.BroadcastRestService.DeleteConferenceRoomV2(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceDeleteConferenceRoomV2ParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceDeleteConferenceRoomV2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceDeleteConferenceRoomV2RoomIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceDeleteConferenceRoomV2RoomIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	roomIdDescription := `Required. the id of the conference room`

	var roomIdFlagName string
	if cmdPrefix == "" {
		roomIdFlagName = "room_id"
	} else {
		roomIdFlagName = fmt.Sprintf("%v.room_id", cmdPrefix)
	}

	var roomIdFlagDefault string

	_ = cmd.PersistentFlags().String(roomIdFlagName, roomIdFlagDefault, roomIdDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceDeleteConferenceRoomV2RoomIDFlag(m *broadcast_rest_service.DeleteConferenceRoomV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("room_id") {

		var roomIdFlagName string
		if cmdPrefix == "" {
			roomIdFlagName = "room_id"
		} else {
			roomIdFlagName = fmt.Sprintf("%v.room_id", cmdPrefix)
		}

		roomIdFlagValue, err := cmd.Flags().GetString(roomIdFlagName)
		if err != nil {
			return err, false
		}
		m.RoomID = roomIdFlagValue

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceDeleteConferenceRoomV2Result parses request result and return the string content
func parseOperationBroadcastRestServiceDeleteConferenceRoomV2Result(resp0 *broadcast_rest_service.DeleteConferenceRoomV2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.DeleteConferenceRoomV2OK)
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
