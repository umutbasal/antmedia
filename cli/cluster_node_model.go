// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for ClusterNode

// register flags to command
func registerModelClusterNodeFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerClusterNodeCPU(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterNodeID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterNodeIP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterNodeLastUpdateTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterNodeMemory(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerClusterNodeStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerClusterNodeCPU(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cpuDescription := ``

	var cpuFlagName string
	if cmdPrefix == "" {
		cpuFlagName = "cpu"
	} else {
		cpuFlagName = fmt.Sprintf("%v.cpu", cmdPrefix)
	}

	var cpuFlagDefault string

	_ = cmd.PersistentFlags().String(cpuFlagName, cpuFlagDefault, cpuDescription)

	return nil
}

func registerClusterNodeID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault string

	_ = cmd.PersistentFlags().String(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerClusterNodeIP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	ipDescription := ``

	var ipFlagName string
	if cmdPrefix == "" {
		ipFlagName = "ip"
	} else {
		ipFlagName = fmt.Sprintf("%v.ip", cmdPrefix)
	}

	var ipFlagDefault string

	_ = cmd.PersistentFlags().String(ipFlagName, ipFlagDefault, ipDescription)

	return nil
}

func registerClusterNodeLastUpdateTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastUpdateTimeDescription := ``

	var lastUpdateTimeFlagName string
	if cmdPrefix == "" {
		lastUpdateTimeFlagName = "lastUpdateTime"
	} else {
		lastUpdateTimeFlagName = fmt.Sprintf("%v.lastUpdateTime", cmdPrefix)
	}

	var lastUpdateTimeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(lastUpdateTimeFlagName, lastUpdateTimeFlagDefault, lastUpdateTimeDescription)

	return nil
}

func registerClusterNodeMemory(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	memoryDescription := ``

	var memoryFlagName string
	if cmdPrefix == "" {
		memoryFlagName = "memory"
	} else {
		memoryFlagName = fmt.Sprintf("%v.memory", cmdPrefix)
	}

	var memoryFlagDefault string

	_ = cmd.PersistentFlags().String(memoryFlagName, memoryFlagDefault, memoryDescription)

	return nil
}

func registerClusterNodeStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := ``

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelClusterNodeFlags(depth int, m *models.ClusterNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, cpuAdded := retrieveClusterNodeCPUFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cpuAdded

	err, idAdded := retrieveClusterNodeIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, ipAdded := retrieveClusterNodeIPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || ipAdded

	err, lastUpdateTimeAdded := retrieveClusterNodeLastUpdateTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastUpdateTimeAdded

	err, memoryAdded := retrieveClusterNodeMemoryFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || memoryAdded

	err, statusAdded := retrieveClusterNodeStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrieveClusterNodeCPUFlags(depth int, m *models.ClusterNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cpuFlagName := fmt.Sprintf("%v.cpu", cmdPrefix)
	if cmd.Flags().Changed(cpuFlagName) {

		var cpuFlagName string
		if cmdPrefix == "" {
			cpuFlagName = "cpu"
		} else {
			cpuFlagName = fmt.Sprintf("%v.cpu", cmdPrefix)
		}

		cpuFlagValue, err := cmd.Flags().GetString(cpuFlagName)
		if err != nil {
			return err, false
		}
		m.CPU = cpuFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterNodeIDFlags(depth int, m *models.ClusterNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetString(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterNodeIPFlags(depth int, m *models.ClusterNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	ipFlagName := fmt.Sprintf("%v.ip", cmdPrefix)
	if cmd.Flags().Changed(ipFlagName) {

		var ipFlagName string
		if cmdPrefix == "" {
			ipFlagName = "ip"
		} else {
			ipFlagName = fmt.Sprintf("%v.ip", cmdPrefix)
		}

		ipFlagValue, err := cmd.Flags().GetString(ipFlagName)
		if err != nil {
			return err, false
		}
		m.IP = ipFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterNodeLastUpdateTimeFlags(depth int, m *models.ClusterNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lastUpdateTimeFlagName := fmt.Sprintf("%v.lastUpdateTime", cmdPrefix)
	if cmd.Flags().Changed(lastUpdateTimeFlagName) {

		var lastUpdateTimeFlagName string
		if cmdPrefix == "" {
			lastUpdateTimeFlagName = "lastUpdateTime"
		} else {
			lastUpdateTimeFlagName = fmt.Sprintf("%v.lastUpdateTime", cmdPrefix)
		}

		lastUpdateTimeFlagValue, err := cmd.Flags().GetInt64(lastUpdateTimeFlagName)
		if err != nil {
			return err, false
		}
		m.LastUpdateTime = lastUpdateTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterNodeMemoryFlags(depth int, m *models.ClusterNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	memoryFlagName := fmt.Sprintf("%v.memory", cmdPrefix)
	if cmd.Flags().Changed(memoryFlagName) {

		var memoryFlagName string
		if cmdPrefix == "" {
			memoryFlagName = "memory"
		} else {
			memoryFlagName = fmt.Sprintf("%v.memory", cmdPrefix)
		}

		memoryFlagValue, err := cmd.Flags().GetString(memoryFlagName)
		if err != nil {
			return err, false
		}
		m.Memory = memoryFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveClusterNodeStatusFlags(depth int, m *models.ClusterNode, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}
