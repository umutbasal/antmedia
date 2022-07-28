// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"

	"antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for Module

// register flags to command
func registerModelModuleFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerModuleAnnotations(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModuleClassLoader(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModuleDeclaredAnnotations(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModuleDescriptor(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModuleLayer(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModuleName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModuleNamed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerModulePackages(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerModuleAnnotations(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: annotations []Annotation array type is not supported by go-swagger cli yet

	return nil
}

func registerModuleClassLoader(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModuleDeclaredAnnotations(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: declaredAnnotations []Annotation array type is not supported by go-swagger cli yet

	return nil
}

func registerModuleDescriptor(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var descriptorFlagName string
	if cmdPrefix == "" {
		descriptorFlagName = "descriptor"
	} else {
		descriptorFlagName = fmt.Sprintf("%v.descriptor", cmdPrefix)
	}

	if err := registerModelModuleDescriptorFlags(depth+1, descriptorFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerModuleLayer(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: layer ModuleLayer map type is not supported by go-swagger cli yet

	return nil
}

func registerModuleName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerModuleNamed(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	namedDescription := ``

	var namedFlagName string
	if cmdPrefix == "" {
		namedFlagName = "named"
	} else {
		namedFlagName = fmt.Sprintf("%v.named", cmdPrefix)
	}

	var namedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(namedFlagName, namedFlagDefault, namedDescription)

	return nil
}

func registerModulePackages(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: packages []string array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelModuleFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, annotationsAdded := retrieveModuleAnnotationsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || annotationsAdded

	err, classLoaderAdded := retrieveModuleClassLoaderFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classLoaderAdded

	err, declaredAnnotationsAdded := retrieveModuleDeclaredAnnotationsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || declaredAnnotationsAdded

	err, descriptorAdded := retrieveModuleDescriptorFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptorAdded

	err, layerAdded := retrieveModuleLayerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || layerAdded

	err, nameAdded := retrieveModuleNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, namedAdded := retrieveModuleNamedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || namedAdded

	err, packagesAdded := retrieveModulePackagesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || packagesAdded

	return nil, retAdded
}

func retrieveModuleAnnotationsFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	annotationsFlagName := fmt.Sprintf("%v.annotations", cmdPrefix)
	if cmd.Flags().Changed(annotationsFlagName) {
		// warning: annotations array type []Annotation is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModuleClassLoaderFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModuleDeclaredAnnotationsFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	declaredAnnotationsFlagName := fmt.Sprintf("%v.declaredAnnotations", cmdPrefix)
	if cmd.Flags().Changed(declaredAnnotationsFlagName) {
		// warning: declaredAnnotations array type []Annotation is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModuleDescriptorFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptorFlagName := fmt.Sprintf("%v.descriptor", cmdPrefix)
	if cmd.Flags().Changed(descriptorFlagName) {
		// info: complex object descriptor ModuleDescriptor is retrieved outside this Changed() block
	}
	descriptorFlagValue := m.Descriptor
	if swag.IsZero(descriptorFlagValue) {
		descriptorFlagValue = &models.ModuleDescriptor{}
	}

	err, descriptorAdded := retrieveModelModuleDescriptorFlags(depth+1, descriptorFlagValue, descriptorFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptorAdded
	if descriptorAdded {
		m.Descriptor = descriptorFlagValue
	}

	return nil, retAdded
}

func retrieveModuleLayerFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	layerFlagName := fmt.Sprintf("%v.layer", cmdPrefix)
	if cmd.Flags().Changed(layerFlagName) {
		// warning: layer map type ModuleLayer is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveModuleNameFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveModuleNamedFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	namedFlagName := fmt.Sprintf("%v.named", cmdPrefix)
	if cmd.Flags().Changed(namedFlagName) {

		var namedFlagName string
		if cmdPrefix == "" {
			namedFlagName = "named"
		} else {
			namedFlagName = fmt.Sprintf("%v.named", cmdPrefix)
		}

		namedFlagValue, err := cmd.Flags().GetBool(namedFlagName)
		if err != nil {
			return err, false
		}
		m.Named = namedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveModulePackagesFlags(depth int, m *models.Module, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	packagesFlagName := fmt.Sprintf("%v.packages", cmdPrefix)
	if cmd.Flags().Changed(packagesFlagName) {
		// warning: packages array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
