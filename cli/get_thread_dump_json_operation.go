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

// makeOperationManagementRestServiceGetThreadDumpJSONCmd returns a cmd to handle operation getThreadDumpJson
func makeOperationManagementRestServiceGetThreadDumpJSONCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getThreadDumpJSON",
		Short: ``,
		RunE:  runOperationManagementRestServiceGetThreadDumpJSON,
	}

	if err := registerOperationManagementRestServiceGetThreadDumpJSONParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceGetThreadDumpJSON uses cmd flags to call endpoint api
func runOperationManagementRestServiceGetThreadDumpJSON(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewGetThreadDumpJSONParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceGetThreadDumpJSONResult(appCli.ManagementRestService.GetThreadDumpJSON(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceGetThreadDumpJSONParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceGetThreadDumpJSONParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationManagementRestServiceGetThreadDumpJSONResult parses request result and return the string content
func parseOperationManagementRestServiceGetThreadDumpJSONResult(resp0 *management_rest_service.GetThreadDumpJSONOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.GetThreadDumpJSONOK)
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
