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

// makeOperationOperationsGetNodeCountCmd returns a cmd to handle operation getNodeCount
func makeOperationOperationsGetNodeCountCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getNodeCount",
		Short: ``,
		RunE:  runOperationOperationsGetNodeCount,
	}

	if err := registerOperationOperationsGetNodeCountParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationOperationsGetNodeCount uses cmd flags to call endpoint api
func runOperationOperationsGetNodeCount(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := operations.NewGetNodeCountParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationOperationsGetNodeCountResult(appCli.Operations.GetNodeCount(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationOperationsGetNodeCountParamFlags registers all flags needed to fill params
func registerOperationOperationsGetNodeCountParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationOperationsGetNodeCountResult parses request result and return the string content
func parseOperationOperationsGetNodeCountResult(resp0 *operations.GetNodeCountOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*operations.GetNodeCountOK)
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
