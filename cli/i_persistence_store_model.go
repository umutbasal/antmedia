// Code generated by go-swagger;

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/umutbasal/antmedia/models"

	"github.com/spf13/cobra"
)

// Schema cli for IPersistenceStore

// register flags to command
func registerModelIPersistenceStoreFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIPersistenceStoreObjectNames(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIPersistenceStoreObjects(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIPersistenceStoreObjectNames(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: objectNames []string array type is not supported by go-swagger cli yet

	return nil
}

func registerIPersistenceStoreObjects(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: objects []*IPersistable array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIPersistenceStoreFlags(depth int, m *models.IPersistenceStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, objectNamesAdded := retrieveIPersistenceStoreObjectNamesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || objectNamesAdded

	err, objectsAdded := retrieveIPersistenceStoreObjectsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || objectsAdded

	return nil, retAdded
}

func retrieveIPersistenceStoreObjectNamesFlags(depth int, m *models.IPersistenceStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	objectNamesFlagName := fmt.Sprintf("%v.objectNames", cmdPrefix)
	if cmd.Flags().Changed(objectNamesFlagName) {
		// warning: objectNames array type []string is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveIPersistenceStoreObjectsFlags(depth int, m *models.IPersistenceStore, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	objectsFlagName := fmt.Sprintf("%v.objects", cmdPrefix)
	if cmd.Flags().Changed(objectsFlagName) {
		// warning: objects array type []*IPersistable is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
