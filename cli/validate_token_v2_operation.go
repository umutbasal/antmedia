// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/client/broadcast_rest_service"
	"antmedia/models"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBroadcastRestServiceValidateTokenV2Cmd returns a cmd to handle operation validateTokenV2
func makeOperationBroadcastRestServiceValidateTokenV2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "validateTokenV2",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceValidateTokenV2,
	}

	if err := registerOperationBroadcastRestServiceValidateTokenV2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceValidateTokenV2 uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceValidateTokenV2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewValidateTokenV2Params()
	if err, _ := retrieveOperationBroadcastRestServiceValidateTokenV2BodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceValidateTokenV2Result(appCli.BroadcastRestService.ValidateTokenV2(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceValidateTokenV2ParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceValidateTokenV2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceValidateTokenV2BodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceValidateTokenV2BodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var bodyFlagName string
	if cmdPrefix == "" {
		bodyFlagName = "body"
	} else {
		bodyFlagName = fmt.Sprintf("%v.body", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(bodyFlagName, "", "Optional json string for [body]. Token to be validated")

	// add flags for body
	if err := registerModelTokenFlags(0, "token", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationBroadcastRestServiceValidateTokenV2BodyFlag(m *broadcast_rest_service.ValidateTokenV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("body") {
		// Read body string from cmd and unmarshal
		bodyValueStr, err := cmd.Flags().GetString("body")
		if err != nil {
			return err, false
		}

		bodyValue := models.Token{}
		if err := json.Unmarshal([]byte(bodyValueStr), &bodyValue); err != nil {
			return fmt.Errorf("cannot unmarshal body string in models.Token: %v", err), false
		}
		m.Body = &bodyValue
	}
	bodyValueModel := m.Body
	if swag.IsZero(bodyValueModel) {
		bodyValueModel = &models.Token{}
	}
	err, added := retrieveModelTokenFlags(0, bodyValueModel, "token", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Body = bodyValueModel
	}
	if dryRun && debug {

		bodyValueDebugBytes, err := json.Marshal(m.Body)
		if err != nil {
			return err, false
		}
		logDebugf("Body dry-run payload: %v", string(bodyValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationBroadcastRestServiceValidateTokenV2Result parses request result and return the string content
func parseOperationBroadcastRestServiceValidateTokenV2Result(resp0 *broadcast_rest_service.ValidateTokenV2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.ValidateTokenV2OK)
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
