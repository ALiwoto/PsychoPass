package sibylValues

import "time"

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
	MarkDownV2 = "markdownv2"
)

const (
	MaxReportCacheTime = 50 * time.Minute
)

const (
	reportStateWaiting reportState = iota
	reportStateAccepted
	reportStateClosed
	reportStateDestroyed
)
