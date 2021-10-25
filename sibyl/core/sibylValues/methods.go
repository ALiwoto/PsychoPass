package sibylValues

func (t *Token) IsInspector() bool {
	return t.Permission == Inspector
}

func (t *Token) IsEnforcer() bool {
	return t.Permission == Enforcer
}

func (t *Token) CanReport() bool {
	return t.Permission > NormalUser
}

func (t *Token) CanBan() bool {
	return t.Permission > Enforcer
}

func (t *Token) CanCreateToken() bool {
	return t.Permission > Inspector
}

func (t *Token) CanRevokeToken() bool {
	return t.Permission > Inspector
}

func (t *Token) CanSeeStats() bool {
	return t.Permission > Enforcer
}

func (t *Token) CanGetToken() bool {
	return t.Permission == Owner
}

func (t *Token) CanPromoteUser() bool {
	return t.Permission > Inspector
}
