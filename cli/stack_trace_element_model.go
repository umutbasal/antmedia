// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for StackTraceElement

// register flags to command
func registerModelStackTraceElementFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStackTraceElementClassLoaderName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStackTraceElementClassName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStackTraceElementFileName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStackTraceElementLineNumber(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStackTraceElementMethodName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStackTraceElementModuleName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStackTraceElementModuleVersion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStackTraceElementNativeMethod(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStackTraceElementClassLoaderName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	classLoaderNameDescription := ``

	var classLoaderNameFlagName string
	if cmdPrefix == "" {
		classLoaderNameFlagName = "classLoaderName"
	} else {
		classLoaderNameFlagName = fmt.Sprintf("%v.classLoaderName", cmdPrefix)
	}

	var classLoaderNameFlagDefault string

	_ = cmd.PersistentFlags().String(classLoaderNameFlagName, classLoaderNameFlagDefault, classLoaderNameDescription)

	return nil
}

func registerStackTraceElementClassName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerStackTraceElementFileName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fileNameDescription := ``

	var fileNameFlagName string
	if cmdPrefix == "" {
		fileNameFlagName = "fileName"
	} else {
		fileNameFlagName = fmt.Sprintf("%v.fileName", cmdPrefix)
	}

	var fileNameFlagDefault string

	_ = cmd.PersistentFlags().String(fileNameFlagName, fileNameFlagDefault, fileNameDescription)

	return nil
}

func registerStackTraceElementLineNumber(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lineNumberDescription := ``

	var lineNumberFlagName string
	if cmdPrefix == "" {
		lineNumberFlagName = "lineNumber"
	} else {
		lineNumberFlagName = fmt.Sprintf("%v.lineNumber", cmdPrefix)
	}

	var lineNumberFlagDefault int32

	_ = cmd.PersistentFlags().Int32(lineNumberFlagName, lineNumberFlagDefault, lineNumberDescription)

	return nil
}

func registerStackTraceElementMethodName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	methodNameDescription := ``

	var methodNameFlagName string
	if cmdPrefix == "" {
		methodNameFlagName = "methodName"
	} else {
		methodNameFlagName = fmt.Sprintf("%v.methodName", cmdPrefix)
	}

	var methodNameFlagDefault string

	_ = cmd.PersistentFlags().String(methodNameFlagName, methodNameFlagDefault, methodNameDescription)

	return nil
}

func registerStackTraceElementModuleName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	moduleNameDescription := ``

	var moduleNameFlagName string
	if cmdPrefix == "" {
		moduleNameFlagName = "moduleName"
	} else {
		moduleNameFlagName = fmt.Sprintf("%v.moduleName", cmdPrefix)
	}

	var moduleNameFlagDefault string

	_ = cmd.PersistentFlags().String(moduleNameFlagName, moduleNameFlagDefault, moduleNameDescription)

	return nil
}

func registerStackTraceElementModuleVersion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	moduleVersionDescription := ``

	var moduleVersionFlagName string
	if cmdPrefix == "" {
		moduleVersionFlagName = "moduleVersion"
	} else {
		moduleVersionFlagName = fmt.Sprintf("%v.moduleVersion", cmdPrefix)
	}

	var moduleVersionFlagDefault string

	_ = cmd.PersistentFlags().String(moduleVersionFlagName, moduleVersionFlagDefault, moduleVersionDescription)

	return nil
}

func registerStackTraceElementNativeMethod(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nativeMethodDescription := ``

	var nativeMethodFlagName string
	if cmdPrefix == "" {
		nativeMethodFlagName = "nativeMethod"
	} else {
		nativeMethodFlagName = fmt.Sprintf("%v.nativeMethod", cmdPrefix)
	}

	var nativeMethodFlagDefault bool

	_ = cmd.PersistentFlags().Bool(nativeMethodFlagName, nativeMethodFlagDefault, nativeMethodDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStackTraceElementFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, classLoaderNameAdded := retrieveStackTraceElementClassLoaderNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classLoaderNameAdded

	err, classNameAdded := retrieveStackTraceElementClassNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || classNameAdded

	err, fileNameAdded := retrieveStackTraceElementFileNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileNameAdded

	err, lineNumberAdded := retrieveStackTraceElementLineNumberFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lineNumberAdded

	err, methodNameAdded := retrieveStackTraceElementMethodNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || methodNameAdded

	err, moduleNameAdded := retrieveStackTraceElementModuleNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || moduleNameAdded

	err, moduleVersionAdded := retrieveStackTraceElementModuleVersionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || moduleVersionAdded

	err, nativeMethodAdded := retrieveStackTraceElementNativeMethodFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nativeMethodAdded

	return nil, retAdded
}

func retrieveStackTraceElementClassLoaderNameFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	classLoaderNameFlagName := fmt.Sprintf("%v.classLoaderName", cmdPrefix)
	if cmd.Flags().Changed(classLoaderNameFlagName) {

		var classLoaderNameFlagName string
		if cmdPrefix == "" {
			classLoaderNameFlagName = "classLoaderName"
		} else {
			classLoaderNameFlagName = fmt.Sprintf("%v.classLoaderName", cmdPrefix)
		}

		classLoaderNameFlagValue, err := cmd.Flags().GetString(classLoaderNameFlagName)
		if err != nil {
			return err, false
		}
		m.ClassLoaderName = classLoaderNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStackTraceElementClassNameFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveStackTraceElementFileNameFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileNameFlagName := fmt.Sprintf("%v.fileName", cmdPrefix)
	if cmd.Flags().Changed(fileNameFlagName) {

		var fileNameFlagName string
		if cmdPrefix == "" {
			fileNameFlagName = "fileName"
		} else {
			fileNameFlagName = fmt.Sprintf("%v.fileName", cmdPrefix)
		}

		fileNameFlagValue, err := cmd.Flags().GetString(fileNameFlagName)
		if err != nil {
			return err, false
		}
		m.FileName = fileNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStackTraceElementLineNumberFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lineNumberFlagName := fmt.Sprintf("%v.lineNumber", cmdPrefix)
	if cmd.Flags().Changed(lineNumberFlagName) {

		var lineNumberFlagName string
		if cmdPrefix == "" {
			lineNumberFlagName = "lineNumber"
		} else {
			lineNumberFlagName = fmt.Sprintf("%v.lineNumber", cmdPrefix)
		}

		lineNumberFlagValue, err := cmd.Flags().GetInt32(lineNumberFlagName)
		if err != nil {
			return err, false
		}
		m.LineNumber = lineNumberFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStackTraceElementMethodNameFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	methodNameFlagName := fmt.Sprintf("%v.methodName", cmdPrefix)
	if cmd.Flags().Changed(methodNameFlagName) {

		var methodNameFlagName string
		if cmdPrefix == "" {
			methodNameFlagName = "methodName"
		} else {
			methodNameFlagName = fmt.Sprintf("%v.methodName", cmdPrefix)
		}

		methodNameFlagValue, err := cmd.Flags().GetString(methodNameFlagName)
		if err != nil {
			return err, false
		}
		m.MethodName = methodNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStackTraceElementModuleNameFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	moduleNameFlagName := fmt.Sprintf("%v.moduleName", cmdPrefix)
	if cmd.Flags().Changed(moduleNameFlagName) {

		var moduleNameFlagName string
		if cmdPrefix == "" {
			moduleNameFlagName = "moduleName"
		} else {
			moduleNameFlagName = fmt.Sprintf("%v.moduleName", cmdPrefix)
		}

		moduleNameFlagValue, err := cmd.Flags().GetString(moduleNameFlagName)
		if err != nil {
			return err, false
		}
		m.ModuleName = moduleNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStackTraceElementModuleVersionFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	moduleVersionFlagName := fmt.Sprintf("%v.moduleVersion", cmdPrefix)
	if cmd.Flags().Changed(moduleVersionFlagName) {

		var moduleVersionFlagName string
		if cmdPrefix == "" {
			moduleVersionFlagName = "moduleVersion"
		} else {
			moduleVersionFlagName = fmt.Sprintf("%v.moduleVersion", cmdPrefix)
		}

		moduleVersionFlagValue, err := cmd.Flags().GetString(moduleVersionFlagName)
		if err != nil {
			return err, false
		}
		m.ModuleVersion = moduleVersionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStackTraceElementNativeMethodFlags(depth int, m *models.StackTraceElement, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nativeMethodFlagName := fmt.Sprintf("%v.nativeMethod", cmdPrefix)
	if cmd.Flags().Changed(nativeMethodFlagName) {

		var nativeMethodFlagName string
		if cmdPrefix == "" {
			nativeMethodFlagName = "nativeMethod"
		} else {
			nativeMethodFlagName = fmt.Sprintf("%v.nativeMethod", cmdPrefix)
		}

		nativeMethodFlagValue, err := cmd.Flags().GetBool(nativeMethodFlagName)
		if err != nil {
			return err, false
		}
		m.NativeMethod = nativeMethodFlagValue

		retAdded = true
	}

	return nil, retAdded
}