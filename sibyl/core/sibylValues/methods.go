package sibylValues

import (
	"strconv"
	"strings"
)

//---------------------------------------------------------

// IsOwner returns true if the token's permission
// is owner.
func (t *Token) IsOwner() bool {
	return t.Permission == Owner
}

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

// HasRole returns true if and only if this token belongs to a
// user which has a role in the Sibyl System (is not a normal user).
func (t *Token) HasRole() bool {
	return t.Permission > NormalUser
}

// CanBan returns true if the token with its current
// permission can ban/unban a user from Sibyl System or not.
func (t *Token) CanBan() bool {
	return t.Permission > Enforcer
}

// CanCreateToken returns true if the token with its current
// permission can create tokens in Sibyl System or not.
func (t *Token) CanCreateToken() bool {
	return t.Permission > Inspector
}

// CanCreateToken returns true if the token with its current
// permission can revoke tokens in Sibyl System or not.
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

func (t *Token) GetStringPermission() string {
	return t.Permission.GetStringPermission()
}

func (t *Token) GetTitleStringPermission() string {
	return strings.Title(t.Permission.GetStringPermission())
}

//---------------------------------------------------------
func (p UserPermission) GetStringPermission() string {
	switch p {
	case NormalUser:
		return "user"
	case Enforcer:
		return "enforcer"
	case Inspector:
		return "inspector"
	case Owner:
		return "owner"
	default:
		return strconv.Itoa(int(p))
	}
}

//---------------------------------------------------------

func (r *Report) CreateUniqueId() {
	if r.uniqueId != 0 || reportMutex == nil {
		return
	}

	reportMutex.Lock()
	r.uniqueId = currentUniqueId
	currentUniqueId++
	reportUniqueMap[r.uniqueId] = r
	reportMutex.Unlock()
}

func (r *Report) MarkAsAccepted() {
	if r.state == reportStateWaiting {
		r.state = reportStateAccepted
	}
}

func (r *Report) MarkAsClosed() {
	if r.state == reportStateWaiting {
		r.state = reportStateClosed
	}
}

func (r *Report) Destroy() {
	if r.uniqueId != 0 || reportMutex == nil {
		return
	}

	reportMutex.Lock()
	if reportUniqueMap[r.uniqueId] != nil {
		delete(reportUniqueMap, r.uniqueId)
	}
	reportMutex.Unlock()

	r.state = reportStateDestroyed
}

func (r *Report) GetUniqueId() int64 {
	return r.uniqueId
}

func (r *Report) IsAccepted() bool {
	return r.state == reportStateAccepted
}
func (r *Report) IsWaiting() bool {
	return r.state == reportStateWaiting
}

func (r *Report) IsClosed() bool {
	return r.state == reportStateClosed
}

func (r *Report) IsDestroyed() bool {
	return r.state == reportStateDestroyed
}

//---------------------------------------------------------
