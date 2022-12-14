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

// Schema cli for ClassLoader

// register flags to command
func registerModelClassLoaderFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerClassLoaderDefinedPackages(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClassLoaderName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClassLoaderParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClassLoaderRegisteredAsParallelCapable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClassLoaderUnnamedModule(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerClassLoaderDefinedPackages(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: definedPackages []*Package array type is not supported by go-swagger cli yet

	return nil
}

func registerClassLoaderName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerClassLoaderParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var parentFlagName string
	if cmdPrefix == "" {
		parentFlagName = "parent"
	} else {
		parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
	}

	if err := registerModelClassLoaderFlags(depth+1, parentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerClassLoaderRegisteredAsParallelCapable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	registeredAsParallelCapableDescription := ``

	var registeredAsParallelCapableFlagName string
	if cmdPrefix == "" {
		registeredAsParallelCapableFlagName = "registeredAsParallelCapable"
	} else {
		registeredAsParallelCapableFlagName = fmt.Sprintf("%v.registeredAsParallelCapable", cmdPrefix)
	}

	var registeredAsParallelCapableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(registeredAsParallelCapableFlagName, registeredAsParallelCapableFlagDefault, registeredAsParallelCapableDescription)

	return nil
}

func registerClassLoaderUnnamedModule(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var unnamedModuleFlagName string
	if cmdPrefix == "" {
		unnamedModuleFlagName = "unnamedModule"
	} else {
		unnamedModuleFlagName = fmt.Sprintf("%v.unnamedModule", cmdPrefix)
	}

	if err := registerModelModuleFlags(depth+1, unnamedModuleFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelClassLoaderFlags(depth int, m *models.ClassLoader, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, definedPackagesAdded := retrieveClassLoaderDefinedPackagesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || definedPackagesAdded

	err, nameAdded := retrieveClassLoaderNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, parentAdded := retrieveClassLoaderParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	err, registeredAsParallelCapableAdded := retrieveClassLoaderRegisteredAsParallelCapableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || registeredAsParallelCapableAdded

	err, unnamedModuleAdded := retrieveClassLoaderUnnamedModuleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unnamedModuleAdded

	return nil, retAdded
}

func retrieveClassLoaderDefinedPackagesFlags(depth int, m *models.ClassLoader, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	definedPackagesFlagName := fmt.Sprintf("%v.definedPackages", cmdPrefix)
	if cmd.Flags().Changed(definedPackagesFlagName) {
		// warning: definedPackages array type []*Package is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveClassLoaderNameFlags(depth int, m *models.ClassLoader, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveClassLoaderParentFlags(depth int, m *models.ClassLoader, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parentFlagName := fmt.Sprintf("%v.parent", cmdPrefix)
	if cmd.Flags().Changed(parentFlagName) {
		// info: complex object parent ClassLoader is retrieved outside this Changed() block
	}
	parentFlagValue := m.Parent
	if swag.IsZero(parentFlagValue) {
		parentFlagValue = &models.ClassLoader{}
	}

	err, parentAdded := retrieveModelClassLoaderFlags(depth+1, parentFlagValue, parentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded
	if parentAdded {
		m.Parent = parentFlagValue
	}

	return nil, retAdded
}

func retrieveClassLoaderRegisteredAsParallelCapableFlags(depth int, m *models.ClassLoader, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	registeredAsParallelCapableFlagName := fmt.Sprintf("%v.registeredAsParallelCapable", cmdPrefix)
	if cmd.Flags().Changed(registeredAsParallelCapableFlagName) {

		var registeredAsParallelCapableFlagName string
		if cmdPrefix == "" {
			registeredAsParallelCapableFlagName = "registeredAsParallelCapable"
		} else {
			registeredAsParallelCapableFlagName = fmt.Sprintf("%v.registeredAsParallelCapable", cmdPrefix)
		}

		registeredAsParallelCapableFlagValue, err := cmd.Flags().GetBool(registeredAsParallelCapableFlagName)
		if err != nil {
			return err, false
		}
		m.RegisteredAsParallelCapable = registeredAsParallelCapableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClassLoaderUnnamedModuleFlags(depth int, m *models.ClassLoader, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	unnamedModuleFlagName := fmt.Sprintf("%v.unnamedModule", cmdPrefix)
	if cmd.Flags().Changed(unnamedModuleFlagName) {
		// info: complex object unnamedModule Module is retrieved outside this Changed() block
	}
	unnamedModuleFlagValue := m.UnnamedModule
	if swag.IsZero(unnamedModuleFlagValue) {
		unnamedModuleFlagValue = &models.Module{}
	}

	err, unnamedModuleAdded := retrieveModelModuleFlags(depth+1, unnamedModuleFlagValue, unnamedModuleFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unnamedModuleAdded
	if unnamedModuleAdded {
		m.UnnamedModule = unnamedModuleFlagValue
	}

	return nil, retAdded
}
