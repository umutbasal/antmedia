// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for AdminApplication

// register flags to command
func registerModelAdminApplicationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAdminApplicationApplicationInfo(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationApplications(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationAttributeNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationAttributes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationChildScopeNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationClientTTL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationClients(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationClusterNotifier(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationConnections(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationContext(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationDataStoreFactory(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationDepth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationGhostConnsCleanupPeriod(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationListeners(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationPath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationPlugins(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationRootScope(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationScheduledJobNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationScope(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationStreamPlaybackSecurity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationStreamPublishSecurity(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationTotalConnectionSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAdminApplicationTotalLiveStreamSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAdminApplicationApplicationInfo(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: applicationInfo []*ApplicationInfo array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationApplications(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: applications []string array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationAttributeNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: attributeNames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationAttributes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: attributes map[string]interface{} map type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationChildScopeNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: childScopeNames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationClientTTL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	clientTtlDescription := ``

	var clientTtlFlagName string
	if cmdPrefix == "" {
		clientTtlFlagName = "clientTTL"
	} else {
		clientTtlFlagName = fmt.Sprintf("%v.clientTTL", cmdPrefix)
	}

	var clientTtlFlagDefault int64

	_ = cmd.PersistentFlags().Int64(clientTtlFlagName, clientTtlFlagDefault, clientTtlDescription)

	return nil
}

func registerAdminApplicationClients(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: clients []*IClient array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationClusterNotifier(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var clusterNotifierFlagName string
	if cmdPrefix == "" {
		clusterNotifierFlagName = "clusterNotifier"
	} else {
		clusterNotifierFlagName = fmt.Sprintf("%v.clusterNotifier", cmdPrefix)
	}

	if err := registerModelIClusterNotifierFlags(depth+1, clusterNotifierFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerAdminApplicationConnections(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: connections [][]*IConnection array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationContext(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerAdminApplicationDataStoreFactory(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var dataStoreFactoryFlagName string
	if cmdPrefix == "" {
		dataStoreFactoryFlagName = "dataStoreFactory"
	} else {
		dataStoreFactoryFlagName = fmt.Sprintf("%v.dataStoreFactory", cmdPrefix)
	}

	if err := registerModelConsoleDataStoreFactoryFlags(depth+1, dataStoreFactoryFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerAdminApplicationDepth(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerAdminApplicationGhostConnsCleanupPeriod(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ghostConnsCleanupPeriodDescription := ``

	var ghostConnsCleanupPeriodFlagName string
	if cmdPrefix == "" {
		ghostConnsCleanupPeriodFlagName = "ghostConnsCleanupPeriod"
	} else {
		ghostConnsCleanupPeriodFlagName = fmt.Sprintf("%v.ghostConnsCleanupPeriod", cmdPrefix)
	}

	var ghostConnsCleanupPeriodFlagDefault int32

	_ = cmd.PersistentFlags().Int32(ghostConnsCleanupPeriodFlagName, ghostConnsCleanupPeriodFlagDefault, ghostConnsCleanupPeriodDescription)

	return nil
}

func registerAdminApplicationListeners(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: listeners []IApplication array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerAdminApplicationParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerAdminApplicationPath(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerAdminApplicationPlugins(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: plugins []*PluginDescriptor array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationRootScope(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var rootScopeFlagName string
	if cmdPrefix == "" {
		rootScopeFlagName = "rootScope"
	} else {
		rootScopeFlagName = fmt.Sprintf("%v.rootScope", cmdPrefix)
	}

	if err := registerModelIScopeFlags(depth+1, rootScopeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerAdminApplicationScheduledJobNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: scheduledJobNames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationScope(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var scopeFlagName string
	if cmdPrefix == "" {
		scopeFlagName = "scope"
	} else {
		scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
	}

	if err := registerModelIScopeFlags(depth+1, scopeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerAdminApplicationStreamPlaybackSecurity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: streamPlaybackSecurity []IStreamPlaybackSecurity array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationStreamPublishSecurity(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: streamPublishSecurity []IStreamPublishSecurity array type is not supported by go-swagger cli yet

	return nil
}

func registerAdminApplicationTotalConnectionSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalConnectionSizeDescription := ``

	var totalConnectionSizeFlagName string
	if cmdPrefix == "" {
		totalConnectionSizeFlagName = "totalConnectionSize"
	} else {
		totalConnectionSizeFlagName = fmt.Sprintf("%v.totalConnectionSize", cmdPrefix)
	}

	var totalConnectionSizeFlagDefault int32

	_ = cmd.PersistentFlags().Int32(totalConnectionSizeFlagName, totalConnectionSizeFlagDefault, totalConnectionSizeDescription)

	return nil
}

func registerAdminApplicationTotalLiveStreamSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalLiveStreamSizeDescription := ``

	var totalLiveStreamSizeFlagName string
	if cmdPrefix == "" {
		totalLiveStreamSizeFlagName = "totalLiveStreamSize"
	} else {
		totalLiveStreamSizeFlagName = fmt.Sprintf("%v.totalLiveStreamSize", cmdPrefix)
	}

	var totalLiveStreamSizeFlagDefault int32

	_ = cmd.PersistentFlags().Int32(totalLiveStreamSizeFlagName, totalLiveStreamSizeFlagDefault, totalLiveStreamSizeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAdminApplicationFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, applicationInfoAdded := retrieveAdminApplicationApplicationInfoFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || applicationInfoAdded

	err, applicationsAdded := retrieveAdminApplicationApplicationsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || applicationsAdded

	err, attributeNamesAdded := retrieveAdminApplicationAttributeNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributeNamesAdded

	err, attributesAdded := retrieveAdminApplicationAttributesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || attributesAdded

	err, childScopeNamesAdded := retrieveAdminApplicationChildScopeNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || childScopeNamesAdded

	err, clientTtlAdded := retrieveAdminApplicationClientTTLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clientTtlAdded

	err, clientsAdded := retrieveAdminApplicationClientsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clientsAdded

	err, clusterNotifierAdded := retrieveAdminApplicationClusterNotifierFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterNotifierAdded

	err, connectionsAdded := retrieveAdminApplicationConnectionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || connectionsAdded

	err, contextAdded := retrieveAdminApplicationContextFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contextAdded

	err, dataStoreFactoryAdded := retrieveAdminApplicationDataStoreFactoryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataStoreFactoryAdded

	err, depthAdded := retrieveAdminApplicationDepthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || depthAdded

	err, ghostConnsCleanupPeriodAdded := retrieveAdminApplicationGhostConnsCleanupPeriodFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ghostConnsCleanupPeriodAdded

	err, listenersAdded := retrieveAdminApplicationListenersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || listenersAdded

	err, nameAdded := retrieveAdminApplicationNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, parentAdded := retrieveAdminApplicationParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	err, pathAdded := retrieveAdminApplicationPathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pathAdded

	err, pluginsAdded := retrieveAdminApplicationPluginsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pluginsAdded

	err, rootScopeAdded := retrieveAdminApplicationRootScopeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rootScopeAdded

	err, scheduledJobNamesAdded := retrieveAdminApplicationScheduledJobNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scheduledJobNamesAdded

	err, scopeAdded := retrieveAdminApplicationScopeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeAdded

	err, streamPlaybackSecurityAdded := retrieveAdminApplicationStreamPlaybackSecurityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamPlaybackSecurityAdded

	err, streamPublishSecurityAdded := retrieveAdminApplicationStreamPublishSecurityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamPublishSecurityAdded

	err, totalConnectionSizeAdded := retrieveAdminApplicationTotalConnectionSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalConnectionSizeAdded

	err, totalLiveStreamSizeAdded := retrieveAdminApplicationTotalLiveStreamSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalLiveStreamSizeAdded

	return nil, retAdded
}

func retrieveAdminApplicationApplicationInfoFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	applicationInfoFlagName := fmt.Sprintf("%v.applicationInfo", cmdPrefix)
	if cmd.Flags().Changed(applicationInfoFlagName) {
		// warning: applicationInfo array type []*ApplicationInfo is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAdminApplicationApplicationsFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	applicationsFlagName := fmt.Sprintf("%v.applications", cmdPrefix)
	if cmd.Flags().Changed(applicationsFlagName) {
		// warning: applications array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAdminApplicationAttributeNamesFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationAttributesFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationChildScopeNamesFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	childScopeNamesFlagName := fmt.Sprintf("%v.childScopeNames", cmdPrefix)
	if cmd.Flags().Changed(childScopeNamesFlagName) {
		// warning: childScopeNames array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAdminApplicationClientTTLFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clientTtlFlagName := fmt.Sprintf("%v.clientTTL", cmdPrefix)
	if cmd.Flags().Changed(clientTtlFlagName) {

		var clientTtlFlagName string
		if cmdPrefix == "" {
			clientTtlFlagName = "clientTTL"
		} else {
			clientTtlFlagName = fmt.Sprintf("%v.clientTTL", cmdPrefix)
		}

		clientTtlFlagValue, err := cmd.Flags().GetInt64(clientTtlFlagName)
		if err != nil {
			return err, false
		}
		m.ClientTTL = clientTtlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAdminApplicationClientsFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationClusterNotifierFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clusterNotifierFlagName := fmt.Sprintf("%v.clusterNotifier", cmdPrefix)
	if cmd.Flags().Changed(clusterNotifierFlagName) {
		// info: complex object clusterNotifier IClusterNotifier is retrieved outside this Changed() block
	}
	clusterNotifierFlagValue := m.ClusterNotifier
	if swag.IsZero(clusterNotifierFlagValue) {
		clusterNotifierFlagValue = &models.IClusterNotifier{}
	}

	err, clusterNotifierAdded := retrieveModelIClusterNotifierFlags(depth+1, clusterNotifierFlagValue, clusterNotifierFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clusterNotifierAdded
	if clusterNotifierAdded {
		m.ClusterNotifier = clusterNotifierFlagValue
	}

	return nil, retAdded
}

func retrieveAdminApplicationConnectionsFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationContextFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationDataStoreFactoryFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dataStoreFactoryFlagName := fmt.Sprintf("%v.dataStoreFactory", cmdPrefix)
	if cmd.Flags().Changed(dataStoreFactoryFlagName) {
		// info: complex object dataStoreFactory ConsoleDataStoreFactory is retrieved outside this Changed() block
	}
	dataStoreFactoryFlagValue := m.DataStoreFactory
	if swag.IsZero(dataStoreFactoryFlagValue) {
		dataStoreFactoryFlagValue = &models.ConsoleDataStoreFactory{}
	}

	err, dataStoreFactoryAdded := retrieveModelConsoleDataStoreFactoryFlags(depth+1, dataStoreFactoryFlagValue, dataStoreFactoryFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dataStoreFactoryAdded
	if dataStoreFactoryAdded {
		m.DataStoreFactory = dataStoreFactoryFlagValue
	}

	return nil, retAdded
}

func retrieveAdminApplicationDepthFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationGhostConnsCleanupPeriodFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ghostConnsCleanupPeriodFlagName := fmt.Sprintf("%v.ghostConnsCleanupPeriod", cmdPrefix)
	if cmd.Flags().Changed(ghostConnsCleanupPeriodFlagName) {

		var ghostConnsCleanupPeriodFlagName string
		if cmdPrefix == "" {
			ghostConnsCleanupPeriodFlagName = "ghostConnsCleanupPeriod"
		} else {
			ghostConnsCleanupPeriodFlagName = fmt.Sprintf("%v.ghostConnsCleanupPeriod", cmdPrefix)
		}

		ghostConnsCleanupPeriodFlagValue, err := cmd.Flags().GetInt32(ghostConnsCleanupPeriodFlagName)
		if err != nil {
			return err, false
		}
		m.GhostConnsCleanupPeriod = ghostConnsCleanupPeriodFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAdminApplicationListenersFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	listenersFlagName := fmt.Sprintf("%v.listeners", cmdPrefix)
	if cmd.Flags().Changed(listenersFlagName) {
		// warning: listeners array type []IApplication is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAdminApplicationNameFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationParentFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationPathFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveAdminApplicationPluginsFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pluginsFlagName := fmt.Sprintf("%v.plugins", cmdPrefix)
	if cmd.Flags().Changed(pluginsFlagName) {
		// warning: plugins array type []*PluginDescriptor is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAdminApplicationRootScopeFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rootScopeFlagName := fmt.Sprintf("%v.rootScope", cmdPrefix)
	if cmd.Flags().Changed(rootScopeFlagName) {
		// info: complex object rootScope IScope is retrieved outside this Changed() block
	}
	rootScopeFlagValue := m.RootScope
	if swag.IsZero(rootScopeFlagValue) {
		rootScopeFlagValue = &models.IScope{}
	}

	err, rootScopeAdded := retrieveModelIScopeFlags(depth+1, rootScopeFlagValue, rootScopeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rootScopeAdded
	if rootScopeAdded {
		m.RootScope = rootScopeFlagValue
	}

	return nil, retAdded
}

func retrieveAdminApplicationScheduledJobNamesFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scheduledJobNamesFlagName := fmt.Sprintf("%v.scheduledJobNames", cmdPrefix)
	if cmd.Flags().Changed(scheduledJobNamesFlagName) {
		// warning: scheduledJobNames array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAdminApplicationScopeFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scopeFlagName := fmt.Sprintf("%v.scope", cmdPrefix)
	if cmd.Flags().Changed(scopeFlagName) {
		// info: complex object scope IScope is retrieved outside this Changed() block
	}
	scopeFlagValue := m.Scope
	if swag.IsZero(scopeFlagValue) {
		scopeFlagValue = &models.IScope{}
	}

	err, scopeAdded := retrieveModelIScopeFlags(depth+1, scopeFlagValue, scopeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeAdded
	if scopeAdded {
		m.Scope = scopeFlagValue
	}

	return nil, retAdded
}

func retrieveAdminApplicationStreamPlaybackSecurityFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamPlaybackSecurityFlagName := fmt.Sprintf("%v.streamPlaybackSecurity", cmdPrefix)
	if cmd.Flags().Changed(streamPlaybackSecurityFlagName) {
		// warning: streamPlaybackSecurity array type []IStreamPlaybackSecurity is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAdminApplicationStreamPublishSecurityFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamPublishSecurityFlagName := fmt.Sprintf("%v.streamPublishSecurity", cmdPrefix)
	if cmd.Flags().Changed(streamPublishSecurityFlagName) {
		// warning: streamPublishSecurity array type []IStreamPublishSecurity is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveAdminApplicationTotalConnectionSizeFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalConnectionSizeFlagName := fmt.Sprintf("%v.totalConnectionSize", cmdPrefix)
	if cmd.Flags().Changed(totalConnectionSizeFlagName) {

		var totalConnectionSizeFlagName string
		if cmdPrefix == "" {
			totalConnectionSizeFlagName = "totalConnectionSize"
		} else {
			totalConnectionSizeFlagName = fmt.Sprintf("%v.totalConnectionSize", cmdPrefix)
		}

		totalConnectionSizeFlagValue, err := cmd.Flags().GetInt32(totalConnectionSizeFlagName)
		if err != nil {
			return err, false
		}
		m.TotalConnectionSize = totalConnectionSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAdminApplicationTotalLiveStreamSizeFlags(depth int, m *models.AdminApplication, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalLiveStreamSizeFlagName := fmt.Sprintf("%v.totalLiveStreamSize", cmdPrefix)
	if cmd.Flags().Changed(totalLiveStreamSizeFlagName) {

		var totalLiveStreamSizeFlagName string
		if cmdPrefix == "" {
			totalLiveStreamSizeFlagName = "totalLiveStreamSize"
		} else {
			totalLiveStreamSizeFlagName = fmt.Sprintf("%v.totalLiveStreamSize", cmdPrefix)
		}

		totalLiveStreamSizeFlagValue, err := cmd.Flags().GetInt32(totalLiveStreamSizeFlagName)
		if err != nil {
			return err, false
		}
		m.TotalLiveStreamSize = totalLiveStreamSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
