// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for IGlobalScope

// register flags to command
func registerModelIGlobalScopeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIGlobalScopeAttributeNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeClassLoader(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeClientConnections(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeClients(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeConnections(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeContext(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeContextPath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeDepth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeEventListeners(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeHandler(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopePath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeScopeNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeServer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeServiceHandlerNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeStatistics(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeStore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIGlobalScopeValid(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIGlobalScopeAttributeNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: attributeNames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: attributes map[string]interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeClassLoader(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var classLoaderFlagName string
	if cmdPrefix == "" {
		classLoaderFlagName = "classLoader"
	} else {
		classLoaderFlagName = fmt.Sprintf("%v.classLoader", cmdPrefix)
	}

	if err := registerModelClassLoaderFlags(depth+1, classLoaderFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIGlobalScopeClientConnections(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: clientConnections []*IConnection array type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeClients(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: clients []*IClient array type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeConnections(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: connections [][]*IConnection array type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeContext(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var contextFlagName string
	if cmdPrefix == "" {
		contextFlagName = "context"
	} else {
		contextFlagName = fmt.Sprintf("%v.context", cmdPrefix)
	}

	if err := registerModelIContextFlags(depth+1, contextFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIGlobalScopeContextPath(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	contextPathDescription := ``

	var contextPathFlagName string
	if cmdPrefix == "" {
		contextPathFlagName = "contextPath"
	} else {
		contextPathFlagName = fmt.Sprintf("%v.contextPath", cmdPrefix)
	}

	var contextPathFlagDefault string

	_ = cmd.PersistentFlags().String(contextPathFlagName, contextPathFlagDefault, contextPathDescription)

	return nil
}

func registerIGlobalScopeDepth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	depthDescription := ``

	var depthFlagName string
	if cmdPrefix == "" {
		depthFlagName = "depth"
	} else {
		depthFlagName = fmt.Sprintf("%v.depth", cmdPrefix)
	}

	var depthFlagDefault int32

	_ = cmd.PersistentFlags().Int32(depthFlagName, depthFlagDefault, depthDescription)

	return nil
}

func registerIGlobalScopeEventListeners(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: eventListeners []IEventListener array type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeHandler(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: handler IScopeHandler map type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerIGlobalScopeParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var parentFlagName string
	if cmdPrefix == "" {
		parentFlagName = "parent"
	} else {
		parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
	}

	if err := registerModelIScopeFlags(depth+1, parentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIGlobalScopePath(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pathDescription := ``

	var pathFlagName string
	if cmdPrefix == "" {
		pathFlagName = "path"
	} else {
		pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
	}

	var pathFlagDefault string

	_ = cmd.PersistentFlags().String(pathFlagName, pathFlagDefault, pathDescription)

	return nil
}

func registerIGlobalScopeScopeNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: scopeNames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeServer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var serverFlagName string
	if cmdPrefix == "" {
		serverFlagName = "server"
	} else {
		serverFlagName = fmt.Sprintf("%v.server", cmdPrefix)
	}

	if err := registerModelIServerFlags(depth+1, serverFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIGlobalScopeServiceHandlerNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: serviceHandlerNames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerIGlobalScopeStatistics(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var statisticsFlagName string
	if cmdPrefix == "" {
		statisticsFlagName = "statistics"
	} else {
		statisticsFlagName = fmt.Sprintf("%v.statistics", cmdPrefix)
	}

	if err := registerModelIScopeStatisticsFlags(depth+1, statisticsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIGlobalScopeStore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var storeFlagName string
	if cmdPrefix == "" {
		storeFlagName = "store"
	} else {
		storeFlagName = fmt.Sprintf("%v.store", cmdPrefix)
	}

	if err := registerModelIPersistenceStoreFlags(depth+1, storeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIGlobalScopeType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `Enum: ["UNDEFINED","GLOBAL","APPLICATION","ROOM","BROADCAST","SHARED_OBJECT"]. `

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
			if err := json.Unmarshal([]byte(`["UNDEFINED","GLOBAL","APPLICATION","ROOM","BROADCAST","SHARED_OBJECT"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

func registerIGlobalScopeValid(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	validDescription := ``

	var validFlagName string
	if cmdPrefix == "" {
		validFlagName = "valid"
	} else {
		validFlagName = fmt.Sprintf("%v.valid", cmdPrefix)
	}

	var validFlagDefault bool

	_ = cmd.PersistentFlags().Bool(validFlagName, validFlagDefault, validDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIGlobalScopeFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, attributeNamesAdded := retrieveIGlobalScopeAttributeNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributeNamesAdded

	err, attributesAdded := retrieveIGlobalScopeAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, classLoaderAdded := retrieveIGlobalScopeClassLoaderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classLoaderAdded

	err, clientConnectionsAdded := retrieveIGlobalScopeClientConnectionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clientConnectionsAdded

	err, clientsAdded := retrieveIGlobalScopeClientsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clientsAdded

	err, connectionsAdded := retrieveIGlobalScopeConnectionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionsAdded

	err, contextAdded := retrieveIGlobalScopeContextFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contextAdded

	err, contextPathAdded := retrieveIGlobalScopeContextPathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contextPathAdded

	err, depthAdded := retrieveIGlobalScopeDepthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || depthAdded

	err, eventListenersAdded := retrieveIGlobalScopeEventListenersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || eventListenersAdded

	err, handlerAdded := retrieveIGlobalScopeHandlerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || handlerAdded

	err, nameAdded := retrieveIGlobalScopeNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, parentAdded := retrieveIGlobalScopeParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	err, pathAdded := retrieveIGlobalScopePathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pathAdded

	err, scopeNamesAdded := retrieveIGlobalScopeScopeNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeNamesAdded

	err, serverAdded := retrieveIGlobalScopeServerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverAdded

	err, serviceHandlerNamesAdded := retrieveIGlobalScopeServiceHandlerNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serviceHandlerNamesAdded

	err, statisticsAdded := retrieveIGlobalScopeStatisticsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statisticsAdded

	err, storeAdded := retrieveIGlobalScopeStoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || storeAdded

	err, typeAdded := retrieveIGlobalScopeTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, validAdded := retrieveIGlobalScopeValidFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || validAdded

	return nil, retAdded
}

func retrieveIGlobalScopeAttributeNamesFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributeNamesFlagName := fmt.Sprintf("%v.attributeNames", cmdPrefix)
	if cmd.Flags().Changed(attributeNamesFlagName) {
		// warning: attributeNames array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeAttributesFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	attributesFlagName := fmt.Sprintf("%v.attributes", cmdPrefix)
	if cmd.Flags().Changed(attributesFlagName) {
		// warning: attributes map type map[string]interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeClassLoaderFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	classLoaderFlagName := fmt.Sprintf("%v.classLoader", cmdPrefix)
	if cmd.Flags().Changed(classLoaderFlagName) {
		// info: complex object classLoader ClassLoader is retrieved outside this Changed() block
	}
	classLoaderFlagValue := m.ClassLoader
	if swag.IsZero(classLoaderFlagValue) {
		classLoaderFlagValue = &models.ClassLoader{}
	}

	err, classLoaderAdded := retrieveModelClassLoaderFlags(depth+1, classLoaderFlagValue, classLoaderFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classLoaderAdded
	if classLoaderAdded {
		m.ClassLoader = classLoaderFlagValue
	}

	return nil, retAdded
}

func retrieveIGlobalScopeClientConnectionsFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clientConnectionsFlagName := fmt.Sprintf("%v.clientConnections", cmdPrefix)
	if cmd.Flags().Changed(clientConnectionsFlagName) {
		// warning: clientConnections array type []*IConnection is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeClientsFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clientsFlagName := fmt.Sprintf("%v.clients", cmdPrefix)
	if cmd.Flags().Changed(clientsFlagName) {
		// warning: clients array type []*IClient is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeConnectionsFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	connectionsFlagName := fmt.Sprintf("%v.connections", cmdPrefix)
	if cmd.Flags().Changed(connectionsFlagName) {
		// warning: connections array type [][]*IConnection is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeContextFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contextFlagName := fmt.Sprintf("%v.context", cmdPrefix)
	if cmd.Flags().Changed(contextFlagName) {
		// info: complex object context IContext is retrieved outside this Changed() block
	}
	contextFlagValue := m.Context
	if swag.IsZero(contextFlagValue) {
		contextFlagValue = &models.IContext{}
	}

	err, contextAdded := retrieveModelIContextFlags(depth+1, contextFlagValue, contextFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contextAdded
	if contextAdded {
		m.Context = contextFlagValue
	}

	return nil, retAdded
}

func retrieveIGlobalScopeContextPathFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contextPathFlagName := fmt.Sprintf("%v.contextPath", cmdPrefix)
	if cmd.Flags().Changed(contextPathFlagName) {

		var contextPathFlagName string
		if cmdPrefix == "" {
			contextPathFlagName = "contextPath"
		} else {
			contextPathFlagName = fmt.Sprintf("%v.contextPath", cmdPrefix)
		}

		contextPathFlagValue, err := cmd.Flags().GetString(contextPathFlagName)
		if err != nil {
			return err, false
		}
		m.ContextPath = contextPathFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIGlobalScopeDepthFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	depthFlagName := fmt.Sprintf("%v.depth", cmdPrefix)
	if cmd.Flags().Changed(depthFlagName) {

		var depthFlagName string
		if cmdPrefix == "" {
			depthFlagName = "depth"
		} else {
			depthFlagName = fmt.Sprintf("%v.depth", cmdPrefix)
		}

		depthFlagValue, err := cmd.Flags().GetInt32(depthFlagName)
		if err != nil {
			return err, false
		}
		m.Depth = depthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIGlobalScopeEventListenersFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	eventListenersFlagName := fmt.Sprintf("%v.eventListeners", cmdPrefix)
	if cmd.Flags().Changed(eventListenersFlagName) {
		// warning: eventListeners array type []IEventListener is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeHandlerFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	handlerFlagName := fmt.Sprintf("%v.handler", cmdPrefix)
	if cmd.Flags().Changed(handlerFlagName) {
		// warning: handler map type IScopeHandler is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeNameFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIGlobalScopeParentFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parentFlagName := fmt.Sprintf("%v.parent", cmdPrefix)
	if cmd.Flags().Changed(parentFlagName) {
		// info: complex object parent IScope is retrieved outside this Changed() block
	}
	parentFlagValue := m.Parent
	if swag.IsZero(parentFlagValue) {
		parentFlagValue = &models.IScope{}
	}

	err, parentAdded := retrieveModelIScopeFlags(depth+1, parentFlagValue, parentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded
	if parentAdded {
		m.Parent = parentFlagValue
	}

	return nil, retAdded
}

func retrieveIGlobalScopePathFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pathFlagName := fmt.Sprintf("%v.path", cmdPrefix)
	if cmd.Flags().Changed(pathFlagName) {

		var pathFlagName string
		if cmdPrefix == "" {
			pathFlagName = "path"
		} else {
			pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
		}

		pathFlagValue, err := cmd.Flags().GetString(pathFlagName)
		if err != nil {
			return err, false
		}
		m.Path = pathFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIGlobalScopeScopeNamesFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scopeNamesFlagName := fmt.Sprintf("%v.scopeNames", cmdPrefix)
	if cmd.Flags().Changed(scopeNamesFlagName) {
		// warning: scopeNames array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeServerFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverFlagName := fmt.Sprintf("%v.server", cmdPrefix)
	if cmd.Flags().Changed(serverFlagName) {
		// info: complex object server IServer is retrieved outside this Changed() block
	}
	serverFlagValue := m.Server
	if swag.IsZero(serverFlagValue) {
		serverFlagValue = &models.IServer{}
	}

	err, serverAdded := retrieveModelIServerFlags(depth+1, serverFlagValue, serverFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverAdded
	if serverAdded {
		m.Server = serverFlagValue
	}

	return nil, retAdded
}

func retrieveIGlobalScopeServiceHandlerNamesFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serviceHandlerNamesFlagName := fmt.Sprintf("%v.serviceHandlerNames", cmdPrefix)
	if cmd.Flags().Changed(serviceHandlerNamesFlagName) {
		// warning: serviceHandlerNames array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIGlobalScopeStatisticsFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statisticsFlagName := fmt.Sprintf("%v.statistics", cmdPrefix)
	if cmd.Flags().Changed(statisticsFlagName) {
		// info: complex object statistics IScopeStatistics is retrieved outside this Changed() block
	}
	statisticsFlagValue := m.Statistics
	if swag.IsZero(statisticsFlagValue) {
		statisticsFlagValue = &models.IScopeStatistics{}
	}

	err, statisticsAdded := retrieveModelIScopeStatisticsFlags(depth+1, statisticsFlagValue, statisticsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statisticsAdded
	if statisticsAdded {
		m.Statistics = statisticsFlagValue
	}

	return nil, retAdded
}

func retrieveIGlobalScopeStoreFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	storeFlagName := fmt.Sprintf("%v.store", cmdPrefix)
	if cmd.Flags().Changed(storeFlagName) {
		// info: complex object store IPersistenceStore is retrieved outside this Changed() block
	}
	storeFlagValue := m.Store
	if swag.IsZero(storeFlagValue) {
		storeFlagValue = &models.IPersistenceStore{}
	}

	err, storeAdded := retrieveModelIPersistenceStoreFlags(depth+1, storeFlagValue, storeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || storeAdded
	if storeAdded {
		m.Store = storeFlagValue
	}

	return nil, retAdded
}

func retrieveIGlobalScopeTypeFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveIGlobalScopeValidFlags(depth int, m *models.IGlobalScope, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	validFlagName := fmt.Sprintf("%v.valid", cmdPrefix)
	if cmd.Flags().Changed(validFlagName) {

		var validFlagName string
		if cmdPrefix == "" {
			validFlagName = "valid"
		} else {
			validFlagName = fmt.Sprintf("%v.valid", cmdPrefix)
		}

		validFlagValue, err := cmd.Flags().GetBool(validFlagName)
		if err != nil {
			return err, false
		}
		m.Valid = validFlagValue

		retAdded = true
	}

	return nil, retAdded
}
