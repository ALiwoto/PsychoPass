package sibylValues

// IsInspector returns true if the token's permission
// is inspector.
func (t *Token) IsInspector() bool {
	return t.Permission == Inspector
}

// IsEnforcer returns true if the token's permission
// is enforcer.
func (t *Token) IsEnforcer() bool {
	return t.Permission == Enforcer
}

// CanReport returns true if the token with its current
// permission can report a user to sibyl system or not.
func (t *Token) CanReport() bool {
	return t.Permission > NormalUser
}

// CanBan returns true if the token with its current
// permission can ban/unban a user from sibyl system or not.
func (t *Token) CanBan() bool {
	return t.Permission > Enforcer
}

// CanCreateToken returns true if the token with its current
// permission can create tokens in sibyl system or not.
func (t *Token) CanCreateToken() bool {
	return t.Permission > Inspector
}

// CanCreateToken returns true if the token with its current
// permission can revoke tokens in sibyl system or not.
func (t *Token) CanRevokeToken() bool {
	return t.Permission > Inspector
}

// CanCreateToken returns true if the token with its current
// permission can see stats of another tokens or not.
func (t *Token) CanSeeStats() bool {
	return t.Permission > Enforcer
}

// CanGetToken returns true if the token with its current
// permission can get the token of another user using their id
// or not.
func (t *Token) CanGetToken() bool {
	return t.Permission == Owner
}

// CanGetToken returns true if the token with its current
// permission can change permission of another tokens or not.
func (t *Token) CanChangePermission() bool {
	return t.Permission > Inspector
}
