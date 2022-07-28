// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/go-openapi/swag"

	"github.com/spf13/cobra"
)

// Schema cli for ThreadGroup

// register flags to command
func registerModelThreadGroupFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerThreadGroupDaemon(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerThreadGroupDestroyed(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerThreadGroupMaxPriority(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerThreadGroupName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerThreadGroupParent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerThreadGroupDaemon(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	daemonDescription := ``

	var daemonFlagName string
	if cmdPrefix == "" {
		daemonFlagName = "daemon"
	} else {
		daemonFlagName = fmt.Sprintf("%v.daemon", cmdPrefix)
	}

	var daemonFlagDefault bool

	_ = cmd.PersistentFlags().Bool(daemonFlagName, daemonFlagDefault, daemonDescription)

	return nil
}

func registerThreadGroupDestroyed(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	destroyedDescription := ``

	var destroyedFlagName string
	if cmdPrefix == "" {
		destroyedFlagName = "destroyed"
	} else {
		destroyedFlagName = fmt.Sprintf("%v.destroyed", cmdPrefix)
	}

	var destroyedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(destroyedFlagName, destroyedFlagDefault, destroyedDescription)

	return nil
}

func registerThreadGroupMaxPriority(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxPriorityDescription := ``

	var maxPriorityFlagName string
	if cmdPrefix == "" {
		maxPriorityFlagName = "maxPriority"
	} else {
		maxPriorityFlagName = fmt.Sprintf("%v.maxPriority", cmdPrefix)
	}

	var maxPriorityFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxPriorityFlagName, maxPriorityFlagDefault, maxPriorityDescription)

	return nil
}

func registerThreadGroupName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerThreadGroupParent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var parentFlagName string
	if cmdPrefix == "" {
		parentFlagName = "parent"
	} else {
		parentFlagName = fmt.Sprintf("%v.parent", cmdPrefix)
	}

	if err := registerModelThreadGroupFlags(depth+1, parentFlagName, cmd); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelThreadGroupFlags(depth int, m *models.ThreadGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, daemonAdded := retrieveThreadGroupDaemonFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || daemonAdded

	err, destroyedAdded := retrieveThreadGroupDestroyedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || destroyedAdded

	err, maxPriorityAdded := retrieveThreadGroupMaxPriorityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxPriorityAdded

	err, nameAdded := retrieveThreadGroupNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, parentAdded := retrieveThreadGroupParentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded

	return nil, retAdded
}

func retrieveThreadGroupDaemonFlags(depth int, m *models.ThreadGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	daemonFlagName := fmt.Sprintf("%v.daemon", cmdPrefix)
	if cmd.Flags().Changed(daemonFlagName) {

		var daemonFlagName string
		if cmdPrefix == "" {
			daemonFlagName = "daemon"
		} else {
			daemonFlagName = fmt.Sprintf("%v.daemon", cmdPrefix)
		}

		daemonFlagValue, err := cmd.Flags().GetBool(daemonFlagName)
		if err != nil {
			return err, false
		}
		m.Daemon = daemonFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveThreadGroupDestroyedFlags(depth int, m *models.ThreadGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	destroyedFlagName := fmt.Sprintf("%v.destroyed", cmdPrefix)
	if cmd.Flags().Changed(destroyedFlagName) {

		var destroyedFlagName string
		if cmdPrefix == "" {
			destroyedFlagName = "destroyed"
		} else {
			destroyedFlagName = fmt.Sprintf("%v.destroyed", cmdPrefix)
		}

		destroyedFlagValue, err := cmd.Flags().GetBool(destroyedFlagName)
		if err != nil {
			return err, false
		}
		m.Destroyed = destroyedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveThreadGroupMaxPriorityFlags(depth int, m *models.ThreadGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxPriorityFlagName := fmt.Sprintf("%v.maxPriority", cmdPrefix)
	if cmd.Flags().Changed(maxPriorityFlagName) {

		var maxPriorityFlagName string
		if cmdPrefix == "" {
			maxPriorityFlagName = "maxPriority"
		} else {
			maxPriorityFlagName = fmt.Sprintf("%v.maxPriority", cmdPrefix)
		}

		maxPriorityFlagValue, err := cmd.Flags().GetInt32(maxPriorityFlagName)
		if err != nil {
			return err, false
		}
		m.MaxPriority = maxPriorityFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveThreadGroupNameFlags(depth int, m *models.ThreadGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveThreadGroupParentFlags(depth int, m *models.ThreadGroup, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	parentFlagName := fmt.Sprintf("%v.parent", cmdPrefix)
	if cmd.Flags().Changed(parentFlagName) {
		// info: complex object parent ThreadGroup is retrieved outside this Changed() block
	}
	parentFlagValue := m.Parent
	if swag.IsZero(parentFlagValue) {
		parentFlagValue = &models.ThreadGroup{}
	}

	err, parentAdded := retrieveModelThreadGroupFlags(depth+1, parentFlagValue, parentFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || parentAdded
	if parentAdded {
		m.Parent = parentFlagValue
	}

	return nil, retAdded
}