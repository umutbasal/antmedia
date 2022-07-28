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

// makeOperationManagementRestServiceGetSettingsCmd returns a cmd to handle operation getSettings
func makeOperationManagementRestServiceGetSettingsCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSettings",
		Short: ``,
		RunE:  runOperationManagementRestServiceGetSettings,
	}

	if err := registerOperationManagementRestServiceGetSettingsParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceGetSettings uses cmd flags to call endpoint api
func runOperationManagementRestServiceGetSettings(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewGetSettingsParams()
	if err, _ := retrieveOperationManagementRestServiceGetSettingsAppnameFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceGetSettingsResult(appCli.ManagementRestService.GetSettings(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceGetSettingsParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceGetSettingsParamFlags(cmd *cobra.Command) error {
	if err := registerOperationManagementRestServiceGetSettingsAppnameParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationManagementRestServiceGetSettingsAppnameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationManagementRestServiceGetSettingsAppnameFlag(m *management_rest_service.GetSettingsParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationManagementRestServiceGetSettingsResult parses request result and return the string content
func parseOperationManagementRestServiceGetSettingsResult(resp0 *management_rest_service.GetSettingsOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.GetSettingsOK)
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
