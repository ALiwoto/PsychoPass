package entryPoints

const (
	// ErrPermissionDenied is the error message when the user
	// does not have permission to perform the action.
	ErrPermissionDenied = "Permission denied"

	// ErrInvalidUserId is the error message when the target's user id
	// is invalid.
	ErrInvalidUserId = "Invalid user-id provided"

	// ErrUserNotFound is the error message when the target's user id
	// cannot be found in the sibyl database.
	ErrUserNotFound = "User not found"

	// ErrInvalidToken is the error message when the current user's
	// token is invalid.
	ErrInvalidToken = "Invalid token"
	ErrNoToken      = "Token is required for this action"

	ErrInternalServerError = "Internal server error; incidents have been reported."
)
