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

// makeOperationManagementRestServiceDeleteApplicationCmd returns a cmd to handle operation deleteApplication
func makeOperationManagementRestServiceDeleteApplicationCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteApplication",
		Short: ``,
		RunE:  runOperationManagementRestServiceDeleteApplication,
	}

	if err := registerOperationManagementRestServiceDeleteApplicationParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationManagementRestServiceDeleteApplication uses cmd flags to call endpoint api
func runOperationManagementRestServiceDeleteApplication(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := management_rest_service.NewDeleteApplicationParams()
	if err, _ := retrieveOperationManagementRestServiceDeleteApplicationAppNameFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationManagementRestServiceDeleteApplicationDeleteDBFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationManagementRestServiceDeleteApplicationResult(appCli.ManagementRestService.DeleteApplication(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationManagementRestServiceDeleteApplicationParamFlags registers all flags needed to fill params
func registerOperationManagementRestServiceDeleteApplicationParamFlags(cmd *cobra.Command) error {
	if err := registerOperationManagementRestServiceDeleteApplicationAppNameParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationManagementRestServiceDeleteApplicationDeleteDBParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationManagementRestServiceDeleteApplicationAppNameParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	appNameDescription := `Required. Name of the application to delete`

	var appNameFlagName string
	if cmdPrefix == "" {
		appNameFlagName = "appName"
	} else {
		appNameFlagName = fmt.Sprintf("%v.appName", cmdPrefix)
	}

	var appNameFlagDefault string

	_ = cmd.PersistentFlags().String(appNameFlagName, appNameFlagDefault, appNameDescription)

	return nil
}
func registerOperationManagementRestServiceDeleteApplicationDeleteDBParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	deleteDBDescription := ``

	var deleteDBFlagName string
	if cmdPrefix == "" {
		deleteDBFlagName = "deleteDB"
	} else {
		deleteDBFlagName = fmt.Sprintf("%v.deleteDB", cmdPrefix)
	}

	var deleteDBFlagDefault bool

	_ = cmd.PersistentFlags().Bool(deleteDBFlagName, deleteDBFlagDefault, deleteDBDescription)

	return nil
}

func retrieveOperationManagementRestServiceDeleteApplicationAppNameFlag(m *management_rest_service.DeleteApplicationParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("appName") {

		var appNameFlagName string
		if cmdPrefix == "" {
			appNameFlagName = "appName"
		} else {
			appNameFlagName = fmt.Sprintf("%v.appName", cmdPrefix)
		}

		appNameFlagValue, err := cmd.Flags().GetString(appNameFlagName)
		if err != nil {
			return err, false
		}
		m.AppName = appNameFlagValue

	}
	return nil, retAdded
}
func retrieveOperationManagementRestServiceDeleteApplicationDeleteDBFlag(m *management_rest_service.DeleteApplicationParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("deleteDB") {

		var deleteDBFlagName string
		if cmdPrefix == "" {
			deleteDBFlagName = "deleteDB"
		} else {
			deleteDBFlagName = fmt.Sprintf("%v.deleteDB", cmdPrefix)
		}

		deleteDBFlagValue, err := cmd.Flags().GetBool(deleteDBFlagName)
		if err != nil {
			return err, false
		}
		m.DeleteDB = &deleteDBFlagValue

	}
	return nil, retAdded
}

// parseOperationManagementRestServiceDeleteApplicationResult parses request result and return the string content
func parseOperationManagementRestServiceDeleteApplicationResult(resp0 *management_rest_service.DeleteApplicationOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*management_rest_service.DeleteApplicationOK)
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
