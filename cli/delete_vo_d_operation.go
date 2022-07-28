// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/client/vo_d_rest_service"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVodRestServiceDeleteVoDCmd returns a cmd to handle operation deleteVoD
func makeOperationVodRestServiceDeleteVoDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "deleteVoD",
		Short: ``,
		RunE:  runOperationVodRestServiceDeleteVoD,
	}

	if err := registerOperationVodRestServiceDeleteVoDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVodRestServiceDeleteVoD uses cmd flags to call endpoint api
func runOperationVodRestServiceDeleteVoD(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := vo_d_rest_service.NewDeleteVoDParams()
	if err, _ := retrieveOperationVodRestServiceDeleteVoDIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVodRestServiceDeleteVoDResult(appCli.VodRestService.DeleteVoD(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationVodRestServiceDeleteVoDParamFlags registers all flags needed to fill params
func registerOperationVodRestServiceDeleteVoDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationVodRestServiceDeleteVoDIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationVodRestServiceDeleteVoDIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. the id of the VoD file`

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

func retrieveOperationVodRestServiceDeleteVoDIDFlag(m *vo_d_rest_service.DeleteVoDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationVodRestServiceDeleteVoDResult parses request result and return the string content
func parseOperationVodRestServiceDeleteVoDResult(resp0 *vo_d_rest_service.DeleteVoDOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*vo_d_rest_service.DeleteVoDOK)
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