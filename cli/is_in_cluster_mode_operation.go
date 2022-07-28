// Code generated by go-swagger;

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

// makeOperationManagementRestServiceIsInClusterModeCmd returns a cmd to handle operation isInClusterMode
func makeOperationManagementRestServiceIsInClusterModeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "isInClusterMode",
		Short: ``,
		RunE:  runOperationManagementRestServiceIsInClusterMode,
	}

	if err := registerOperationManagementRestServiceIsInClusterModeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceIsInClusterMode uses cmd flags to call endpoint api
func runOperationManagementRestServiceIsInClusterMode(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewIsInClusterModeParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceIsInClusterModeResult(appCli.ManagementRestService.IsInClusterMode(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceIsInClusterModeParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceIsInClusterModeParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationManagementRestServiceIsInClusterModeResult parses request result and return the string content
func parseOperationManagementRestServiceIsInClusterModeResult(resp0 *management_rest_service.IsInClusterModeOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.IsInClusterModeOK)
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
