package sibylLogging

import sv "github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues"

// FullRevertInfo struct holds information about a full-revert
// request coming from user.
type FullRevertInfo struct {
	// AgentId is the agent who sent this full-revert request.
	AgentId int64

	// TargetId is the user-id of the target, which got full-reverted.
	TargetId int64

	// SourceUrl is the source url of the full-revert request.
	SourceUrl string

	// CrimeCoefficient field is the cc of the target user before getting
	// reverted.
	CrimeCoefficient int
}

type FullRevertLogHandler func(info *FullRevertInfo)

type AssaultDominatorHandler func(r *sv.AssaultDominatorData)
