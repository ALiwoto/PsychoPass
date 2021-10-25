package entryPoints

const (
	// ErrPermissionDenied is the error message when the user
	// does not have permission to perform the action.
	ErrPermissionDenied = "Permission denied"

	// ErrInvalidUserId is the error message when the target's user id
	// is invalid.
	ErrInvalidUserId = "Invalid user-id"

	// ErrUserNotFound is the error message when the target's user id
	// cannot be found in the sibyl database.
	ErrUserNotFound = "User not found"

	// ErrInvalidToken is the error message when the current user's
	// token is invalid.
	ErrInvalidToken = "Invalid token"

	ErrInternalServerError = "Internal server error. Incidents have been reported."
)

const (
	MessageBanned   = "User was banned"
	MessageUnbanned = "User was unbanned"
)
