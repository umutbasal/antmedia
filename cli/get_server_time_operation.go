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

// makeOperationManagementRestServiceGetServerTimeCmd returns a cmd to handle operation getServerTime
func makeOperationManagementRestServiceGetServerTimeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getServerTime",
		Short: ``,
		RunE:  runOperationManagementRestServiceGetServerTime,
	}

	if err := registerOperationManagementRestServiceGetServerTimeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceGetServerTime uses cmd flags to call endpoint api
func runOperationManagementRestServiceGetServerTime(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewGetServerTimeParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceGetServerTimeResult(appCli.ManagementRestService.GetServerTime(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceGetServerTimeParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceGetServerTimeParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationManagementRestServiceGetServerTimeResult parses request result and return the string content
func parseOperationManagementRestServiceGetServerTimeResult(resp0 *management_rest_service.GetServerTimeOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.GetServerTimeOK)
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
