// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppSettings app settings
//
// swagger:model AppSettings
type AppSettings struct {

	// aac encoding enabled
	AacEncodingEnabled bool `json:"aacEncodingEnabled,omitempty"`

	// accept only rooms in data store
	AcceptOnlyRoomsInDataStore bool `json:"acceptOnlyRoomsInDataStore,omitempty"`

	// accept only streams in data store
	AcceptOnlyStreamsInDataStore bool `json:"acceptOnlyStreamsInDataStore,omitempty"`

	// add date time to mp4 file name
	AddDateTimeToMp4FileName bool `json:"addDateTimeToMp4FileName,omitempty"`

	// allowed publisher c ID r
	AllowedPublisherCIDR string `json:"allowedPublisherCIDR,omitempty"`

	// app name
	AppName string `json:"appName,omitempty"`

	// audio bitrate s f u
	AudioBitrateSFU int32 `json:"audioBitrateSFU,omitempty"`

	// collect social media activity
	CollectSocialMediaActivity bool `json:"collectSocialMediaActivity,omitempty"`

	// constant rate factor
	ConstantRateFactor string `json:"constantRateFactor,omitempty"`

	// create preview period
	CreatePreviewPeriod int32 `json:"createPreviewPeriod,omitempty"`

	// dash extra window size
	DashExtraWindowSize string `json:"dashExtraWindowSize,omitempty"`

	// dash fragment duration
	DashFragmentDuration string `json:"dashFragmentDuration,omitempty"`

	// dash Http endpoint
	DashHTTPEndpoint string `json:"dashHttpEndpoint,omitempty"`

	// dash Http streaming
	DashHTTPStreaming bool `json:"dashHttpStreaming,omitempty"`

	// dash muxing enabled
	DashMuxingEnabled bool `json:"dashMuxingEnabled,omitempty"`

	// dash seg duration
	DashSegDuration string `json:"dashSegDuration,omitempty"`

	// dash window size
	DashWindowSize string `json:"dashWindowSize,omitempty"`

	// data channel enabled
	DataChannelEnabled bool `json:"dataChannelEnabled,omitempty"`

	// data channel player distribution
	DataChannelPlayerDistribution string `json:"dataChannelPlayerDistribution,omitempty"`

	// data channel web hook
	DataChannelWebHook string `json:"dataChannelWebHook,omitempty"`

	// default decoders enabled
	DefaultDecodersEnabled bool `json:"defaultDecodersEnabled,omitempty"`

	// delete d a s h files on ended
	DeleteDASHFilesOnEnded bool `json:"deleteDASHFilesOnEnded,omitempty"`

	// delete h l s files on ended
	DeleteHLSFilesOnEnded bool `json:"deleteHLSFilesOnEnded,omitempty"`

	// disable IPv6 candidates
	DisableIPV6Candidates bool `json:"disableIPv6Candidates,omitempty"`

	// enable time token for play
	EnableTimeTokenForPlay bool `json:"enableTimeTokenForPlay,omitempty"`

	// enable time token for publish
	EnableTimeTokenForPublish bool `json:"enableTimeTokenForPublish,omitempty"`

	// encoder level
	EncoderLevel string `json:"encoderLevel,omitempty"`

	// encoder name
	EncoderName string `json:"encoderName,omitempty"`

	// encoder preset
	EncoderPreset string `json:"encoderPreset,omitempty"`

	// encoder profile
	EncoderProfile string `json:"encoderProfile,omitempty"`

	// encoder rc
	EncoderRc string `json:"encoderRc,omitempty"`

	// encoder selection preference
	EncoderSelectionPreference string `json:"encoderSelectionPreference,omitempty"`

	// encoder settings
	EncoderSettings []*EncoderSettings `json:"encoderSettings"`

	// encoder settings string
	EncoderSettingsString string `json:"encoderSettingsString,omitempty"`

	// encoder specific
	EncoderSpecific string `json:"encoderSpecific,omitempty"`

	// encoder thread count
	EncoderThreadCount int32 `json:"encoderThreadCount,omitempty"`

	// encoder thread type
	EncoderThreadType int32 `json:"encoderThreadType,omitempty"`

	// encoding timeout
	EncodingTimeout int32 `json:"encodingTimeout,omitempty"`

	// endpoint health check period ms
	EndpointHealthCheckPeriodMs int32 `json:"endpointHealthCheckPeriodMs,omitempty"`

	// endpoint republish limit
	EndpointRepublishLimit int32 `json:"endpointRepublishLimit,omitempty"`

	// excessive bandwidth algorithm enabled
	ExcessiveBandwidthAlgorithmEnabled bool `json:"excessiveBandwidthAlgorithmEnabled,omitempty"`

	// excessive bandwidth call threshold
	ExcessiveBandwidthCallThreshold int32 `json:"excessiveBandwidthCallThreshold,omitempty"`

	// excessive bandwidth value
	ExcessiveBandwidthValue int32 `json:"excessiveBandwidthValue,omitempty"`

	// excessive bandwith try count before switchback
	ExcessiveBandwithTryCountBeforeSwitchback int32 `json:"excessiveBandwithTryCountBeforeSwitchback,omitempty"`

	// file name format
	FileNameFormat string `json:"fileNameFormat,omitempty"`

	// force aspect ratio in transcoding
	ForceAspectRatioInTranscoding bool `json:"forceAspectRatioInTranscoding,omitempty"`

	// force decoding
	ForceDecoding bool `json:"forceDecoding,omitempty"`

	// generate preview
	GeneratePreview bool `json:"generatePreview,omitempty"`

	// gop size
	GopSize int32 `json:"gopSize,omitempty"`

	// h264 enabled
	H264Enabled bool `json:"h264Enabled,omitempty"`

	// h265 enabled
	H265Enabled bool `json:"h265Enabled,omitempty"`

	// h265 encoder level
	H265EncoderLevel string `json:"h265EncoderLevel,omitempty"`

	// h265 encoder preset
	H265EncoderPreset string `json:"h265EncoderPreset,omitempty"`

	// h265 encoder profile
	H265EncoderProfile string `json:"h265EncoderProfile,omitempty"`

	// h265 encoder rc
	H265EncoderRc string `json:"h265EncoderRc,omitempty"`

	// h265 encoder specific
	H265EncoderSpecific string `json:"h265EncoderSpecific,omitempty"`

	// hash control play enabled
	HashControlPlayEnabled bool `json:"hashControlPlayEnabled,omitempty"`

	// hash control publish enabled
	HashControlPublishEnabled bool `json:"hashControlPublishEnabled,omitempty"`

	// height rtmp forwarding
	HeightRtmpForwarding int32 `json:"heightRtmpForwarding,omitempty"`

	// hls enabled via dash
	HlsEnabledViaDash bool `json:"hlsEnabledViaDash,omitempty"`

	// hls encryption key info file
	HlsEncryptionKeyInfoFile string `json:"hlsEncryptionKeyInfoFile,omitempty"`

	// hls flags
	HlsFlags string `json:"hlsFlags,omitempty"`

	// hls list size
	HlsListSize string `json:"hlsListSize,omitempty"`

	// hls muxing enabled
	HlsMuxingEnabled bool `json:"hlsMuxingEnabled,omitempty"`

	// hls play list type
	HlsPlayListType string `json:"hlsPlayListType,omitempty"`

	// hls time
	HlsTime string `json:"hlsTime,omitempty"`

	// http forwarding base URL
	HTTPForwardingBaseURL string `json:"httpForwardingBaseURL,omitempty"`

	// http forwarding extension
	HTTPForwardingExtension string `json:"httpForwardingExtension,omitempty"`

	// ingesting stream limit
	IngestingStreamLimit int32 `json:"ingestingStreamLimit,omitempty"`

	// ip filter enabled
	IPFilterEnabled bool `json:"ipFilterEnabled,omitempty"`

	// isl l dash enabled
	IslLDashEnabled bool `json:"islLDashEnabled,omitempty"`

	// isl l h l s enabled
	IslLHLSEnabled bool `json:"islLHLSEnabled,omitempty"`

	// jwks URL
	JwksURL string `json:"jwksURL,omitempty"`

	// jwt control enabled
	JwtControlEnabled bool `json:"jwtControlEnabled,omitempty"`

	// jwt secret key
	JwtSecretKey string `json:"jwtSecretKey,omitempty"`

	// jwt stream secret key
	JwtStreamSecretKey string `json:"jwtStreamSecretKey,omitempty"`

	// listener hook URL
	ListenerHookURL string `json:"listenerHookURL,omitempty"`

	// max analyze duration m s
	MaxAnalyzeDurationMS int32 `json:"maxAnalyzeDurationMS,omitempty"`

	// max resolution accept
	MaxResolutionAccept int32 `json:"maxResolutionAccept,omitempty"`

	// mp4 muxing enabled
	Mp4MuxingEnabled bool `json:"mp4MuxingEnabled,omitempty"`

	// muxer finish script
	MuxerFinishScript string `json:"muxerFinishScript,omitempty"`

	// my Sql client path
	MySQLClientPath string `json:"mySqlClientPath,omitempty"`

	// object detection enabled
	ObjectDetectionEnabled bool `json:"objectDetectionEnabled,omitempty"`

	// packet loss diff threshold for switchback
	PacketLossDiffThresholdForSwitchback int32 `json:"packetLossDiffThresholdForSwitchback,omitempty"`

	// play jwt control enabled
	PlayJwtControlEnabled bool `json:"playJwtControlEnabled,omitempty"`

	// play token control enabled
	PlayTokenControlEnabled bool `json:"playTokenControlEnabled,omitempty"`

	// port allocator flags
	PortAllocatorFlags int32 `json:"portAllocatorFlags,omitempty"`

	// preview height
	PreviewHeight int32 `json:"previewHeight,omitempty"`

	// preview overwrite
	PreviewOverwrite bool `json:"previewOverwrite,omitempty"`

	// publish jwt control enabled
	PublishJwtControlEnabled bool `json:"publishJwtControlEnabled,omitempty"`

	// publish token control enabled
	PublishTokenControlEnabled bool `json:"publishTokenControlEnabled,omitempty"`

	// pull war file
	PullWarFile bool `json:"pullWarFile,omitempty"`

	// remote allowed c ID r
	RemoteAllowedCIDR string `json:"remoteAllowedCIDR,omitempty"`

	// replace candidate addr with server addr
	ReplaceCandidateAddrWithServerAddr bool `json:"replaceCandidateAddrWithServerAddr,omitempty"`

	// restart stream fetcher period
	RestartStreamFetcherPeriod int32 `json:"restartStreamFetcherPeriod,omitempty"`

	// rtmp ingest buffer time ms
	RtmpIngestBufferTimeMs int64 `json:"rtmpIngestBufferTimeMs,omitempty"`

	// rtsp pull transport type
	RtspPullTransportType string `json:"rtspPullTransportType,omitempty"`

	// rtsp timeout duration ms
	RtspTimeoutDurationMs int32 `json:"rtspTimeoutDurationMs,omitempty"`

	// rtt measurement diff threshold for switchback
	RttMeasurementDiffThresholdForSwitchback int32 `json:"rttMeasurementDiffThresholdForSwitchback,omitempty"`

	// s3 access key
	S3AccessKey string `json:"s3AccessKey,omitempty"`

	// s3 bucket name
	S3BucketName string `json:"s3BucketName,omitempty"`

	// s3 endpoint
	S3Endpoint string `json:"s3Endpoint,omitempty"`

	// s3 permission
	S3Permission string `json:"s3Permission,omitempty"`

	// s3 previews folder path
	S3PreviewsFolderPath string `json:"s3PreviewsFolderPath,omitempty"`

	// s3 recording enabled
	S3RecordingEnabled bool `json:"s3RecordingEnabled,omitempty"`

	// s3 region name
	S3RegionName string `json:"s3RegionName,omitempty"`

	// s3 secret key
	S3SecretKey string `json:"s3SecretKey,omitempty"`

	// s3 storage class
	S3StorageClass string `json:"s3StorageClass,omitempty"`

	// s3 streams folder path
	S3StreamsFolderPath string `json:"s3StreamsFolderPath,omitempty"`

	// signaling address
	SignalingAddress string `json:"signalingAddress,omitempty"`

	// signaling enabled
	SignalingEnabled bool `json:"signalingEnabled,omitempty"`

	// stalker d b password
	StalkerDBPassword string `json:"stalkerDBPassword,omitempty"`

	// stalker d b server
	StalkerDBServer string `json:"stalkerDBServer,omitempty"`

	// stalker d b username
	StalkerDBUsername string `json:"stalkerDBUsername,omitempty"`

	// start stream fetcher automatically
	StartStreamFetcherAutomatically bool `json:"startStreamFetcherAutomatically,omitempty"`

	// stream fetcher buffer time
	StreamFetcherBufferTime int32 `json:"streamFetcherBufferTime,omitempty"`

	// stun server URI
	StunServerURI string `json:"stunServerURI,omitempty"`

	// target latency
	TargetLatency string `json:"targetLatency,omitempty"`

	// time token period
	TimeTokenPeriod int32 `json:"timeTokenPeriod,omitempty"`

	// time token subscriber only
	TimeTokenSubscriberOnly bool `json:"timeTokenSubscriberOnly,omitempty"`

	// to be deleted
	ToBeDeleted bool `json:"toBeDeleted,omitempty"`

	// token hash secret
	TokenHashSecret string `json:"tokenHashSecret,omitempty"`

	// update time
	UpdateTime int64 `json:"updateTime,omitempty"`

	// upload extensions to s3
	UploadExtensionsToS3 int32 `json:"uploadExtensionsToS3,omitempty"`

	// use original web r t c enabled
	UseOriginalWebRTCEnabled bool `json:"useOriginalWebRTCEnabled,omitempty"`

	// use timeline dash muxing
	UseTimelineDashMuxing bool `json:"useTimelineDashMuxing,omitempty"`

	// vod finish script
	VodFinishScript string `json:"vodFinishScript,omitempty"`

	// vod folder
	VodFolder string `json:"vodFolder,omitempty"`

	// vp8 enabled
	Vp8Enabled bool `json:"vp8Enabled,omitempty"`

	// vp8 encoder deadline
	Vp8EncoderDeadline string `json:"vp8EncoderDeadline,omitempty"`

	// vp8 encoder speed
	Vp8EncoderSpeed int32 `json:"vp8EncoderSpeed,omitempty"`

	// vp8 encoder thread count
	Vp8EncoderThreadCount int32 `json:"vp8EncoderThreadCount,omitempty"`

	// war file origin server address
	WarFileOriginServerAddress string `json:"warFileOriginServerAddress,omitempty"`

	// web m muxing enabled
	WebMMuxingEnabled bool `json:"webMMuxingEnabled,omitempty"`

	// web r t c client start timeout ms
	WebRTCClientStartTimeoutMs int32 `json:"webRTCClientStartTimeoutMs,omitempty"`

	// web r t c enabled
	WebRTCEnabled bool `json:"webRTCEnabled,omitempty"`

	// web r t c frame rate
	WebRTCFrameRate int32 `json:"webRTCFrameRate,omitempty"`

	// web r t c keyframe time
	WebRTCKeyframeTime int32 `json:"webRTCKeyframeTime,omitempty"`

	// web r t c port range max
	WebRTCPortRangeMax int32 `json:"webRTCPortRangeMax,omitempty"`

	// web r t c port range min
	WebRTCPortRangeMin int32 `json:"webRTCPortRangeMin,omitempty"`

	// web r t c sdp semantics
	WebRTCSdpSemantics string `json:"webRTCSdpSemantics,omitempty"`

	// web r t c Tcp candidates enabled
	WebRTCTCPCandidatesEnabled bool `json:"webRTCTcpCandidatesEnabled,omitempty"`

	// web r t c viewer limit
	WebRTCViewerLimit int32 `json:"webRTCViewerLimit,omitempty"`

	// webhook authenticate URL
	WebhookAuthenticateURL string `json:"webhookAuthenticateURL,omitempty"`

	// write stats to datastore
	WriteStatsToDatastore bool `json:"writeStatsToDatastore,omitempty"`
}

// Validate validates this app settings
func (m *AppSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncoderSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppSettings) validateEncoderSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.EncoderSettings) { // not required
		return nil
	}

	for i := 0; i < len(m.EncoderSettings); i++ {
		if swag.IsZero(m.EncoderSettings[i]) { // not required
			continue
		}

		if m.EncoderSettings[i] != nil {
			if err := m.EncoderSettings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("encoderSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("encoderSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this app settings based on the context it is used
func (m *AppSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEncoderSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppSettings) contextValidateEncoderSettings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EncoderSettings); i++ {

		if m.EncoderSettings[i] != nil {
			if err := m.EncoderSettings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("encoderSettings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("encoderSettings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppSettings) UnmarshalBinary(b []byte) error {
	var res AppSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
