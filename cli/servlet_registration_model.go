// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for ServletRegistration

// register flags to command
func registerModelServletRegistrationFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServletRegistrationClassName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServletRegistrationInitParameters(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServletRegistrationMappings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServletRegistrationName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServletRegistrationRunAsRole(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServletRegistrationClassName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerServletRegistrationInitParameters(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: initParameters map[string]string map type is not supported by go-swagger cli yet

	return nil
}

func registerServletRegistrationMappings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: mappings []string array type is not supported by go-swagger cli yet

	return nil
}

func registerServletRegistrationName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerServletRegistrationRunAsRole(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	runAsRoleDescription := ``

	var runAsRoleFlagName string
	if cmdPrefix == "" {
		runAsRoleFlagName = "runAsRole"
	} else {
		runAsRoleFlagName = fmt.Sprintf("%v.runAsRole", cmdPrefix)
	}

	var runAsRoleFlagDefault string

	_ = cmd.PersistentFlags().String(runAsRoleFlagName, runAsRoleFlagDefault, runAsRoleDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServletRegistrationFlags(depth int, m *models.ServletRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, classNameAdded := retrieveServletRegistrationClassNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classNameAdded

	err, initParametersAdded := retrieveServletRegistrationInitParametersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || initParametersAdded

	err, mappingsAdded := retrieveServletRegistrationMappingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || mappingsAdded

	err, nameAdded := retrieveServletRegistrationNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, runAsRoleAdded := retrieveServletRegistrationRunAsRoleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || runAsRoleAdded

	return nil, retAdded
}

func retrieveServletRegistrationClassNameFlags(depth int, m *models.ServletRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveServletRegistrationInitParametersFlags(depth int, m *models.ServletRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveServletRegistrationMappingsFlags(depth int, m *models.ServletRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	mappingsFlagName := fmt.Sprintf("%v.mappings", cmdPrefix)
	if cmd.Flags().Changed(mappingsFlagName) {
		// warning: mappings array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveServletRegistrationNameFlags(depth int, m *models.ServletRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveServletRegistrationRunAsRoleFlags(depth int, m *models.ServletRegistration, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	runAsRoleFlagName := fmt.Sprintf("%v.runAsRole", cmdPrefix)
	if cmd.Flags().Changed(runAsRoleFlagName) {

		var runAsRoleFlagName string
		if cmdPrefix == "" {
			runAsRoleFlagName = "runAsRole"
		} else {
			runAsRoleFlagName = fmt.Sprintf("%v.runAsRole", cmdPrefix)
		}

		runAsRoleFlagValue, err := cmd.Flags().GetString(runAsRoleFlagName)
		if err != nil {
			return err, false
		}
		m.RunAsRole = runAsRoleFlagValue

		retAdded = true
	}

	return nil, retAdded
}
