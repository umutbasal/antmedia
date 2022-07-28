// Code generated by go-swagger;

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

// makeOperationBroadcastRestServiceDeleteStreamFromTheRoomCmd returns a cmd to handle operation deleteStreamFromTheRoom
func makeOperationBroadcastRestServiceDeleteStreamFromTheRoomCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteStreamFromTheRoom",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceDeleteStreamFromTheRoom,
	}

	if err := registerOperationBroadcastRestServiceDeleteStreamFromTheRoomParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceDeleteStreamFromTheRoom uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceDeleteStreamFromTheRoom(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewDeleteStreamFromTheRoomParams()
	if err, _ := retrieveOperationBroadcastRestServiceDeleteStreamFromTheRoomRoomIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceDeleteStreamFromTheRoomStreamIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceDeleteStreamFromTheRoomResult(appCli.BroadcastRestService.DeleteStreamFromTheRoom(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceDeleteStreamFromTheRoomParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceDeleteStreamFromTheRoomParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceDeleteStreamFromTheRoomRoomIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceDeleteStreamFromTheRoomStreamIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceDeleteStreamFromTheRoomRoomIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	roomIdDescription := `Required. Room id`

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
func registerOperationBroadcastRestServiceDeleteStreamFromTheRoomStreamIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	streamIdDescription := `Required. Stream id to delete from the conference room`

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

func retrieveOperationBroadcastRestServiceDeleteStreamFromTheRoomRoomIDFlag(m *broadcast_rest_service.DeleteStreamFromTheRoomParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceDeleteStreamFromTheRoomStreamIDFlag(m *broadcast_rest_service.DeleteStreamFromTheRoomParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("streamId") {

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

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceDeleteStreamFromTheRoomResult parses request result and return the string content
func parseOperationBroadcastRestServiceDeleteStreamFromTheRoomResult(resp0 *broadcast_rest_service.DeleteStreamFromTheRoomOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.DeleteStreamFromTheRoomOK)
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
