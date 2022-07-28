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

// makeOperationBroadcastRestServiceListSubscriberStatsV2Cmd returns a cmd to handle operation listSubscriberStatsV2
func makeOperationBroadcastRestServiceListSubscriberStatsV2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "listSubscriberStatsV2",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceListSubscriberStatsV2,
	}

	if err := registerOperationBroadcastRestServiceListSubscriberStatsV2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceListSubscriberStatsV2 uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceListSubscriberStatsV2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewListSubscriberStatsV2Params()
	if err, _ := retrieveOperationBroadcastRestServiceListSubscriberStatsV2IDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceListSubscriberStatsV2OffsetFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceListSubscriberStatsV2SizeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceListSubscriberStatsV2Result(appCli.BroadcastRestService.ListSubscriberStatsV2(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceListSubscriberStatsV2ParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceListSubscriberStatsV2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceListSubscriberStatsV2IDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceListSubscriberStatsV2OffsetParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceListSubscriberStatsV2SizeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceListSubscriberStatsV2IDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBroadcastRestServiceListSubscriberStatsV2OffsetParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	offsetDescription := `Required. the starting point of the list`

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
func registerOperationBroadcastRestServiceListSubscriberStatsV2SizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sizeDescription := `Required. size of the return list (max:50 )`

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

func retrieveOperationBroadcastRestServiceListSubscriberStatsV2IDFlag(m *broadcast_rest_service.ListSubscriberStatsV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceListSubscriberStatsV2OffsetFlag(m *broadcast_rest_service.ListSubscriberStatsV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceListSubscriberStatsV2SizeFlag(m *broadcast_rest_service.ListSubscriberStatsV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBroadcastRestServiceListSubscriberStatsV2Result parses request result and return the string content
func parseOperationBroadcastRestServiceListSubscriberStatsV2Result(resp0 *broadcast_rest_service.ListSubscriberStatsV2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.ListSubscriberStatsV2OK)
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
