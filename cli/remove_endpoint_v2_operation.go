// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/client/broadcast_rest_service"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBroadcastRestServiceRemoveEndpointV2Cmd returns a cmd to handle operation removeEndpointV2
func makeOperationBroadcastRestServiceRemoveEndpointV2Cmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "removeEndpointV2",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceRemoveEndpointV2,
	}

	if err := registerOperationBroadcastRestServiceRemoveEndpointV2ParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceRemoveEndpointV2 uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceRemoveEndpointV2(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewRemoveEndpointV2Params()
	if err, _ := retrieveOperationBroadcastRestServiceRemoveEndpointV2EndpointServiceIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceRemoveEndpointV2IDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceRemoveEndpointV2ResolutionHeightFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceRemoveEndpointV2Result(appCli.BroadcastRestService.RemoveEndpointV2(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceRemoveEndpointV2ParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceRemoveEndpointV2ParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceRemoveEndpointV2EndpointServiceIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceRemoveEndpointV2IDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceRemoveEndpointV2ResolutionHeightParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceRemoveEndpointV2EndpointServiceIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	endpointServiceIdDescription := `Required. RTMP url of the endpoint that will be stopped.`

	var endpointServiceIdFlagName string
	if cmdPrefix == "" {
		endpointServiceIdFlagName = "endpointServiceId"
	} else {
		endpointServiceIdFlagName = fmt.Sprintf("%v.endpointServiceId", cmdPrefix)
	}

	var endpointServiceIdFlagDefault string

	_ = cmd.PersistentFlags().String(endpointServiceIdFlagName, endpointServiceIdFlagDefault, endpointServiceIdDescription)

	return nil
}
func registerOperationBroadcastRestServiceRemoveEndpointV2IDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	idDescription := `Required. Broadcast id`

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
func registerOperationBroadcastRestServiceRemoveEndpointV2ResolutionHeightParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	resolutionHeightDescription := `Required. Resolution specifier if endpoint has been added with resolution. Only applicable if user added RTMP endpoint with a resolution speficier. Otherwise won't work and won't remove the endpoint.`

	var resolutionHeightFlagName string
	if cmdPrefix == "" {
		resolutionHeightFlagName = "resolutionHeight"
	} else {
		resolutionHeightFlagName = fmt.Sprintf("%v.resolutionHeight", cmdPrefix)
	}

	var resolutionHeightFlagDefault int32

	_ = cmd.PersistentFlags().Int32(resolutionHeightFlagName, resolutionHeightFlagDefault, resolutionHeightDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceRemoveEndpointV2EndpointServiceIDFlag(m *broadcast_rest_service.RemoveEndpointV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("endpointServiceId") {

		var endpointServiceIdFlagName string
		if cmdPrefix == "" {
			endpointServiceIdFlagName = "endpointServiceId"
		} else {
			endpointServiceIdFlagName = fmt.Sprintf("%v.endpointServiceId", cmdPrefix)
		}

		endpointServiceIdFlagValue, err := cmd.Flags().GetString(endpointServiceIdFlagName)
		if err != nil {
			return err, false
		}
		m.EndpointServiceID = endpointServiceIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBroadcastRestServiceRemoveEndpointV2IDFlag(m *broadcast_rest_service.RemoveEndpointV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceRemoveEndpointV2ResolutionHeightFlag(m *broadcast_rest_service.RemoveEndpointV2Params, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("resolutionHeight") {

		var resolutionHeightFlagName string
		if cmdPrefix == "" {
			resolutionHeightFlagName = "resolutionHeight"
		} else {
			resolutionHeightFlagName = fmt.Sprintf("%v.resolutionHeight", cmdPrefix)
		}

		resolutionHeightFlagValue, err := cmd.Flags().GetInt32(resolutionHeightFlagName)
		if err != nil {
			return err, false
		}
		m.ResolutionHeight = resolutionHeightFlagValue

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceRemoveEndpointV2Result parses request result and return the string content
func parseOperationBroadcastRestServiceRemoveEndpointV2Result(resp0 *broadcast_rest_service.RemoveEndpointV2OK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.RemoveEndpointV2OK)
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
