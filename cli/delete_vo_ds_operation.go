// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/client/vo_d_rest_service"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVodRestServiceDeleteVoDsCmd returns a cmd to handle operation deleteVoDs
func makeOperationVodRestServiceDeleteVoDsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteVoDs",
		Short: ``,
		RunE:  runOperationVodRestServiceDeleteVoDs,
	}

	if err := registerOperationVodRestServiceDeleteVoDsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVodRestServiceDeleteVoDs uses cmd flags to call endpoint api
func runOperationVodRestServiceDeleteVoDs(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := vo_d_rest_service.NewDeleteVoDsParams()
	if err, _ := retrieveOperationVodRestServiceDeleteVoDsBodyFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVodRestServiceDeleteVoDsResult(appCli.VodRestService.DeleteVoDs(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationVodRestServiceDeleteVoDsParamFlags registers all flags needed to fill params
func registerOperationVodRestServiceDeleteVoDsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVodRestServiceDeleteVoDsBodyParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVodRestServiceDeleteVoDsBodyParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	bodyDescription := `Required. the ids of the VoD file`

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

func retrieveOperationVodRestServiceDeleteVoDsBodyFlag(m *vo_d_rest_service.DeleteVoDsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationVodRestServiceDeleteVoDsResult parses request result and return the string content
func parseOperationVodRestServiceDeleteVoDsResult(resp0 *vo_d_rest_service.DeleteVoDsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*vo_d_rest_service.DeleteVoDsOK)
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
