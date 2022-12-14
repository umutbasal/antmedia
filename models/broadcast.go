// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Broadcast The basic broadcast class
//
// swagger:model Broadcast
type Broadcast struct {

	// Absolute start time in milliseconds - unix timestamp. It's used for measuring the absolute latency
	AbsoluteStartTimeMs int64 `json:"absoluteStartTimeMs,omitempty"`

	// altitude of the broadcasting location
	Altitude string `json:"altitude,omitempty"`

	// the received bytes / duration
	Bitrate int64 `json:"bitrate,omitempty"`

	// the category of the stream
	Category string `json:"category,omitempty"`

	// Current playing index for playlist types
	CurrentPlayIndex int32 `json:"currentPlayIndex,omitempty"`

	// the date when record is created in milliseconds
	Date int64 `json:"date,omitempty"`

	// the description of the stream
	Description string `json:"description,omitempty"`

	// the duration of the stream in milliseconds
	Duration int64 `json:"duration,omitempty"`

	// the list of endpoints such as Facebook, Twitter or custom RTMP endpoints
	EndPointList []*Endpoint `json:"endPointList"`

	// the expire time in milliseconds For instance if this value is 10000 then broadcast should be started in 10 seconds after it is created.If expire duration is 0, then stream will never expire
	ExpireDurationMS int32 `json:"expireDurationMS,omitempty"`

	// the number of HLS viewers of the stream
	HlsViewerCount int32 `json:"hlsViewerCount,omitempty"`

	// Number of the allowed maximum HLS viewers for the broadcast
	HlsViewerLimit int32 `json:"hlsViewerLimit,omitempty"`

	// the IP Address of the IP Camera or publisher
	IPAddr string `json:"ipAddr,omitempty"`

	// the identifier of whether stream is 360 or not
	Is360 bool `json:"is360,omitempty"`

	// latitude of the broadcasting location
	Latitude string `json:"latitude,omitempty"`

	// the url that will be notified when stream is published, ended and muxing finished
	ListenerHookURL string `json:"listenerHookURL,omitempty"`

	// longitude of the broadcasting location
	Longitude string `json:"longitude,omitempty"`

	// If this broadcast is a track of a WebRTC stream. This variable is Id of that stream.
	MainTrackStreamID string `json:"mainTrackStreamId,omitempty"`

	// Meta data filed for the custom usage
	MetaData string `json:"metaData,omitempty"`

	// MP4 muxing whether enabled or not for the stream, 1 means enabled, -1 means disabled, 0 means no settings for the stream
	Mp4Enabled int32 `json:"mp4Enabled,omitempty"`

	// the name of the stream
	Name string `json:"name,omitempty"`

	// the origin address server broadcasting
	OriginAdress string `json:"originAdress,omitempty"`

	// the password of the IP Camera
	Password string `json:"password,omitempty"`

	// the number of audio and video packets that is being pending to be encoded in the queue
	PendingPacketSize int32 `json:"pendingPacketSize,omitempty"`

	// the planned end date
	PlannedEndDate int64 `json:"plannedEndDate,omitempty"`

	// the planned start date
	PlannedStartDate int64 `json:"plannedStartDate,omitempty"`

	// the list broadcasts of Playlist Items. This list has values when the broadcast type is playlist
	PlayListItemList []*PlayListItem `json:"playListItemList"`

	// The status of the playlist. It's usable if type is playlist
	// Enum: [finished broadcasting created]
	PlayListStatus string `json:"playListStatus,omitempty"`

	// the identifier of playlist loop status
	PlaylistLoopEnabled bool `json:"playlistLoopEnabled,omitempty"`

	// the identifier of whether stream is public or not
	PublicStream bool `json:"publicStream,omitempty"`

	// it is a video filter for the service, this value is controlled by the user, default value is true in the db
	Publish bool `json:"publish,omitempty"`

	// the publish type of the stream
	// Enum: [WebRTC RTMP Pull]
	PublishType string `json:"publishType,omitempty"`

	// the quality of the incoming stream during publishing
	Quality string `json:"quality,omitempty"`

	// the received bytes until now
	ReceivedBytes int64 `json:"receivedBytes,omitempty"`

	// the RTMP URL where to publish live stream to
	RtmpURL string `json:"rtmpURL,omitempty"`

	// the number of RTMP viewers of the stream
	RtmpViewerCount int32 `json:"rtmpViewerCount,omitempty"`

	// the speed of the incoming stream, for better quality and performance it should be around 1.00
	Speed float64 `json:"speed,omitempty"`

	// the publishing start time of the stream
	StartTime int64 `json:"startTime,omitempty"`

	// the status of the stream
	// Enum: [finished broadcasting created]
	Status string `json:"status,omitempty"`

	// the id of the stream
	StreamID string `json:"streamId,omitempty"`

	// the stream URL for fetching stream, especially should be defined for IP Cameras or Cloud streams
	StreamURL string `json:"streamUrl,omitempty"`

	// Name of the subfolder that will contain stream files
	SubFolder string `json:"subFolder,omitempty"`

	// If this broadcast is main track. This variable hold sub track ids.
	SubTrackStreamIds []string `json:"subTrackStreamIds"`

	// the type of the stream
	// Enum: [liveStream ipCamera streamSource VoD playlist]
	Type string `json:"type,omitempty"`

	// User - Agent
	UserAgent string `json:"userAgent,omitempty"`

	// the user name of the IP Camera
	Username string `json:"username,omitempty"`

	// WebM muxing whether enabled or not for the stream, 1 means enabled, -1 means disabled, 0 means no settings for the stream
	WebMEnabled int32 `json:"webMEnabled,omitempty"`

	// the number of WebRTC viewers of the stream
	WebRTCViewerCount int32 `json:"webRTCViewerCount,omitempty"`

	// Number of the allowed maximum WebRTC viewers for the broadcast
	WebRTCViewerLimit int32 `json:"webRTCViewerLimit,omitempty"`

	// is true, if a broadcast that is not added to data store through rest service or management console It is false by default
	Zombi bool `json:"zombi,omitempty"`
}

// Validate validates this broadcast
func (m *Broadcast) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndPointList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayListItemList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlayListStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Broadcast) validateEndPointList(formats strfmt.Registry) error {
	if swag.IsZero(m.EndPointList) { // not required
		return nil
	}

	for i := 0; i < len(m.EndPointList); i++ {
		if swag.IsZero(m.EndPointList[i]) { // not required
			continue
		}

		if m.EndPointList[i] != nil {
			if err := m.EndPointList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endPointList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endPointList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Broadcast) validatePlayListItemList(formats strfmt.Registry) error {
	if swag.IsZero(m.PlayListItemList) { // not required
		return nil
	}

	for i := 0; i < len(m.PlayListItemList); i++ {
		if swag.IsZero(m.PlayListItemList[i]) { // not required
			continue
		}

		if m.PlayListItemList[i] != nil {
			if err := m.PlayListItemList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("playListItemList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("playListItemList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var broadcastTypePlayListStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["finished","broadcasting","created"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		broadcastTypePlayListStatusPropEnum = append(broadcastTypePlayListStatusPropEnum, v)
	}
}

const (

	// BroadcastPlayListStatusFinished captures enum value "finished"
	BroadcastPlayListStatusFinished string = "finished"

	// BroadcastPlayListStatusBroadcasting captures enum value "broadcasting"
	BroadcastPlayListStatusBroadcasting string = "broadcasting"

	// BroadcastPlayListStatusCreated captures enum value "created"
	BroadcastPlayListStatusCreated string = "created"
)

// prop value enum
func (m *Broadcast) validatePlayListStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, broadcastTypePlayListStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Broadcast) validatePlayListStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PlayListStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlayListStatusEnum("playListStatus", "body", m.PlayListStatus); err != nil {
		return err
	}

	return nil
}

var broadcastTypePublishTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["WebRTC","RTMP","Pull"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		broadcastTypePublishTypePropEnum = append(broadcastTypePublishTypePropEnum, v)
	}
}

const (

	// BroadcastPublishTypeWebRTC captures enum value "WebRTC"
	BroadcastPublishTypeWebRTC string = "WebRTC"

	// BroadcastPublishTypeRTMP captures enum value "RTMP"
	BroadcastPublishTypeRTMP string = "RTMP"

	// BroadcastPublishTypePull captures enum value "Pull"
	BroadcastPublishTypePull string = "Pull"
)

// prop value enum
func (m *Broadcast) validatePublishTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, broadcastTypePublishTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Broadcast) validatePublishType(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePublishTypeEnum("publishType", "body", m.PublishType); err != nil {
		return err
	}

	return nil
}

var broadcastTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["finished","broadcasting","created"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		broadcastTypeStatusPropEnum = append(broadcastTypeStatusPropEnum, v)
	}
}

const (

	// BroadcastStatusFinished captures enum value "finished"
	BroadcastStatusFinished string = "finished"

	// BroadcastStatusBroadcasting captures enum value "broadcasting"
	BroadcastStatusBroadcasting string = "broadcasting"

	// BroadcastStatusCreated captures enum value "created"
	BroadcastStatusCreated string = "created"
)

// prop value enum
func (m *Broadcast) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, broadcastTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Broadcast) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var broadcastTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["liveStream","ipCamera","streamSource","VoD","playlist"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		broadcastTypeTypePropEnum = append(broadcastTypeTypePropEnum, v)
	}
}

const (

	// BroadcastTypeLiveStream captures enum value "liveStream"
	BroadcastTypeLiveStream string = "liveStream"

	// BroadcastTypeIPCamera captures enum value "ipCamera"
	BroadcastTypeIPCamera string = "ipCamera"

	// BroadcastTypeStreamSource captures enum value "streamSource"
	BroadcastTypeStreamSource string = "streamSource"

	// BroadcastTypeVoD captures enum value "VoD"
	BroadcastTypeVoD string = "VoD"

	// BroadcastTypePlaylist captures enum value "playlist"
	BroadcastTypePlaylist string = "playlist"
)

// prop value enum
func (m *Broadcast) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, broadcastTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Broadcast) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this broadcast based on the context it is used
func (m *Broadcast) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndPointList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlayListItemList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Broadcast) contextValidateEndPointList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EndPointList); i++ {

		if m.EndPointList[i] != nil {
			if err := m.EndPointList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("endPointList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("endPointList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Broadcast) contextValidatePlayListItemList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PlayListItemList); i++ {

		if m.PlayListItemList[i] != nil {
			if err := m.PlayListItemList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("playListItemList" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("playListItemList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Broadcast) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Broadcast) UnmarshalBinary(b []byte) error {
	var res Broadcast
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
