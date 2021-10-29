package entryPoints

const (
	// ErrPermissionDenied is the error message that should be sent
	// when the user does not have enough permissions to perform the action.
	ErrPermissionDenied = "Permission denied"

	// ErrInvalidUserId is the error message that should be sent when
	// the target's user id is invalid. (contains invalid characters)
	ErrInvalidUserId = "Invalid user-id provided"

	// ErrUserNotFound is the error message that should be sent when the
	// target's user id cannot be found in the sibyl database.
	ErrUserNotFound = "User not found"

	// ErrUserNotBanned is the error message that should be sent when the
	// target's user id can be found in the sibyl database but
	// it's already unbanned.
	ErrUserNotBanned = "User is not banned"

	// ErrUserAlreadyBanned is the error message that should be sent when
	// the target's user id can be found in the sibyl database but
	// it's already banned with the exact same parameters.
	ErrUserAlreadyBanned = "User is already banned with the same parameters"

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
