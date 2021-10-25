package sibylValues

const (
	// Can read from the Sibyl System.
	NormalUser UserPermission = iota
	// Can only report to the Sibyl System.
	Enforcer
	// Can read/write directly to the Sibyl System.
	Inspector
	// Can create/revoke tokens.
	Owner
)
