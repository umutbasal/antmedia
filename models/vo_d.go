// Code generated by go-swagger;

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VoD The recorded video-on-demand object class
//
// swagger:model VoD
type VoD struct {

	// the creation of the VoD
	CreationDate int64 `json:"creationDate,omitempty"`

	// the duration of the VoD
	Duration int64 `json:"duration,omitempty"`

	// the path of the VoD
	FilePath string `json:"filePath,omitempty"`

	// the size of the VoD file in bytes
	FileSize int64 `json:"fileSize,omitempty"`

	// the type of the VoD, such as userVod, streamVod, uploadedVod
	PreviewFilePath string `json:"previewFilePath,omitempty"`

	// the time when the VoD is being started to get recorded in milliseconds(UTC- Unix epoch)
	StartTime int64 `json:"startTime,omitempty"`

	// the stream id of the VoD
	StreamID string `json:"streamId,omitempty"`

	// the object id of the VoD
	StreamName string `json:"streamName,omitempty"`

	// the type of the VoD, such as userVod, streamVod, uploadedVod
	Type string `json:"type,omitempty"`

	// the id of the VoD
	VodID string `json:"vodId,omitempty"`

	// the name of the VoD
	VodName string `json:"vodName,omitempty"`
}

// Validate validates this vo d
func (m *VoD) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vo d based on context it is used
func (m *VoD) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VoD) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoD) UnmarshalBinary(b []byte) error {
	var res VoD
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
