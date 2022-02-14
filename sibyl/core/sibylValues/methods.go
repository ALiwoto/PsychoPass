package sibylValues

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ALiwoto/mdparser/mdparser"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/sibylValues/whatColor"
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

// IsRegistered returns true if the owner of this token is considered as
// a valid registered user in the system.
func (t *Token) IsRegistered() bool {
	return t != nil && t.Permission > NormalUser
}

// CanReport returns true if the token with its current
// permission can report a user to sibyl system or not.
func (t *Token) CanReport() bool {
	return t.Permission > NormalUser
}

// CanBeReported returns true if the token with its current
// permission can be reported to sibyl system or not.
func (t *Token) CanBeReported(agentPerm UserPermission) bool {
	return t.Permission < Inspector && t.Permission <= agentPerm
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

// CanGetGeneralInfo returns true if the token with its current
// permission can get general info of a registered user using their id
// or not.
func (t *Token) CanGetGeneralInfo() bool {
	return t.Permission > NormalUser
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

// CanBeRevoked returns true if this token can be revoked; otherwise false.
func (t *Token) CanBeRevoked() bool {
	if t.Permission >= Inspector {
		return true
	}

	if time.Since(t.LastRevokeDate) < 24*time.Hour {
		if t.RevokeCount >= MaxTokenRevokeCount {
			return false
		}
	}

	return true
}

// CanChangePermission returns true if the token with its current
// permission can change permission of another tokens or not.
func (t *Token) CanChangePermission(pre, target UserPermission) bool {
	return !(t.Permission < Inspector || pre >= t.Permission ||
		target >= t.Permission)
}

// CanTryChangePermission returns true if the token with its current
// permission can try to change permission of another tokens or not.
func (t *Token) CanTryChangePermission(direct bool) bool {
	if direct {
		return t.Permission > Inspector
	}

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

func (t *Token) IsExpired(d time.Duration) bool {
	return time.Since(t.cacheDate) > d
}

func (t *Token) GetFormatedCreatedDate() string {
	return t.CreatedAt.Format("2006-01-02 at 15:04:05")
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

func (p UserPermission) GetNormString() string {
	switch p {
	case Enforcer:
		return "Enforcer"
	case Inspector, Owner:
		return "Inspector"
	default:
		return "Not registered"
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

func (p UserPermission) ToString() string {
	return strconv.Itoa(int(p))
}

//---------------------------------------------------------.

func (t EntityType) IsInvalid() bool {
	switch t {
	case EntityTypeUser,
		EntityTypeBot,
		EntityTypeAdmin,
		EntityTypeOwner:
		return false
	default:
		return true
	}
}

func (t EntityType) ToString() string {
	switch t {
	case EntityTypeUser:
		return "user"
	case EntityTypeBot:
		return "bot"
	case EntityTypeAdmin:
		return "admin"
	case EntityTypeOwner:
		return "owner"
	default:
		return "unknown"
	}
}

func (t EntityType) ToTitle() string {
	switch t {
	case EntityTypeUser:
		return "User"
	case EntityTypeBot:
		return "Bot"
	case EntityTypeAdmin:
		return "Admin"
	case EntityTypeOwner:
		return "Owner"
	case EntityTypeChannel:
		return "Channel"
	case EntityTypeGroup:
		return "Group"
	default:
		return "Unknown"
	}
}

func (t EntityType) IsNormal() bool {
	return t == EntityTypeUser
}

func (t EntityType) IsBot() bool {
	return t == EntityTypeBot
}

func (t EntityType) IsAdmin() bool {
	return t == EntityTypeAdmin
}

func (t EntityType) IsOwner() bool {
	return t == EntityTypeOwner
}

func (t EntityType) IsChannel() bool {
	return t == EntityTypeChannel
}

//---------------------------------------------------------

func (r *Report) getNameById(id int64) string {
	chat, err := HelperBot.GetChat(id)
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

func (r *Report) GetTargetType() string {
	if r.TargetUser < 0 {
		return "Chat"
	}

	return r.TargetType.ToTitle()
}

func (r *Report) IsPending() bool {
	return r.ScanStatus == ScanPending
}

func (r *Report) IsApproved() bool {
	return r.ScanStatus == ScanApproved
}

func (r *Report) IsRejected() bool {
	return r.ScanStatus == ScanRejected
}

func (r *Report) IsClosed() bool {
	return r.ScanStatus == ScanClosed
}

func (r *Report) CanBeChanged() bool {
	return r.ScanStatus == ScanPending
}

func (r *Report) GetStatusString() string {
	switch r.ScanStatus {
	case ScanPending:
		return "pending"
	case ScanApproved:
		return "approved"
	case ScanRejected:
		return "rejected"
	case ScanClosed:
		return "closed"
	default:
		return "unknown"
	}
}

func (r *Report) Approve(agentId int64, newReason string) {
	r.ScanStatus = ScanApproved
	r.AgentDate = time.Now()
	r.AgentId = agentId
	if len(newReason) > 0 {
		r.ReportReason = newReason
	}
}

func (r *Report) Reject(agentId int64, reason string) {
	r.ScanStatus = ScanRejected
	r.AgentDate = time.Now()
	r.AgentId = agentId
	r.AgentReason = reason
}

func (r *Report) Close(agentId int64, reason string) {
	r.ScanStatus = ScanClosed
	r.AgentDate = time.Now()
	r.AgentId = agentId
	r.AgentReason = reason
}

func (r *Report) GetMaxMessage() int {
	if r.ScanStatus == ScanPending {
		return 512
	} else {
		return 128
	}
}

func (r *Report) ParseAsMd() mdparser.WMarkDown {
	md := mdparser.GetNormal("\u200D#SCAN:\n")
	agentId := strconv.FormatInt(r.ReporterId, 10)
	targetId := strconv.FormatInt(r.TargetUser, 10)
	agent := r.getReporterName()
	if len(agent) > 22 {
		// truncate the name if it's just too long
		agent = agent[:22] + "..."
	}

	target := r.getTargetName()
	if len(target) > 22 {
		// truncate the name if it's just too long
		target = target[:22] + "..."
	}

	maxMessage := r.GetMaxMessage()
	var theScanMessage string
	if len(r.ReportMessage) > maxMessage {
		// truncate the message if it's just too long
		theScanMessage = r.ReportMessage[:maxMessage] + "..."
	} else {
		theScanMessage = r.ReportMessage
	}

	var theReason string
	if len(r.ReportReason) > 126 {
		// truncate the message if it's just too long
		theReason = r.ReportReason[:126] + "..."
	} else {
		theReason = r.ReportReason
	}

	md.Bold("・" + r.ReporterPermission.GetNormString() + ": ")

	if len(agent) != 0 {
		md.Mention(agent, r.ReporterId)
		md.Normal(" [").Mono(agentId).Normal("]")
	} else {
		md.Mention("\u200D", r.ReporterId)
		md.Mono(agentId)
	}

	md.Normal("\n")
	md.Bold("・Target: ")

	if len(target) != 0 {
		md.Mention(target, r.TargetUser)
		md.Normal(" [").Mono(targetId).Normal("]")
	} else {
		md.Mention("\u200D", r.TargetUser)
		md.Mono(targetId)
	}

	md.Bold("\n・Scan reason: ")
	md.Mono(theReason)
	md.Bold("\n・Date: ")
	md.Mono(r.ReportDate)
	md.Bold("\n・Type: ")
	md.Mono(r.GetTargetType())
	md.Bold("\n・Scan source: ")
	md.Normal(r.ScanSourceLink)
	//md.Bold("\n・Unique ID: ")
	//md.Mono(r.UniqueId)
	md.Bold("\n・Message: ")
	md.Mono(theScanMessage)

	if !r.IsPending() && r.AgentUser != nil {
		md.Normal("\n\n Scan has been " + r.GetStatusString() + " by ")
		md.Mention(r.AgentUser.FirstName, r.AgentUser.Id)
		md.Normal(" at ").Mono(r.AgentDate.Format("2006-01-02 15:04:05"))
	}

	return md
}

func (r *Report) SetUniqueId() {
	if r.UniqueId != "" {
		return
	}

	r.UniqueId = strconv.FormatInt(time.Now().Unix(), 16)
	r.UniqueId += "-" + strconv.FormatInt(r.ReporterId, 16)
	r.UniqueId += "-" + strconv.FormatInt(r.TargetUser, 16)
}

func (r *Report) IsInvalid() bool {
	return r.UniqueId == ""
}

func (r *Report) GetCacheDate() time.Time {
	return r.cacheDate
}

func (r *Report) SetCacheDate() {
	r.cacheDate = time.Now()
}

func (r *Report) IsExpired(d time.Duration) bool {
	return time.Since(r.cacheDate) > d
}

//---------------------------------------------------------

func (u *User) IsCCValid(t *Token) bool {
	if u.Banned || t == nil {
		return true
	}

	switch t.Permission {
	case Owner, Inspector:
		return u.CrimeCoefficient == 0
	case Enforcer:
		return RangeEnforcer.IsInRange(u.CrimeCoefficient)
	case NormalUser:
		if u.IsPastBanned() {
			return RangeRestored.IsInRange(u.CrimeCoefficient)
		}
		return RangeCivilian.IsInRange(u.CrimeCoefficient)
	}

	// impossible to reach;
	// added for backward compatibility.
	return false
}

func (u *User) IsPastBanned() bool {
	return RangeRestored.IsInRange(u.CrimeCoefficient) || u.BanCount > 0
}

func (u *User) IsCivilian() bool {
	return RangeCivilian.IsInRange(u.CrimeCoefficient)
}

func (u *User) IsRestored() bool {
	return RangeRestored.IsInRange(u.CrimeCoefficient)
}

func (u *User) GetCacheDate() time.Time {
	return u.cacheDate
}

func (u *User) SetCacheDate() {
	u.cacheDate = time.Now()
}

func (u *User) IsExpired(d time.Duration) bool {
	return time.Since(u.cacheDate) > d
}

func (u *User) setHueColor() {
	u.HueColor = whatColor.GetHueColor(u.CrimeCoefficient)
}

func (u *User) GetCrimeCoefficientRange() *CrimeCoefficientRange {
	return GetCrimeCoefficientRange(u.CrimeCoefficient)
}

func (u *User) SetAsRestored(clearHistory bool) {
	u.invalidateFlags()
	u.Banned = false
	u.Reason = ""
	u.Message = ""
	u.BanSourceUrl = ""
	u.BannedBy = 0
	u.Date = time.Now()
	u.CrimeCoefficient = RangeRestored.GetRandom()
	if !clearHistory {
		// internal usage only; not meant to be seen by users.
		// this field is for auto-appeal system; please don't use it as seeing
		// how many times this user has been banned. it shows the past history, not
		// the current status.
		// when a user becomes banned, this field will be 0, so there is a chance
		// for auto-appeal system to work.
		// this ban count will increase each this user is marked as `restored`.
		u.BanCount++
	} else {
		u.BanCount = 0x0
	}
}

func (u *User) ClearHistory() {
	u.BanCount = 0x0
}

func (u *User) IsInvalid() bool {
	return IsInvalidID(u.UserID)
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
	if u.Banned {
		return
	}

	if p == Owner || p == Inspector {
		u.CrimeCoefficient = 0
		return
	}

	switch p {
	case Enforcer:
		u.CrimeCoefficient = RangeEnforcer.GetRandom()
	case NormalUser:
		if u.IsPastBanned() {
			u.CrimeCoefficient = RangeRestored.GetRandom()
			return
		}
		u.CrimeCoefficient = RangeCivilian.GetRandom()
	}
}

func (u *User) CanTryAppealing() bool {
	return u.BanCount < MaxAppealCount
}

func (u *User) CanAppeal() bool {
	return u.CrimeCoefficient <= MaxAppealCoefficient &&
		!u.HasCustomFlag() && !u.IsPerma()
}

func (u *User) IsPerma() bool {
	return strings.Contains(u.Reason, "perma")
}

func (u *User) HasCustomFlag() bool {
	return len(u.BanFlags) != 0 && u.BanFlags[0x0] == BanFlagCustom
}

func (u *User) SetAsBanReason(reason string) {
	u.Reason = reason
}

func (u *User) FormatBanDate() {
	u.BanDate = time.Now().Format("2006-01-02 at 15:04:05")
}

func (u *User) GetDateAsShort() string {
	return u.Date.Format(AppealLogDateFormat)
}

func (u *User) EstimateCrimeCoefficient() string {
	c := u.CrimeCoefficient
	if c > 100 {
		str := strconv.Itoa(c)
		return "over " + str[:len(str)-2] + "00"
	}
	return "under 100"
}

func (u *User) GetStringCrimeCoefficient() string {
	return strconv.Itoa(u.CrimeCoefficient)
}

func (u *User) FormatFlags() mdparser.WMarkDown {
	md := mdparser.GetEmpty()
	if len(u.BanFlags) == 0 {
		return md
	}

	for i, current := range u.BanFlags {
		if i != 0 {
			md.Normal(", ")
		}
		md.Mono(string(current))
	}

	return md
}

func (u *User) FormatCuteFlags() mdparser.WMarkDown {
	md := mdparser.GetEmpty()
	if len(u.BanFlags) == 0 {
		return md
	} else if len(u.BanFlags) == 1 {
		return md.Normal(strings.ToLower(string(u.BanFlags[0x0])))
	}

	for i, current := range u.BanFlags {
		if i != 0 && i != len(u.BanFlags)-1 {
			md.Normal(", ")
		} else if i == len(u.BanFlags)-1 {
			md.Normal(" and ")
		}
		md.Normal(strings.ToLower(string(current)))
	}

	return md
}

func (u *User) FormatDetailStrings(showPrefixes bool) mdparser.WMarkDown {
	md := mdparser.GetEmpty()
	if len(u.BanFlags) == 0 {
		return md
	}

	var details string
	if showPrefixes {
		md.Normal(".\n\n")
	}

	for _, current := range u.BanFlags {
		details = _detailsString[current]
		if len(details) == 0 {
			continue
		}
		md.Normal("• " + details + "\n\n")
	}

	return md
}

func (u *User) EstimateCrimeCoefficientSep() (string, string) {
	c := u.CrimeCoefficient
	if c > 100 {
		str := strconv.Itoa(c)
		return "over ", str[:len(str)-2] + "00"
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
	u.setHueColor()
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

func (u *User) ShouldSaveInDB() bool {
	return u.Banned || !RangeCivilian.IsInRange(u.CrimeCoefficient)
}

func (u *User) ToDominatorData(isBan bool) *AssaultDominatorData {
	var t string
	if isBan {
		t = "ban"
	} else {
		t = "revert"
	}
	return &AssaultDominatorData{
		Type:         t,
		TargetUser:   u.UserID,
		ShortReasons: u.BanFlags,
		LongReason:   u.Reason,
		ScannedBy:    u.BannedBy,
		SrcUrl:       u.BanSourceUrl,
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

func (s *StatValue) IsExpired(max time.Duration) bool {
	return time.Since(s.cacheTime) > max
}
func (s *StatValue) SetCachedTime() {
	s.cacheTime = time.Now()
}

//---------------------------------------------------------

func (m *MultiBanUserInfo) IsInvalid(by int64) bool {
	return m.UserId == by || IsInvalidID(m.UserId) ||
		len(m.Reason) == 0
}

//---------------------------------------------------------

func (m *MultiScanRawData) GenerateID() {
	m.AssociationBanId =
		AssociationScanPrefix + strconv.FormatInt(time.Now().Unix(), 34)
}

func (m *MultiScanRawData) SetCacheDate() {
	m.cacheDate = time.Now()
}

func (m *MultiScanRawData) IsExpired(d time.Duration) bool {
	return time.Since(m.cacheDate) > d
}

func (m *MultiScanRawData) getNameById(id int64) string {
	chat, err := HelperBot.GetChat(id)
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

func (m *MultiScanRawData) getReporterName() string {
	return m.getNameById(m.ReporterId)
}

func (m *MultiScanRawData) GetSingleReason() string {
	for _, current := range m.Users {
		if current.Reason != "" {
			if len(current.Reason) > 256 {
				return current.Reason[:256]
			}

			return current.Reason
		}
	}

	return ""
}

func (m *MultiScanRawData) IsPending() bool {
	return m.Status == ScanPending
}

func (m *MultiScanRawData) IsApproved() bool {
	return m.Status == ScanApproved
}

func (m *MultiScanRawData) IsRejected() bool {
	return m.Status == ScanRejected
}

func (m *MultiScanRawData) IsClosed() bool {
	return m.Status == ScanClosed
}

func (m *MultiScanRawData) CanBeChanged() bool {
	return m.Status == ScanPending
}

func (m *MultiScanRawData) GetStatusString() string {
	switch m.Status {
	case ScanPending:
		return "pending"
	case ScanApproved:
		return "approved"
	case ScanRejected:
		return "rejected"
	case ScanClosed:
		return "closed"
	default:
		return "unknown"
	}
}

func (m *MultiScanRawData) ParseAsMd() mdparser.WMarkDown {
	md := mdparser.GetNormal("#ASSOCIATION_SCAN:\n")
	agentId := strconv.FormatInt(m.ReporterId, 10)
	agent := m.getReporterName()
	if len(agent) > 22 {
		// truncate the name if it's just too long
		agent = agent[:22] + "..."
	}

	var theReason = m.GetSingleReason()

	md.Bold("・" + m.ReporterPermission.GetNormString() + ": ")

	if len(agent) != 0 {
		md.Mention(agent, m.ReporterId)
		md.Normal(" [").Mono(agentId).Normal("]")
	} else {
		md.Mention("\u200D", m.ReporterId)
		md.Mono(agentId)
	}

	md.Bold("\n・Scan reason: ")
	md.Mono(theReason)
	md.Bold("\n・Users: \n")

	for _, current := range m.Users {
		md.Mono(strconv.FormatInt(current.UserId, 10) + "\n")
	}

	md.Bold("\n・Date: ")
	md.Mono(time.Now().Format("2006-01-02 15:04:05"))
	md.Bold("\n・Scan source: ")
	md.Normal(m.Source)
	//md.Bold("\n・Unique ID: ")
	//md.Mono(r.UniqueId)

	if !m.IsPending() && m.AgentUser != nil {
		md.Normal("\n\n Scan has been " + m.GetStatusString() + " by ")
		md.Mention(m.AgentUser.FirstName, m.AgentUser.Id)
		md.Normal(" at ").Mono(m.AgentDate.Format("2006-01-02 15:04:05"))
	}

	return md
}

//---------------------------------------------------------

func (d *AssaultDominatorData) ParseAsText() string {
	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return ""
	}

	return mdparser.GetMono(string(b)).ToString()
}
