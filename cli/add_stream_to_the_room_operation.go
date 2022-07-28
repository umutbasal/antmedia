// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/umutbasal/antmedia/client/broadcast_rest_service"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBroadcastRestServiceAddStreamToTheRoomCmd returns a cmd to handle operation addStreamToTheRoom
func makeOperationBroadcastRestServiceAddStreamToTheRoomCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "addStreamToTheRoom",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceAddStreamToTheRoom,
	}

	if err := registerOperationBroadcastRestServiceAddStreamToTheRoomParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceAddStreamToTheRoom uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceAddStreamToTheRoom(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewAddStreamToTheRoomParams()
	if err, _ := retrieveOperationBroadcastRestServiceAddStreamToTheRoomRoomIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceAddStreamToTheRoomStreamIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceAddStreamToTheRoomResult(appCli.BroadcastRestService.AddStreamToTheRoom(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceAddStreamToTheRoomParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceAddStreamToTheRoomParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceAddStreamToTheRoomRoomIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceAddStreamToTheRoomStreamIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceAddStreamToTheRoomRoomIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBroadcastRestServiceAddStreamToTheRoomStreamIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	streamIdDescription := `Required. Stream id to add to the conference room`

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

func retrieveOperationBroadcastRestServiceAddStreamToTheRoomRoomIDFlag(m *broadcast_rest_service.AddStreamToTheRoomParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceAddStreamToTheRoomStreamIDFlag(m *broadcast_rest_service.AddStreamToTheRoomParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBroadcastRestServiceAddStreamToTheRoomResult parses request result and return the string content
func parseOperationBroadcastRestServiceAddStreamToTheRoomResult(resp0 *broadcast_rest_service.AddStreamToTheRoomOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.AddStreamToTheRoomOK)
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
