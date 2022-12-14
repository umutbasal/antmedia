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

// Schema cli for StreamAcceptFilter

// register flags to command
func registerModelStreamAcceptFilterFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStreamAcceptFilterAppSettings(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStreamAcceptFilterMaxBitrate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStreamAcceptFilterMaxFps(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStreamAcceptFilterMaxResolution(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStreamAcceptFilterAppSettings(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var appSettingsFlagName string
	if cmdPrefix == "" {
		appSettingsFlagName = "appSettings"
	} else {
		appSettingsFlagName = fmt.Sprintf("%v.appSettings", cmdPrefix)
	}

	if err := registerModelAppSettingsFlags(depth+1, appSettingsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerStreamAcceptFilterMaxBitrate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxBitrateDescription := ``

	var maxBitrateFlagName string
	if cmdPrefix == "" {
		maxBitrateFlagName = "maxBitrate"
	} else {
		maxBitrateFlagName = fmt.Sprintf("%v.maxBitrate", cmdPrefix)
	}

	var maxBitrateFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxBitrateFlagName, maxBitrateFlagDefault, maxBitrateDescription)

	return nil
}

func registerStreamAcceptFilterMaxFps(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxFpsDescription := ``

	var maxFpsFlagName string
	if cmdPrefix == "" {
		maxFpsFlagName = "maxFps"
	} else {
		maxFpsFlagName = fmt.Sprintf("%v.maxFps", cmdPrefix)
	}

	var maxFpsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxFpsFlagName, maxFpsFlagDefault, maxFpsDescription)

	return nil
}

func registerStreamAcceptFilterMaxResolution(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxResolutionDescription := ``

	var maxResolutionFlagName string
	if cmdPrefix == "" {
		maxResolutionFlagName = "maxResolution"
	} else {
		maxResolutionFlagName = fmt.Sprintf("%v.maxResolution", cmdPrefix)
	}

	var maxResolutionFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxResolutionFlagName, maxResolutionFlagDefault, maxResolutionDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStreamAcceptFilterFlags(depth int, m *models.StreamAcceptFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, appSettingsAdded := retrieveStreamAcceptFilterAppSettingsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || appSettingsAdded

	err, maxBitrateAdded := retrieveStreamAcceptFilterMaxBitrateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxBitrateAdded

	err, maxFpsAdded := retrieveStreamAcceptFilterMaxFpsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxFpsAdded

	err, maxResolutionAdded := retrieveStreamAcceptFilterMaxResolutionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxResolutionAdded

	return nil, retAdded
}

func retrieveStreamAcceptFilterAppSettingsFlags(depth int, m *models.StreamAcceptFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	appSettingsFlagName := fmt.Sprintf("%v.appSettings", cmdPrefix)
	if cmd.Flags().Changed(appSettingsFlagName) {
		// info: complex object appSettings AppSettings is retrieved outside this Changed() block
	}
	appSettingsFlagValue := m.AppSettings
	if swag.IsZero(appSettingsFlagValue) {
		appSettingsFlagValue = &models.AppSettings{}
	}

	err, appSettingsAdded := retrieveModelAppSettingsFlags(depth+1, appSettingsFlagValue, appSettingsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || appSettingsAdded
	if appSettingsAdded {
		m.AppSettings = appSettingsFlagValue
	}

	return nil, retAdded
}

func retrieveStreamAcceptFilterMaxBitrateFlags(depth int, m *models.StreamAcceptFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxBitrateFlagName := fmt.Sprintf("%v.maxBitrate", cmdPrefix)
	if cmd.Flags().Changed(maxBitrateFlagName) {

		var maxBitrateFlagName string
		if cmdPrefix == "" {
			maxBitrateFlagName = "maxBitrate"
		} else {
			maxBitrateFlagName = fmt.Sprintf("%v.maxBitrate", cmdPrefix)
		}

		maxBitrateFlagValue, err := cmd.Flags().GetInt32(maxBitrateFlagName)
		if err != nil {
			return err, false
		}
		m.MaxBitrate = maxBitrateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStreamAcceptFilterMaxFpsFlags(depth int, m *models.StreamAcceptFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxFpsFlagName := fmt.Sprintf("%v.maxFps", cmdPrefix)
	if cmd.Flags().Changed(maxFpsFlagName) {

		var maxFpsFlagName string
		if cmdPrefix == "" {
			maxFpsFlagName = "maxFps"
		} else {
			maxFpsFlagName = fmt.Sprintf("%v.maxFps", cmdPrefix)
		}

		maxFpsFlagValue, err := cmd.Flags().GetInt32(maxFpsFlagName)
		if err != nil {
			return err, false
		}
		m.MaxFps = maxFpsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStreamAcceptFilterMaxResolutionFlags(depth int, m *models.StreamAcceptFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxResolutionFlagName := fmt.Sprintf("%v.maxResolution", cmdPrefix)
	if cmd.Flags().Changed(maxResolutionFlagName) {

		var maxResolutionFlagName string
		if cmdPrefix == "" {
			maxResolutionFlagName = "maxResolution"
		} else {
			maxResolutionFlagName = fmt.Sprintf("%v.maxResolution", cmdPrefix)
		}

		maxResolutionFlagValue, err := cmd.Flags().GetInt32(maxResolutionFlagName)
		if err != nil {
			return err, false
		}
		m.MaxResolution = maxResolutionFlagValue

		retAdded = true
	}

	return nil, retAdded
}
