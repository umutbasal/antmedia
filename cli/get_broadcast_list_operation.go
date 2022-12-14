// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/umutbasal/antmedia/client/broadcast_rest_service"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationBroadcastRestServiceGetBroadcastListCmd returns a cmd to handle operation getBroadcastList
func makeOperationBroadcastRestServiceGetBroadcastListCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getBroadcastList",
		Short: ``,
		RunE:  runOperationBroadcastRestServiceGetBroadcastList,
	}

	if err := registerOperationBroadcastRestServiceGetBroadcastListParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationBroadcastRestServiceGetBroadcastList uses cmd flags to call endpoint api
func runOperationBroadcastRestServiceGetBroadcastList(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := broadcast_rest_service.NewGetBroadcastListParams()
	if err, _ := retrieveOperationBroadcastRestServiceGetBroadcastListOffsetFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetBroadcastListOrderByFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetBroadcastListSearchFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetBroadcastListSizeFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetBroadcastListSortByFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationBroadcastRestServiceGetBroadcastListTypeByFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationBroadcastRestServiceGetBroadcastListResult(appCli.BroadcastRestService.GetBroadcastList(params))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationBroadcastRestServiceGetBroadcastListParamFlags registers all flags needed to fill params
func registerOperationBroadcastRestServiceGetBroadcastListParamFlags(cmd *cobra.Command) error {
	if err := registerOperationBroadcastRestServiceGetBroadcastListOffsetParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetBroadcastListOrderByParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetBroadcastListSearchParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetBroadcastListSizeParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetBroadcastListSortByParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationBroadcastRestServiceGetBroadcastListTypeByParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationBroadcastRestServiceGetBroadcastListOffsetParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBroadcastRestServiceGetBroadcastListOrderByParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	orderByDescription := `"asc" for Ascending, "desc" Descending order`

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
func registerOperationBroadcastRestServiceGetBroadcastListSearchParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBroadcastRestServiceGetBroadcastListSizeParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationBroadcastRestServiceGetBroadcastListSortByParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	sortByDescription := `Field to sort. Possible values are "name", "date", "status"`

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
func registerOperationBroadcastRestServiceGetBroadcastListTypeByParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	typeByDescription := `Type of the stream. Possible values are "liveStream", "ipCamera", "streamSource", "VoD"`

	var typeByFlagName string
	if cmdPrefix == "" {
		typeByFlagName = "type_by"
	} else {
		typeByFlagName = fmt.Sprintf("%v.type_by", cmdPrefix)
	}

	var typeByFlagDefault string

	_ = cmd.PersistentFlags().String(typeByFlagName, typeByFlagDefault, typeByDescription)

	return nil
}

func retrieveOperationBroadcastRestServiceGetBroadcastListOffsetFlag(m *broadcast_rest_service.GetBroadcastListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceGetBroadcastListOrderByFlag(m *broadcast_rest_service.GetBroadcastListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceGetBroadcastListSearchFlag(m *broadcast_rest_service.GetBroadcastListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceGetBroadcastListSizeFlag(m *broadcast_rest_service.GetBroadcastListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceGetBroadcastListSortByFlag(m *broadcast_rest_service.GetBroadcastListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationBroadcastRestServiceGetBroadcastListTypeByFlag(m *broadcast_rest_service.GetBroadcastListParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("type_by") {

		var typeByFlagName string
		if cmdPrefix == "" {
			typeByFlagName = "type_by"
		} else {
			typeByFlagName = fmt.Sprintf("%v.type_by", cmdPrefix)
		}

		typeByFlagValue, err := cmd.Flags().GetString(typeByFlagName)
		if err != nil {
			return err, false
		}
		m.TypeBy = &typeByFlagValue

	}
	return nil, retAdded
}

// parseOperationBroadcastRestServiceGetBroadcastListResult parses request result and return the string content
func parseOperationBroadcastRestServiceGetBroadcastListResult(resp0 *broadcast_rest_service.GetBroadcastListOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*broadcast_rest_service.GetBroadcastListOK)
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
