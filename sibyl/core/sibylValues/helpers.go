/*
 * This file is part of PsychoPass Project (https://github.com/MinistryOfWelfare/PsychoPass).
 * Copyright (c) 2021-2022 PsychoPass Authors, Ministry of welfare.
 */
package sibylValues

import (
	"strconv"
	"strings"
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/MinistryOfWelfare/PsychoPass/sibyl/core/utils/timeUtils"
)

func NewReport(reason, message, link string, target, reporter int64,
	reporterPerm UserPermission, targetType EntityType) *Report {

	r := &Report{
		ReportReason:       reason,
		ReportMessage:      message,
		ScanSourceLink:     link,
		TargetUser:         target,
		TargetType:         targetType,
		ReporterId:         reporter,
		ReportDate:         timeUtils.GenerateCurrentDateTime(),
		ReporterPermission: reporterPerm,
	}

	r.SetUniqueId()
	return r
}

func ConvertToPermission(value string) (UserPermission, error) {
	value = strings.ToLower(strings.TrimSpace(value))
	// first of all check and see if value is an integer or not
	valueInt, err := strconv.Atoi(value)
	if err == nil {
		perm := UserPermission(valueInt)
		if perm.IsValid() {
			return perm, nil
		}
		// we already know that the value is a valid integer, so there is no
		// chance that the value is a valid permission in string format.
		return NormalUser, ErrInvalidPerm
	}

	switch value {
	case "user", "civilian":
		return NormalUser, nil
	case "enforcer":
		return Enforcer, nil
	case "inspector":
		return Inspector, nil
	case "owner":
		return Owner, nil
	default:
		return NormalUser, ErrInvalidPerm
	}
}

func IsInvalidID(id int64) bool {
	return _invalidUserIDs[id]
}

func GetCrimeCoefficientRange(value int) *CrimeCoefficientRange {
	/*
		for seeing all possible values of crime coefficient ranges
		and related values, please refer to https://t.me/SibylSystem/4
	*/
	switch {
	case value < 0:
		return nil
	case RangeCivilian.IsInRange(value):
		return RangeCivilian
	case RangeRestored.IsInRange(value):
		return RangeRestored
	case RangeEnforcer.IsInRange(value):
		return RangeEnforcer
	case RangeTrolling.IsInRange(value):
		return RangeTrolling
	case RangeSpam.IsInRange(value):
		return RangeSpam
	case RangeEvade.IsInRange(value):
		return RangeEvade
	case RangeCustom.IsInRange(value):
		return RangeCustom
	case RangePsychoHazard.IsInRange(value):
		return RangePsychoHazard
	case RangeMalImp.IsInRange(value):
		return RangeMalImp
	case RangeNSFW.IsInRange(value):
		return RangeNSFW
	case RangeRaid.IsInRange(value):
		return RangeRaid
	case RangeSpamBot.IsInRange(value):
		return RangeSpamBot
	case RangeMassAdd.IsInRange(value):
		return RangeMassAdd
	}

	return nil
}

func RegisterNewPollingValue(ownerId int64, uniqueId PollingUniqueId, timeout time.Duration) *RegisteredPollingValue {
	pValue := &RegisteredPollingValue{
		OwnerId:    ownerId,
		UniqueId:   uniqueId,
		theChannel: make(chan *PollingUserUpdate, 1), // buffered channel with len of 1
	}

	pValue.GenerateContext(timeout)

	registeredPollingValues.Add(uniqueId, pValue)
	return pValue
}

func RegisterNewPersistancePollingValue(ownerId int64, uniqueId PollingUniqueId) *RegisteredPollingValue {
	pValue := &RegisteredPollingValue{
		OwnerId:       ownerId,
		UniqueId:      uniqueId,
		theChannel:    make(chan *PollingUserUpdate, 1), // buffered channel with len of 1
		isPersistance: true,
	}

	registeredPollingValues.Add(uniqueId, pValue)
	return pValue
}

func GetPollingValueByUniqueId(pollingId *PollingIdentifier, timeout time.Duration) *RegisteredPollingValue {
	pValue := registeredPollingValues.Get(pollingId.PollingUniqueId)
	if pValue == nil || pValue.AccessHash != pollingId.PollingAccessHash {
		return nil
	}

	pValue.GenerateContext(timeout)
	return pValue
}

func UnregisterPollingValue(withContext bool, pValue *RegisteredPollingValue) {
	pValue.MarkAsInvalid(withContext)
	registeredPollingValues.Delete(pValue.UniqueId)
}

func BroadcastUpdate(updateValue *PollingUserUpdate, pollingId *PollingIdentifier) {
	if registeredPollingValues == nil || registeredPollingValues.Length() == 0 {
		// no one is listening anyway
		return
	}

	if pollingId != nil {
		// local broadcast.
		pValue := registeredPollingValues.Get(pollingId.PollingUniqueId)
		if !pValue.IsInvalid() && pValue.AccessHash == pollingId.PollingAccessHash {
			// there is no need to use a new goroutine in here, because this function itself
			// should be called in a separated goroutine, even if `pValue.SendUpdate` method
			// blocks the execution, we will have to return after that, so doesn't really
			// matter.
			pValue.SendUpdate(updateValue)
			return
		}
	}

	registeredPollingValues.ForEach(func(_ PollingUniqueId, pValue *RegisteredPollingValue) bool {
		if pValue.IsInvalid() {
			// let the map remove the invalid value from itself.
			return true
		}

		// the reason we are calling this function in a separated goroutine is because
		// if it blocks our execution, the mutex in registeredPollingValues will remain
		// locked, now imagine this blocking happens for all of the registered polling,
		// server **might** get into a deadlock state.
		go pValue.SendUpdate(updateValue)

		if !pValue.IsPersistance() {
			// temporary polling values should be removed from the
			// memory once they get an update, if a polling-value is
			// persistance, it should remain there as long as possible (unless
			// it gets removed from the memory cache because of inactivity)
			pValue.MarkAsInvalid(false)
			return true
		}

		return false
	})
}

func GetCCRangeByString(value string) []*CrimeCoefficientRange {
	value = fixReasonString(strings.ToLower(strings.TrimSpace(value)))
	values := ssg.Split(value, " ", "\n", ",", "|", "\t", ";",
		".", "..", "...", "....", "-", "--", "---")
	var tmp *CrimeCoefficientRange
	var result []*CrimeCoefficientRange
	exists := func(c *CrimeCoefficientRange) bool {
		if len(result) == 0 {
			return false
		}

		for _, v := range result {
			if v != nil && v.IsValueInRange(c) {
				return true
			}
		}
		return false
	}

	for _, current := range values {
		tmp = getCCRangeByString(current)
		if tmp != nil && !exists(tmp) {
			result = append(result, tmp)
		}
	}

	if len(result) == 0 {
		result = append(result, RangeCustom)
	}
	return result
}

func getCCRangeByString(value string) *CrimeCoefficientRange {
	/*
		for seeing all possible values of crime coefficient ranges
		and related values, please refer to https://t.me/SibylSystem/4
	*/
	switch {
	case canMatchStringArray(value, ReasonTrolling):
		return RangeTrolling
	case canMatchStringArray(value, ReasonSpamBot):
		return RangeSpamBot
	case canMatchStringArray(value, ReasonSpam):
		return RangeSpam
	case canMatchStringArray(value, ReasonEvade):
		return RangeEvade
	case canMatchStringArray(value, ReasonMalImp):
		return RangeMalImp
	case canMatchStringArray(value, ReasonPsychoHazard):
		return RangePsychoHazard
	case canMatchStringArray(value, ReasonNSFW):
		return RangeNSFW
	case canMatchStringArray(value, ReasonRaid):
		return RangeRaid
	case canMatchStringArray(value, ReasonMassAdd):
		return RangeMassAdd
	}

	return nil
}

func fixReasonString(value string) string {
	/*
		for seeing all possible values of crime coefficient ranges
		and related values, please refer to https://t.me/SibylSystem/4
	*/
	if len(ReasonTrolling) > 0 {
		value = strings.ReplaceAll(value, "fooling around", ReasonTrolling[0])
	}

	if len(ReasonMassAdd) > 0 {
		value = strings.ReplaceAll(value, "mass add", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "mass-add", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "member scrap", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "member add", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "spam add", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "spam-add", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "bulk add", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "bulk-add", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "group add", ReasonMassAdd[0])
		value = strings.ReplaceAll(value, "group spam add", ReasonMassAdd[0])
	}

	if len(ReasonNSFW) > 0 {
		value = strings.ReplaceAll(value, "n.s.f.w", ReasonNSFW[0])
	}

	if len(ReasonEvade) > 0 {
		value = strings.ReplaceAll(value, "ban evade", ReasonEvade[0])
		value = strings.ReplaceAll(value, "ban evasion", ReasonEvade[0])
		value = strings.ReplaceAll(value, "ban evading", ReasonEvade[0])
	}

	if len(ReasonPsychoHazard) > 0 {
		value = strings.ReplaceAll(value, "psycho hazard", ReasonPsychoHazard[0])
		value = strings.ReplaceAll(value, "psycho-hazard", ReasonPsychoHazard[0])
	}

	if len(ReasonMalImp) > 0 {
		value = strings.ReplaceAll(value, "fake profile", ReasonMalImp[0])
		value = strings.ReplaceAll(value, "fake name", ReasonMalImp[0])
		value = strings.ReplaceAll(value, "fake username", ReasonMalImp[0])
		value = strings.ReplaceAll(value, "fake alt", ReasonMalImp[0])
		value = strings.ReplaceAll(value, "fake id", ReasonMalImp[0])
	}

	return value
}

func canMatchStringArray(value string, array []string) bool {
	for _, v := range array {
		if strings.HasPrefix(value, v) {
			return true
		}
	}
	return false
}

func GetPrettyUptime() string {
	return timeUtils.GetPrettyTimeDuration(time.Since(ServerStartTime), true)
}

// IsForbiddenID function checks if the given ID is forbidden
// or not. forbidden IDs cannot be looked up by any of the API endpoints.
// If they try to interact with these IDs, they will get a 403 forbidden
// response.
func IsForbiddenID(id int64) bool {
	if HelperBot != nil && HelperBot.Id == id {
		return true
	}
	return false
}

// IsPollingTimeoutInvalid method returns true if the value is in the correct
// range of polling timeout.
func IsPollingTimeoutInvalid(value int64) bool {
	return value >= MinPollingTimeout && value <= MaxPollingTimeout
}

func _getRegisteredPollingValues() *ssg.SafeEMap[PollingUniqueId, RegisteredPollingValue] {
	m := ssg.NewSafeEMap[PollingUniqueId, RegisteredPollingValue]()
	m.SetInterval(20 * time.Minute)
	m.SetExpiration(40 * time.Minute)
	m.SetOnExpired(func(key PollingUniqueId, value RegisteredPollingValue) {
		defer func() {
			_ = recover()
		}()
		value.MarkAsInvalid(true)
	})
	m.EnableChecking()

	return m
}
