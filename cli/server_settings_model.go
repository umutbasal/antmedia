// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for ServerSettings

// register flags to command
func registerModelServerSettingsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerServerSettingsAllowedDashboardCIDR(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsBuildForMarket(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsCPUMeasurementPeriodMs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsCPUMeasurementWindowSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsDefaultHTTPPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsHeartbeatEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsHostAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsJwksURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsJwtServerControlEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsJwtServerSecretKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsLicenceKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsLogLevel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsMarketplace(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsNativeLogLevel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsNodeGroup(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsOriginServerPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsProxyAddress(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsServerName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsSrtPort(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsUseGlobalIP(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerServerSettingsWebRTCLogLevel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerServerSettingsAllowedDashboardCIDR(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	allowedDashboardCIdRDescription := ``

	var allowedDashboardCIdRFlagName string
	if cmdPrefix == "" {
		allowedDashboardCIdRFlagName = "allowedDashboardCIDR"
	} else {
		allowedDashboardCIdRFlagName = fmt.Sprintf("%v.allowedDashboardCIDR", cmdPrefix)
	}

	var allowedDashboardCIdRFlagDefault string

	_ = cmd.PersistentFlags().String(allowedDashboardCIdRFlagName, allowedDashboardCIdRFlagDefault, allowedDashboardCIdRDescription)

	return nil
}

func registerServerSettingsBuildForMarket(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	buildForMarketDescription := ``

	var buildForMarketFlagName string
	if cmdPrefix == "" {
		buildForMarketFlagName = "buildForMarket"
	} else {
		buildForMarketFlagName = fmt.Sprintf("%v.buildForMarket", cmdPrefix)
	}

	var buildForMarketFlagDefault bool

	_ = cmd.PersistentFlags().Bool(buildForMarketFlagName, buildForMarketFlagDefault, buildForMarketDescription)

	return nil
}

func registerServerSettingsCPUMeasurementPeriodMs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cpuMeasurementPeriodMsDescription := ``

	var cpuMeasurementPeriodMsFlagName string
	if cmdPrefix == "" {
		cpuMeasurementPeriodMsFlagName = "cpuMeasurementPeriodMs"
	} else {
		cpuMeasurementPeriodMsFlagName = fmt.Sprintf("%v.cpuMeasurementPeriodMs", cmdPrefix)
	}

	var cpuMeasurementPeriodMsFlagDefault int32

	_ = cmd.PersistentFlags().Int32(cpuMeasurementPeriodMsFlagName, cpuMeasurementPeriodMsFlagDefault, cpuMeasurementPeriodMsDescription)

	return nil
}

func registerServerSettingsCPUMeasurementWindowSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	cpuMeasurementWindowSizeDescription := ``

	var cpuMeasurementWindowSizeFlagName string
	if cmdPrefix == "" {
		cpuMeasurementWindowSizeFlagName = "cpuMeasurementWindowSize"
	} else {
		cpuMeasurementWindowSizeFlagName = fmt.Sprintf("%v.cpuMeasurementWindowSize", cmdPrefix)
	}

	var cpuMeasurementWindowSizeFlagDefault int32

	_ = cmd.PersistentFlags().Int32(cpuMeasurementWindowSizeFlagName, cpuMeasurementWindowSizeFlagDefault, cpuMeasurementWindowSizeDescription)

	return nil
}

func registerServerSettingsDefaultHTTPPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	defaultHttpPortDescription := ``

	var defaultHttpPortFlagName string
	if cmdPrefix == "" {
		defaultHttpPortFlagName = "defaultHttpPort"
	} else {
		defaultHttpPortFlagName = fmt.Sprintf("%v.defaultHttpPort", cmdPrefix)
	}

	var defaultHttpPortFlagDefault int32

	_ = cmd.PersistentFlags().Int32(defaultHttpPortFlagName, defaultHttpPortFlagDefault, defaultHttpPortDescription)

	return nil
}

func registerServerSettingsHeartbeatEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	heartbeatEnabledDescription := ``

	var heartbeatEnabledFlagName string
	if cmdPrefix == "" {
		heartbeatEnabledFlagName = "heartbeatEnabled"
	} else {
		heartbeatEnabledFlagName = fmt.Sprintf("%v.heartbeatEnabled", cmdPrefix)
	}

	var heartbeatEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(heartbeatEnabledFlagName, heartbeatEnabledFlagDefault, heartbeatEnabledDescription)

	return nil
}

func registerServerSettingsHostAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	hostAddressDescription := ``

	var hostAddressFlagName string
	if cmdPrefix == "" {
		hostAddressFlagName = "hostAddress"
	} else {
		hostAddressFlagName = fmt.Sprintf("%v.hostAddress", cmdPrefix)
	}

	var hostAddressFlagDefault string

	_ = cmd.PersistentFlags().String(hostAddressFlagName, hostAddressFlagDefault, hostAddressDescription)

	return nil
}

func registerServerSettingsJwksURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	jwksUrlDescription := ``

	var jwksUrlFlagName string
	if cmdPrefix == "" {
		jwksUrlFlagName = "jwksURL"
	} else {
		jwksUrlFlagName = fmt.Sprintf("%v.jwksURL", cmdPrefix)
	}

	var jwksUrlFlagDefault string

	_ = cmd.PersistentFlags().String(jwksUrlFlagName, jwksUrlFlagDefault, jwksUrlDescription)

	return nil
}

func registerServerSettingsJwtServerControlEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	jwtServerControlEnabledDescription := ``

	var jwtServerControlEnabledFlagName string
	if cmdPrefix == "" {
		jwtServerControlEnabledFlagName = "jwtServerControlEnabled"
	} else {
		jwtServerControlEnabledFlagName = fmt.Sprintf("%v.jwtServerControlEnabled", cmdPrefix)
	}

	var jwtServerControlEnabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(jwtServerControlEnabledFlagName, jwtServerControlEnabledFlagDefault, jwtServerControlEnabledDescription)

	return nil
}

func registerServerSettingsJwtServerSecretKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	jwtServerSecretKeyDescription := ``

	var jwtServerSecretKeyFlagName string
	if cmdPrefix == "" {
		jwtServerSecretKeyFlagName = "jwtServerSecretKey"
	} else {
		jwtServerSecretKeyFlagName = fmt.Sprintf("%v.jwtServerSecretKey", cmdPrefix)
	}

	var jwtServerSecretKeyFlagDefault string

	_ = cmd.PersistentFlags().String(jwtServerSecretKeyFlagName, jwtServerSecretKeyFlagDefault, jwtServerSecretKeyDescription)

	return nil
}

func registerServerSettingsLicenceKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	licenceKeyDescription := ``

	var licenceKeyFlagName string
	if cmdPrefix == "" {
		licenceKeyFlagName = "licenceKey"
	} else {
		licenceKeyFlagName = fmt.Sprintf("%v.licenceKey", cmdPrefix)
	}

	var licenceKeyFlagDefault string

	_ = cmd.PersistentFlags().String(licenceKeyFlagName, licenceKeyFlagDefault, licenceKeyDescription)

	return nil
}

func registerServerSettingsLogLevel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	logLevelDescription := ``

	var logLevelFlagName string
	if cmdPrefix == "" {
		logLevelFlagName = "logLevel"
	} else {
		logLevelFlagName = fmt.Sprintf("%v.logLevel", cmdPrefix)
	}

	var logLevelFlagDefault string

	_ = cmd.PersistentFlags().String(logLevelFlagName, logLevelFlagDefault, logLevelDescription)

	return nil
}

func registerServerSettingsMarketplace(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	marketplaceDescription := ``

	var marketplaceFlagName string
	if cmdPrefix == "" {
		marketplaceFlagName = "marketplace"
	} else {
		marketplaceFlagName = fmt.Sprintf("%v.marketplace", cmdPrefix)
	}

	var marketplaceFlagDefault string

	_ = cmd.PersistentFlags().String(marketplaceFlagName, marketplaceFlagDefault, marketplaceDescription)

	return nil
}

func registerServerSettingsNativeLogLevel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nativeLogLevelDescription := ``

	var nativeLogLevelFlagName string
	if cmdPrefix == "" {
		nativeLogLevelFlagName = "nativeLogLevel"
	} else {
		nativeLogLevelFlagName = fmt.Sprintf("%v.nativeLogLevel", cmdPrefix)
	}

	var nativeLogLevelFlagDefault string

	_ = cmd.PersistentFlags().String(nativeLogLevelFlagName, nativeLogLevelFlagDefault, nativeLogLevelDescription)

	return nil
}

func registerServerSettingsNodeGroup(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nodeGroupDescription := ``

	var nodeGroupFlagName string
	if cmdPrefix == "" {
		nodeGroupFlagName = "nodeGroup"
	} else {
		nodeGroupFlagName = fmt.Sprintf("%v.nodeGroup", cmdPrefix)
	}

	var nodeGroupFlagDefault string

	_ = cmd.PersistentFlags().String(nodeGroupFlagName, nodeGroupFlagDefault, nodeGroupDescription)

	return nil
}

func registerServerSettingsOriginServerPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	originServerPortDescription := ``

	var originServerPortFlagName string
	if cmdPrefix == "" {
		originServerPortFlagName = "originServerPort"
	} else {
		originServerPortFlagName = fmt.Sprintf("%v.originServerPort", cmdPrefix)
	}

	var originServerPortFlagDefault int32

	_ = cmd.PersistentFlags().Int32(originServerPortFlagName, originServerPortFlagDefault, originServerPortDescription)

	return nil
}

func registerServerSettingsProxyAddress(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	proxyAddressDescription := ``

	var proxyAddressFlagName string
	if cmdPrefix == "" {
		proxyAddressFlagName = "proxyAddress"
	} else {
		proxyAddressFlagName = fmt.Sprintf("%v.proxyAddress", cmdPrefix)
	}

	var proxyAddressFlagDefault string

	_ = cmd.PersistentFlags().String(proxyAddressFlagName, proxyAddressFlagDefault, proxyAddressDescription)

	return nil
}

func registerServerSettingsServerName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	serverNameDescription := ``

	var serverNameFlagName string
	if cmdPrefix == "" {
		serverNameFlagName = "serverName"
	} else {
		serverNameFlagName = fmt.Sprintf("%v.serverName", cmdPrefix)
	}

	var serverNameFlagDefault string

	_ = cmd.PersistentFlags().String(serverNameFlagName, serverNameFlagDefault, serverNameDescription)

	return nil
}

func registerServerSettingsSrtPort(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	srtPortDescription := ``

	var srtPortFlagName string
	if cmdPrefix == "" {
		srtPortFlagName = "srtPort"
	} else {
		srtPortFlagName = fmt.Sprintf("%v.srtPort", cmdPrefix)
	}

	var srtPortFlagDefault int32

	_ = cmd.PersistentFlags().Int32(srtPortFlagName, srtPortFlagDefault, srtPortDescription)

	return nil
}

func registerServerSettingsUseGlobalIP(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	useGlobalIpDescription := ``

	var useGlobalIpFlagName string
	if cmdPrefix == "" {
		useGlobalIpFlagName = "useGlobalIp"
	} else {
		useGlobalIpFlagName = fmt.Sprintf("%v.useGlobalIp", cmdPrefix)
	}

	var useGlobalIpFlagDefault bool

	_ = cmd.PersistentFlags().Bool(useGlobalIpFlagName, useGlobalIpFlagDefault, useGlobalIpDescription)

	return nil
}

func registerServerSettingsWebRTCLogLevel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	webRTCLogLevelDescription := `Enum: ["LS_VERBOSE","LS_INFO","LS_WARNING","LS_ERROR","LS_NONE"]. `

	var webRTCLogLevelFlagName string
	if cmdPrefix == "" {
		webRTCLogLevelFlagName = "webRTCLogLevel"
	} else {
		webRTCLogLevelFlagName = fmt.Sprintf("%v.webRTCLogLevel", cmdPrefix)
	}

	var webRTCLogLevelFlagDefault string

	_ = cmd.PersistentFlags().String(webRTCLogLevelFlagName, webRTCLogLevelFlagDefault, webRTCLogLevelDescription)

	if err := cmd.RegisterFlagCompletionFunc(webRTCLogLevelFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["LS_VERBOSE","LS_INFO","LS_WARNING","LS_ERROR","LS_NONE"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelServerSettingsFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, allowedDashboardCIdRAdded := retrieveServerSettingsAllowedDashboardCIDRFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || allowedDashboardCIdRAdded

	err, buildForMarketAdded := retrieveServerSettingsBuildForMarketFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || buildForMarketAdded

	err, cpuMeasurementPeriodMsAdded := retrieveServerSettingsCPUMeasurementPeriodMsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cpuMeasurementPeriodMsAdded

	err, cpuMeasurementWindowSizeAdded := retrieveServerSettingsCPUMeasurementWindowSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || cpuMeasurementWindowSizeAdded

	err, defaultHttpPortAdded := retrieveServerSettingsDefaultHTTPPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || defaultHttpPortAdded

	err, heartbeatEnabledAdded := retrieveServerSettingsHeartbeatEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || heartbeatEnabledAdded

	err, hostAddressAdded := retrieveServerSettingsHostAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || hostAddressAdded

	err, jwksUrlAdded := retrieveServerSettingsJwksURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || jwksUrlAdded

	err, jwtServerControlEnabledAdded := retrieveServerSettingsJwtServerControlEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || jwtServerControlEnabledAdded

	err, jwtServerSecretKeyAdded := retrieveServerSettingsJwtServerSecretKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || jwtServerSecretKeyAdded

	err, licenceKeyAdded := retrieveServerSettingsLicenceKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || licenceKeyAdded

	err, logLevelAdded := retrieveServerSettingsLogLevelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || logLevelAdded

	err, marketplaceAdded := retrieveServerSettingsMarketplaceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || marketplaceAdded

	err, nativeLogLevelAdded := retrieveServerSettingsNativeLogLevelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nativeLogLevelAdded

	err, nodeGroupAdded := retrieveServerSettingsNodeGroupFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nodeGroupAdded

	err, originServerPortAdded := retrieveServerSettingsOriginServerPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || originServerPortAdded

	err, proxyAddressAdded := retrieveServerSettingsProxyAddressFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || proxyAddressAdded

	err, serverNameAdded := retrieveServerSettingsServerNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || serverNameAdded

	err, srtPortAdded := retrieveServerSettingsSrtPortFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || srtPortAdded

	err, useGlobalIpAdded := retrieveServerSettingsUseGlobalIPFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || useGlobalIpAdded

	err, webRTCLogLevelAdded := retrieveServerSettingsWebRTCLogLevelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || webRTCLogLevelAdded

	return nil, retAdded
}

func retrieveServerSettingsAllowedDashboardCIDRFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	allowedDashboardCIdRFlagName := fmt.Sprintf("%v.allowedDashboardCIDR", cmdPrefix)
	if cmd.Flags().Changed(allowedDashboardCIdRFlagName) {

		var allowedDashboardCIdRFlagName string
		if cmdPrefix == "" {
			allowedDashboardCIdRFlagName = "allowedDashboardCIDR"
		} else {
			allowedDashboardCIdRFlagName = fmt.Sprintf("%v.allowedDashboardCIDR", cmdPrefix)
		}

		allowedDashboardCIdRFlagValue, err := cmd.Flags().GetString(allowedDashboardCIdRFlagName)
		if err != nil {
			return err, false
		}
		m.AllowedDashboardCIDR = allowedDashboardCIdRFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsBuildForMarketFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	buildForMarketFlagName := fmt.Sprintf("%v.buildForMarket", cmdPrefix)
	if cmd.Flags().Changed(buildForMarketFlagName) {

		var buildForMarketFlagName string
		if cmdPrefix == "" {
			buildForMarketFlagName = "buildForMarket"
		} else {
			buildForMarketFlagName = fmt.Sprintf("%v.buildForMarket", cmdPrefix)
		}

		buildForMarketFlagValue, err := cmd.Flags().GetBool(buildForMarketFlagName)
		if err != nil {
			return err, false
		}
		m.BuildForMarket = buildForMarketFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsCPUMeasurementPeriodMsFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cpuMeasurementPeriodMsFlagName := fmt.Sprintf("%v.cpuMeasurementPeriodMs", cmdPrefix)
	if cmd.Flags().Changed(cpuMeasurementPeriodMsFlagName) {

		var cpuMeasurementPeriodMsFlagName string
		if cmdPrefix == "" {
			cpuMeasurementPeriodMsFlagName = "cpuMeasurementPeriodMs"
		} else {
			cpuMeasurementPeriodMsFlagName = fmt.Sprintf("%v.cpuMeasurementPeriodMs", cmdPrefix)
		}

		cpuMeasurementPeriodMsFlagValue, err := cmd.Flags().GetInt32(cpuMeasurementPeriodMsFlagName)
		if err != nil {
			return err, false
		}
		m.CPUMeasurementPeriodMs = cpuMeasurementPeriodMsFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsCPUMeasurementWindowSizeFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	cpuMeasurementWindowSizeFlagName := fmt.Sprintf("%v.cpuMeasurementWindowSize", cmdPrefix)
	if cmd.Flags().Changed(cpuMeasurementWindowSizeFlagName) {

		var cpuMeasurementWindowSizeFlagName string
		if cmdPrefix == "" {
			cpuMeasurementWindowSizeFlagName = "cpuMeasurementWindowSize"
		} else {
			cpuMeasurementWindowSizeFlagName = fmt.Sprintf("%v.cpuMeasurementWindowSize", cmdPrefix)
		}

		cpuMeasurementWindowSizeFlagValue, err := cmd.Flags().GetInt32(cpuMeasurementWindowSizeFlagName)
		if err != nil {
			return err, false
		}
		m.CPUMeasurementWindowSize = cpuMeasurementWindowSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsDefaultHTTPPortFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	defaultHttpPortFlagName := fmt.Sprintf("%v.defaultHttpPort", cmdPrefix)
	if cmd.Flags().Changed(defaultHttpPortFlagName) {

		var defaultHttpPortFlagName string
		if cmdPrefix == "" {
			defaultHttpPortFlagName = "defaultHttpPort"
		} else {
			defaultHttpPortFlagName = fmt.Sprintf("%v.defaultHttpPort", cmdPrefix)
		}

		defaultHttpPortFlagValue, err := cmd.Flags().GetInt32(defaultHttpPortFlagName)
		if err != nil {
			return err, false
		}
		m.DefaultHTTPPort = defaultHttpPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsHeartbeatEnabledFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	heartbeatEnabledFlagName := fmt.Sprintf("%v.heartbeatEnabled", cmdPrefix)
	if cmd.Flags().Changed(heartbeatEnabledFlagName) {

		var heartbeatEnabledFlagName string
		if cmdPrefix == "" {
			heartbeatEnabledFlagName = "heartbeatEnabled"
		} else {
			heartbeatEnabledFlagName = fmt.Sprintf("%v.heartbeatEnabled", cmdPrefix)
		}

		heartbeatEnabledFlagValue, err := cmd.Flags().GetBool(heartbeatEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.HeartbeatEnabled = heartbeatEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsHostAddressFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	hostAddressFlagName := fmt.Sprintf("%v.hostAddress", cmdPrefix)
	if cmd.Flags().Changed(hostAddressFlagName) {

		var hostAddressFlagName string
		if cmdPrefix == "" {
			hostAddressFlagName = "hostAddress"
		} else {
			hostAddressFlagName = fmt.Sprintf("%v.hostAddress", cmdPrefix)
		}

		hostAddressFlagValue, err := cmd.Flags().GetString(hostAddressFlagName)
		if err != nil {
			return err, false
		}
		m.HostAddress = hostAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsJwksURLFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	jwksUrlFlagName := fmt.Sprintf("%v.jwksURL", cmdPrefix)
	if cmd.Flags().Changed(jwksUrlFlagName) {

		var jwksUrlFlagName string
		if cmdPrefix == "" {
			jwksUrlFlagName = "jwksURL"
		} else {
			jwksUrlFlagName = fmt.Sprintf("%v.jwksURL", cmdPrefix)
		}

		jwksUrlFlagValue, err := cmd.Flags().GetString(jwksUrlFlagName)
		if err != nil {
			return err, false
		}
		m.JwksURL = jwksUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsJwtServerControlEnabledFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	jwtServerControlEnabledFlagName := fmt.Sprintf("%v.jwtServerControlEnabled", cmdPrefix)
	if cmd.Flags().Changed(jwtServerControlEnabledFlagName) {

		var jwtServerControlEnabledFlagName string
		if cmdPrefix == "" {
			jwtServerControlEnabledFlagName = "jwtServerControlEnabled"
		} else {
			jwtServerControlEnabledFlagName = fmt.Sprintf("%v.jwtServerControlEnabled", cmdPrefix)
		}

		jwtServerControlEnabledFlagValue, err := cmd.Flags().GetBool(jwtServerControlEnabledFlagName)
		if err != nil {
			return err, false
		}
		m.JwtServerControlEnabled = jwtServerControlEnabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsJwtServerSecretKeyFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	jwtServerSecretKeyFlagName := fmt.Sprintf("%v.jwtServerSecretKey", cmdPrefix)
	if cmd.Flags().Changed(jwtServerSecretKeyFlagName) {

		var jwtServerSecretKeyFlagName string
		if cmdPrefix == "" {
			jwtServerSecretKeyFlagName = "jwtServerSecretKey"
		} else {
			jwtServerSecretKeyFlagName = fmt.Sprintf("%v.jwtServerSecretKey", cmdPrefix)
		}

		jwtServerSecretKeyFlagValue, err := cmd.Flags().GetString(jwtServerSecretKeyFlagName)
		if err != nil {
			return err, false
		}
		m.JwtServerSecretKey = jwtServerSecretKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsLicenceKeyFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	licenceKeyFlagName := fmt.Sprintf("%v.licenceKey", cmdPrefix)
	if cmd.Flags().Changed(licenceKeyFlagName) {

		var licenceKeyFlagName string
		if cmdPrefix == "" {
			licenceKeyFlagName = "licenceKey"
		} else {
			licenceKeyFlagName = fmt.Sprintf("%v.licenceKey", cmdPrefix)
		}

		licenceKeyFlagValue, err := cmd.Flags().GetString(licenceKeyFlagName)
		if err != nil {
			return err, false
		}
		m.LicenceKey = licenceKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsLogLevelFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	logLevelFlagName := fmt.Sprintf("%v.logLevel", cmdPrefix)
	if cmd.Flags().Changed(logLevelFlagName) {

		var logLevelFlagName string
		if cmdPrefix == "" {
			logLevelFlagName = "logLevel"
		} else {
			logLevelFlagName = fmt.Sprintf("%v.logLevel", cmdPrefix)
		}

		logLevelFlagValue, err := cmd.Flags().GetString(logLevelFlagName)
		if err != nil {
			return err, false
		}
		m.LogLevel = logLevelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsMarketplaceFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	marketplaceFlagName := fmt.Sprintf("%v.marketplace", cmdPrefix)
	if cmd.Flags().Changed(marketplaceFlagName) {

		var marketplaceFlagName string
		if cmdPrefix == "" {
			marketplaceFlagName = "marketplace"
		} else {
			marketplaceFlagName = fmt.Sprintf("%v.marketplace", cmdPrefix)
		}

		marketplaceFlagValue, err := cmd.Flags().GetString(marketplaceFlagName)
		if err != nil {
			return err, false
		}
		m.Marketplace = marketplaceFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsNativeLogLevelFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nativeLogLevelFlagName := fmt.Sprintf("%v.nativeLogLevel", cmdPrefix)
	if cmd.Flags().Changed(nativeLogLevelFlagName) {

		var nativeLogLevelFlagName string
		if cmdPrefix == "" {
			nativeLogLevelFlagName = "nativeLogLevel"
		} else {
			nativeLogLevelFlagName = fmt.Sprintf("%v.nativeLogLevel", cmdPrefix)
		}

		nativeLogLevelFlagValue, err := cmd.Flags().GetString(nativeLogLevelFlagName)
		if err != nil {
			return err, false
		}
		m.NativeLogLevel = nativeLogLevelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsNodeGroupFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nodeGroupFlagName := fmt.Sprintf("%v.nodeGroup", cmdPrefix)
	if cmd.Flags().Changed(nodeGroupFlagName) {

		var nodeGroupFlagName string
		if cmdPrefix == "" {
			nodeGroupFlagName = "nodeGroup"
		} else {
			nodeGroupFlagName = fmt.Sprintf("%v.nodeGroup", cmdPrefix)
		}

		nodeGroupFlagValue, err := cmd.Flags().GetString(nodeGroupFlagName)
		if err != nil {
			return err, false
		}
		m.NodeGroup = nodeGroupFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsOriginServerPortFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	originServerPortFlagName := fmt.Sprintf("%v.originServerPort", cmdPrefix)
	if cmd.Flags().Changed(originServerPortFlagName) {

		var originServerPortFlagName string
		if cmdPrefix == "" {
			originServerPortFlagName = "originServerPort"
		} else {
			originServerPortFlagName = fmt.Sprintf("%v.originServerPort", cmdPrefix)
		}

		originServerPortFlagValue, err := cmd.Flags().GetInt32(originServerPortFlagName)
		if err != nil {
			return err, false
		}
		m.OriginServerPort = originServerPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsProxyAddressFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	proxyAddressFlagName := fmt.Sprintf("%v.proxyAddress", cmdPrefix)
	if cmd.Flags().Changed(proxyAddressFlagName) {

		var proxyAddressFlagName string
		if cmdPrefix == "" {
			proxyAddressFlagName = "proxyAddress"
		} else {
			proxyAddressFlagName = fmt.Sprintf("%v.proxyAddress", cmdPrefix)
		}

		proxyAddressFlagValue, err := cmd.Flags().GetString(proxyAddressFlagName)
		if err != nil {
			return err, false
		}
		m.ProxyAddress = proxyAddressFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsServerNameFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	serverNameFlagName := fmt.Sprintf("%v.serverName", cmdPrefix)
	if cmd.Flags().Changed(serverNameFlagName) {

		var serverNameFlagName string
		if cmdPrefix == "" {
			serverNameFlagName = "serverName"
		} else {
			serverNameFlagName = fmt.Sprintf("%v.serverName", cmdPrefix)
		}

		serverNameFlagValue, err := cmd.Flags().GetString(serverNameFlagName)
		if err != nil {
			return err, false
		}
		m.ServerName = serverNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsSrtPortFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	srtPortFlagName := fmt.Sprintf("%v.srtPort", cmdPrefix)
	if cmd.Flags().Changed(srtPortFlagName) {

		var srtPortFlagName string
		if cmdPrefix == "" {
			srtPortFlagName = "srtPort"
		} else {
			srtPortFlagName = fmt.Sprintf("%v.srtPort", cmdPrefix)
		}

		srtPortFlagValue, err := cmd.Flags().GetInt32(srtPortFlagName)
		if err != nil {
			return err, false
		}
		m.SrtPort = srtPortFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsUseGlobalIPFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	useGlobalIpFlagName := fmt.Sprintf("%v.useGlobalIp", cmdPrefix)
	if cmd.Flags().Changed(useGlobalIpFlagName) {

		var useGlobalIpFlagName string
		if cmdPrefix == "" {
			useGlobalIpFlagName = "useGlobalIp"
		} else {
			useGlobalIpFlagName = fmt.Sprintf("%v.useGlobalIp", cmdPrefix)
		}

		useGlobalIpFlagValue, err := cmd.Flags().GetBool(useGlobalIpFlagName)
		if err != nil {
			return err, false
		}
		m.UseGlobalIP = useGlobalIpFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveServerSettingsWebRTCLogLevelFlags(depth int, m *models.ServerSettings, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	webRTCLogLevelFlagName := fmt.Sprintf("%v.webRTCLogLevel", cmdPrefix)
	if cmd.Flags().Changed(webRTCLogLevelFlagName) {

		var webRTCLogLevelFlagName string
		if cmdPrefix == "" {
			webRTCLogLevelFlagName = "webRTCLogLevel"
		} else {
			webRTCLogLevelFlagName = fmt.Sprintf("%v.webRTCLogLevel", cmdPrefix)
		}

		webRTCLogLevelFlagValue, err := cmd.Flags().GetString(webRTCLogLevelFlagName)
		if err != nil {
			return err, false
		}
		m.WebRTCLogLevel = webRTCLogLevelFlagValue

		retAdded = true
	}

	return nil, retAdded
}