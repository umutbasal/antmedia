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

// makeOperationBroadcastRestServiceDeleteBroadcastsCmd returns a cmd to handle operation deleteBroadcasts
func makeOperationBroadcastRestServiceDeleteBroadcastsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteBroadcasts",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceDeleteBroadcasts,
	}

	if err := registerOperationBroadcastRestServiceDeleteBroadcastsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceDeleteBroadcasts uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceDeleteBroadcasts(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewDeleteBroadcastsParams()
	if err, _ := retrieveOperationBroadcastRestServiceDeleteBroadcastsBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceDeleteBroadcastsResult(appCli.BroadcastRestService.DeleteBroadcasts(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceDeleteBroadcastsParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceDeleteBroadcastsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceDeleteBroadcastsBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceDeleteBroadcastsBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	bodyDescription := `Required.  Id of the broadcast`

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	var bodyFlagDefault []string

	_ = cmd.PersistentFlags().StringSlice(bodyFlagName, bodyFlagDefault, bodyDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceDeleteBroadcastsBodyFlag(m *broadcast_rest_service.DeleteBroadcastsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {

		var bodyFlagName string
		if cmdPrefix == "" {
			bodyFlagName = "body"
		} else {
			bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
		}

		bodyFlagValues, err := cmd.Flags().GetStringSlice(bodyFlagName)
		if err != nil {
			return err, false
		}
		m.Body = bodyFlagValues

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceDeleteBroadcastsResult parses request result and return the string content
func parseOperationBroadcastRestServiceDeleteBroadcastsResult(resp0 *broadcast_rest_service.DeleteBroadcastsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.DeleteBroadcastsOK)
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
