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

// makeOperationBroadcastRestServiceGetObjectDetectedTotalCmd returns a cmd to handle operation getObjectDetectedTotal
func makeOperationBroadcastRestServiceGetObjectDetectedTotalCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getObjectDetectedTotal",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetObjectDetectedTotal,
	}

	if err := registerOperationBroadcastRestServiceGetObjectDetectedTotalParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetObjectDetectedTotal uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetObjectDetectedTotal(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetObjectDetectedTotalParams()
	if err, _ := retrieveOperationBroadcastRestServiceGetObjectDetectedTotalIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetObjectDetectedTotalResult(appCli.BroadcastRestService.GetObjectDetectedTotal(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetObjectDetectedTotalParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetObjectDetectedTotalParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetObjectDetectedTotalIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetObjectDetectedTotalIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. id of the stream`

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

func retrieveOperationBroadcastRestServiceGetObjectDetectedTotalIDFlag(m *broadcast_rest_service.GetObjectDetectedTotalParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBroadcastRestServiceGetObjectDetectedTotalResult parses request result and return the string content
func parseOperationBroadcastRestServiceGetObjectDetectedTotalResult(resp0 *broadcast_rest_service.GetObjectDetectedTotalOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.GetObjectDetectedTotalOK)
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
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}