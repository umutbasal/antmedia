// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/client/management_rest_service"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationManagementRestServiceGetApplicationInfoCmd returns a cmd to handle operation getApplicationInfo
func makeOperationManagementRestServiceGetApplicationInfoCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getApplicationInfo",
		Short: ``,
		RunE:  runOperationManagementRestServiceGetApplicationInfo,
	}

	if err := registerOperationManagementRestServiceGetApplicationInfoParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceGetApplicationInfo uses cmd flags to call endpoint api
func runOperationManagementRestServiceGetApplicationInfo(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewGetApplicationInfoParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceGetApplicationInfoResult(appCli.ManagementRestService.GetApplicationInfo(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceGetApplicationInfoParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceGetApplicationInfoParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationManagementRestServiceGetApplicationInfoResult parses request result and return the string content
func parseOperationManagementRestServiceGetApplicationInfoResult(resp0 *management_rest_service.GetApplicationInfoOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.GetApplicationInfoOK)
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
