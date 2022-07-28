// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for AtomicBoolean

// register flags to command
func registerModelAtomicBooleanFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAtomicBooleanAcquire(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAtomicBooleanOpaque(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAtomicBooleanPlain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAtomicBooleanAcquire(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	acquireDescription := ``

	var acquireFlagName string
	if cmdPrefix == "" {
		acquireFlagName = "acquire"
	} else {
		acquireFlagName = fmt.Sprintf("%v.acquire", cmdPrefix)
	}

	var acquireFlagDefault bool

	_ = cmd.PersistentFlags().Bool(acquireFlagName, acquireFlagDefault, acquireDescription)

	return nil
}

func registerAtomicBooleanOpaque(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	opaqueDescription := ``

	var opaqueFlagName string
	if cmdPrefix == "" {
		opaqueFlagName = "opaque"
	} else {
		opaqueFlagName = fmt.Sprintf("%v.opaque", cmdPrefix)
	}

	var opaqueFlagDefault bool

	_ = cmd.PersistentFlags().Bool(opaqueFlagName, opaqueFlagDefault, opaqueDescription)

	return nil
}

func registerAtomicBooleanPlain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	plainDescription := ``

	var plainFlagName string
	if cmdPrefix == "" {
		plainFlagName = "plain"
	} else {
		plainFlagName = fmt.Sprintf("%v.plain", cmdPrefix)
	}

	var plainFlagDefault bool

	_ = cmd.PersistentFlags().Bool(plainFlagName, plainFlagDefault, plainDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAtomicBooleanFlags(depth int, m *models.AtomicBoolean, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, acquireAdded := retrieveAtomicBooleanAcquireFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || acquireAdded

	err, opaqueAdded := retrieveAtomicBooleanOpaqueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || opaqueAdded

	err, plainAdded := retrieveAtomicBooleanPlainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || plainAdded

	return nil, retAdded
}

func retrieveAtomicBooleanAcquireFlags(depth int, m *models.AtomicBoolean, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	acquireFlagName := fmt.Sprintf("%v.acquire", cmdPrefix)
	if cmd.Flags().Changed(acquireFlagName) {

		var acquireFlagName string
		if cmdPrefix == "" {
			acquireFlagName = "acquire"
		} else {
			acquireFlagName = fmt.Sprintf("%v.acquire", cmdPrefix)
		}

		acquireFlagValue, err := cmd.Flags().GetBool(acquireFlagName)
		if err != nil {
			return err, false
		}
		m.Acquire = acquireFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAtomicBooleanOpaqueFlags(depth int, m *models.AtomicBoolean, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	opaqueFlagName := fmt.Sprintf("%v.opaque", cmdPrefix)
	if cmd.Flags().Changed(opaqueFlagName) {

		var opaqueFlagName string
		if cmdPrefix == "" {
			opaqueFlagName = "opaque"
		} else {
			opaqueFlagName = fmt.Sprintf("%v.opaque", cmdPrefix)
		}

		opaqueFlagValue, err := cmd.Flags().GetBool(opaqueFlagName)
		if err != nil {
			return err, false
		}
		m.Opaque = opaqueFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAtomicBooleanPlainFlags(depth int, m *models.AtomicBoolean, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	plainFlagName := fmt.Sprintf("%v.plain", cmdPrefix)
	if cmd.Flags().Changed(plainFlagName) {

		var plainFlagName string
		if cmdPrefix == "" {
			plainFlagName = "plain"
		} else {
			plainFlagName = fmt.Sprintf("%v.plain", cmdPrefix)
		}

		plainFlagValue, err := cmd.Flags().GetBool(plainFlagName)
		if err != nil {
			return err, false
		}
		m.Plain = plainFlagValue

		retAdded = true
	}

	return nil, retAdded
}
