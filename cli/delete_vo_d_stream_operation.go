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

// makeOperationOperationsDeleteVoDStreamCmd returns a cmd to handle operation deleteVoDStream
func makeOperationOperationsDeleteVoDStreamCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteVoDStream",
		Short: ``,
		RunE:  runOperationOperationsDeleteVoDStream,
	}

	if err := registerOperationOperationsDeleteVoDStreamParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationOperationsDeleteVoDStream uses cmd flags to call endpoint api
func runOperationOperationsDeleteVoDStream(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := operations.NewDeleteVoDStreamParams()
	if err, _ := retrieveOperationOperationsDeleteVoDStreamAppnameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationOperationsDeleteVoDStreamStreamNameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationOperationsDeleteVoDStreamResult(appCli.Operations.DeleteVoDStream(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationOperationsDeleteVoDStreamParamFlags registers all flags needed to fill params
func registerOperationOperationsDeleteVoDStreamParamFlags(cmd *cobra.Command) error {
	if err := registerOperationOperationsDeleteVoDStreamAppnameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationOperationsDeleteVoDStreamStreamNameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationOperationsDeleteVoDStreamAppnameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	appnameDescription := `Required. `

	var appnameFlagName string
	if cmdPrefix == "" {
		appnameFlagName = "appname"
	} else {
		appnameFlagName = fmt.Sprintf("%v.appname", cmdPrefix)
	}

	var appnameFlagDefault string

	_ = cmd.PersistentFlags().String(appnameFlagName, appnameFlagDefault, appnameDescription)

	return nil
}
func registerOperationOperationsDeleteVoDStreamStreamNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	streamNameDescription := ``

	var streamNameFlagName string
	if cmdPrefix == "" {
		streamNameFlagName = "streamName"
	} else {
		streamNameFlagName = fmt.Sprintf("%v.streamName", cmdPrefix)
	}

	var streamNameFlagDefault string

	_ = cmd.PersistentFlags().String(streamNameFlagName, streamNameFlagDefault, streamNameDescription)

	return nil
}

func retrieveOperationOperationsDeleteVoDStreamAppnameFlag(m *operations.DeleteVoDStreamParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("appname") {

		var appnameFlagName string
		if cmdPrefix == "" {
			appnameFlagName = "appname"
		} else {
			appnameFlagName = fmt.Sprintf("%v.appname", cmdPrefix)
		}

		appnameFlagValue, err := cmd.Flags().GetString(appnameFlagName)
		if err != nil {
			return err, false
		}
		m.Appname = appnameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationOperationsDeleteVoDStreamStreamNameFlag(m *operations.DeleteVoDStreamParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("streamName") {

		var streamNameFlagName string
		if cmdPrefix == "" {
			streamNameFlagName = "streamName"
		} else {
			streamNameFlagName = fmt.Sprintf("%v.streamName", cmdPrefix)
		}

		streamNameFlagValue, err := cmd.Flags().GetString(streamNameFlagName)
		if err != nil {
			return err, false
		}
		m.StreamName = &streamNameFlagValue

	}
	return nil, retAdded
}

// parseOperationOperationsDeleteVoDStreamResult parses request result and return the string content
func parseOperationOperationsDeleteVoDStreamResult(resp0 *operations.DeleteVoDStreamOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*operations.DeleteVoDStreamOK)
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
		msgStr := fmt.Sprintf("%v", resp0.Payload)
		return string(msgStr), nil
	}

	return "", nil
}
