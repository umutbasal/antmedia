// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for Notify

// register flags to command
func registerModelNotifyFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerNotifyAction(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyCall(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyConnectionParams(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyData(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyDataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyHeader(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyObject(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifySource(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifySourceType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyTimestamp(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyTransactionID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerNotifyType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerNotifyAction(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	actionDescription := ``

	var actionFlagName string
	if cmdPrefix == "" {
		actionFlagName = "action"
	} else {
		actionFlagName = fmt.Sprintf("%v.action", cmdPrefix)
	}

	var actionFlagDefault string

	_ = cmd.PersistentFlags().String(actionFlagName, actionFlagDefault, actionDescription)

	return nil
}

func registerNotifyCall(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var callFlagName string
	if cmdPrefix == "" {
		callFlagName = "call"
	} else {
		callFlagName = fmt.Sprintf("%v.call", cmdPrefix)
	}

	if err := registerModelIServiceCallFlags(depth+1, callFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNotifyConnectionParams(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: connectionParams map[string]interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerNotifyData(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataFlagName string
	if cmdPrefix == "" {
		dataFlagName = "data"
	} else {
		dataFlagName = fmt.Sprintf("%v.data", cmdPrefix)
	}

	if err := registerModelIoBufferFlags(depth+1, dataFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNotifyDataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive dataType strfmt.Base64 is not supported by go-swagger cli yet

	return nil
}

func registerNotifyHeader(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var headerFlagName string
	if cmdPrefix == "" {
		headerFlagName = "header"
	} else {
		headerFlagName = fmt.Sprintf("%v.header", cmdPrefix)
	}

	if err := registerModelHeaderFlags(depth+1, headerFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerNotifyObject(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: object interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerNotifySource(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: source IEventListener map type is not supported by go-swagger cli yet

	return nil
}

func registerNotifySourceType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: primitive sourceType strfmt.Base64 is not supported by go-swagger cli yet

	return nil
}

func registerNotifyTimestamp(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	timestampDescription := ``

	var timestampFlagName string
	if cmdPrefix == "" {
		timestampFlagName = "timestamp"
	} else {
		timestampFlagName = fmt.Sprintf("%v.timestamp", cmdPrefix)
	}

	var timestampFlagDefault int32

	_ = cmd.PersistentFlags().Int32(timestampFlagName, timestampFlagDefault, timestampDescription)

	return nil
}

func registerNotifyTransactionID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	transactionIdDescription := ``

	var transactionIdFlagName string
	if cmdPrefix == "" {
		transactionIdFlagName = "transactionId"
	} else {
		transactionIdFlagName = fmt.Sprintf("%v.transactionId", cmdPrefix)
	}

	var transactionIdFlagDefault int32

	_ = cmd.PersistentFlags().Int32(transactionIdFlagName, transactionIdFlagDefault, transactionIdDescription)

	return nil
}

func registerNotifyType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["SYSTEM","STATUS","SERVICE_CALL","SHARED_OBJECT","STREAM_ACTION","STREAM_CONTROL","STREAM_DATA","CLIENT","CLIENT_INVOKE","CLIENT_NOTIFY","SERVER"]. `

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	if err := cmd.RegisterFlagCompletionFunc(typeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["SYSTEM","STATUS","SERVICE_CALL","SHARED_OBJECT","STREAM_ACTION","STREAM_CONTROL","STREAM_DATA","CLIENT","CLIENT_INVOKE","CLIENT_NOTIFY","SERVER"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelNotifyFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, actionAdded := retrieveNotifyActionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || actionAdded

	err, callAdded := retrieveNotifyCallFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || callAdded

	err, connectionParamsAdded := retrieveNotifyConnectionParamsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionParamsAdded

	err, dataAdded := retrieveNotifyDataFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded

	err, dataTypeAdded := retrieveNotifyDataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataTypeAdded

	err, headerAdded := retrieveNotifyHeaderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || headerAdded

	err, objectAdded := retrieveNotifyObjectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || objectAdded

	err, sourceAdded := retrieveNotifySourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceAdded

	err, sourceTypeAdded := retrieveNotifySourceTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || sourceTypeAdded

	err, timestampAdded := retrieveNotifyTimestampFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || timestampAdded

	err, transactionIdAdded := retrieveNotifyTransactionIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || transactionIdAdded

	err, typeAdded := retrieveNotifyTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveNotifyActionFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	actionFlagName := fmt.Sprintf("%v.action", cmdPrefix)
	if cmd.Flags().Changed(actionFlagName) {

		var actionFlagName string
		if cmdPrefix == "" {
			actionFlagName = "action"
		} else {
			actionFlagName = fmt.Sprintf("%v.action", cmdPrefix)
		}

		actionFlagValue, err := cmd.Flags().GetString(actionFlagName)
		if err != nil {
			return err, false
		}
		m.Action = actionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNotifyCallFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	callFlagName := fmt.Sprintf("%v.call", cmdPrefix)
	if cmd.Flags().Changed(callFlagName) {
		// info: complex object call IServiceCall is retrieved outside this Changed() block
	}
	callFlagValue := m.Call
	if swag.IsZero(callFlagValue) {
		callFlagValue = &models.IServiceCall{}
	}

	err, callAdded := retrieveModelIServiceCallFlags(depth+1, callFlagValue, callFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || callAdded
	if callAdded {
		m.Call = callFlagValue
	}

	return nil, retAdded
}

func retrieveNotifyConnectionParamsFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectionParamsFlagName := fmt.Sprintf("%v.connectionParams", cmdPrefix)
	if cmd.Flags().Changed(connectionParamsFlagName) {
		// warning: connectionParams map type map[string]interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNotifyDataFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataFlagName := fmt.Sprintf("%v.data", cmdPrefix)
	if cmd.Flags().Changed(dataFlagName) {
		// info: complex object data IoBuffer is retrieved outside this Changed() block
	}
	dataFlagValue := m.Data
	if swag.IsZero(dataFlagValue) {
		dataFlagValue = &models.IoBuffer{}
	}

	err, dataAdded := retrieveModelIoBufferFlags(depth+1, dataFlagValue, dataFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataAdded
	if dataAdded {
		m.Data = dataFlagValue
	}

	return nil, retAdded
}

func retrieveNotifyDataTypeFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataTypeFlagName := fmt.Sprintf("%v.dataType", cmdPrefix)
	if cmd.Flags().Changed(dataTypeFlagName) {

		// warning: primitive dataType strfmt.Base64 is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNotifyHeaderFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	headerFlagName := fmt.Sprintf("%v.header", cmdPrefix)
	if cmd.Flags().Changed(headerFlagName) {
		// info: complex object header Header is retrieved outside this Changed() block
	}
	headerFlagValue := m.Header
	if swag.IsZero(headerFlagValue) {
		headerFlagValue = &models.Header{}
	}

	err, headerAdded := retrieveModelHeaderFlags(depth+1, headerFlagValue, headerFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || headerAdded
	if headerAdded {
		m.Header = headerFlagValue
	}

	return nil, retAdded
}

func retrieveNotifyObjectFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	objectFlagName := fmt.Sprintf("%v.object", cmdPrefix)
	if cmd.Flags().Changed(objectFlagName) {
		// warning: object map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNotifySourceFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sourceFlagName := fmt.Sprintf("%v.source", cmdPrefix)
	if cmd.Flags().Changed(sourceFlagName) {
		// warning: source map type IEventListener is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveNotifySourceTypeFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	sourceTypeFlagName := fmt.Sprintf("%v.sourceType", cmdPrefix)
	if cmd.Flags().Changed(sourceTypeFlagName) {

		// warning: primitive sourceType strfmt.Base64 is not supported by go-swagger cli yet

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNotifyTimestampFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	timestampFlagName := fmt.Sprintf("%v.timestamp", cmdPrefix)
	if cmd.Flags().Changed(timestampFlagName) {

		var timestampFlagName string
		if cmdPrefix == "" {
			timestampFlagName = "timestamp"
		} else {
			timestampFlagName = fmt.Sprintf("%v.timestamp", cmdPrefix)
		}

		timestampFlagValue, err := cmd.Flags().GetInt32(timestampFlagName)
		if err != nil {
			return err, false
		}
		m.Timestamp = timestampFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNotifyTransactionIDFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	transactionIdFlagName := fmt.Sprintf("%v.transactionId", cmdPrefix)
	if cmd.Flags().Changed(transactionIdFlagName) {

		var transactionIdFlagName string
		if cmdPrefix == "" {
			transactionIdFlagName = "transactionId"
		} else {
			transactionIdFlagName = fmt.Sprintf("%v.transactionId", cmdPrefix)
		}

		transactionIdFlagValue, err := cmd.Flags().GetInt32(transactionIdFlagName)
		if err != nil {
			return err, false
		}
		m.TransactionID = transactionIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveNotifyTypeFlags(depth int, m *models.Notify, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
