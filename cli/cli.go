// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/umutbasal/antmedia/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debug flag indicating that cli should output debug logs
var debug bool

// config file location
var configFile string

// dry run flag
var dryRun bool

// name of the executable
var exeName string = filepath.Base(os.Args[0])

// logDebugf writes debug log to stdout
func logDebugf(format string, v ...interface{}) {
	if !debug {
		return
	}
	log.Printf(format, v...)
}

// depth of recursion to construct model flags
var maxDepth int = 5

// makeClient constructs a client object
func makeClient(cmd *cobra.Command, args []string) (*client.AntMediaServerRESTAPIReference, error) {
	hostname := viper.GetString("hostname")
	scheme := viper.GetString("scheme")
	base := viper.GetString("base")

	r := httptransport.New(hostname, base, []string{scheme})
	r.SetDebug(debug)
	// set custom producer and consumer to use the default ones

	r.Consumers["application/json"] = runtime.JSONConsumer()

	// warning: consumes multipart/form-data is not supported by go-swagger cli yet

	// warning: consumes application/x-www-form-urlencoded is not supported by go-swagger cli yet

	// warning: produces application/octet-stream is not supported by go-swagger cli yet

	r.Producers["application/json"] = runtime.JSONProducer()

	// warning: produces text/plain is not supported by go-swagger cli yet

	appCli := client.New(r, strfmt.Default)
	logDebugf("Server url: %v://%v/%v", scheme, hostname, base)
	return appCli, nil
}

// MakeRootCmd returns the root cmd
func MakeRootCmd() (*cobra.Command, error) {
	cobra.OnInitialize(initViperConfigs)

	// Use executable name as the command name
	rootCmd := &cobra.Command{
		Use: exeName,
	}

	// register basic flags
	rootCmd.PersistentFlags().String("hostname", client.DefaultHost, "hostname of the service")
	viper.BindPFlag("hostname", rootCmd.PersistentFlags().Lookup("hostname"))
	rootCmd.PersistentFlags().String("scheme", client.DefaultSchemes[0], fmt.Sprintf("Choose from: %v", client.DefaultSchemes))
	viper.BindPFlag("scheme", rootCmd.PersistentFlags().Lookup("scheme"))
	rootCmd.PersistentFlags().String("base", client.DefaultBasePath, fmt.Sprintf("Choose from: %v", client.DefaultBasePath))
	viper.BindPFlag("base", rootCmd.PersistentFlags().Lookup("base"))

	// configure debug flag
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "output debug logs")
	// configure config location
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file path")
	// configure dry run flag
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "do not send the request to server")

	// register security flags
	// add all operation groups
	operationGroupBroadcastRestServiceCmd, err := makeOperationGroupBroadcastRestServiceCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupBroadcastRestServiceCmd)

	operationGroupManagementRestServiceCmd, err := makeOperationGroupManagementRestServiceCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupManagementRestServiceCmd)

	operationGroupOperationsCmd, err := makeOperationGroupOperationsCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupOperationsCmd)

	operationGroupVodRestServiceCmd, err := makeOperationGroupVodRestServiceCmd()
	if err != nil {
		return nil, err
	}
	rootCmd.AddCommand(operationGroupVodRestServiceCmd)

	// add cobra completion
	rootCmd.AddCommand(makeGenCompletionCmd())

	return rootCmd, nil
}

// initViperConfigs initialize viper config using config file in '$HOME/.config/<cli name>/config.<json|yaml...>'
// currently hostname, scheme and auth tokens can be specified in this config file.
func initViperConfigs() {
	if configFile != "" {
		// use user specified config file location
		viper.SetConfigFile(configFile)
	} else {
		// look for default config
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(path.Join(home, ".config", exeName))
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		logDebugf("Error: loading config file: %v", err)
		return
	}
	logDebugf("Using config file: %v", viper.ConfigFileUsed())
}

func makeOperationGroupBroadcastRestServiceCmd() (*cobra.Command, error) {
	operationGroupBroadcastRestServiceCmd := &cobra.Command{
		Use:  "broadcast_rest_service",
		Long: ``,
	}

	operationAddEndpointV2Cmd, err := makeOperationBroadcastRestServiceAddEndpointV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationAddEndpointV2Cmd)

	operationAddEndpointV3Cmd, err := makeOperationBroadcastRestServiceAddEndpointV3Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationAddEndpointV3Cmd)

	operationAddStreamToTheRoomCmd, err := makeOperationBroadcastRestServiceAddStreamToTheRoomCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationAddStreamToTheRoomCmd)

	operationAddSubTrackCmd, err := makeOperationBroadcastRestServiceAddSubTrackCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationAddSubTrackCmd)

	operationAddSubscriberCmd, err := makeOperationBroadcastRestServiceAddSubscriberCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationAddSubscriberCmd)

	operationCreateBroadcastCmd, err := makeOperationBroadcastRestServiceCreateBroadcastCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationCreateBroadcastCmd)

	operationCreateConferenceRoomV2Cmd, err := makeOperationBroadcastRestServiceCreateConferenceRoomV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationCreateConferenceRoomV2Cmd)

	operationDeleteBroadcastCmd, err := makeOperationBroadcastRestServiceDeleteBroadcastCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationDeleteBroadcastCmd)

	operationDeleteBroadcastsCmd, err := makeOperationBroadcastRestServiceDeleteBroadcastsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationDeleteBroadcastsCmd)

	operationDeleteConferenceRoomV2Cmd, err := makeOperationBroadcastRestServiceDeleteConferenceRoomV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationDeleteConferenceRoomV2Cmd)

	operationDeleteStreamFromTheRoomCmd, err := makeOperationBroadcastRestServiceDeleteStreamFromTheRoomCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationDeleteStreamFromTheRoomCmd)

	operationDeleteSubscriberCmd, err := makeOperationBroadcastRestServiceDeleteSubscriberCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationDeleteSubscriberCmd)

	operationEditConferenceRoomCmd, err := makeOperationBroadcastRestServiceEditConferenceRoomCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationEditConferenceRoomCmd)

	operationEnableRecordingCmd, err := makeOperationBroadcastRestServiceEnableRecordingCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationEnableRecordingCmd)

	operationFilterBroadcastListV2Cmd, err := makeOperationBroadcastRestServiceFilterBroadcastListV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationFilterBroadcastListV2Cmd)

	operationGetAppLiveStatisticsCmd, err := makeOperationBroadcastRestServiceGetAppLiveStatisticsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetAppLiveStatisticsCmd)

	operationGetBroadcastCmd, err := makeOperationBroadcastRestServiceGetBroadcastCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetBroadcastCmd)

	operationGetBroadcastListCmd, err := makeOperationBroadcastRestServiceGetBroadcastListCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetBroadcastListCmd)

	operationGetBroadcastStatisticsCmd, err := makeOperationBroadcastRestServiceGetBroadcastStatisticsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetBroadcastStatisticsCmd)

	operationGetBroadcastTotalStatisticsCmd, err := makeOperationBroadcastRestServiceGetBroadcastTotalStatisticsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetBroadcastTotalStatisticsCmd)

	operationGetCameraErrorV2Cmd, err := makeOperationBroadcastRestServiceGetCameraErrorV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetCameraErrorV2Cmd)

	operationGetConferenceRoomCmd, err := makeOperationBroadcastRestServiceGetConferenceRoomCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetConferenceRoomCmd)

	operationGetConferenceRoomListCmd, err := makeOperationBroadcastRestServiceGetConferenceRoomListCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetConferenceRoomListCmd)

	operationGetDetectionListV2Cmd, err := makeOperationBroadcastRestServiceGetDetectionListV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetDetectionListV2Cmd)

	operationGetJwtTokenV2Cmd, err := makeOperationBroadcastRestServiceGetJwtTokenV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetJwtTokenV2Cmd)

	operationGetObjectDetectedTotalCmd, err := makeOperationBroadcastRestServiceGetObjectDetectedTotalCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetObjectDetectedTotalCmd)

	operationGetRTMPToWebRTCStatsCmd, err := makeOperationBroadcastRestServiceGetRTMPToWebRTCStatsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetRTMPToWebRTCStatsCmd)

	operationGetRoomInfoCmd, err := makeOperationBroadcastRestServiceGetRoomInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetRoomInfoCmd)

	operationGetStreamInfoCmd, err := makeOperationBroadcastRestServiceGetStreamInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetStreamInfoCmd)

	operationGetTokenV2Cmd, err := makeOperationBroadcastRestServiceGetTokenV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetTokenV2Cmd)

	operationGetTotalBroadcastNumberV2Cmd, err := makeOperationBroadcastRestServiceGetTotalBroadcastNumberV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetTotalBroadcastNumberV2Cmd)

	operationGetWebRTCClientStatsListV2Cmd, err := makeOperationBroadcastRestServiceGetWebRTCClientStatsListV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetWebRTCClientStatsListV2Cmd)

	operationGetWebRTCLowLevelReceiveStatsCmd, err := makeOperationBroadcastRestServiceGetWebRTCLowLevelReceiveStatsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetWebRTCLowLevelReceiveStatsCmd)

	operationGetWebRTCLowLevelSendStatsCmd, err := makeOperationBroadcastRestServiceGetWebRTCLowLevelSendStatsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetWebRTCLowLevelSendStatsCmd)

	operationGetWebRTCViewerListCmd, err := makeOperationBroadcastRestServiceGetWebRTCViewerListCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationGetWebRTCViewerListCmd)

	operationImportLiveStreams2StalkerV2Cmd, err := makeOperationBroadcastRestServiceImportLiveStreams2StalkerV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationImportLiveStreams2StalkerV2Cmd)

	operationListSubscriberStatsV2Cmd, err := makeOperationBroadcastRestServiceListSubscriberStatsV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationListSubscriberStatsV2Cmd)

	operationListSubscriberV2Cmd, err := makeOperationBroadcastRestServiceListSubscriberV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationListSubscriberV2Cmd)

	operationListTokensV2Cmd, err := makeOperationBroadcastRestServiceListTokensV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationListTokensV2Cmd)

	operationMoveIPCameraCmd, err := makeOperationBroadcastRestServiceMoveIPCameraCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationMoveIPCameraCmd)

	operationRemoveEndpointCmd, err := makeOperationBroadcastRestServiceRemoveEndpointCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationRemoveEndpointCmd)

	operationRemoveEndpointV2Cmd, err := makeOperationBroadcastRestServiceRemoveEndpointV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationRemoveEndpointV2Cmd)

	operationRevokeSubscribersCmd, err := makeOperationBroadcastRestServiceRevokeSubscribersCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationRevokeSubscribersCmd)

	operationRevokeTokensV2Cmd, err := makeOperationBroadcastRestServiceRevokeTokensV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationRevokeTokensV2Cmd)

	operationSearchOnvifDevicesV2Cmd, err := makeOperationBroadcastRestServiceSearchOnvifDevicesV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationSearchOnvifDevicesV2Cmd)

	operationSendMessageCmd, err := makeOperationBroadcastRestServiceSendMessageCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationSendMessageCmd)

	operationStartStreamSourceV2Cmd, err := makeOperationBroadcastRestServiceStartStreamSourceV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationStartStreamSourceV2Cmd)

	operationStopMoveCmd, err := makeOperationBroadcastRestServiceStopMoveCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationStopMoveCmd)

	operationStopPlayingCmd, err := makeOperationBroadcastRestServiceStopPlayingCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationStopPlayingCmd)

	operationStopStreamingV2Cmd, err := makeOperationBroadcastRestServiceStopStreamingV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationStopStreamingV2Cmd)

	operationUpdateBroadcastCmd, err := makeOperationBroadcastRestServiceUpdateBroadcastCmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationUpdateBroadcastCmd)

	operationValidateTokenV2Cmd, err := makeOperationBroadcastRestServiceValidateTokenV2Cmd()
	if err != nil {
		return nil, err
	}
	operationGroupBroadcastRestServiceCmd.AddCommand(operationValidateTokenV2Cmd)

	return operationGroupBroadcastRestServiceCmd, nil
}
func makeOperationGroupManagementRestServiceCmd() (*cobra.Command, error) {
	operationGroupManagementRestServiceCmd := &cobra.Command{
		Use:  "management_rest_service",
		Long: ``,
	}

	operationPutV2UsersCmd, err := makeOperationManagementRestServicePutV2UsersCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationPutV2UsersCmd)

	operationAddInitialUserCmd, err := makeOperationManagementRestServiceAddInitialUserCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationAddInitialUserCmd)

	operationAddUserCmd, err := makeOperationManagementRestServiceAddUserCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationAddUserCmd)

	operationAuthenticateUserCmd, err := makeOperationManagementRestServiceAuthenticateUserCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationAuthenticateUserCmd)

	operationChangeServerSettingsCmd, err := makeOperationManagementRestServiceChangeServerSettingsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationChangeServerSettingsCmd)

	operationChangeSettingsCmd, err := makeOperationManagementRestServiceChangeSettingsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationChangeSettingsCmd)

	operationChangeUserPasswordCmd, err := makeOperationManagementRestServiceChangeUserPasswordCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationChangeUserPasswordCmd)

	operationCreateApplicationCmd, err := makeOperationManagementRestServiceCreateApplicationCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationCreateApplicationCmd)

	operationDeleteApplicationCmd, err := makeOperationManagementRestServiceDeleteApplicationCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationDeleteApplicationCmd)

	operationDeleteUserCmd, err := makeOperationManagementRestServiceDeleteUserCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationDeleteUserCmd)

	operationGetAppLiveStreamsCmd, err := makeOperationManagementRestServiceGetAppLiveStreamsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetAppLiveStreamsCmd)

	operationGetApplicationInfoCmd, err := makeOperationManagementRestServiceGetApplicationInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetApplicationInfoCmd)

	operationGetApplicationsCmd, err := makeOperationManagementRestServiceGetApplicationsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetApplicationsCmd)

	operationGetBlockedStatusCmd, err := makeOperationManagementRestServiceGetBlockedStatusCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetBlockedStatusCmd)

	operationGetCPUInfoCmd, err := makeOperationManagementRestServiceGetCPUInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetCPUInfoCmd)

	operationGetFileSystemInfoCmd, err := makeOperationManagementRestServiceGetFileSystemInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetFileSystemInfoCmd)

	operationGetGPUInfoCmd, err := makeOperationManagementRestServiceGetGPUInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetGPUInfoCmd)

	operationGetHeapDumpCmd, err := makeOperationManagementRestServiceGetHeapDumpCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetHeapDumpCmd)

	operationGetJVMMemoryInfoCmd, err := makeOperationManagementRestServiceGetJVMMemoryInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetJVMMemoryInfoCmd)

	operationGetLicenceStatusCmd, err := makeOperationManagementRestServiceGetLicenceStatusCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetLicenceStatusCmd)

	operationGetLiveClientsSizeCmd, err := makeOperationManagementRestServiceGetLiveClientsSizeCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetLiveClientsSizeCmd)

	operationGetLogFileCmd, err := makeOperationManagementRestServiceGetLogFileCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetLogFileCmd)

	operationGetServerSettingsCmd, err := makeOperationManagementRestServiceGetServerSettingsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetServerSettingsCmd)

	operationGetServerTimeCmd, err := makeOperationManagementRestServiceGetServerTimeCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetServerTimeCmd)

	operationGetSettingsCmd, err := makeOperationManagementRestServiceGetSettingsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetSettingsCmd)

	operationGetSystemInfoCmd, err := makeOperationManagementRestServiceGetSystemInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetSystemInfoCmd)

	operationGetSystemMemoryInfoCmd, err := makeOperationManagementRestServiceGetSystemMemoryInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetSystemMemoryInfoCmd)

	operationGetSystemResourcesInfoCmd, err := makeOperationManagementRestServiceGetSystemResourcesInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetSystemResourcesInfoCmd)

	operationGetThreadDumpJSONCmd, err := makeOperationManagementRestServiceGetThreadDumpJSONCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetThreadDumpJSONCmd)

	operationGetThreadsInfoCmd, err := makeOperationManagementRestServiceGetThreadsInfoCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetThreadsInfoCmd)

	operationGetUserListCmd, err := makeOperationManagementRestServiceGetUserListCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetUserListCmd)

	operationGetVersionCmd, err := makeOperationManagementRestServiceGetVersionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationGetVersionCmd)

	operationIsAdminCmd, err := makeOperationManagementRestServiceIsAdminCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationIsAdminCmd)

	operationIsAuthenticatedRestCmd, err := makeOperationManagementRestServiceIsAuthenticatedRestCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationIsAuthenticatedRestCmd)

	operationIsEnterpriseEditionCmd, err := makeOperationManagementRestServiceIsEnterpriseEditionCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationIsEnterpriseEditionCmd)

	operationIsFirstLoginCmd, err := makeOperationManagementRestServiceIsFirstLoginCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationIsFirstLoginCmd)

	operationIsInClusterModeCmd, err := makeOperationManagementRestServiceIsInClusterModeCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationIsInClusterModeCmd)

	operationIsShutdownProperlyCmd, err := makeOperationManagementRestServiceIsShutdownProperlyCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationIsShutdownProperlyCmd)

	operationResetBroadcastCmd, err := makeOperationManagementRestServiceResetBroadcastCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationResetBroadcastCmd)

	operationSetShutdownStatusCmd, err := makeOperationManagementRestServiceSetShutdownStatusCmd()
	if err != nil {
		return nil, err
	}
	operationGroupManagementRestServiceCmd.AddCommand(operationSetShutdownStatusCmd)

	return operationGroupManagementRestServiceCmd, nil
}
func makeOperationGroupOperationsCmd() (*cobra.Command, error) {
	operationGroupOperationsCmd := &cobra.Command{
		Use:  "operations",
		Long: ``,
	}

	operationGetClusterDeleteNodeIDCmd, err := makeOperationOperationsGetClusterDeleteNodeIDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationGetClusterDeleteNodeIDCmd)

	operationChangeLogSettingsCmd, err := makeOperationOperationsChangeLogSettingsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationChangeLogSettingsCmd)

	operationDeleteNodeCmd, err := makeOperationOperationsDeleteNodeCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationDeleteNodeCmd)

	operationDeleteVoDStreamCmd, err := makeOperationOperationsDeleteVoDStreamCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationDeleteVoDStreamCmd)

	operationEditUserCmd, err := makeOperationOperationsEditUserCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationEditUserCmd)

	operationGetNodeCountCmd, err := makeOperationOperationsGetNodeCountCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationGetNodeCountCmd)

	operationGetNodeListCmd, err := makeOperationOperationsGetNodeListCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationGetNodeListCmd)

	operationGetShutdownStatusCmd, err := makeOperationOperationsGetShutdownStatusCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationGetShutdownStatusCmd)

	operationGetThreadDumpCmd, err := makeOperationOperationsGetThreadDumpCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationGetThreadDumpCmd)

	operationSendSupportRequestCmd, err := makeOperationOperationsSendSupportRequestCmd()
	if err != nil {
		return nil, err
	}
	operationGroupOperationsCmd.AddCommand(operationSendSupportRequestCmd)

	return operationGroupOperationsCmd, nil
}
func makeOperationGroupVodRestServiceCmd() (*cobra.Command, error) {
	operationGroupVodRestServiceCmd := &cobra.Command{
		Use:  "vo_d_rest_service",
		Long: ``,
	}

	operationDeleteVoDCmd, err := makeOperationVodRestServiceDeleteVoDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVodRestServiceCmd.AddCommand(operationDeleteVoDCmd)

	operationDeleteVoDsCmd, err := makeOperationVodRestServiceDeleteVoDsCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVodRestServiceCmd.AddCommand(operationDeleteVoDsCmd)

	operationGetTotalVodNumberCmd, err := makeOperationVodRestServiceGetTotalVodNumberCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVodRestServiceCmd.AddCommand(operationGetTotalVodNumberCmd)

	operationGetVoDCmd, err := makeOperationVodRestServiceGetVoDCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVodRestServiceCmd.AddCommand(operationGetVoDCmd)

	operationGetVodListCmd, err := makeOperationVodRestServiceGetVodListCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVodRestServiceCmd.AddCommand(operationGetVodListCmd)

	operationImportVoDsToStalkerCmd, err := makeOperationVodRestServiceImportVoDsToStalkerCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVodRestServiceCmd.AddCommand(operationImportVoDsToStalkerCmd)

	operationSynchUserVodListCmd, err := makeOperationVodRestServiceSynchUserVodListCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVodRestServiceCmd.AddCommand(operationSynchUserVodListCmd)

	operationUploadVoDFileCmd, err := makeOperationVodRestServiceUploadVoDFileCmd()
	if err != nil {
		return nil, err
	}
	operationGroupVodRestServiceCmd.AddCommand(operationUploadVoDFileCmd)

	return operationGroupVodRestServiceCmd, nil
}
