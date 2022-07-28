// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/umutbasal/antmedia/client/vo_d_rest_service"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationVodRestServiceSynchUserVodListCmd returns a cmd to handle operation synchUserVodList
func makeOperationVodRestServiceSynchUserVodListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "synchUserVodList",
		Short: `Notes here`,
		RunE:  runOperationVodRestServiceSynchUserVodList,
	}

	if err := registerOperationVodRestServiceSynchUserVodListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationVodRestServiceSynchUserVodList uses cmd flags to call endpoint api
func runOperationVodRestServiceSynchUserVodList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := vo_d_rest_service.NewSynchUserVodListParams()
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationVodRestServiceSynchUserVodListResult(appCli.VodRestService.SynchUserVodList(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationVodRestServiceSynchUserVodListParamFlags registers all flags needed to fill params
func registerOperationVodRestServiceSynchUserVodListParamFlags(cmd *cobra.Command) error {
	return nil
}

// parseOperationVodRestServiceSynchUserVodListResult parses request result and return the string content
func parseOperationVodRestServiceSynchUserVodListResult(resp0 *vo_d_rest_service.SynchUserVodListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*vo_d_rest_service.SynchUserVodListOK)
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
