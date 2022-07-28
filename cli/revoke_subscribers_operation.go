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

// makeOperationBroadcastRestServiceRevokeSubscribersCmd returns a cmd to handle operation revokeSubscribers
func makeOperationBroadcastRestServiceRevokeSubscribersCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "revokeSubscribers",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceRevokeSubscribers,
	}

	if err := registerOperationBroadcastRestServiceRevokeSubscribersParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceRevokeSubscribers uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceRevokeSubscribers(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewRevokeSubscribersParams()
	if err, _ := retrieveOperationBroadcastRestServiceRevokeSubscribersIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceRevokeSubscribersResult(appCli.BroadcastRestService.RevokeSubscribers(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceRevokeSubscribersParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceRevokeSubscribersParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceRevokeSubscribersIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceRevokeSubscribersIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationBroadcastRestServiceRevokeSubscribersIDFlag(m *broadcast_rest_service.RevokeSubscribersParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBroadcastRestServiceRevokeSubscribersResult parses request result and return the string content
func parseOperationBroadcastRestServiceRevokeSubscribersResult(resp0 *broadcast_rest_service.RevokeSubscribersOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.RevokeSubscribersOK)
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
