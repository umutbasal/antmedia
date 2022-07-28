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

// makeOperationManagementRestServiceGetLiveClientsSizeCmd returns a cmd to handle operation getLiveClientsSize
func makeOperationManagementRestServiceGetLiveClientsSizeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getLiveClientsSize",
		Short: ``,
		RunE:  runOperationManagementRestServiceGetLiveClientsSize,
	}

	if err := registerOperationManagementRestServiceGetLiveClientsSizeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceGetLiveClientsSize uses cmd flags to call endpoint api
func runOperationManagementRestServiceGetLiveClientsSize(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewGetLiveClientsSizeParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceGetLiveClientsSizeResult(appCli.ManagementRestService.GetLiveClientsSize(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceGetLiveClientsSizeParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceGetLiveClientsSizeParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationManagementRestServiceGetLiveClientsSizeResult parses request result and return the string content
func parseOperationManagementRestServiceGetLiveClientsSizeResult(resp0 *management_rest_service.GetLiveClientsSizeOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.GetLiveClientsSizeOK)
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
