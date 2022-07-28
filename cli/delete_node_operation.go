// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/client/operations"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationOperationsDeleteNodeCmd returns a cmd to handle operation deleteNode
func makeOperationOperationsDeleteNodeCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteNode",
		Short: ``,
		RunE:  runOperationOperationsDeleteNode,
	}

	if err := registerOperationOperationsDeleteNodeParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationOperationsDeleteNode uses cmd flags to call endpoint api
func runOperationOperationsDeleteNode(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := operations.NewDeleteNodeParams()
	if err, _ := retrieveOperationOperationsDeleteNodeIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationOperationsDeleteNodeResult(appCli.Operations.DeleteNode(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationOperationsDeleteNodeParamFlags registers all flags needed to fill params
func registerOperationOperationsDeleteNodeParamFlags(cmd *cobra.Command) error {
	if err := registerOperationOperationsDeleteNodeIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationOperationsDeleteNodeIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. `

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func retrieveOperationOperationsDeleteNodeIDFlag(m *operations.DeleteNodeParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("id") {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

	}
	return nil, retAdded
}

// parseOperationOperationsDeleteNodeResult parses request result and return the string content
func parseOperationOperationsDeleteNodeResult(resp0 *operations.DeleteNodeOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*operations.DeleteNodeOK)
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
