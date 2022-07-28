// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/umutbasal/antmedia/client/operations"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationOperationsSendSupportRequestCmd returns a cmd to handle operation sendSupportRequest
func makeOperationOperationsSendSupportRequestCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "sendSupportRequest",
		Short: ``,
		RunE:  runOperationOperationsSendSupportRequest,
	}

	if err := registerOperationOperationsSendSupportRequestParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationOperationsSendSupportRequest uses cmd flags to call endpoint api
func runOperationOperationsSendSupportRequest(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := operations.NewSendSupportRequestParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationOperationsSendSupportRequestResult(appCli.Operations.SendSupportRequest(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationOperationsSendSupportRequestParamFlags registers all flags needed to fill params
func registerOperationOperationsSendSupportRequestParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationOperationsSendSupportRequestResult parses request result and return the string content
func parseOperationOperationsSendSupportRequestResult(resp0 *operations.SendSupportRequestOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*operations.SendSupportRequestOK)
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
