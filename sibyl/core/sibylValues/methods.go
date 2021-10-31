package sibylValues

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ALiwoto/mdparser/mdparser"
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

// CanBeReported returns true if the token with its current
// permission can be reported to sibyl system or not.
func (t *Token) CanBeReported() bool {
	return t.Permission == NormalUser
}

// CanBeBanned returns true if the token with its current
// permission can be banned on sibyl system or not.
func (t *Token) CanBeBanned() bool {
	return t.Permission == NormalUser
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

// CanRevokeToken returns true if the token with its current
// permission can revoke tokens in Sibyl System or not.
func (t *Token) CanRevokeToken() bool {
	return t.Permission > Inspector
}

// CanSeeStats returns true if the token with its current
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

// CanChangePermission returns true if the token with its current
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

func (t *Token) GetCacheDate() time.Time {
	return t.cacheDate
}

func (t *Token) SetCacheDate() {
	t.cacheDate = time.Now()
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

func (p UserPermission) IsValid() bool {
	switch p {
	case NormalUser, Enforcer, Inspector, Owner:
		return true
	default:
		return false
	}
}

func (p UserPermission) IsOwner() bool {
	return p == Owner
}

//---------------------------------------------------------

func (r *Report) ParseAsMd() mdparser.WMarkDown {
	md := mdparser.GetNormal("\u200D#Report Event:\n")
	md.AppendBoldThis("・User:").AppendNormalThis(" ")
	md.AppendMentionThis(strconv.FormatInt(r.ReporterId, 10), r.ReporterId)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・By " + r.ReporterPermission).AppendNormalThis(" ")
	md.AppendMentionThis(strconv.FormatInt(r.ReporterId, 10), r.ReporterId)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・Reason:").AppendNormalThis(" ")
	md.AppendMonoThis(r.ReportReason)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・Date:").AppendNormalThis(" ")
	md.AppendItalicThis(r.ReportDate)
	md.AppendNormalThis("\n\n")
	md.AppendBoldThis("・Message:").AppendNormalThis(" ")
	md.AppendNormalThis(r.ReportMessage)
	return md
}

//---------------------------------------------------------

func (u *User) GetCacheDate() time.Time {
	return u.cacheDate
}

func (u *User) SetCacheDate() {
	u.cacheDate = time.Now()
}

func (u *User) GetCrimeCoefficientRange() *CrimeCoefficientRange {
	return GetCrimeCoefficientRange(u.CrimeCoefficient)
}

func (u *User) SetAsPastBan() {
	u.CrimeCoefficient = RangePastBanned.GetRandom()
}

func (u *User) IncreaseCrimeCoefficient(reason string) {
	ranges := GetCCRangeByString(reason)
	u.IncreaseCrimeCoefficientByRanges(ranges...)
}

func (u *User) IncreaseCrimeCoefficientAuto() {
	u.IncreaseCrimeCoefficient(u.Reason)
}

func (u *User) IncreaseCrimeCoefficientByRanges(ranges ...*CrimeCoefficientRange) {
	for _, r := range ranges {
		u.CrimeCoefficient += r.GetRandom()
	}
}

//---------------------------------------------------------

func (c *CrimeCoefficientRange) IsInRange(value int) bool {
	return c.start <= value && c.end >= value
}

func (c *CrimeCoefficientRange) IsValueInRange(value *CrimeCoefficientRange) bool {
	if value == nil {
		return false
	}

	return value == c ||
		(c.start <= value.start && c.end >= value.end)
}

func (c *CrimeCoefficientRange) GetRandom() int {
	return rand.Intn(c.end-c.start) + c.start
}

//---------------------------------------------------------
