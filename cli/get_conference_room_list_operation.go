// Code generated by go-swagger; DO NOT EDIT.

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

// makeOperationBroadcastRestServiceGetConferenceRoomListCmd returns a cmd to handle operation getConferenceRoomList
func makeOperationBroadcastRestServiceGetConferenceRoomListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getConferenceRoomList",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetConferenceRoomList,
	}

	if err := registerOperationBroadcastRestServiceGetConferenceRoomListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetConferenceRoomList uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetConferenceRoomList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetConferenceRoomListParams()
	if err, _ := retrieveOperationBroadcastRestServiceGetConferenceRoomListOffsetFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetConferenceRoomListOrderByFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetConferenceRoomListSearchFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetConferenceRoomListSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetConferenceRoomListSortByFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetConferenceRoomListResult(appCli.BroadcastRestService.GetConferenceRoomList(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetConferenceRoomListParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetConferenceRoomListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetConferenceRoomListOffsetParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetConferenceRoomListOrderByParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetConferenceRoomListSearchParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetConferenceRoomListSizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetConferenceRoomListSortByParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetConferenceRoomListOffsetParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	offsetDescription := `Required. This is the offset of the list, it is useful for pagination. If you want to use sort mechanism, we recommend using Mongo DB.`

	var offsetFlagName string
	if cmdPrefix == "" {
		offsetFlagName = "offset"
	} else {
		offsetFlagName = fmt.Sprintf("%v.offset", cmdPrefix)
	}

	var offsetFlagDefault int32

	_ = cmd.PersistentFlags().Int32(offsetFlagName, offsetFlagDefault, offsetDescription)

	return nil
}
func registerOperationBroadcastRestServiceGetConferenceRoomListOrderByParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	orderByDescription := `asc for Ascending, desc Descending order`

	var orderByFlagName string
	if cmdPrefix == "" {
		orderByFlagName = "order_by"
	} else {
		orderByFlagName = fmt.Sprintf("%v.order_by", cmdPrefix)
	}

	var orderByFlagDefault string

	_ = cmd.PersistentFlags().String(orderByFlagName, orderByFlagDefault, orderByDescription)

	return nil
}
func registerOperationBroadcastRestServiceGetConferenceRoomListSearchParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	searchDescription := `Search parameter, returns specific items that contains search string`

	var searchFlagName string
	if cmdPrefix == "" {
		searchFlagName = "search"
	} else {
		searchFlagName = fmt.Sprintf("%v.search", cmdPrefix)
	}

	var searchFlagDefault string

	_ = cmd.PersistentFlags().String(searchFlagName, searchFlagDefault, searchDescription)

	return nil
}
func registerOperationBroadcastRestServiceGetConferenceRoomListSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sizeDescription := `Required. Number of items that will be fetched. If there is not enough item in the datastore, returned list size may less then this value`

	var sizeFlagName string
	if cmdPrefix == "" {
		sizeFlagName = "size"
	} else {
		sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
	}

	var sizeFlagDefault int32

	_ = cmd.PersistentFlags().Int32(sizeFlagName, sizeFlagDefault, sizeDescription)

	return nil
}
func registerOperationBroadcastRestServiceGetConferenceRoomListSortByParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sortByDescription := `field to sort`

	var sortByFlagName string
	if cmdPrefix == "" {
		sortByFlagName = "sort_by"
	} else {
		sortByFlagName = fmt.Sprintf("%v.sort_by", cmdPrefix)
	}

	var sortByFlagDefault string

	_ = cmd.PersistentFlags().String(sortByFlagName, sortByFlagDefault, sortByDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceGetConferenceRoomListOffsetFlag(m *broadcast_rest_service.GetConferenceRoomListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("offset") {

		var offsetFlagName string
		if cmdPrefix == "" {
			offsetFlagName = "offset"
		} else {
			offsetFlagName = fmt.Sprintf("%v.offset", cmdPrefix)
		}

		offsetFlagValue, err := cmd.Flags().GetInt32(offsetFlagName)
		if err != nil {
			return err, false
		}
		m.Offset = offsetFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBroadcastRestServiceGetConferenceRoomListOrderByFlag(m *broadcast_rest_service.GetConferenceRoomListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("order_by") {

		var orderByFlagName string
		if cmdPrefix == "" {
			orderByFlagName = "order_by"
		} else {
			orderByFlagName = fmt.Sprintf("%v.order_by", cmdPrefix)
		}

		orderByFlagValue, err := cmd.Flags().GetString(orderByFlagName)
		if err != nil {
			return err, false
		}
		m.OrderBy = &orderByFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBroadcastRestServiceGetConferenceRoomListSearchFlag(m *broadcast_rest_service.GetConferenceRoomListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("search") {

		var searchFlagName string
		if cmdPrefix == "" {
			searchFlagName = "search"
		} else {
			searchFlagName = fmt.Sprintf("%v.search", cmdPrefix)
		}

		searchFlagValue, err := cmd.Flags().GetString(searchFlagName)
		if err != nil {
			return err, false
		}
		m.Search = &searchFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBroadcastRestServiceGetConferenceRoomListSizeFlag(m *broadcast_rest_service.GetConferenceRoomListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("size") {

		var sizeFlagName string
		if cmdPrefix == "" {
			sizeFlagName = "size"
		} else {
			sizeFlagName = fmt.Sprintf("%v.size", cmdPrefix)
		}

		sizeFlagValue, err := cmd.Flags().GetInt32(sizeFlagName)
		if err != nil {
			return err, false
		}
		m.Size = sizeFlagValue

	}
	return nil, retAdded
}
func retrieveOperationBroadcastRestServiceGetConferenceRoomListSortByFlag(m *broadcast_rest_service.GetConferenceRoomListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("sort_by") {

		var sortByFlagName string
		if cmdPrefix == "" {
			sortByFlagName = "sort_by"
		} else {
			sortByFlagName = fmt.Sprintf("%v.sort_by", cmdPrefix)
		}

		sortByFlagValue, err := cmd.Flags().GetString(sortByFlagName)
		if err != nil {
			return err, false
		}
		m.SortBy = &sortByFlagValue

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceGetConferenceRoomListResult parses request result and return the string content
func parseOperationBroadcastRestServiceGetConferenceRoomListResult(resp0 *broadcast_rest_service.GetConferenceRoomListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.GetConferenceRoomListOK)
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
