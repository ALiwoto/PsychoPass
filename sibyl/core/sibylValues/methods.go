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

// CanGetAllBans returns true if the token with its current
// permission can get all the banned users.
func (t *Token) CanGetAllBans() bool {
	return t.Permission > NormalUser
}

// CanGetRegisteredList returns true if the token with its current
// permission can get all the registered users.
func (t *Token) CanGetRegisteredList() bool {
	return t.Permission > NormalUser
}

// CanChangePermission returns true if the token with its current
// permission can change permission of another tokens or not.
func (t *Token) CanChangePermission(pre, target UserPermission) bool {
	return !(t.Permission < Inspector || pre >= t.Permission ||
		target >= t.Permission)
}

// CanTryChangePermission returns true if the token with its current
// permission can try to change permission of another tokens or not.
func (t *Token) CanTryChangePermission() bool {
	return t.Permission > Enforcer
}

// CanGetStats returns true if the token with its current
// permission can get all stats of sibyl system or not.
func (t *Token) CanGetStats() bool {
	return t.Permission > Enforcer
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
		return "civilian"
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

func (r *Report) getNameById(id int64) string {
	chat, err := HelperBot.GetChat(r.ReporterId)
	if err != nil || chat == nil {
		return ""
	}
	if len(chat.FirstName) > 0 {
		return chat.FirstName
	}
	if len(chat.LastName) > 0 {
		return chat.LastName
	}
	return chat.Title
}

func (r *Report) getReporterName() string {
	return r.getNameById(r.ReporterId)
}
func (r *Report) getTargetName() string {
	return r.getNameById(r.TargetUser)
}

func (r *Report) ParseAsMd() mdparser.WMarkDown {
	md := mdparser.GetNormal("\u200D#REPORT:\n")
	md.AppendBoldThis("・User: ")
	agent := r.getReporterName()
	target := r.getTargetName()
	if len(target) != 0 {
		md.AppendMentionThis(target, r.TargetUser)
	} else {
		md.AppendMentionThis("\u200D", r.TargetUser)
		md.AppendMonoThis(strconv.FormatInt(r.TargetUser, 10))
	}
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・By " + r.ReporterPermission + " ")

	if len(agent) != 0 {
		md.AppendMentionThis(target, r.ReporterId)
	} else {
		md.AppendMentionThis("\u200D", r.ReporterId)
		md.AppendMonoThis(strconv.FormatInt(r.ReporterId, 10))
	}

	md.AppendNormalThis("\n")
	md.AppendBoldThis("・Reason: ")
	md.AppendMonoThis(r.ReportReason)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・Date: ")
	md.AppendItalicThis(r.ReportDate)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・Report Source: ")
	md.AppendNormalThis(r.ScanSourceLink)
	md.AppendNormalThis("\n")
	md.AppendBoldThis("・Target Message: ")
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
	u.invalidateFlags()
	u.Banned = false
	u.Reason = ""
	u.Message = ""
	u.BanSourceUrl = ""
	u.BannedBy = 0
	u.Date = time.Now()
	u.CrimeCoefficient = RangePastBanned.GetRandom()
}

func (u *User) IncreaseCrimeCoefficient(reason string) {
	ranges := GetCCRangeByString(reason)
	u.IncreaseCrimeCoefficientByRanges(ranges...)
	u.SetBanFlags()
}

func (u *User) IncreaseCrimeCoefficientAuto() {
	u.IncreaseCrimeCoefficient(u.Reason)
}

func (u *User) IncreaseCrimeCoefficientByPerm(p UserPermission) {
	if p == Owner || p == Inspector || u.Banned {
		return
	} else if p == Enforcer {
		u.CrimeCoefficient = rand.Intn(10)
	} else if p == NormalUser {
		u.CrimeCoefficient = RangeCivilian.GetRandom() / 4
	}
}

func (u *User) SetAsBanReason(reason string) {
	u.Reason = reason
}

func (u *User) FormatBanDate() {
	u.BanDate = time.Now().Format("2006-01-02 at 15:04:05")
}

func (u *User) EstimateCrimeCoefficient() string {
	c := u.CrimeCoefficient
	if c > 100 {
		str := strconv.Itoa(c)
		return "over " + str[:len(str)-2] + "00"
	}
	if c < 10 {
		return "under 10"
	}
	return "under 100"
}

func (u *User) EstimateCrimeCoefficientSep() (string, string) {
	c := u.CrimeCoefficient
	if c > 100 {
		str := strconv.Itoa(c)
		return "over ", str[:len(str)-2] + "00"
	}
	if c < 10 {
		return "under ", "10"
	}
	return "under ", "100"
}

func (u *User) SetBanFlags() {
	/*
		FlagTrolling     bool      `json:"-"`
		FlagSpam         bool      `json:"-"`
		FlagEvade        bool      `json:"-"`
		FlagCustom       bool      `json:"-"`
		FlagPsychoHazard bool      `json:"-"`
		FlagMalImp       bool      `json:"-"`
		FlagNsfw         bool      `json:"-"`
		FlagRaid         bool      `json:"-"`
		FlagSpamBot      bool      `json:"-"`
		FlagMassAdd      bool      `json:"-"`
	*/
	u.BanFlags = nil
	if !u.Banned {
		return
	}

	if u.FlagTrolling {
		u.BanFlags = append(u.BanFlags, BanFlagTrolling)
	}
	if u.FlagSpam {
		u.BanFlags = append(u.BanFlags, BanFlagSpam)
	}
	if u.FlagEvade {
		u.BanFlags = append(u.BanFlags, BanFlagEvade)
	}
	if u.FlagCustom {
		u.BanFlags = append(u.BanFlags, BanFlagCustom)
	}
	if u.FlagPsychoHazard {
		u.BanFlags = append(u.BanFlags, BanFlagPsychoHazard)
	}
	if u.FlagMalImp {
		u.BanFlags = append(u.BanFlags, BanFlagMalImp)
	}
	if u.FlagNsfw {
		u.BanFlags = append(u.BanFlags, BanFlagNSFW)
	}
	if u.FlagRaid {
		u.BanFlags = append(u.BanFlags, BanFlagRaid)
	}
	if u.FlagSpamBot {
		u.BanFlags = append(u.BanFlags, BanFlagSpamBot)
	}
	if u.FlagMassAdd {
		u.BanFlags = append(u.BanFlags, BanFlagMassAdd)
	}

	if len(u.BanFlags) == 0 {
		u.BanFlags = append(u.BanFlags, BanFlagCustom)
	}
}

func (u *User) IncreaseCrimeCoefficientByRanges(ranges ...*CrimeCoefficientRange) {
	var cc int
	u.invalidateFlags()
	for _, r := range ranges {
		if r == nil || r.IsValueInRange(RangeCivilian) {
			// ignore civilian
			continue
		}
		cc += r.GetRandom()
		u.validateFlags(r)
	}
	u.CrimeCoefficient = cc
}

func (u *User) validateFlags(r *CrimeCoefficientRange) {
	if r.IsValueInRange(RangeTrolling) && !u.FlagTrolling {
		u.FlagTrolling = true
	}

	if r.IsValueInRange(RangeSpam) && !u.FlagSpam {
		u.FlagSpam = true
	}

	if r.IsValueInRange(RangeEvade) && !u.FlagEvade {
		u.FlagEvade = true
	}

	if r.IsValueInRange(RangeCustom) && !u.FlagCustom {
		u.FlagCustom = true
	}

	if r.IsValueInRange(RangePsychoHazard) && !u.FlagPsychoHazard {
		u.FlagPsychoHazard = true
	}

	if r.IsValueInRange(RangeMalImp) && !u.FlagMalImp {
		u.FlagMalImp = true
	}

	if r.IsValueInRange(RangeNSFW) && !u.FlagNsfw {
		u.FlagNsfw = true
	}

	if r.IsValueInRange(RangeRaid) && !u.FlagRaid {
		u.FlagRaid = true
	}

	if r.IsValueInRange(RangeSpamBot) && !u.FlagSpamBot {
		u.FlagSpamBot = true
	}

	if r.IsValueInRange(RangeMassAdd) && !u.FlagMassAdd {
		u.FlagMassAdd = true
	}
}

func (u *User) invalidateFlags() {
	if len(u.BanFlags) > 0 {
		u.BanFlags = nil
	}
	u.FlagTrolling = false
	u.FlagSpam = false
	u.FlagEvade = false
	u.FlagCustom = false
	u.FlagPsychoHazard = false
	u.FlagMalImp = false
	u.FlagNsfw = false
	u.FlagRaid = false
	u.FlagSpamBot = false
	u.FlagMassAdd = false
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

func (s *StatValue) GetBannedCountString() string {
	return strconv.FormatInt(s.BannedCount, 10)
}

func (s *StatValue) GetTrollingCountString() string {
	return strconv.FormatInt(s.TrollingBanCount, 10)
}

func (s *StatValue) GetSpamCountString() string {
	return strconv.FormatInt(s.SpamBanCount, 10)
}

func (s *StatValue) GetEvadeCountString() string {
	return strconv.FormatInt(s.EvadeBanCount, 10)
}

func (s *StatValue) GetCustomCountString() string {
	return strconv.FormatInt(s.CustomBanCount, 10)
}

func (s *StatValue) GetPsychoHazardCountString() string {
	return strconv.FormatInt(s.PsychoHazardBanCount, 10)
}

func (s *StatValue) GetMalImpBanCountString() string {
	return strconv.FormatInt(s.MalImpBanCount, 10)
}

func (s *StatValue) GetNSFWBanCountString() string {
	return strconv.FormatInt(s.NSFWBanCount, 10)
}

func (s *StatValue) GetRaidBanCountString() string {
	return strconv.FormatInt(s.RaidBanCount, 10)
}

func (s *StatValue) GetMassAddBanCountString() string {
	return strconv.FormatInt(s.MassAddBanCount, 10)
}

func (s *StatValue) GetCloudyCountString() string {
	return strconv.FormatInt(s.CloudyCount, 10)
}

func (s *StatValue) GetTokenCountString() string {
	return strconv.FormatInt(s.TokenCount, 10)
}

func (s *StatValue) GetInspectorsCountString() string {
	return strconv.FormatInt(s.InspectorsCount, 10)
}

func (s *StatValue) GetEnforcesCountString() string {
	return strconv.FormatInt(s.EnforcesCount, 10)
}

//---------------------------------------------------------
