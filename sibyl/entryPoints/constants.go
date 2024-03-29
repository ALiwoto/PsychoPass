/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package entryPoints

const (
	// ErrPermissionDenied is the error message that should be sent
	// when the user does not have enough permissions to perform the action.
	ErrPermissionDenied = "Permission denied"

	// ErrInvalidUserId is the error message that should be sent when
	// the target's user id is invalid. (contains invalid characters)
	ErrInvalidUserId = "Invalid user-id provided"

	// ErrInvalidUniqueId is the error message that should be sent when
	// the provided unique-id is invalid.
	ErrInvalidUniqueId = "Invalid unique-id provided"

	// ErrInvalidPerm is the error message that should be sent when
	// the target's permission param is invalid.
	ErrInvalidPerm = "Invalid permission provided"

	// ErrCannotChangePerm is the error message that should be sent when
	// there is a problem in changing someone's permission; for example
	// the target user is an owner. as owner users can't change another owners'
	// permission.
	ErrCannotChangePerm = "Can't change target's permission"

	// ErrSamePerm is the error message that should be sent when
	// the user already has the same as requested permission.
	ErrSamePerm = "Target already has the same permission"

	// ErrUserNotFound is the error message that should be sent when the
	// target's user id cannot be found in the sibyl database.
	ErrUserNotFound = "User not found"

	// ErrRestoredOnly is the error message that should be sent when user is
	// trying to use a method which can only be used on a target user with `Restored`
	// status, but the target user doesn't have that status currently.
	ErrRestoredOnly = "This feature can only be used on users with Restored status"

	// ErrUserNotRegistered is the error message that should be sent when the
	// target's user is not considered as a registered user.
	ErrUserNotRegistered = "User not registered"

	// ErrUserNotBanned is the error message that should be sent when the
	// target's user id can be found in the sibyl database but
	// it's already unbanned (or never banned).
	ErrUserNotBanned = "User is not banned"

	// ErrNoData is the error message that should be sent when the
	// inspector hasn't provided any data for us.
	ErrNoData = "No raw data provided"

	// ErrBadData is the error message that should be sent when the
	// inspector has provided malformatted data.
	ErrBadData = "Raw data should be in JSON format"

	// ErrCannotBeRevoked is the error message that should be sent when the
	// target token has been revoked too many times.
	ErrCannotBeRevoked = "Token cannot be revoked anymore"

	// ErrTooManyUsers is the error message that should be sent when the
	// amount of users in multi* endpoints are too many.
	// Edit sibyl/entryPoints/banHandlers/constants.go if this is edited
	ErrTooManyUsers = "Too many users, only %d users are permitted"

	// ErrNoReason is the error message that should be sent when the user
	// has sent a request without providing any reason.
	ErrNoReason = "Reason is required for this action"

	// ErrNoMessage is the error message that should be sent when the user
	// has sent a request without providing any message.
	ErrNoMessage = "Message is required for this action"

	// ErrNoSource is the error message that should be sent when the user
	// has sent a request without providing any source.
	ErrNoSource = "Scan source is required for this action"

	// ErrCannotBeReported is the error message that should be sent when the user
	// has sent a report request for a user that cannot be reported.
	ErrCannotBeReported = "User cannot be scanned"

	// ErrCannotBeReported is the error message that should be sent when the user
	// has sent a report request for a user that cannot be reported.
	ErrCannotBeBanned = "User cannot be banned"

	// ErrCannotReportYourself is the error message that should be sent when the user
	// has sent a report request for itself.
	ErrCannotReportYourself = "You can't scan yourself"

	// ErrCannotBanYourself is the error message that should be sent when the user
	// has sent a report request for itself.
	ErrCannotBanYourself = "You can't ban yourself"

	// ErrUserAlreadyBanned is the error message that should be sent when
	// the target's user id can be found in the sibyl database but
	// it's already banned with the exact same parameters.
	ErrUserAlreadyBanned = "User is already banned with the same parameters: "

	// ErrInvalidToken is the error message that should be sent when the
	// current user's token is invalid.
	ErrInvalidToken = "Invalid token"

	// ErrNoToken is the error message that should be sent when the user
	// has sent a request without providing any token.
	// (token-related headers are empty)
	ErrNoToken = "Token is required for this action"

	// ErrInternalServerError is the error message that should be sent when there is nothing
	// wrong from client-side; rather something unexpected had happened on
	// server-side.
	ErrInternalServerError = "Internal server error; incidents have been reported."
)
