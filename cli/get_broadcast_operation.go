// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/client/broadcast_rest_service"

	"github.com/spf13/cobra"
)

// makeOperationBroadcastRestServiceGetBroadcastCmd returns a cmd to handle operation getBroadcast
func makeOperationBroadcastRestServiceGetBroadcastCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getBroadcast",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetBroadcast,
	}

	if err := registerOperationBroadcastRestServiceGetBroadcastParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetBroadcast uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetBroadcast(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetBroadcastParams()
	if err, _ := retrieveOperationBroadcastRestServiceGetBroadcastIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetBroadcastResult(appCli.BroadcastRestService.GetBroadcast(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetBroadcastParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetBroadcastParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetBroadcastIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetBroadcastIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. id of the broadcast`

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

func retrieveOperationBroadcastRestServiceGetBroadcastIDFlag(m *broadcast_rest_service.GetBroadcastParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationBroadcastRestServiceGetBroadcastResult parses request result and return the string content
func parseOperationBroadcastRestServiceGetBroadcastResult(resp0 *broadcast_rest_service.GetBroadcastOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning getBroadcastOK is not supported

		// Non schema case: warning getBroadcastNotFound is not supported

		return "", respErr
	}

	// warning: non schema response getBroadcastOK is not supported by go-swagger cli yet.

	return "", nil
}
