// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for Endpoint

// register flags to command
func registerModelEndpointFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerEndpointEndpointServiceID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointRtmpURL(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerEndpointType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerEndpointEndpointServiceID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endpointServiceIdDescription := `the endpoint service id, this field holds the id of the endpoint`

	var endpointServiceIdFlagName string
	if cmdPrefix == "" {
		endpointServiceIdFlagName = "endpointServiceId"
	} else {
		endpointServiceIdFlagName = fmt.Sprintf("%v.endpointServiceId", cmdPrefix)
	}

	var endpointServiceIdFlagDefault string

	_ = cmd.PersistentFlags().String(endpointServiceIdFlagName, endpointServiceIdFlagDefault, endpointServiceIdDescription)

	return nil
}

func registerEndpointRtmpURL(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	rtmpUrlDescription := `the RTMP URL of the endpoint`

	var rtmpUrlFlagName string
	if cmdPrefix == "" {
		rtmpUrlFlagName = "rtmpUrl"
	} else {
		rtmpUrlFlagName = fmt.Sprintf("%v.rtmpUrl", cmdPrefix)
	}

	var rtmpUrlFlagDefault string

	_ = cmd.PersistentFlags().String(rtmpUrlFlagName, rtmpUrlFlagDefault, rtmpUrlDescription)

	return nil
}

func registerEndpointStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := `Status of the RTMP muxer, possible values are started, finished, failed, broadcasting, {@link IAntMediaStreamHandler#BROADCAST_STATUS_*}`

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

func registerEndpointType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `the service name like facebook, periscope, youtube or generic`

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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelEndpointFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endpointServiceIdAdded := retrieveEndpointEndpointServiceIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointServiceIdAdded

	err, rtmpUrlAdded := retrieveEndpointRtmpURLFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || rtmpUrlAdded

	err, statusAdded := retrieveEndpointStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, typeAdded := retrieveEndpointTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	return nil, retAdded
}

func retrieveEndpointEndpointServiceIDFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointServiceIdFlagName := fmt.Sprintf("%v.endpointServiceId", cmdPrefix)
	if cmd.Flags().Changed(endpointServiceIdFlagName) {

		var endpointServiceIdFlagName string
		if cmdPrefix == "" {
			endpointServiceIdFlagName = "endpointServiceId"
		} else {
			endpointServiceIdFlagName = fmt.Sprintf("%v.endpointServiceId", cmdPrefix)
		}

		endpointServiceIdFlagValue, err := cmd.Flags().GetString(endpointServiceIdFlagName)
		if err != nil {
			return err, false
		}
		m.EndpointServiceID = endpointServiceIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointRtmpURLFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	rtmpUrlFlagName := fmt.Sprintf("%v.rtmpUrl", cmdPrefix)
	if cmd.Flags().Changed(rtmpUrlFlagName) {

		var rtmpUrlFlagName string
		if cmdPrefix == "" {
			rtmpUrlFlagName = "rtmpUrl"
		} else {
			rtmpUrlFlagName = fmt.Sprintf("%v.rtmpUrl", cmdPrefix)
		}

		rtmpUrlFlagValue, err := cmd.Flags().GetString(rtmpUrlFlagName)
		if err != nil {
			return err, false
		}
		m.RtmpURL = rtmpUrlFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveEndpointStatusFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveEndpointTypeFlags(depth int, m *models.Endpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
