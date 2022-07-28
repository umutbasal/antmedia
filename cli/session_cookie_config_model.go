// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for SessionCookieConfig

// register flags to command
func registerModelSessionCookieConfigFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerSessionCookieConfigComment(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSessionCookieConfigDomain(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSessionCookieConfigHTTPOnly(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSessionCookieConfigMaxAge(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSessionCookieConfigName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSessionCookieConfigPath(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerSessionCookieConfigSecure(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerSessionCookieConfigComment(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	commentDescription := ``

	var commentFlagName string
	if cmdPrefix == "" {
		commentFlagName = "comment"
	} else {
		commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
	}

	var commentFlagDefault string

	_ = cmd.PersistentFlags().String(commentFlagName, commentFlagDefault, commentDescription)

	return nil
}

func registerSessionCookieConfigDomain(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	domainDescription := ``

	var domainFlagName string
	if cmdPrefix == "" {
		domainFlagName = "domain"
	} else {
		domainFlagName = fmt.Sprintf("%v.domain", cmdPrefix)
	}

	var domainFlagDefault string

	_ = cmd.PersistentFlags().String(domainFlagName, domainFlagDefault, domainDescription)

	return nil
}

func registerSessionCookieConfigHTTPOnly(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	httpOnlyDescription := ``

	var httpOnlyFlagName string
	if cmdPrefix == "" {
		httpOnlyFlagName = "httpOnly"
	} else {
		httpOnlyFlagName = fmt.Sprintf("%v.httpOnly", cmdPrefix)
	}

	var httpOnlyFlagDefault bool

	_ = cmd.PersistentFlags().Bool(httpOnlyFlagName, httpOnlyFlagDefault, httpOnlyDescription)

	return nil
}

func registerSessionCookieConfigMaxAge(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	maxAgeDescription := ``

	var maxAgeFlagName string
	if cmdPrefix == "" {
		maxAgeFlagName = "maxAge"
	} else {
		maxAgeFlagName = fmt.Sprintf("%v.maxAge", cmdPrefix)
	}

	var maxAgeFlagDefault int32

	_ = cmd.PersistentFlags().Int32(maxAgeFlagName, maxAgeFlagDefault, maxAgeDescription)

	return nil
}

func registerSessionCookieConfigName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := ``

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerSessionCookieConfigPath(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pathDescription := ``

	var pathFlagName string
	if cmdPrefix == "" {
		pathFlagName = "path"
	} else {
		pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
	}

	var pathFlagDefault string

	_ = cmd.PersistentFlags().String(pathFlagName, pathFlagDefault, pathDescription)

	return nil
}

func registerSessionCookieConfigSecure(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	secureDescription := ``

	var secureFlagName string
	if cmdPrefix == "" {
		secureFlagName = "secure"
	} else {
		secureFlagName = fmt.Sprintf("%v.secure", cmdPrefix)
	}

	var secureFlagDefault bool

	_ = cmd.PersistentFlags().Bool(secureFlagName, secureFlagDefault, secureDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelSessionCookieConfigFlags(depth int, m *models.SessionCookieConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, commentAdded := retrieveSessionCookieConfigCommentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || commentAdded

	err, domainAdded := retrieveSessionCookieConfigDomainFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || domainAdded

	err, httpOnlyAdded := retrieveSessionCookieConfigHTTPOnlyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || httpOnlyAdded

	err, maxAgeAdded := retrieveSessionCookieConfigMaxAgeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || maxAgeAdded

	err, nameAdded := retrieveSessionCookieConfigNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, pathAdded := retrieveSessionCookieConfigPathFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pathAdded

	err, secureAdded := retrieveSessionCookieConfigSecureFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secureAdded

	return nil, retAdded
}

func retrieveSessionCookieConfigCommentFlags(depth int, m *models.SessionCookieConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	commentFlagName := fmt.Sprintf("%v.comment", cmdPrefix)
	if cmd.Flags().Changed(commentFlagName) {

		var commentFlagName string
		if cmdPrefix == "" {
			commentFlagName = "comment"
		} else {
			commentFlagName = fmt.Sprintf("%v.comment", cmdPrefix)
		}

		commentFlagValue, err := cmd.Flags().GetString(commentFlagName)
		if err != nil {
			return err, false
		}
		m.Comment = commentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSessionCookieConfigDomainFlags(depth int, m *models.SessionCookieConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	domainFlagName := fmt.Sprintf("%v.domain", cmdPrefix)
	if cmd.Flags().Changed(domainFlagName) {

		var domainFlagName string
		if cmdPrefix == "" {
			domainFlagName = "domain"
		} else {
			domainFlagName = fmt.Sprintf("%v.domain", cmdPrefix)
		}

		domainFlagValue, err := cmd.Flags().GetString(domainFlagName)
		if err != nil {
			return err, false
		}
		m.Domain = domainFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSessionCookieConfigHTTPOnlyFlags(depth int, m *models.SessionCookieConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	httpOnlyFlagName := fmt.Sprintf("%v.httpOnly", cmdPrefix)
	if cmd.Flags().Changed(httpOnlyFlagName) {

		var httpOnlyFlagName string
		if cmdPrefix == "" {
			httpOnlyFlagName = "httpOnly"
		} else {
			httpOnlyFlagName = fmt.Sprintf("%v.httpOnly", cmdPrefix)
		}

		httpOnlyFlagValue, err := cmd.Flags().GetBool(httpOnlyFlagName)
		if err != nil {
			return err, false
		}
		m.HTTPOnly = httpOnlyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSessionCookieConfigMaxAgeFlags(depth int, m *models.SessionCookieConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	maxAgeFlagName := fmt.Sprintf("%v.maxAge", cmdPrefix)
	if cmd.Flags().Changed(maxAgeFlagName) {

		var maxAgeFlagName string
		if cmdPrefix == "" {
			maxAgeFlagName = "maxAge"
		} else {
			maxAgeFlagName = fmt.Sprintf("%v.maxAge", cmdPrefix)
		}

		maxAgeFlagValue, err := cmd.Flags().GetInt32(maxAgeFlagName)
		if err != nil {
			return err, false
		}
		m.MaxAge = maxAgeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSessionCookieConfigNameFlags(depth int, m *models.SessionCookieConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSessionCookieConfigPathFlags(depth int, m *models.SessionCookieConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pathFlagName := fmt.Sprintf("%v.path", cmdPrefix)
	if cmd.Flags().Changed(pathFlagName) {

		var pathFlagName string
		if cmdPrefix == "" {
			pathFlagName = "path"
		} else {
			pathFlagName = fmt.Sprintf("%v.path", cmdPrefix)
		}

		pathFlagValue, err := cmd.Flags().GetString(pathFlagName)
		if err != nil {
			return err, false
		}
		m.Path = pathFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveSessionCookieConfigSecureFlags(depth int, m *models.SessionCookieConfig, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	secureFlagName := fmt.Sprintf("%v.secure", cmdPrefix)
	if cmd.Flags().Changed(secureFlagName) {

		var secureFlagName string
		if cmdPrefix == "" {
			secureFlagName = "secure"
		} else {
			secureFlagName = fmt.Sprintf("%v.secure", cmdPrefix)
		}

		secureFlagValue, err := cmd.Flags().GetBool(secureFlagName)
		if err != nil {
			return err, false
		}
		m.Secure = secureFlagValue

		retAdded = true
	}

	return nil, retAdded
}
