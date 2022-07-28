// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for IStatsCollector

// register flags to command
func registerModelIStatsCollectorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIStatsCollectorCPULimit(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIStatsCollectorCPULoad(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIStatsCollectorFreeRAM(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIStatsCollectorMinFreeRAMSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIStatsCollectorCPULimit(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cpuLimitDescription := ``

	var cpuLimitFlagName string
	if cmdPrefix == "" {
		cpuLimitFlagName = "cpuLimit"
	} else {
		cpuLimitFlagName = fmt.Sprintf("%v.cpuLimit", cmdPrefix)
	}

	var cpuLimitFlagDefault int32

	_ = cmd.PersistentFlags().Int32(cpuLimitFlagName, cpuLimitFlagDefault, cpuLimitDescription)

	return nil
}

func registerIStatsCollectorCPULoad(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cpuLoadDescription := ``

	var cpuLoadFlagName string
	if cmdPrefix == "" {
		cpuLoadFlagName = "cpuLoad"
	} else {
		cpuLoadFlagName = fmt.Sprintf("%v.cpuLoad", cmdPrefix)
	}

	var cpuLoadFlagDefault int32

	_ = cmd.PersistentFlags().Int32(cpuLoadFlagName, cpuLoadFlagDefault, cpuLoadDescription)

	return nil
}

func registerIStatsCollectorFreeRAM(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	freeRamDescription := ``

	var freeRamFlagName string
	if cmdPrefix == "" {
		freeRamFlagName = "freeRam"
	} else {
		freeRamFlagName = fmt.Sprintf("%v.freeRam", cmdPrefix)
	}

	var freeRamFlagDefault int32

	_ = cmd.PersistentFlags().Int32(freeRamFlagName, freeRamFlagDefault, freeRamDescription)

	return nil
}

func registerIStatsCollectorMinFreeRAMSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	minFreeRamSizeDescription := ``

	var minFreeRamSizeFlagName string
	if cmdPrefix == "" {
		minFreeRamSizeFlagName = "minFreeRamSize"
	} else {
		minFreeRamSizeFlagName = fmt.Sprintf("%v.minFreeRamSize", cmdPrefix)
	}

	var minFreeRamSizeFlagDefault int32

	_ = cmd.PersistentFlags().Int32(minFreeRamSizeFlagName, minFreeRamSizeFlagDefault, minFreeRamSizeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIStatsCollectorFlags(depth int, m *models.IStatsCollector, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, cpuLimitAdded := retrieveIStatsCollectorCPULimitFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cpuLimitAdded

	err, cpuLoadAdded := retrieveIStatsCollectorCPULoadFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cpuLoadAdded

	err, freeRamAdded := retrieveIStatsCollectorFreeRAMFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || freeRamAdded

	err, minFreeRamSizeAdded := retrieveIStatsCollectorMinFreeRAMSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || minFreeRamSizeAdded

	return nil, retAdded
}

func retrieveIStatsCollectorCPULimitFlags(depth int, m *models.IStatsCollector, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cpuLimitFlagName := fmt.Sprintf("%v.cpuLimit", cmdPrefix)
	if cmd.Flags().Changed(cpuLimitFlagName) {

		var cpuLimitFlagName string
		if cmdPrefix == "" {
			cpuLimitFlagName = "cpuLimit"
		} else {
			cpuLimitFlagName = fmt.Sprintf("%v.cpuLimit", cmdPrefix)
		}

		cpuLimitFlagValue, err := cmd.Flags().GetInt32(cpuLimitFlagName)
		if err != nil {
			return err, false
		}
		m.CPULimit = cpuLimitFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIStatsCollectorCPULoadFlags(depth int, m *models.IStatsCollector, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cpuLoadFlagName := fmt.Sprintf("%v.cpuLoad", cmdPrefix)
	if cmd.Flags().Changed(cpuLoadFlagName) {

		var cpuLoadFlagName string
		if cmdPrefix == "" {
			cpuLoadFlagName = "cpuLoad"
		} else {
			cpuLoadFlagName = fmt.Sprintf("%v.cpuLoad", cmdPrefix)
		}

		cpuLoadFlagValue, err := cmd.Flags().GetInt32(cpuLoadFlagName)
		if err != nil {
			return err, false
		}
		m.CPULoad = cpuLoadFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIStatsCollectorFreeRAMFlags(depth int, m *models.IStatsCollector, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	freeRamFlagName := fmt.Sprintf("%v.freeRam", cmdPrefix)
	if cmd.Flags().Changed(freeRamFlagName) {

		var freeRamFlagName string
		if cmdPrefix == "" {
			freeRamFlagName = "freeRam"
		} else {
			freeRamFlagName = fmt.Sprintf("%v.freeRam", cmdPrefix)
		}

		freeRamFlagValue, err := cmd.Flags().GetInt32(freeRamFlagName)
		if err != nil {
			return err, false
		}
		m.FreeRAM = freeRamFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIStatsCollectorMinFreeRAMSizeFlags(depth int, m *models.IStatsCollector, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	minFreeRamSizeFlagName := fmt.Sprintf("%v.minFreeRamSize", cmdPrefix)
	if cmd.Flags().Changed(minFreeRamSizeFlagName) {

		var minFreeRamSizeFlagName string
		if cmdPrefix == "" {
			minFreeRamSizeFlagName = "minFreeRamSize"
		} else {
			minFreeRamSizeFlagName = fmt.Sprintf("%v.minFreeRamSize", cmdPrefix)
		}

		minFreeRamSizeFlagValue, err := cmd.Flags().GetInt32(minFreeRamSizeFlagName)
		if err != nil {
			return err, false
		}
		m.MinFreeRAMSize = minFreeRamSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}