// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for VoD

// register flags to command
func registerModelVoDFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerVoDCreationDate(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDDuration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDFilePath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDFileSize(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDPreviewFilePath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDStartTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDStreamID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDStreamName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDVodID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerVoDVodName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerVoDCreationDate(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	creationDateDescription := `the creation of the VoD`

	var creationDateFlagName string
	if cmdPrefix == "" {
		creationDateFlagName = "creationDate"
	} else {
		creationDateFlagName = fmt.Sprintf("%v.creationDate", cmdPrefix)
	}

	var creationDateFlagDefault int64

	_ = cmd.PersistentFlags().Int64(creationDateFlagName, creationDateFlagDefault, creationDateDescription)

	return nil
}

func registerVoDDuration(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	durationDescription := `the duration of the VoD`

	var durationFlagName string
	if cmdPrefix == "" {
		durationFlagName = "duration"
	} else {
		durationFlagName = fmt.Sprintf("%v.duration", cmdPrefix)
	}

	var durationFlagDefault int64

	_ = cmd.PersistentFlags().Int64(durationFlagName, durationFlagDefault, durationDescription)

	return nil
}

func registerVoDFilePath(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	filePathDescription := `the path of the VoD`

	var filePathFlagName string
	if cmdPrefix == "" {
		filePathFlagName = "filePath"
	} else {
		filePathFlagName = fmt.Sprintf("%v.filePath", cmdPrefix)
	}

	var filePathFlagDefault string

	_ = cmd.PersistentFlags().String(filePathFlagName, filePathFlagDefault, filePathDescription)

	return nil
}

func registerVoDFileSize(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fileSizeDescription := `the size of the VoD file in bytes`

	var fileSizeFlagName string
	if cmdPrefix == "" {
		fileSizeFlagName = "fileSize"
	} else {
		fileSizeFlagName = fmt.Sprintf("%v.fileSize", cmdPrefix)
	}

	var fileSizeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(fileSizeFlagName, fileSizeFlagDefault, fileSizeDescription)

	return nil
}

func registerVoDPreviewFilePath(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	previewFilePathDescription := `the type of the VoD, such as userVod, streamVod, uploadedVod`

	var previewFilePathFlagName string
	if cmdPrefix == "" {
		previewFilePathFlagName = "previewFilePath"
	} else {
		previewFilePathFlagName = fmt.Sprintf("%v.previewFilePath", cmdPrefix)
	}

	var previewFilePathFlagDefault string

	_ = cmd.PersistentFlags().String(previewFilePathFlagName, previewFilePathFlagDefault, previewFilePathDescription)

	return nil
}

func registerVoDStartTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startTimeDescription := `the time when the VoD is being started to get recorded in milliseconds(UTC- Unix epoch)`

	var startTimeFlagName string
	if cmdPrefix == "" {
		startTimeFlagName = "startTime"
	} else {
		startTimeFlagName = fmt.Sprintf("%v.startTime", cmdPrefix)
	}

	var startTimeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(startTimeFlagName, startTimeFlagDefault, startTimeDescription)

	return nil
}

func registerVoDStreamID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	streamIdDescription := `the stream id of the VoD`

	var streamIdFlagName string
	if cmdPrefix == "" {
		streamIdFlagName = "streamId"
	} else {
		streamIdFlagName = fmt.Sprintf("%v.streamId", cmdPrefix)
	}

	var streamIdFlagDefault string

	_ = cmd.PersistentFlags().String(streamIdFlagName, streamIdFlagDefault, streamIdDescription)

	return nil
}

func registerVoDStreamName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	streamNameDescription := `the object id of the VoD`

	var streamNameFlagName string
	if cmdPrefix == "" {
		streamNameFlagName = "streamName"
	} else {
		streamNameFlagName = fmt.Sprintf("%v.streamName", cmdPrefix)
	}

	var streamNameFlagDefault string

	_ = cmd.PersistentFlags().String(streamNameFlagName, streamNameFlagDefault, streamNameDescription)

	return nil
}

func registerVoDType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `the type of the VoD, such as userVod, streamVod, uploadedVod`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

func registerVoDVodID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vodIdDescription := `the id of the VoD`

	var vodIdFlagName string
	if cmdPrefix == "" {
		vodIdFlagName = "vodId"
	} else {
		vodIdFlagName = fmt.Sprintf("%v.vodId", cmdPrefix)
	}

	var vodIdFlagDefault string

	_ = cmd.PersistentFlags().String(vodIdFlagName, vodIdFlagDefault, vodIdDescription)

	return nil
}

func registerVoDVodName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vodNameDescription := `the name of the VoD`

	var vodNameFlagName string
	if cmdPrefix == "" {
		vodNameFlagName = "vodName"
	} else {
		vodNameFlagName = fmt.Sprintf("%v.vodName", cmdPrefix)
	}

	var vodNameFlagDefault string

	_ = cmd.PersistentFlags().String(vodNameFlagName, vodNameFlagDefault, vodNameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelVoDFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, creationDateAdded := retrieveVoDCreationDateFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || creationDateAdded

	err, durationAdded := retrieveVoDDurationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || durationAdded

	err, filePathAdded := retrieveVoDFilePathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || filePathAdded

	err, fileSizeAdded := retrieveVoDFileSizeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fileSizeAdded

	err, previewFilePathAdded := retrieveVoDPreviewFilePathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || previewFilePathAdded

	err, startTimeAdded := retrieveVoDStartTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startTimeAdded

	err, streamIdAdded := retrieveVoDStreamIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamIdAdded

	err, streamNameAdded := retrieveVoDStreamNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || streamNameAdded

	err, typeAdded := retrieveVoDTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, vodIdAdded := retrieveVoDVodIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vodIdAdded

	err, vodNameAdded := retrieveVoDVodNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vodNameAdded

	return nil, retAdded
}

func retrieveVoDCreationDateFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	creationDateFlagName := fmt.Sprintf("%v.creationDate", cmdPrefix)
	if cmd.Flags().Changed(creationDateFlagName) {

		var creationDateFlagName string
		if cmdPrefix == "" {
			creationDateFlagName = "creationDate"
		} else {
			creationDateFlagName = fmt.Sprintf("%v.creationDate", cmdPrefix)
		}

		creationDateFlagValue, err := cmd.Flags().GetInt64(creationDateFlagName)
		if err != nil {
			return err, false
		}
		m.CreationDate = creationDateFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDDurationFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	durationFlagName := fmt.Sprintf("%v.duration", cmdPrefix)
	if cmd.Flags().Changed(durationFlagName) {

		var durationFlagName string
		if cmdPrefix == "" {
			durationFlagName = "duration"
		} else {
			durationFlagName = fmt.Sprintf("%v.duration", cmdPrefix)
		}

		durationFlagValue, err := cmd.Flags().GetInt64(durationFlagName)
		if err != nil {
			return err, false
		}
		m.Duration = durationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDFilePathFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	filePathFlagName := fmt.Sprintf("%v.filePath", cmdPrefix)
	if cmd.Flags().Changed(filePathFlagName) {

		var filePathFlagName string
		if cmdPrefix == "" {
			filePathFlagName = "filePath"
		} else {
			filePathFlagName = fmt.Sprintf("%v.filePath", cmdPrefix)
		}

		filePathFlagValue, err := cmd.Flags().GetString(filePathFlagName)
		if err != nil {
			return err, false
		}
		m.FilePath = filePathFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDFileSizeFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fileSizeFlagName := fmt.Sprintf("%v.fileSize", cmdPrefix)
	if cmd.Flags().Changed(fileSizeFlagName) {

		var fileSizeFlagName string
		if cmdPrefix == "" {
			fileSizeFlagName = "fileSize"
		} else {
			fileSizeFlagName = fmt.Sprintf("%v.fileSize", cmdPrefix)
		}

		fileSizeFlagValue, err := cmd.Flags().GetInt64(fileSizeFlagName)
		if err != nil {
			return err, false
		}
		m.FileSize = fileSizeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDPreviewFilePathFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	previewFilePathFlagName := fmt.Sprintf("%v.previewFilePath", cmdPrefix)
	if cmd.Flags().Changed(previewFilePathFlagName) {

		var previewFilePathFlagName string
		if cmdPrefix == "" {
			previewFilePathFlagName = "previewFilePath"
		} else {
			previewFilePathFlagName = fmt.Sprintf("%v.previewFilePath", cmdPrefix)
		}

		previewFilePathFlagValue, err := cmd.Flags().GetString(previewFilePathFlagName)
		if err != nil {
			return err, false
		}
		m.PreviewFilePath = previewFilePathFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDStartTimeFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startTimeFlagName := fmt.Sprintf("%v.startTime", cmdPrefix)
	if cmd.Flags().Changed(startTimeFlagName) {

		var startTimeFlagName string
		if cmdPrefix == "" {
			startTimeFlagName = "startTime"
		} else {
			startTimeFlagName = fmt.Sprintf("%v.startTime", cmdPrefix)
		}

		startTimeFlagValue, err := cmd.Flags().GetInt64(startTimeFlagName)
		if err != nil {
			return err, false
		}
		m.StartTime = startTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDStreamIDFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamIdFlagName := fmt.Sprintf("%v.streamId", cmdPrefix)
	if cmd.Flags().Changed(streamIdFlagName) {

		var streamIdFlagName string
		if cmdPrefix == "" {
			streamIdFlagName = "streamId"
		} else {
			streamIdFlagName = fmt.Sprintf("%v.streamId", cmdPrefix)
		}

		streamIdFlagValue, err := cmd.Flags().GetString(streamIdFlagName)
		if err != nil {
			return err, false
		}
		m.StreamID = streamIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDStreamNameFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	streamNameFlagName := fmt.Sprintf("%v.streamName", cmdPrefix)
	if cmd.Flags().Changed(streamNameFlagName) {

		var streamNameFlagName string
		if cmdPrefix == "" {
			streamNameFlagName = "streamName"
		} else {
			streamNameFlagName = fmt.Sprintf("%v.streamName", cmdPrefix)
		}

		streamNameFlagValue, err := cmd.Flags().GetString(streamNameFlagName)
		if err != nil {
			return err, false
		}
		m.StreamName = streamNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDTypeFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDVodIDFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vodIdFlagName := fmt.Sprintf("%v.vodId", cmdPrefix)
	if cmd.Flags().Changed(vodIdFlagName) {

		var vodIdFlagName string
		if cmdPrefix == "" {
			vodIdFlagName = "vodId"
		} else {
			vodIdFlagName = fmt.Sprintf("%v.vodId", cmdPrefix)
		}

		vodIdFlagValue, err := cmd.Flags().GetString(vodIdFlagName)
		if err != nil {
			return err, false
		}
		m.VodID = vodIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveVoDVodNameFlags(depth int, m *models.VoD, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vodNameFlagName := fmt.Sprintf("%v.vodName", cmdPrefix)
	if cmd.Flags().Changed(vodNameFlagName) {

		var vodNameFlagName string
		if cmdPrefix == "" {
			vodNameFlagName = "vodName"
		} else {
			vodNameFlagName = fmt.Sprintf("%v.vodName", cmdPrefix)
		}

		vodNameFlagValue, err := cmd.Flags().GetString(vodNameFlagName)
		if err != nil {
			return err, false
		}
		m.VodName = vodNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
