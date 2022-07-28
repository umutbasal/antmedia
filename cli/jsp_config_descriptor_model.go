// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for JspConfigDescriptor

// register flags to command
func registerModelJspConfigDescriptorFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerJspConfigDescriptorJspPropertyGroups(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerJspConfigDescriptorTaglibs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerJspConfigDescriptorJspPropertyGroups(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: jspPropertyGroups []*JspPropertyGroupDescriptor array type is not supported by go-swagger cli yet

	return nil
}

func registerJspConfigDescriptorTaglibs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: taglibs []*TaglibDescriptor array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelJspConfigDescriptorFlags(depth int, m *models.JspConfigDescriptor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, jspPropertyGroupsAdded := retrieveJspConfigDescriptorJspPropertyGroupsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || jspPropertyGroupsAdded

	err, taglibsAdded := retrieveJspConfigDescriptorTaglibsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || taglibsAdded

	return nil, retAdded
}

func retrieveJspConfigDescriptorJspPropertyGroupsFlags(depth int, m *models.JspConfigDescriptor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	jspPropertyGroupsFlagName := fmt.Sprintf("%v.jspPropertyGroups", cmdPrefix)
	if cmd.Flags().Changed(jspPropertyGroupsFlagName) {
		// warning: jspPropertyGroups array type []*JspPropertyGroupDescriptor is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveJspConfigDescriptorTaglibsFlags(depth int, m *models.JspConfigDescriptor, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	taglibsFlagName := fmt.Sprintf("%v.taglibs", cmdPrefix)
	if cmd.Flags().Changed(taglibsFlagName) {
		// warning: taglibs array type []*TaglibDescriptor is not supported by go-swagger cli yet
	}

	return nil, retAdded
}