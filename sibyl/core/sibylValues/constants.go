package sibylValues

const (
	// NormalUser Can read from the Sibyl System.
	NormalUser UserPermission = iota
	// Enforcer Can only report to the Sibyl System.
	Enforcer
	// Inspector Can read/write directly to the Sibyl System.
	Inspector
	// Owner Can create/revoke tokens.
	Owner
)

const (
	MarkDownV2          = "markdownv2"
	OwnersTokenFileName = "owners.token"
)

// Factor constants
const (
	LowerCloudyFactor = 80
	UpperCloudyFactor = 100
)
