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

// makeOperationBroadcastRestServiceGetStreamInfoCmd returns a cmd to handle operation getStreamInfo
func makeOperationBroadcastRestServiceGetStreamInfoCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getStreamInfo",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetStreamInfo,
	}

	if err := registerOperationBroadcastRestServiceGetStreamInfoParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetStreamInfo uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetStreamInfo(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetStreamInfoParams()
	if err, _ := retrieveOperationBroadcastRestServiceGetStreamInfoIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetStreamInfoResult(appCli.BroadcastRestService.GetStreamInfo(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetStreamInfoParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetStreamInfoParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetStreamInfoIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetStreamInfoIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

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

func retrieveOperationBroadcastRestServiceGetStreamInfoIDFlag(m *broadcast_rest_service.GetStreamInfoParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBroadcastRestServiceGetStreamInfoResult parses request result and return the string content
func parseOperationBroadcastRestServiceGetStreamInfoResult(resp0 *broadcast_rest_service.GetStreamInfoOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.GetStreamInfoOK)
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
