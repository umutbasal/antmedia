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

// makeOperationManagementRestServiceGetAppLiveStreamsCmd returns a cmd to handle operation getAppLiveStreams
func makeOperationManagementRestServiceGetAppLiveStreamsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getAppLiveStreams",
		Short: ``,
		RunE:  runOperationManagementRestServiceGetAppLiveStreams,
	}

	if err := registerOperationManagementRestServiceGetAppLiveStreamsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceGetAppLiveStreams uses cmd flags to call endpoint api
func runOperationManagementRestServiceGetAppLiveStreams(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewGetAppLiveStreamsParams()
	if err, _ := retrieveOperationManagementRestServiceGetAppLiveStreamsAppnameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceGetAppLiveStreamsResult(appCli.ManagementRestService.GetAppLiveStreams(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceGetAppLiveStreamsParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceGetAppLiveStreamsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationManagementRestServiceGetAppLiveStreamsAppnameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationManagementRestServiceGetAppLiveStreamsAppnameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	appnameDescription := `Required. Application name`

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

func retrieveOperationManagementRestServiceGetAppLiveStreamsAppnameFlag(m *management_rest_service.GetAppLiveStreamsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationManagementRestServiceGetAppLiveStreamsResult parses request result and return the string content
func parseOperationManagementRestServiceGetAppLiveStreamsResult(resp0 *management_rest_service.GetAppLiveStreamsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.GetAppLiveStreamsOK)
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
