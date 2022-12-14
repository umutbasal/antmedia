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

// Schema cli for IContext

// register flags to command
func registerModelIContextFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIContextApplicationContext(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIContextClassLoader(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIContextClientRegistry(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIContextGlobalScope(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIContextMappingStrategy(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIContextPersistanceStore(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIContextServiceInvoker(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIContextApplicationContext(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var applicationContextFlagName string
	if cmdPrefix == "" {
		applicationContextFlagName = "applicationContext"
	} else {
		applicationContextFlagName = fmt.Sprintf("%v.applicationContext", cmdPrefix)
	}

	if err := registerModelApplicationContextFlags(depth+1, applicationContextFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIContextClassLoader(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerIContextClientRegistry(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: clientRegistry IClientRegistry map type is not supported by go-swagger cli yet

	return nil
}

func registerIContextGlobalScope(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var globalScopeFlagName string
	if cmdPrefix == "" {
		globalScopeFlagName = "globalScope"
	} else {
		globalScopeFlagName = fmt.Sprintf("%v.globalScope", cmdPrefix)
	}

	if err := registerModelIGlobalScopeFlags(depth+1, globalScopeFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIContextMappingStrategy(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: mappingStrategy IMappingStrategy map type is not supported by go-swagger cli yet

	return nil
}

func registerIContextPersistanceStore(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var persistanceStoreFlagName string
	if cmdPrefix == "" {
		persistanceStoreFlagName = "persistanceStore"
	} else {
		persistanceStoreFlagName = fmt.Sprintf("%v.persistanceStore", cmdPrefix)
	}

	if err := registerModelIPersistenceStoreFlags(depth+1, persistanceStoreFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerIContextServiceInvoker(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: serviceInvoker IServiceInvoker map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIContextFlags(depth int, m *models.IContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, applicationContextAdded := retrieveIContextApplicationContextFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || applicationContextAdded

	err, classLoaderAdded := retrieveIContextClassLoaderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classLoaderAdded

	err, clientRegistryAdded := retrieveIContextClientRegistryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || clientRegistryAdded

	err, globalScopeAdded := retrieveIContextGlobalScopeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || globalScopeAdded

	err, mappingStrategyAdded := retrieveIContextMappingStrategyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mappingStrategyAdded

	err, persistanceStoreAdded := retrieveIContextPersistanceStoreFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || persistanceStoreAdded

	err, serviceInvokerAdded := retrieveIContextServiceInvokerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serviceInvokerAdded

	return nil, retAdded
}

func retrieveIContextApplicationContextFlags(depth int, m *models.IContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	applicationContextFlagName := fmt.Sprintf("%v.applicationContext", cmdPrefix)
	if cmd.Flags().Changed(applicationContextFlagName) {
		// info: complex object applicationContext ApplicationContext is retrieved outside this Changed() block
	}
	applicationContextFlagValue := m.ApplicationContext
	if swag.IsZero(applicationContextFlagValue) {
		applicationContextFlagValue = &models.ApplicationContext{}
	}

	err, applicationContextAdded := retrieveModelApplicationContextFlags(depth+1, applicationContextFlagValue, applicationContextFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || applicationContextAdded
	if applicationContextAdded {
		m.ApplicationContext = applicationContextFlagValue
	}

	return nil, retAdded
}

func retrieveIContextClassLoaderFlags(depth int, m *models.IContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveIContextClientRegistryFlags(depth int, m *models.IContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	clientRegistryFlagName := fmt.Sprintf("%v.clientRegistry", cmdPrefix)
	if cmd.Flags().Changed(clientRegistryFlagName) {
		// warning: clientRegistry map type IClientRegistry is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIContextGlobalScopeFlags(depth int, m *models.IContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	globalScopeFlagName := fmt.Sprintf("%v.globalScope", cmdPrefix)
	if cmd.Flags().Changed(globalScopeFlagName) {
		// info: complex object globalScope IGlobalScope is retrieved outside this Changed() block
	}
	globalScopeFlagValue := m.GlobalScope
	if swag.IsZero(globalScopeFlagValue) {
		globalScopeFlagValue = &models.IGlobalScope{}
	}

	err, globalScopeAdded := retrieveModelIGlobalScopeFlags(depth+1, globalScopeFlagValue, globalScopeFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || globalScopeAdded
	if globalScopeAdded {
		m.GlobalScope = globalScopeFlagValue
	}

	return nil, retAdded
}

func retrieveIContextMappingStrategyFlags(depth int, m *models.IContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mappingStrategyFlagName := fmt.Sprintf("%v.mappingStrategy", cmdPrefix)
	if cmd.Flags().Changed(mappingStrategyFlagName) {
		// warning: mappingStrategy map type IMappingStrategy is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIContextPersistanceStoreFlags(depth int, m *models.IContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	persistanceStoreFlagName := fmt.Sprintf("%v.persistanceStore", cmdPrefix)
	if cmd.Flags().Changed(persistanceStoreFlagName) {
		// info: complex object persistanceStore IPersistenceStore is retrieved outside this Changed() block
	}
	persistanceStoreFlagValue := m.PersistanceStore
	if swag.IsZero(persistanceStoreFlagValue) {
		persistanceStoreFlagValue = &models.IPersistenceStore{}
	}

	err, persistanceStoreAdded := retrieveModelIPersistenceStoreFlags(depth+1, persistanceStoreFlagValue, persistanceStoreFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || persistanceStoreAdded
	if persistanceStoreAdded {
		m.PersistanceStore = persistanceStoreFlagValue
	}

	return nil, retAdded
}

func retrieveIContextServiceInvokerFlags(depth int, m *models.IContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serviceInvokerFlagName := fmt.Sprintf("%v.serviceInvoker", cmdPrefix)
	if cmd.Flags().Changed(serviceInvokerFlagName) {
		// warning: serviceInvoker map type IServiceInvoker is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
