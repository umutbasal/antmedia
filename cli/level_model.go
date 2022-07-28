// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for Level

// register flags to command
func registerModelLevelFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerLevelLevelInt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLevelLevelStr(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerLevelLevelInt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	levelIntDescription := ``

	var levelIntFlagName string
	if cmdPrefix == "" {
		levelIntFlagName = "levelInt"
	} else {
		levelIntFlagName = fmt.Sprintf("%v.levelInt", cmdPrefix)
	}

	var levelIntFlagDefault int32

	_ = cmd.PersistentFlags().Int32(levelIntFlagName, levelIntFlagDefault, levelIntDescription)

	return nil
}

func registerLevelLevelStr(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	levelStrDescription := ``

	var levelStrFlagName string
	if cmdPrefix == "" {
		levelStrFlagName = "levelStr"
	} else {
		levelStrFlagName = fmt.Sprintf("%v.levelStr", cmdPrefix)
	}

	var levelStrFlagDefault string

	_ = cmd.PersistentFlags().String(levelStrFlagName, levelStrFlagDefault, levelStrDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelLevelFlags(depth int, m *models.Level, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, levelIntAdded := retrieveLevelLevelIntFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || levelIntAdded

	err, levelStrAdded := retrieveLevelLevelStrFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || levelStrAdded

	return nil, retAdded
}

func retrieveLevelLevelIntFlags(depth int, m *models.Level, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	levelIntFlagName := fmt.Sprintf("%v.levelInt", cmdPrefix)
	if cmd.Flags().Changed(levelIntFlagName) {

		var levelIntFlagName string
		if cmdPrefix == "" {
			levelIntFlagName = "levelInt"
		} else {
			levelIntFlagName = fmt.Sprintf("%v.levelInt", cmdPrefix)
		}

		levelIntFlagValue, err := cmd.Flags().GetInt32(levelIntFlagName)
		if err != nil {
			return err, false
		}
		m.LevelInt = levelIntFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLevelLevelStrFlags(depth int, m *models.Level, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	levelStrFlagName := fmt.Sprintf("%v.levelStr", cmdPrefix)
	if cmd.Flags().Changed(levelStrFlagName) {

		var levelStrFlagName string
		if cmdPrefix == "" {
			levelStrFlagName = "levelStr"
		} else {
			levelStrFlagName = fmt.Sprintf("%v.levelStr", cmdPrefix)
		}

		levelStrFlagValue, err := cmd.Flags().GetString(levelStrFlagName)
		if err != nil {
			return err, false
		}
		m.LevelStr = levelStrFlagValue

		retAdded = true
	}

	return nil, retAdded
}
