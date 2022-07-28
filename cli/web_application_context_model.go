// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for WebApplicationContext

// register flags to command
func registerModelWebApplicationContextFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerWebApplicationContextApplicationName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextAutowireCapableBeanFactory(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextBeanDefinitionCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextBeanDefinitionNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextClassLoader(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextDisplayName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextEnvironment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextParentBeanFactory(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextServletContext(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerWebApplicationContextStartupDate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebApplicationContextApplicationName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	applicationNameDescription := ``

	var applicationNameFlagName string
	if cmdPrefix == "" {
		applicationNameFlagName = "applicationName"
	} else {
		applicationNameFlagName = fmt.Sprintf("%v.applicationName", cmdPrefix)
	}

	var applicationNameFlagDefault string

	_ = cmd.PersistentFlags().String(applicationNameFlagName, applicationNameFlagDefault, applicationNameDescription)

	return nil
}

func registerWebApplicationContextAutowireCapableBeanFactory(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: autowireCapableBeanFactory AutowireCapableBeanFactory map type is not supported by go-swagger cli yet

	return nil
}

func registerWebApplicationContextBeanDefinitionCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	beanDefinitionCountDescription := ``

	var beanDefinitionCountFlagName string
	if cmdPrefix == "" {
		beanDefinitionCountFlagName = "beanDefinitionCount"
	} else {
		beanDefinitionCountFlagName = fmt.Sprintf("%v.beanDefinitionCount", cmdPrefix)
	}

	var beanDefinitionCountFlagDefault int32

	_ = cmd.PersistentFlags().Int32(beanDefinitionCountFlagName, beanDefinitionCountFlagDefault, beanDefinitionCountDescription)

	return nil
}

func registerWebApplicationContextBeanDefinitionNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: beanDefinitionNames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerWebApplicationContextClassLoader(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerWebApplicationContextDisplayName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	displayNameDescription := ``

	var displayNameFlagName string
	if cmdPrefix == "" {
		displayNameFlagName = "displayName"
	} else {
		displayNameFlagName = fmt.Sprintf("%v.displayName", cmdPrefix)
	}

	var displayNameFlagDefault string

	_ = cmd.PersistentFlags().String(displayNameFlagName, displayNameFlagDefault, displayNameDescription)

	return nil
}

func registerWebApplicationContextEnvironment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var environmentFlagName string
	if cmdPrefix == "" {
		environmentFlagName = "environment"
	} else {
		environmentFlagName = fmt.Sprintf("%v.environment", cmdPrefix)
	}

	if err := registerModelEnvironmentFlags(depth+1, environmentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebApplicationContextID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerWebApplicationContextParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var parentFlagName string
	if cmdPrefix == "" {
		parentFlagName = "parent"
	} else {
		parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
	}

	if err := registerModelApplicationContextFlags(depth+1, parentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebApplicationContextParentBeanFactory(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: parentBeanFactory BeanFactory map type is not supported by go-swagger cli yet

	return nil
}

func registerWebApplicationContextServletContext(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var servletContextFlagName string
	if cmdPrefix == "" {
		servletContextFlagName = "servletContext"
	} else {
		servletContextFlagName = fmt.Sprintf("%v.servletContext", cmdPrefix)
	}

	if err := registerModelServletContextFlags(depth+1, servletContextFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerWebApplicationContextStartupDate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startupDateDescription := ``

	var startupDateFlagName string
	if cmdPrefix == "" {
		startupDateFlagName = "startupDate"
	} else {
		startupDateFlagName = fmt.Sprintf("%v.startupDate", cmdPrefix)
	}

	var startupDateFlagDefault int64

	_ = cmd.PersistentFlags().Int64(startupDateFlagName, startupDateFlagDefault, startupDateDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelWebApplicationContextFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, applicationNameAdded := retrieveWebApplicationContextApplicationNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || applicationNameAdded

	err, autowireCapableBeanFactoryAdded := retrieveWebApplicationContextAutowireCapableBeanFactoryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || autowireCapableBeanFactoryAdded

	err, beanDefinitionCountAdded := retrieveWebApplicationContextBeanDefinitionCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || beanDefinitionCountAdded

	err, beanDefinitionNamesAdded := retrieveWebApplicationContextBeanDefinitionNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || beanDefinitionNamesAdded

	err, classLoaderAdded := retrieveWebApplicationContextClassLoaderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classLoaderAdded

	err, displayNameAdded := retrieveWebApplicationContextDisplayNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || displayNameAdded

	err, environmentAdded := retrieveWebApplicationContextEnvironmentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || environmentAdded

	err, idAdded := retrieveWebApplicationContextIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, parentAdded := retrieveWebApplicationContextParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	err, parentBeanFactoryAdded := retrieveWebApplicationContextParentBeanFactoryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentBeanFactoryAdded

	err, servletContextAdded := retrieveWebApplicationContextServletContextFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || servletContextAdded

	err, startupDateAdded := retrieveWebApplicationContextStartupDateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startupDateAdded

	return nil, retAdded
}

func retrieveWebApplicationContextApplicationNameFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	applicationNameFlagName := fmt.Sprintf("%v.applicationName", cmdPrefix)
	if cmd.Flags().Changed(applicationNameFlagName) {

		var applicationNameFlagName string
		if cmdPrefix == "" {
			applicationNameFlagName = "applicationName"
		} else {
			applicationNameFlagName = fmt.Sprintf("%v.applicationName", cmdPrefix)
		}

		applicationNameFlagValue, err := cmd.Flags().GetString(applicationNameFlagName)
		if err != nil {
			return err, false
		}
		m.ApplicationName = applicationNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebApplicationContextAutowireCapableBeanFactoryFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	autowireCapableBeanFactoryFlagName := fmt.Sprintf("%v.autowireCapableBeanFactory", cmdPrefix)
	if cmd.Flags().Changed(autowireCapableBeanFactoryFlagName) {
		// warning: autowireCapableBeanFactory map type AutowireCapableBeanFactory is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveWebApplicationContextBeanDefinitionCountFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	beanDefinitionCountFlagName := fmt.Sprintf("%v.beanDefinitionCount", cmdPrefix)
	if cmd.Flags().Changed(beanDefinitionCountFlagName) {

		var beanDefinitionCountFlagName string
		if cmdPrefix == "" {
			beanDefinitionCountFlagName = "beanDefinitionCount"
		} else {
			beanDefinitionCountFlagName = fmt.Sprintf("%v.beanDefinitionCount", cmdPrefix)
		}

		beanDefinitionCountFlagValue, err := cmd.Flags().GetInt32(beanDefinitionCountFlagName)
		if err != nil {
			return err, false
		}
		m.BeanDefinitionCount = beanDefinitionCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebApplicationContextBeanDefinitionNamesFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	beanDefinitionNamesFlagName := fmt.Sprintf("%v.beanDefinitionNames", cmdPrefix)
	if cmd.Flags().Changed(beanDefinitionNamesFlagName) {
		// warning: beanDefinitionNames array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveWebApplicationContextClassLoaderFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveWebApplicationContextDisplayNameFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	displayNameFlagName := fmt.Sprintf("%v.displayName", cmdPrefix)
	if cmd.Flags().Changed(displayNameFlagName) {

		var displayNameFlagName string
		if cmdPrefix == "" {
			displayNameFlagName = "displayName"
		} else {
			displayNameFlagName = fmt.Sprintf("%v.displayName", cmdPrefix)
		}

		displayNameFlagValue, err := cmd.Flags().GetString(displayNameFlagName)
		if err != nil {
			return err, false
		}
		m.DisplayName = displayNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebApplicationContextEnvironmentFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	environmentFlagName := fmt.Sprintf("%v.environment", cmdPrefix)
	if cmd.Flags().Changed(environmentFlagName) {
		// info: complex object environment Environment is retrieved outside this Changed() block
	}
	environmentFlagValue := m.Environment
	if swag.IsZero(environmentFlagValue) {
		environmentFlagValue = &models.Environment{}
	}

	err, environmentAdded := retrieveModelEnvironmentFlags(depth+1, environmentFlagValue, environmentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || environmentAdded
	if environmentAdded {
		m.Environment = environmentFlagValue
	}

	return nil, retAdded
}

func retrieveWebApplicationContextIDFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

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

		retAdded = true
	}

	return nil, retAdded
}

func retrieveWebApplicationContextParentFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parentFlagName := fmt.Sprintf("%v.parent", cmdPrefix)
	if cmd.Flags().Changed(parentFlagName) {
		// info: complex object parent ApplicationContext is retrieved outside this Changed() block
	}
	parentFlagValue := m.Parent
	if swag.IsZero(parentFlagValue) {
		parentFlagValue = &models.ApplicationContext{}
	}

	err, parentAdded := retrieveModelApplicationContextFlags(depth+1, parentFlagValue, parentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded
	if parentAdded {
		m.Parent = parentFlagValue
	}

	return nil, retAdded
}

func retrieveWebApplicationContextParentBeanFactoryFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parentBeanFactoryFlagName := fmt.Sprintf("%v.parentBeanFactory", cmdPrefix)
	if cmd.Flags().Changed(parentBeanFactoryFlagName) {
		// warning: parentBeanFactory map type BeanFactory is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveWebApplicationContextServletContextFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	servletContextFlagName := fmt.Sprintf("%v.servletContext", cmdPrefix)
	if cmd.Flags().Changed(servletContextFlagName) {
		// info: complex object servletContext ServletContext is retrieved outside this Changed() block
	}
	servletContextFlagValue := m.ServletContext
	if swag.IsZero(servletContextFlagValue) {
		servletContextFlagValue = &models.ServletContext{}
	}

	err, servletContextAdded := retrieveModelServletContextFlags(depth+1, servletContextFlagValue, servletContextFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || servletContextAdded
	if servletContextAdded {
		m.ServletContext = servletContextFlagValue
	}

	return nil, retAdded
}

func retrieveWebApplicationContextStartupDateFlags(depth int, m *models.WebApplicationContext, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startupDateFlagName := fmt.Sprintf("%v.startupDate", cmdPrefix)
	if cmd.Flags().Changed(startupDateFlagName) {

		var startupDateFlagName string
		if cmdPrefix == "" {
			startupDateFlagName = "startupDate"
		} else {
			startupDateFlagName = fmt.Sprintf("%v.startupDate", cmdPrefix)
		}

		startupDateFlagValue, err := cmd.Flags().GetInt64(startupDateFlagName)
		if err != nil {
			return err, false
		}
		m.StartupDate = startupDateFlagValue

		retAdded = true
	}

	return nil, retAdded
}
