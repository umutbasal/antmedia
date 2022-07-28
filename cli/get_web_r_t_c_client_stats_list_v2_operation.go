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

// makeOperationBroadcastRestServiceGetWebRTCClientStatsListV2Cmd returns a cmd to handle operation getWebRTCClientStatsListV2
func makeOperationBroadcastRestServiceGetWebRTCClientStatsListV2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getWebRTCClientStatsListV2",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetWebRTCClientStatsListV2,
	}

	if err := registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetWebRTCClientStatsListV2 uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetWebRTCClientStatsListV2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetWebRTCClientStatsListV2Params()
	if err, _ := retrieveOperationBroadcastRestServiceGetWebRTCClientStatsListV2OffsetFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetWebRTCClientStatsListV2SizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetWebRTCClientStatsListV2StreamIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetWebRTCClientStatsListV2Result(appCli.BroadcastRestService.GetWebRTCClientStatsListV2(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2ParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2OffsetParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2SizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2StreamIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2OffsetParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	offsetDescription := `Required. offset of the list`

	var offsetFlagName string
	if cmdPrefix == "" {
		offsetFlagName = "offset"
	} else {
		offsetFlagName = fmt.Sprintf("%v.offset", cmdPrefix)
	}

	var offsetFlagDefault int32

	_ = cmd.PersistentFlags().Int32(offsetFlagName, offsetFlagDefault, offsetDescription)

	return nil
}
func registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2SizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sizeDescription := `Required. Number of items that will be fetched`

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
	}

	var sizeFlagDefault int32

	_ = cmd.PersistentFlags().Int32(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}
func registerOperationBroadcastRestServiceGetWebRTCClientStatsListV2StreamIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	streamIdDescription := `Required. the id of the stream`

	var streamIdFlagName string
	if cmdPrefix == "" {
		streamIdFlagName = "stream_id"
	} else {
		streamIdFlagName = fmt.Sprintf("%v.stream_id", cmdPrefix)
	}

	var streamIdFlagDefault string

	_ = cmd.PersistentFlags().String(streamIdFlagName, streamIdFlagDefault, streamIdDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceGetWebRTCClientStatsListV2OffsetFlag(m *broadcast_rest_service.GetWebRTCClientStatsListV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("offset") {

		var offsetFlagName string
		if cmdPrefix == "" {
			offsetFlagName = "offset"
		} else {
			offsetFlagName = fmt.Sprintf("%v.offset", cmdPrefix)
		}

		offsetFlagValue, err := cmd.Flags().GetInt32(offsetFlagName)
		if err != nil {
			return err, false
		}
		m.Offset = offsetFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBroadcastRestServiceGetWebRTCClientStatsListV2SizeFlag(m *broadcast_rest_service.GetWebRTCClientStatsListV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("size") {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetInt32(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = sizeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBroadcastRestServiceGetWebRTCClientStatsListV2StreamIDFlag(m *broadcast_rest_service.GetWebRTCClientStatsListV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("stream_id") {

		var streamIdFlagName string
		if cmdPrefix == "" {
			streamIdFlagName = "stream_id"
		} else {
			streamIdFlagName = fmt.Sprintf("%v.stream_id", cmdPrefix)
		}

		streamIdFlagValue, err := cmd.Flags().GetString(streamIdFlagName)
		if err != nil {
			return err, false
		}
		m.StreamID = streamIdFlagValue

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceGetWebRTCClientStatsListV2Result parses request result and return the string content
func parseOperationBroadcastRestServiceGetWebRTCClientStatsListV2Result(resp0 *broadcast_rest_service.GetWebRTCClientStatsListV2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.GetWebRTCClientStatsListV2OK)
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
