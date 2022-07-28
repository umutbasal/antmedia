// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/client/broadcast_rest_service"
	"fmt"

	"github.com/spf13/cobra"
)

// makeOperationBroadcastRestServiceGetConferenceRoomCmd returns a cmd to handle operation getConferenceRoom
func makeOperationBroadcastRestServiceGetConferenceRoomCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getConferenceRoom",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetConferenceRoom,
	}

	if err := registerOperationBroadcastRestServiceGetConferenceRoomParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetConferenceRoom uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetConferenceRoom(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetConferenceRoomParams()
	if err, _ := retrieveOperationBroadcastRestServiceGetConferenceRoomRoomIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetConferenceRoomResult(appCli.BroadcastRestService.GetConferenceRoom(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetConferenceRoomParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetConferenceRoomParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetConferenceRoomRoomIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetConferenceRoomRoomIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	roomIdDescription := `Required. id of the room`

	var roomIdFlagName string
	if cmdPrefix == "" {
		roomIdFlagName = "roomId"
	} else {
		roomIdFlagName = fmt.Sprintf("%v.roomId", cmdPrefix)
	}

	var roomIdFlagDefault string

	_ = cmd.PersistentFlags().String(roomIdFlagName, roomIdFlagDefault, roomIdDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceGetConferenceRoomRoomIDFlag(m *broadcast_rest_service.GetConferenceRoomParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("roomId") {

		var roomIdFlagName string
		if cmdPrefix == "" {
			roomIdFlagName = "roomId"
		} else {
			roomIdFlagName = fmt.Sprintf("%v.roomId", cmdPrefix)
		}

		roomIdFlagValue, err := cmd.Flags().GetString(roomIdFlagName)
		if err != nil {
			return err, false
		}
		m.RoomID = roomIdFlagValue

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceGetConferenceRoomResult parses request result and return the string content
func parseOperationBroadcastRestServiceGetConferenceRoomResult(resp0 *broadcast_rest_service.GetConferenceRoomOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning getConferenceRoomOK is not supported

		// Non schema case: warning getConferenceRoomNotFound is not supported

		return "", respErr
	}

	// warning: non schema response getConferenceRoomOK is not supported by go-swagger cli yet.

	return "", nil
}
