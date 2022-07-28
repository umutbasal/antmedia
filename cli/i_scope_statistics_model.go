// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for IScopeStatistics

// register flags to command
func registerModelIScopeStatisticsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIScopeStatisticsActiveClients(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsActiveConnections(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsActiveSubscopes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsCreationTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsDepth(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsMaxClients(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsMaxConnections(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsMaxSubscopes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsPath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsTotalClients(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsTotalConnections(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIScopeStatisticsTotalSubscopes(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIScopeStatisticsActiveClients(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	activeClientsDescription := ``

	var activeClientsFlagName string
	if cmdPrefix == "" {
		activeClientsFlagName = "activeClients"
	} else {
		activeClientsFlagName = fmt.Sprintf("%v.activeClients", cmdPrefix)
	}

	var activeClientsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(activeClientsFlagName, activeClientsFlagDefault, activeClientsDescription)

	return nil
}

func registerIScopeStatisticsActiveConnections(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	activeConnectionsDescription := ``

	var activeConnectionsFlagName string
	if cmdPrefix == "" {
		activeConnectionsFlagName = "activeConnections"
	} else {
		activeConnectionsFlagName = fmt.Sprintf("%v.activeConnections", cmdPrefix)
	}

	var activeConnectionsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(activeConnectionsFlagName, activeConnectionsFlagDefault, activeConnectionsDescription)

	return nil
}

func registerIScopeStatisticsActiveSubscopes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	activeSubscopesDescription := ``

	var activeSubscopesFlagName string
	if cmdPrefix == "" {
		activeSubscopesFlagName = "activeSubscopes"
	} else {
		activeSubscopesFlagName = fmt.Sprintf("%v.activeSubscopes", cmdPrefix)
	}

	var activeSubscopesFlagDefault int32

	_ = cmd.PersistentFlags().Int32(activeSubscopesFlagName, activeSubscopesFlagDefault, activeSubscopesDescription)

	return nil
}

func registerIScopeStatisticsCreationTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	creationTimeDescription := ``

	var creationTimeFlagName string
	if cmdPrefix == "" {
		creationTimeFlagName = "creationTime"
	} else {
		creationTimeFlagName = fmt.Sprintf("%v.creationTime", cmdPrefix)
	}

	var creationTimeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(creationTimeFlagName, creationTimeFlagDefault, creationTimeDescription)

	return nil
}

func registerIScopeStatisticsDepth(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	depthDescription := ``

	var depthFlagName string
	if cmdPrefix == "" {
		depthFlagName = "depth"
	} else {
		depthFlagName = fmt.Sprintf("%v.depth", cmdPrefix)
	}

	var depthFlagDefault int32

	_ = cmd.PersistentFlags().Int32(depthFlagName, depthFlagDefault, depthDescription)

	return nil
}

func registerIScopeStatisticsMaxClients(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxClientsDescription := ``

	var maxClientsFlagName string
	if cmdPrefix == "" {
		maxClientsFlagName = "maxClients"
	} else {
		maxClientsFlagName = fmt.Sprintf("%v.maxClients", cmdPrefix)
	}

	var maxClientsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxClientsFlagName, maxClientsFlagDefault, maxClientsDescription)

	return nil
}

func registerIScopeStatisticsMaxConnections(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxConnectionsDescription := ``

	var maxConnectionsFlagName string
	if cmdPrefix == "" {
		maxConnectionsFlagName = "maxConnections"
	} else {
		maxConnectionsFlagName = fmt.Sprintf("%v.maxConnections", cmdPrefix)
	}

	var maxConnectionsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxConnectionsFlagName, maxConnectionsFlagDefault, maxConnectionsDescription)

	return nil
}

func registerIScopeStatisticsMaxSubscopes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxSubscopesDescription := ``

	var maxSubscopesFlagName string
	if cmdPrefix == "" {
		maxSubscopesFlagName = "maxSubscopes"
	} else {
		maxSubscopesFlagName = fmt.Sprintf("%v.maxSubscopes", cmdPrefix)
	}

	var maxSubscopesFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxSubscopesFlagName, maxSubscopesFlagDefault, maxSubscopesDescription)

	return nil
}

func registerIScopeStatisticsName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerIScopeStatisticsPath(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pathDescription := ``

	var pathFlagName string
	if cmdPrefix == "" {
		pathFlagName = "path"
	} else {
		pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
	}

	var pathFlagDefault string

	_ = cmd.PersistentFlags().String(pathFlagName, pathFlagDefault, pathDescription)

	return nil
}

func registerIScopeStatisticsTotalClients(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalClientsDescription := ``

	var totalClientsFlagName string
	if cmdPrefix == "" {
		totalClientsFlagName = "totalClients"
	} else {
		totalClientsFlagName = fmt.Sprintf("%v.totalClients", cmdPrefix)
	}

	var totalClientsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(totalClientsFlagName, totalClientsFlagDefault, totalClientsDescription)

	return nil
}

func registerIScopeStatisticsTotalConnections(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalConnectionsDescription := ``

	var totalConnectionsFlagName string
	if cmdPrefix == "" {
		totalConnectionsFlagName = "totalConnections"
	} else {
		totalConnectionsFlagName = fmt.Sprintf("%v.totalConnections", cmdPrefix)
	}

	var totalConnectionsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(totalConnectionsFlagName, totalConnectionsFlagDefault, totalConnectionsDescription)

	return nil
}

func registerIScopeStatisticsTotalSubscopes(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	totalSubscopesDescription := ``

	var totalSubscopesFlagName string
	if cmdPrefix == "" {
		totalSubscopesFlagName = "totalSubscopes"
	} else {
		totalSubscopesFlagName = fmt.Sprintf("%v.totalSubscopes", cmdPrefix)
	}

	var totalSubscopesFlagDefault int32

	_ = cmd.PersistentFlags().Int32(totalSubscopesFlagName, totalSubscopesFlagDefault, totalSubscopesDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIScopeStatisticsFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, activeClientsAdded := retrieveIScopeStatisticsActiveClientsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || activeClientsAdded

	err, activeConnectionsAdded := retrieveIScopeStatisticsActiveConnectionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || activeConnectionsAdded

	err, activeSubscopesAdded := retrieveIScopeStatisticsActiveSubscopesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || activeSubscopesAdded

	err, creationTimeAdded := retrieveIScopeStatisticsCreationTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || creationTimeAdded

	err, depthAdded := retrieveIScopeStatisticsDepthFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || depthAdded

	err, maxClientsAdded := retrieveIScopeStatisticsMaxClientsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxClientsAdded

	err, maxConnectionsAdded := retrieveIScopeStatisticsMaxConnectionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxConnectionsAdded

	err, maxSubscopesAdded := retrieveIScopeStatisticsMaxSubscopesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxSubscopesAdded

	err, nameAdded := retrieveIScopeStatisticsNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, pathAdded := retrieveIScopeStatisticsPathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pathAdded

	err, totalClientsAdded := retrieveIScopeStatisticsTotalClientsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalClientsAdded

	err, totalConnectionsAdded := retrieveIScopeStatisticsTotalConnectionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalConnectionsAdded

	err, totalSubscopesAdded := retrieveIScopeStatisticsTotalSubscopesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || totalSubscopesAdded

	return nil, retAdded
}

func retrieveIScopeStatisticsActiveClientsFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	activeClientsFlagName := fmt.Sprintf("%v.activeClients", cmdPrefix)
	if cmd.Flags().Changed(activeClientsFlagName) {

		var activeClientsFlagName string
		if cmdPrefix == "" {
			activeClientsFlagName = "activeClients"
		} else {
			activeClientsFlagName = fmt.Sprintf("%v.activeClients", cmdPrefix)
		}

		activeClientsFlagValue, err := cmd.Flags().GetInt32(activeClientsFlagName)
		if err != nil {
			return err, false
		}
		m.ActiveClients = activeClientsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsActiveConnectionsFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	activeConnectionsFlagName := fmt.Sprintf("%v.activeConnections", cmdPrefix)
	if cmd.Flags().Changed(activeConnectionsFlagName) {

		var activeConnectionsFlagName string
		if cmdPrefix == "" {
			activeConnectionsFlagName = "activeConnections"
		} else {
			activeConnectionsFlagName = fmt.Sprintf("%v.activeConnections", cmdPrefix)
		}

		activeConnectionsFlagValue, err := cmd.Flags().GetInt32(activeConnectionsFlagName)
		if err != nil {
			return err, false
		}
		m.ActiveConnections = activeConnectionsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsActiveSubscopesFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	activeSubscopesFlagName := fmt.Sprintf("%v.activeSubscopes", cmdPrefix)
	if cmd.Flags().Changed(activeSubscopesFlagName) {

		var activeSubscopesFlagName string
		if cmdPrefix == "" {
			activeSubscopesFlagName = "activeSubscopes"
		} else {
			activeSubscopesFlagName = fmt.Sprintf("%v.activeSubscopes", cmdPrefix)
		}

		activeSubscopesFlagValue, err := cmd.Flags().GetInt32(activeSubscopesFlagName)
		if err != nil {
			return err, false
		}
		m.ActiveSubscopes = activeSubscopesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsCreationTimeFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	creationTimeFlagName := fmt.Sprintf("%v.creationTime", cmdPrefix)
	if cmd.Flags().Changed(creationTimeFlagName) {

		var creationTimeFlagName string
		if cmdPrefix == "" {
			creationTimeFlagName = "creationTime"
		} else {
			creationTimeFlagName = fmt.Sprintf("%v.creationTime", cmdPrefix)
		}

		creationTimeFlagValue, err := cmd.Flags().GetInt64(creationTimeFlagName)
		if err != nil {
			return err, false
		}
		m.CreationTime = creationTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsDepthFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	depthFlagName := fmt.Sprintf("%v.depth", cmdPrefix)
	if cmd.Flags().Changed(depthFlagName) {

		var depthFlagName string
		if cmdPrefix == "" {
			depthFlagName = "depth"
		} else {
			depthFlagName = fmt.Sprintf("%v.depth", cmdPrefix)
		}

		depthFlagValue, err := cmd.Flags().GetInt32(depthFlagName)
		if err != nil {
			return err, false
		}
		m.Depth = depthFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsMaxClientsFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxClientsFlagName := fmt.Sprintf("%v.maxClients", cmdPrefix)
	if cmd.Flags().Changed(maxClientsFlagName) {

		var maxClientsFlagName string
		if cmdPrefix == "" {
			maxClientsFlagName = "maxClients"
		} else {
			maxClientsFlagName = fmt.Sprintf("%v.maxClients", cmdPrefix)
		}

		maxClientsFlagValue, err := cmd.Flags().GetInt32(maxClientsFlagName)
		if err != nil {
			return err, false
		}
		m.MaxClients = maxClientsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsMaxConnectionsFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxConnectionsFlagName := fmt.Sprintf("%v.maxConnections", cmdPrefix)
	if cmd.Flags().Changed(maxConnectionsFlagName) {

		var maxConnectionsFlagName string
		if cmdPrefix == "" {
			maxConnectionsFlagName = "maxConnections"
		} else {
			maxConnectionsFlagName = fmt.Sprintf("%v.maxConnections", cmdPrefix)
		}

		maxConnectionsFlagValue, err := cmd.Flags().GetInt32(maxConnectionsFlagName)
		if err != nil {
			return err, false
		}
		m.MaxConnections = maxConnectionsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsMaxSubscopesFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxSubscopesFlagName := fmt.Sprintf("%v.maxSubscopes", cmdPrefix)
	if cmd.Flags().Changed(maxSubscopesFlagName) {

		var maxSubscopesFlagName string
		if cmdPrefix == "" {
			maxSubscopesFlagName = "maxSubscopes"
		} else {
			maxSubscopesFlagName = fmt.Sprintf("%v.maxSubscopes", cmdPrefix)
		}

		maxSubscopesFlagValue, err := cmd.Flags().GetInt32(maxSubscopesFlagName)
		if err != nil {
			return err, false
		}
		m.MaxSubscopes = maxSubscopesFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsNameFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveIScopeStatisticsPathFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pathFlagName := fmt.Sprintf("%v.path", cmdPrefix)
	if cmd.Flags().Changed(pathFlagName) {

		var pathFlagName string
		if cmdPrefix == "" {
			pathFlagName = "path"
		} else {
			pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
		}

		pathFlagValue, err := cmd.Flags().GetString(pathFlagName)
		if err != nil {
			return err, false
		}
		m.Path = pathFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsTotalClientsFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalClientsFlagName := fmt.Sprintf("%v.totalClients", cmdPrefix)
	if cmd.Flags().Changed(totalClientsFlagName) {

		var totalClientsFlagName string
		if cmdPrefix == "" {
			totalClientsFlagName = "totalClients"
		} else {
			totalClientsFlagName = fmt.Sprintf("%v.totalClients", cmdPrefix)
		}

		totalClientsFlagValue, err := cmd.Flags().GetInt32(totalClientsFlagName)
		if err != nil {
			return err, false
		}
		m.TotalClients = totalClientsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsTotalConnectionsFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalConnectionsFlagName := fmt.Sprintf("%v.totalConnections", cmdPrefix)
	if cmd.Flags().Changed(totalConnectionsFlagName) {

		var totalConnectionsFlagName string
		if cmdPrefix == "" {
			totalConnectionsFlagName = "totalConnections"
		} else {
			totalConnectionsFlagName = fmt.Sprintf("%v.totalConnections", cmdPrefix)
		}

		totalConnectionsFlagValue, err := cmd.Flags().GetInt32(totalConnectionsFlagName)
		if err != nil {
			return err, false
		}
		m.TotalConnections = totalConnectionsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIScopeStatisticsTotalSubscopesFlags(depth int, m *models.IScopeStatistics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	totalSubscopesFlagName := fmt.Sprintf("%v.totalSubscopes", cmdPrefix)
	if cmd.Flags().Changed(totalSubscopesFlagName) {

		var totalSubscopesFlagName string
		if cmdPrefix == "" {
			totalSubscopesFlagName = "totalSubscopes"
		} else {
			totalSubscopesFlagName = fmt.Sprintf("%v.totalSubscopes", cmdPrefix)
		}

		totalSubscopesFlagValue, err := cmd.Flags().GetInt32(totalSubscopesFlagName)
		if err != nil {
			return err, false
		}
		m.TotalSubscopes = totalSubscopesFlagValue

		retAdded = true
	}

	return nil, retAdded
}
