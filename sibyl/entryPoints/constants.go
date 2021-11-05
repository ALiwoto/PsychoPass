package entryPoints

const (
	// ErrPermissionDenied is the error message that should be sent
	// when the user does not have enough permissions to perform the action.
	ErrPermissionDenied = "Permission denied"

	// ErrInvalidUserId is the error message that should be sent when
	// the target's user id is invalid. (contains invalid characters)
	ErrInvalidUserId = "Invalid user-id provided"

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

	// ErrUserNotBanned is the error message that should be sent when the
	// target's user id can be found in the sibyl database but
	// it's already unbanned (or never banned).
	ErrUserNotBanned = "User is not banned"

	// ErrNoReason is the error message that should be sent when the user
	// has sent a request without providing any reason.
	ErrNoReason = "Reason is required for this action"

	// ErrCannotBeReported is the error message that should be sent when the user
	// has sent a report request for a user that cannot be reported.
	ErrCannotBeReported = "User cannot be reported"

	// ErrCannotBeReported is the error message that should be sent when the user
	// has sent a report request for a user that cannot be reported.
	ErrCannotBeBanned = "User cannot be banned"

	// ErrCannotReportYourself is the error message that should be sent when the user
	// has sent a report request for itself.
	ErrCannotReportYourself = "You can't report yourself"

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
