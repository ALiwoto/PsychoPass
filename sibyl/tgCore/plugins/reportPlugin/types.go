package reportPlugin

// FullRevertInfo struct holds information about a full-revert
// request coming from user.
type FullRevertInfo struct {
	// AgentId is the agent who sent this full-revert request.
	AgentId int64

	// TargetId is the user-id of the target, which got full-reverted.
	TargetId int64

	// SourceUrl is the source url of the full-revert request.
	SourceUrl string
}
