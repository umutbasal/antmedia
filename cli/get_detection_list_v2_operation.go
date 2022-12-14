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

// makeOperationBroadcastRestServiceGetDetectionListV2Cmd returns a cmd to handle operation getDetectionListV2
func makeOperationBroadcastRestServiceGetDetectionListV2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getDetectionListV2",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetDetectionListV2,
	}

	if err := registerOperationBroadcastRestServiceGetDetectionListV2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetDetectionListV2 uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetDetectionListV2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetDetectionListV2Params()
	if err, _ := retrieveOperationBroadcastRestServiceGetDetectionListV2IDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetDetectionListV2OffsetFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetDetectionListV2SizeFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetDetectionListV2Result(appCli.BroadcastRestService.GetDetectionListV2(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetDetectionListV2ParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetDetectionListV2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetDetectionListV2IDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetDetectionListV2OffsetParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetDetectionListV2SizeParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetDetectionListV2IDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBroadcastRestServiceGetDetectionListV2OffsetParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	offsetDescription := `Required. starting point of the list`

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
func registerOperationBroadcastRestServiceGetDetectionListV2SizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sizeDescription := `Required. total size of the return list`

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

func retrieveOperationBroadcastRestServiceGetDetectionListV2IDFlag(m *broadcast_rest_service.GetDetectionListV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceGetDetectionListV2OffsetFlag(m *broadcast_rest_service.GetDetectionListV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceGetDetectionListV2SizeFlag(m *broadcast_rest_service.GetDetectionListV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBroadcastRestServiceGetDetectionListV2Result parses request result and return the string content
func parseOperationBroadcastRestServiceGetDetectionListV2Result(resp0 *broadcast_rest_service.GetDetectionListV2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.GetDetectionListV2OK)
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
