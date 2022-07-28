// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/umutbasal/antmedia/client/management_rest_service"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationManagementRestServiceGetSystemMemoryInfoCmd returns a cmd to handle operation getSystemMemoryInfo
func makeOperationManagementRestServiceGetSystemMemoryInfoCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSystemMemoryInfo",
		Short: ``,
		RunE:  runOperationManagementRestServiceGetSystemMemoryInfo,
	}

	if err := registerOperationManagementRestServiceGetSystemMemoryInfoParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceGetSystemMemoryInfo uses cmd flags to call endpoint api
func runOperationManagementRestServiceGetSystemMemoryInfo(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewGetSystemMemoryInfoParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceGetSystemMemoryInfoResult(appCli.ManagementRestService.GetSystemMemoryInfo(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceGetSystemMemoryInfoParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceGetSystemMemoryInfoParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationManagementRestServiceGetSystemMemoryInfoResult parses request result and return the string content
func parseOperationManagementRestServiceGetSystemMemoryInfoResult(resp0 *management_rest_service.GetSystemMemoryInfoOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.GetSystemMemoryInfoOK)
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
