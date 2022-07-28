// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for TensorFlowObject

// register flags to command
func registerModelTensorFlowObjectFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerTensorFlowObjectDetectionTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTensorFlowObjectImageID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTensorFlowObjectMaxX(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTensorFlowObjectMaxY(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTensorFlowObjectMinX(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTensorFlowObjectMinY(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTensorFlowObjectObjectName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerTensorFlowObjectProbability(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerTensorFlowObjectDetectionTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	detectionTimeDescription := `the time of the detected object`

	var detectionTimeFlagName string
	if cmdPrefix == "" {
		detectionTimeFlagName = "detectionTime"
	} else {
		detectionTimeFlagName = fmt.Sprintf("%v.detectionTime", cmdPrefix)
	}

	var detectionTimeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(detectionTimeFlagName, detectionTimeFlagDefault, detectionTimeDescription)

	return nil
}

func registerTensorFlowObjectImageID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	imageIdDescription := `the id of the detected image`

	var imageIdFlagName string
	if cmdPrefix == "" {
		imageIdFlagName = "imageId"
	} else {
		imageIdFlagName = fmt.Sprintf("%v.imageId", cmdPrefix)
	}

	var imageIdFlagDefault string

	_ = cmd.PersistentFlags().String(imageIdFlagName, imageIdFlagDefault, imageIdDescription)

	return nil
}

func registerTensorFlowObjectMaxX(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxXDescription := `the x coordinate of lower-right corner of detected object frame`

	var maxXFlagName string
	if cmdPrefix == "" {
		maxXFlagName = "maxX"
	} else {
		maxXFlagName = fmt.Sprintf("%v.maxX", cmdPrefix)
	}

	var maxXFlagDefault float64

	_ = cmd.PersistentFlags().Float64(maxXFlagName, maxXFlagDefault, maxXDescription)

	return nil
}

func registerTensorFlowObjectMaxY(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxYDescription := `the y coordinate of lower-right corner of detected object frame`

	var maxYFlagName string
	if cmdPrefix == "" {
		maxYFlagName = "maxY"
	} else {
		maxYFlagName = fmt.Sprintf("%v.maxY", cmdPrefix)
	}

	var maxYFlagDefault float64

	_ = cmd.PersistentFlags().Float64(maxYFlagName, maxYFlagDefault, maxYDescription)

	return nil
}

func registerTensorFlowObjectMinX(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	minXDescription := `the x coordinate of upper-left corner of detected object frame`

	var minXFlagName string
	if cmdPrefix == "" {
		minXFlagName = "minX"
	} else {
		minXFlagName = fmt.Sprintf("%v.minX", cmdPrefix)
	}

	var minXFlagDefault float64

	_ = cmd.PersistentFlags().Float64(minXFlagName, minXFlagDefault, minXDescription)

	return nil
}

func registerTensorFlowObjectMinY(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	minYDescription := `the y coordinate of upper-left corner of detected object frame`

	var minYFlagName string
	if cmdPrefix == "" {
		minYFlagName = "minY"
	} else {
		minYFlagName = fmt.Sprintf("%v.minY", cmdPrefix)
	}

	var minYFlagDefault float64

	_ = cmd.PersistentFlags().Float64(minYFlagName, minYFlagDefault, minYDescription)

	return nil
}

func registerTensorFlowObjectObjectName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	objectNameDescription := `the name of the detected object`

	var objectNameFlagName string
	if cmdPrefix == "" {
		objectNameFlagName = "objectName"
	} else {
		objectNameFlagName = fmt.Sprintf("%v.objectName", cmdPrefix)
	}

	var objectNameFlagDefault string

	_ = cmd.PersistentFlags().String(objectNameFlagName, objectNameFlagDefault, objectNameDescription)

	return nil
}

func registerTensorFlowObjectProbability(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	probabilityDescription := `the probablity of the detected object`

	var probabilityFlagName string
	if cmdPrefix == "" {
		probabilityFlagName = "probability"
	} else {
		probabilityFlagName = fmt.Sprintf("%v.probability", cmdPrefix)
	}

	var probabilityFlagDefault float32

	_ = cmd.PersistentFlags().Float32(probabilityFlagName, probabilityFlagDefault, probabilityDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelTensorFlowObjectFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, detectionTimeAdded := retrieveTensorFlowObjectDetectionTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || detectionTimeAdded

	err, imageIdAdded := retrieveTensorFlowObjectImageIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || imageIdAdded

	err, maxXAdded := retrieveTensorFlowObjectMaxXFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxXAdded

	err, maxYAdded := retrieveTensorFlowObjectMaxYFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxYAdded

	err, minXAdded := retrieveTensorFlowObjectMinXFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || minXAdded

	err, minYAdded := retrieveTensorFlowObjectMinYFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || minYAdded

	err, objectNameAdded := retrieveTensorFlowObjectObjectNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || objectNameAdded

	err, probabilityAdded := retrieveTensorFlowObjectProbabilityFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || probabilityAdded

	return nil, retAdded
}

func retrieveTensorFlowObjectDetectionTimeFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	detectionTimeFlagName := fmt.Sprintf("%v.detectionTime", cmdPrefix)
	if cmd.Flags().Changed(detectionTimeFlagName) {

		var detectionTimeFlagName string
		if cmdPrefix == "" {
			detectionTimeFlagName = "detectionTime"
		} else {
			detectionTimeFlagName = fmt.Sprintf("%v.detectionTime", cmdPrefix)
		}

		detectionTimeFlagValue, err := cmd.Flags().GetInt64(detectionTimeFlagName)
		if err != nil {
			return err, false
		}
		m.DetectionTime = detectionTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTensorFlowObjectImageIDFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	imageIdFlagName := fmt.Sprintf("%v.imageId", cmdPrefix)
	if cmd.Flags().Changed(imageIdFlagName) {

		var imageIdFlagName string
		if cmdPrefix == "" {
			imageIdFlagName = "imageId"
		} else {
			imageIdFlagName = fmt.Sprintf("%v.imageId", cmdPrefix)
		}

		imageIdFlagValue, err := cmd.Flags().GetString(imageIdFlagName)
		if err != nil {
			return err, false
		}
		m.ImageID = imageIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTensorFlowObjectMaxXFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxXFlagName := fmt.Sprintf("%v.maxX", cmdPrefix)
	if cmd.Flags().Changed(maxXFlagName) {

		var maxXFlagName string
		if cmdPrefix == "" {
			maxXFlagName = "maxX"
		} else {
			maxXFlagName = fmt.Sprintf("%v.maxX", cmdPrefix)
		}

		maxXFlagValue, err := cmd.Flags().GetFloat64(maxXFlagName)
		if err != nil {
			return err, false
		}
		m.MaxX = maxXFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTensorFlowObjectMaxYFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxYFlagName := fmt.Sprintf("%v.maxY", cmdPrefix)
	if cmd.Flags().Changed(maxYFlagName) {

		var maxYFlagName string
		if cmdPrefix == "" {
			maxYFlagName = "maxY"
		} else {
			maxYFlagName = fmt.Sprintf("%v.maxY", cmdPrefix)
		}

		maxYFlagValue, err := cmd.Flags().GetFloat64(maxYFlagName)
		if err != nil {
			return err, false
		}
		m.MaxY = maxYFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTensorFlowObjectMinXFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	minXFlagName := fmt.Sprintf("%v.minX", cmdPrefix)
	if cmd.Flags().Changed(minXFlagName) {

		var minXFlagName string
		if cmdPrefix == "" {
			minXFlagName = "minX"
		} else {
			minXFlagName = fmt.Sprintf("%v.minX", cmdPrefix)
		}

		minXFlagValue, err := cmd.Flags().GetFloat64(minXFlagName)
		if err != nil {
			return err, false
		}
		m.MinX = minXFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTensorFlowObjectMinYFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	minYFlagName := fmt.Sprintf("%v.minY", cmdPrefix)
	if cmd.Flags().Changed(minYFlagName) {

		var minYFlagName string
		if cmdPrefix == "" {
			minYFlagName = "minY"
		} else {
			minYFlagName = fmt.Sprintf("%v.minY", cmdPrefix)
		}

		minYFlagValue, err := cmd.Flags().GetFloat64(minYFlagName)
		if err != nil {
			return err, false
		}
		m.MinY = minYFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTensorFlowObjectObjectNameFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	objectNameFlagName := fmt.Sprintf("%v.objectName", cmdPrefix)
	if cmd.Flags().Changed(objectNameFlagName) {

		var objectNameFlagName string
		if cmdPrefix == "" {
			objectNameFlagName = "objectName"
		} else {
			objectNameFlagName = fmt.Sprintf("%v.objectName", cmdPrefix)
		}

		objectNameFlagValue, err := cmd.Flags().GetString(objectNameFlagName)
		if err != nil {
			return err, false
		}
		m.ObjectName = objectNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveTensorFlowObjectProbabilityFlags(depth int, m *models.TensorFlowObject, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	probabilityFlagName := fmt.Sprintf("%v.probability", cmdPrefix)
	if cmd.Flags().Changed(probabilityFlagName) {

		var probabilityFlagName string
		if cmdPrefix == "" {
			probabilityFlagName = "probability"
		} else {
			probabilityFlagName = fmt.Sprintf("%v.probability", cmdPrefix)
		}

		probabilityFlagValue, err := cmd.Flags().GetFloat32(probabilityFlagName)
		if err != nil {
			return err, false
		}
		m.Probability = probabilityFlagValue

		retAdded = true
	}

	return nil, retAdded
}
