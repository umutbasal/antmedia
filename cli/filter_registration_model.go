// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for FilterRegistration

// register flags to command
func registerModelFilterRegistrationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerFilterRegistrationClassName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterRegistrationInitParameters(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterRegistrationName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterRegistrationServletNameMappings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerFilterRegistrationURLPatternMappings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerFilterRegistrationClassName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	classNameDescription := ``

	var classNameFlagName string
	if cmdPrefix == "" {
		classNameFlagName = "className"
	} else {
		classNameFlagName = fmt.Sprintf("%v.className", cmdPrefix)
	}

	var classNameFlagDefault string

	_ = cmd.PersistentFlags().String(classNameFlagName, classNameFlagDefault, classNameDescription)

	return nil
}

func registerFilterRegistrationInitParameters(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: initParameters map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerFilterRegistrationName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerFilterRegistrationServletNameMappings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: servletNameMappings []string array type is not supported by go-swagger cli yet

	return nil
}

func registerFilterRegistrationURLPatternMappings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: urlPatternMappings []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelFilterRegistrationFlags(depth int, m *models.FilterRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, classNameAdded := retrieveFilterRegistrationClassNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classNameAdded

	err, initParametersAdded := retrieveFilterRegistrationInitParametersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || initParametersAdded

	err, nameAdded := retrieveFilterRegistrationNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, servletNameMappingsAdded := retrieveFilterRegistrationServletNameMappingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || servletNameMappingsAdded

	err, urlPatternMappingsAdded := retrieveFilterRegistrationURLPatternMappingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || urlPatternMappingsAdded

	return nil, retAdded
}

func retrieveFilterRegistrationClassNameFlags(depth int, m *models.FilterRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	classNameFlagName := fmt.Sprintf("%v.className", cmdPrefix)
	if cmd.Flags().Changed(classNameFlagName) {

		var classNameFlagName string
		if cmdPrefix == "" {
			classNameFlagName = "className"
		} else {
			classNameFlagName = fmt.Sprintf("%v.className", cmdPrefix)
		}

		classNameFlagValue, err := cmd.Flags().GetString(classNameFlagName)
		if err != nil {
			return err, false
		}
		m.ClassName = classNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveFilterRegistrationInitParametersFlags(depth int, m *models.FilterRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	initParametersFlagName := fmt.Sprintf("%v.initParameters", cmdPrefix)
	if cmd.Flags().Changed(initParametersFlagName) {
		// warning: initParameters map type map[string]string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFilterRegistrationNameFlags(depth int, m *models.FilterRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveFilterRegistrationServletNameMappingsFlags(depth int, m *models.FilterRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	servletNameMappingsFlagName := fmt.Sprintf("%v.servletNameMappings", cmdPrefix)
	if cmd.Flags().Changed(servletNameMappingsFlagName) {
		// warning: servletNameMappings array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveFilterRegistrationURLPatternMappingsFlags(depth int, m *models.FilterRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	urlPatternMappingsFlagName := fmt.Sprintf("%v.urlPatternMappings", cmdPrefix)
	if cmd.Flags().Changed(urlPatternMappingsFlagName) {
		// warning: urlPatternMappings array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
