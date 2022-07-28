// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for StorageClient

// register flags to command
func registerModelStorageClientFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerStorageClientAccessKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStorageClientEnabled(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStorageClientEndpoint(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStorageClientPermission(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStorageClientRegion(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStorageClientSecretKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStorageClientStorageClass(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerStorageClientStorageName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerStorageClientAccessKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	accessKeyDescription := ``

	var accessKeyFlagName string
	if cmdPrefix == "" {
		accessKeyFlagName = "accessKey"
	} else {
		accessKeyFlagName = fmt.Sprintf("%v.accessKey", cmdPrefix)
	}

	var accessKeyFlagDefault string

	_ = cmd.PersistentFlags().String(accessKeyFlagName, accessKeyFlagDefault, accessKeyDescription)

	return nil
}

func registerStorageClientEnabled(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	enabledDescription := ``

	var enabledFlagName string
	if cmdPrefix == "" {
		enabledFlagName = "enabled"
	} else {
		enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
	}

	var enabledFlagDefault bool

	_ = cmd.PersistentFlags().Bool(enabledFlagName, enabledFlagDefault, enabledDescription)

	return nil
}

func registerStorageClientEndpoint(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endpointDescription := ``

	var endpointFlagName string
	if cmdPrefix == "" {
		endpointFlagName = "endpoint"
	} else {
		endpointFlagName = fmt.Sprintf("%v.endpoint", cmdPrefix)
	}

	var endpointFlagDefault string

	_ = cmd.PersistentFlags().String(endpointFlagName, endpointFlagDefault, endpointDescription)

	return nil
}

func registerStorageClientPermission(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	permissionDescription := ``

	var permissionFlagName string
	if cmdPrefix == "" {
		permissionFlagName = "permission"
	} else {
		permissionFlagName = fmt.Sprintf("%v.permission", cmdPrefix)
	}

	var permissionFlagDefault string

	_ = cmd.PersistentFlags().String(permissionFlagName, permissionFlagDefault, permissionDescription)

	return nil
}

func registerStorageClientRegion(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	regionDescription := ``

	var regionFlagName string
	if cmdPrefix == "" {
		regionFlagName = "region"
	} else {
		regionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
	}

	var regionFlagDefault string

	_ = cmd.PersistentFlags().String(regionFlagName, regionFlagDefault, regionDescription)

	return nil
}

func registerStorageClientSecretKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	secretKeyDescription := ``

	var secretKeyFlagName string
	if cmdPrefix == "" {
		secretKeyFlagName = "secretKey"
	} else {
		secretKeyFlagName = fmt.Sprintf("%v.secretKey", cmdPrefix)
	}

	var secretKeyFlagDefault string

	_ = cmd.PersistentFlags().String(secretKeyFlagName, secretKeyFlagDefault, secretKeyDescription)

	return nil
}

func registerStorageClientStorageClass(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	storageClassDescription := ``

	var storageClassFlagName string
	if cmdPrefix == "" {
		storageClassFlagName = "storageClass"
	} else {
		storageClassFlagName = fmt.Sprintf("%v.storageClass", cmdPrefix)
	}

	var storageClassFlagDefault string

	_ = cmd.PersistentFlags().String(storageClassFlagName, storageClassFlagDefault, storageClassDescription)

	return nil
}

func registerStorageClientStorageName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	storageNameDescription := ``

	var storageNameFlagName string
	if cmdPrefix == "" {
		storageNameFlagName = "storageName"
	} else {
		storageNameFlagName = fmt.Sprintf("%v.storageName", cmdPrefix)
	}

	var storageNameFlagDefault string

	_ = cmd.PersistentFlags().String(storageNameFlagName, storageNameFlagDefault, storageNameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelStorageClientFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, accessKeyAdded := retrieveStorageClientAccessKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || accessKeyAdded

	err, enabledAdded := retrieveStorageClientEnabledFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || enabledAdded

	err, endpointAdded := retrieveStorageClientEndpointFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointAdded

	err, permissionAdded := retrieveStorageClientPermissionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || permissionAdded

	err, regionAdded := retrieveStorageClientRegionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || regionAdded

	err, secretKeyAdded := retrieveStorageClientSecretKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secretKeyAdded

	err, storageClassAdded := retrieveStorageClientStorageClassFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || storageClassAdded

	err, storageNameAdded := retrieveStorageClientStorageNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || storageNameAdded

	return nil, retAdded
}

func retrieveStorageClientAccessKeyFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	accessKeyFlagName := fmt.Sprintf("%v.accessKey", cmdPrefix)
	if cmd.Flags().Changed(accessKeyFlagName) {

		var accessKeyFlagName string
		if cmdPrefix == "" {
			accessKeyFlagName = "accessKey"
		} else {
			accessKeyFlagName = fmt.Sprintf("%v.accessKey", cmdPrefix)
		}

		accessKeyFlagValue, err := cmd.Flags().GetString(accessKeyFlagName)
		if err != nil {
			return err, false
		}
		m.AccessKey = accessKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStorageClientEnabledFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	enabledFlagName := fmt.Sprintf("%v.enabled", cmdPrefix)
	if cmd.Flags().Changed(enabledFlagName) {

		var enabledFlagName string
		if cmdPrefix == "" {
			enabledFlagName = "enabled"
		} else {
			enabledFlagName = fmt.Sprintf("%v.enabled", cmdPrefix)
		}

		enabledFlagValue, err := cmd.Flags().GetBool(enabledFlagName)
		if err != nil {
			return err, false
		}
		m.Enabled = enabledFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStorageClientEndpointFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointFlagName := fmt.Sprintf("%v.endpoint", cmdPrefix)
	if cmd.Flags().Changed(endpointFlagName) {

		var endpointFlagName string
		if cmdPrefix == "" {
			endpointFlagName = "endpoint"
		} else {
			endpointFlagName = fmt.Sprintf("%v.endpoint", cmdPrefix)
		}

		endpointFlagValue, err := cmd.Flags().GetString(endpointFlagName)
		if err != nil {
			return err, false
		}
		m.Endpoint = endpointFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStorageClientPermissionFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	permissionFlagName := fmt.Sprintf("%v.permission", cmdPrefix)
	if cmd.Flags().Changed(permissionFlagName) {

		var permissionFlagName string
		if cmdPrefix == "" {
			permissionFlagName = "permission"
		} else {
			permissionFlagName = fmt.Sprintf("%v.permission", cmdPrefix)
		}

		permissionFlagValue, err := cmd.Flags().GetString(permissionFlagName)
		if err != nil {
			return err, false
		}
		m.Permission = permissionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStorageClientRegionFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	regionFlagName := fmt.Sprintf("%v.region", cmdPrefix)
	if cmd.Flags().Changed(regionFlagName) {

		var regionFlagName string
		if cmdPrefix == "" {
			regionFlagName = "region"
		} else {
			regionFlagName = fmt.Sprintf("%v.region", cmdPrefix)
		}

		regionFlagValue, err := cmd.Flags().GetString(regionFlagName)
		if err != nil {
			return err, false
		}
		m.Region = regionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStorageClientSecretKeyFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	secretKeyFlagName := fmt.Sprintf("%v.secretKey", cmdPrefix)
	if cmd.Flags().Changed(secretKeyFlagName) {

		var secretKeyFlagName string
		if cmdPrefix == "" {
			secretKeyFlagName = "secretKey"
		} else {
			secretKeyFlagName = fmt.Sprintf("%v.secretKey", cmdPrefix)
		}

		secretKeyFlagValue, err := cmd.Flags().GetString(secretKeyFlagName)
		if err != nil {
			return err, false
		}
		m.SecretKey = secretKeyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStorageClientStorageClassFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	storageClassFlagName := fmt.Sprintf("%v.storageClass", cmdPrefix)
	if cmd.Flags().Changed(storageClassFlagName) {

		var storageClassFlagName string
		if cmdPrefix == "" {
			storageClassFlagName = "storageClass"
		} else {
			storageClassFlagName = fmt.Sprintf("%v.storageClass", cmdPrefix)
		}

		storageClassFlagValue, err := cmd.Flags().GetString(storageClassFlagName)
		if err != nil {
			return err, false
		}
		m.StorageClass = storageClassFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveStorageClientStorageNameFlags(depth int, m *models.StorageClient, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	storageNameFlagName := fmt.Sprintf("%v.storageName", cmdPrefix)
	if cmd.Flags().Changed(storageNameFlagName) {

		var storageNameFlagName string
		if cmdPrefix == "" {
			storageNameFlagName = "storageName"
		} else {
			storageNameFlagName = fmt.Sprintf("%v.storageName", cmdPrefix)
		}

		storageNameFlagValue, err := cmd.Flags().GetString(storageNameFlagName)
		if err != nil {
			return err, false
		}
		m.StorageName = storageNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
