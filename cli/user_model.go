// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"antmedia/models"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// Schema cli for User

// register flags to command
func registerModelUserFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserFirstName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserFullName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserLastName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserPicture(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserScope(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserUserType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	emailDescription := `the email of the user`

	var emailFlagName string
	if cmdPrefix == "" {
		emailFlagName = "email"
	} else {
		emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
	}

	var emailFlagDefault string

	_ = cmd.PersistentFlags().String(emailFlagName, emailFlagDefault, emailDescription)

	return nil
}

func registerUserFirstName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	firstNameDescription := `Fist name of the user`

	var firstNameFlagName string
	if cmdPrefix == "" {
		firstNameFlagName = "firstName"
	} else {
		firstNameFlagName = fmt.Sprintf("%v.firstName", cmdPrefix)
	}

	var firstNameFlagDefault string

	_ = cmd.PersistentFlags().String(firstNameFlagName, firstNameFlagDefault, firstNameDescription)

	return nil
}

func registerUserFullName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	fullNameDescription := `the name of the user`

	var fullNameFlagName string
	if cmdPrefix == "" {
		fullNameFlagName = "fullName"
	} else {
		fullNameFlagName = fmt.Sprintf("%v.fullName", cmdPrefix)
	}

	var fullNameFlagDefault string

	_ = cmd.PersistentFlags().String(fullNameFlagName, fullNameFlagDefault, fullNameDescription)

	return nil
}

func registerUserLastName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	lastNameDescription := `last name of the user`

	var lastNameFlagName string
	if cmdPrefix == "" {
		lastNameFlagName = "lastName"
	} else {
		lastNameFlagName = fmt.Sprintf("%v.lastName", cmdPrefix)
	}

	var lastNameFlagDefault string

	_ = cmd.PersistentFlags().String(lastNameFlagName, lastNameFlagDefault, lastNameDescription)

	return nil
}

func registerUserPicture(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pictureDescription := `the URL of the user picture`

	var pictureFlagName string
	if cmdPrefix == "" {
		pictureFlagName = "picture"
	} else {
		pictureFlagName = fmt.Sprintf("%v.picture", cmdPrefix)
	}

	var pictureFlagDefault string

	_ = cmd.PersistentFlags().String(pictureFlagName, pictureFlagDefault, pictureDescription)

	return nil
}

func registerUserScope(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	scopeDescription := `Scope can be 'system' or name of the application. Scope of the user. If it's scope is system, it can access the stuff in system-level. If it's scope is an application, it can access the stuff in application-levelIt makes more sense with UserType`

	var scopeFlagName string
	if cmdPrefix == "" {
		scopeFlagName = "scope"
	} else {
		scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
	}

	var scopeFlagDefault string

	_ = cmd.PersistentFlags().String(scopeFlagName, scopeFlagDefault, scopeDescription)

	return nil
}

func registerUserUserType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	userTypeDescription := `Enum: ["ADMIN","READ-ONLY","USER"]. the type of the user`

	var userTypeFlagName string
	if cmdPrefix == "" {
		userTypeFlagName = "userType"
	} else {
		userTypeFlagName = fmt.Sprintf("%v.userType", cmdPrefix)
	}

	var userTypeFlagDefault string

	_ = cmd.PersistentFlags().String(userTypeFlagName, userTypeFlagDefault, userTypeDescription)

	if err := cmd.RegisterFlagCompletionFunc(userTypeFlagName,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			var res []string
			if err := json.Unmarshal([]byte(`["ADMIN","READ-ONLY","USER"]`), &res); err != nil {
				panic(err)
			}
			return res, cobra.ShellCompDirectiveDefault
		}); err != nil {
		return err
	}

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserFlags(depth int, m *models.User, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, emailAdded := retrieveUserEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, firstNameAdded := retrieveUserFirstNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || firstNameAdded

	err, fullNameAdded := retrieveUserFullNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || fullNameAdded

	err, lastNameAdded := retrieveUserLastNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || lastNameAdded

	err, pictureAdded := retrieveUserPictureFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pictureAdded

	err, scopeAdded := retrieveUserScopeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeAdded

	err, userTypeAdded := retrieveUserUserTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userTypeAdded

	return nil, retAdded
}

func retrieveUserEmailFlags(depth int, m *models.User, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	emailFlagName := fmt.Sprintf("%v.email", cmdPrefix)
	if cmd.Flags().Changed(emailFlagName) {

		var emailFlagName string
		if cmdPrefix == "" {
			emailFlagName = "email"
		} else {
			emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
		}

		emailFlagValue, err := cmd.Flags().GetString(emailFlagName)
		if err != nil {
			return err, false
		}
		m.Email = emailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserFirstNameFlags(depth int, m *models.User, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	firstNameFlagName := fmt.Sprintf("%v.firstName", cmdPrefix)
	if cmd.Flags().Changed(firstNameFlagName) {

		var firstNameFlagName string
		if cmdPrefix == "" {
			firstNameFlagName = "firstName"
		} else {
			firstNameFlagName = fmt.Sprintf("%v.firstName", cmdPrefix)
		}

		firstNameFlagValue, err := cmd.Flags().GetString(firstNameFlagName)
		if err != nil {
			return err, false
		}
		m.FirstName = firstNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserFullNameFlags(depth int, m *models.User, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	fullNameFlagName := fmt.Sprintf("%v.fullName", cmdPrefix)
	if cmd.Flags().Changed(fullNameFlagName) {

		var fullNameFlagName string
		if cmdPrefix == "" {
			fullNameFlagName = "fullName"
		} else {
			fullNameFlagName = fmt.Sprintf("%v.fullName", cmdPrefix)
		}

		fullNameFlagValue, err := cmd.Flags().GetString(fullNameFlagName)
		if err != nil {
			return err, false
		}
		m.FullName = fullNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserLastNameFlags(depth int, m *models.User, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	lastNameFlagName := fmt.Sprintf("%v.lastName", cmdPrefix)
	if cmd.Flags().Changed(lastNameFlagName) {

		var lastNameFlagName string
		if cmdPrefix == "" {
			lastNameFlagName = "lastName"
		} else {
			lastNameFlagName = fmt.Sprintf("%v.lastName", cmdPrefix)
		}

		lastNameFlagValue, err := cmd.Flags().GetString(lastNameFlagName)
		if err != nil {
			return err, false
		}
		m.LastName = lastNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserPictureFlags(depth int, m *models.User, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pictureFlagName := fmt.Sprintf("%v.picture", cmdPrefix)
	if cmd.Flags().Changed(pictureFlagName) {

		var pictureFlagName string
		if cmdPrefix == "" {
			pictureFlagName = "picture"
		} else {
			pictureFlagName = fmt.Sprintf("%v.picture", cmdPrefix)
		}

		pictureFlagValue, err := cmd.Flags().GetString(pictureFlagName)
		if err != nil {
			return err, false
		}
		m.Picture = pictureFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserScopeFlags(depth int, m *models.User, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scopeFlagName := fmt.Sprintf("%v.scope", cmdPrefix)
	if cmd.Flags().Changed(scopeFlagName) {

		var scopeFlagName string
		if cmdPrefix == "" {
			scopeFlagName = "scope"
		} else {
			scopeFlagName = fmt.Sprintf("%v.scope", cmdPrefix)
		}

		scopeFlagValue, err := cmd.Flags().GetString(scopeFlagName)
		if err != nil {
			return err, false
		}
		m.Scope = scopeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserUserTypeFlags(depth int, m *models.User, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userTypeFlagName := fmt.Sprintf("%v.userType", cmdPrefix)
	if cmd.Flags().Changed(userTypeFlagName) {

		var userTypeFlagName string
		if cmdPrefix == "" {
			userTypeFlagName = "userType"
		} else {
			userTypeFlagName = fmt.Sprintf("%v.userType", cmdPrefix)
		}

		userTypeFlagValue, err := cmd.Flags().GetString(userTypeFlagName)
		if err != nil {
			return err, false
		}
		m.UserType = userTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
