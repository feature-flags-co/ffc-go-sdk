package ffc

import (
	"encoding/json"
	"github.com/feature-flags-co/ffc-go-sdk/data"
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"log"
	"regexp"
	"strings"
)

type Evaluator struct {
	dataStorage data.DataStorage
}

func NewEvaluator(dataStorage data.DataStorage) Evaluator {
	return Evaluator{
		dataStorage: dataStorage,
	}
}

func (e *Evaluator) Evaluate(flag data.FeatureFlag, user model.FFCUser, event data.Event) *data.EvalResult {
	if len(user.UserName) == 0 || len(flag.Id) == 0 {
		return nil
	}
	return e.matchUserVariation(flag, user, event)
}

func (e *Evaluator) matchUserVariation(flag data.FeatureFlag, user model.FFCUser, event data.Event) *data.EvalResult {

	// return a value when flag is off or not match prerequisite rule
	var er *data.EvalResult
	er = e.matchFeatureFlagDisabledUserVariation(flag, user, event)
	if er != nil {
		return er
	}

	//return the value of target user
	er = e.matchTargetedUserVariation(flag, user)
	if er != nil {
		return er
	}

	//return the value of matched rule
	er = e.matchConditionedUserVariation(flag, user)
	if er != nil {
		return er
	}
	//get value from default rule
	er = e.matchDefaultUserVariation(flag, user)
	if er != nil {
		return er
	}

	defer func() {
		if er == nil {
			log.Printf("FFC GO SDK:user %v Feature Flag %v, Flag Value %v", user.Key, flag.Info.KeyName, er.Value)
			if event != nil {
				event.Add(data.OfFlagEventVariation(flag.Info.KeyName, er))
			}
		}
	}()
	return er
}

func (e *Evaluator) matchDefaultUserVariation(flag data.FeatureFlag, user model.FFCUser) *data.EvalResult {

	return getRollOutVariationOption(flag.Info.DefaultRulePercentageRollouts,
		user,
		model.EvaReasonFallthrough,
		flag.ExptIncludeAllRules,
		flag.Info.IsDefaultRulePercentageRolloutsIncludedInExpt,
		flag.Info.KeyName,
		flag.Info.Name)
}

func (e *Evaluator) matchConditionedUserVariation(flag data.FeatureFlag, user model.FFCUser) *data.EvalResult {

	targetRules := flag.Rules
	var rule data.TargetRule
	for _, v := range targetRules {
		if e.ifUserMatchRule(user, v.RuleJsonContent) {
			rule = v
			break
		}
	}
	return getRollOutVariationOption(rule.ValueOptionsVariationRuleValues,
		user,
		model.EvaReasonRuleMatch,
		flag.ExptIncludeAllRules,
		rule.IsIncludedInExpt,
		flag.Info.KeyName,
		flag.Info.Name)
}

func getRollOutVariationOption(rollouts []data.VariationOptionPercentageRollout,
	user model.FFCUser,
	reason string,
	exptIncludeAllRules bool,
	ruleIncludedInExperiment bool,
	flagKeyName string,
	flagName string) *data.EvalResult {

	newUserKey := utils.Base64Encode(user.Key)
	if len(rollouts) == 0 {
		return nil
	}
	var out data.VariationOptionPercentageRollout
	for _, v := range rollouts {
		if utils.IfKeyBelongsPercentage(user.Key, v.RolloutPercentage) {
			out = v
			break
		}
	}
	ret := data.NewEvalResultWithOption(out.ValueOption,
		reason,
		isSendToExperiment(newUserKey, out, exptIncludeAllRules, ruleIncludedInExperiment),
		flagKeyName,
		flagName)
	return &ret

}

func isSendToExperiment(userKey string, out data.VariationOptionPercentageRollout, exptIncludeAllRules bool,
	ruleIncludedInExperiment bool) bool {

	sendToExperimentPercentage := out.ExptRollout
	splittingPercentage := out.RolloutPercentage[1] - out.RolloutPercentage[0]
	if sendToExperimentPercentage == 0 || splittingPercentage == 0 {
		return false
	}


	upperBound := sendToExperimentPercentage / splittingPercentage
	if upperBound > 1 {
		upperBound = 1
	}
	rangs := []float64{0, upperBound}
	return utils.IfKeyBelongsPercentage(userKey, rangs)
}

func (e *Evaluator) matchTargetedUserVariation(flag data.FeatureFlag, user model.FFCUser) *data.EvalResult {

	targets := flag.Targets
	if len(targets) == 0 {
		return nil
	}
	var tg data.TargetIndividuals
	for _, v := range targets {
		if v.IsTargeted(user.Key) {
			tg = v
			break
		}
	}
	ret := data.NewEvalResultWithOption(tg.ValueOption,
		model.EvaReasonTargetMatch,
		flag.ExptIncludeAllRules,
		flag.Info.KeyName,
		flag.Info.Name)
	return &ret

}
func (e *Evaluator) matchFeatureFlagDisabledUserVariation(flag data.FeatureFlag, user model.FFCUser,
	event data.Event) *data.EvalResult {

	if flag.Info.Status == model.EvaFlagDisableStats {
		ret := data.NewEvalResultWithOption(flag.Info.VariationOptionWhenDisabled,
			model.EvaReasonFlagOff,
			false,
			flag.Info.KeyName,
			flag.Info.Name)
		return &ret
	}

	visits := flag.Prerequisites
	var ffp data.FeatureFlagPrerequisite
	for _, v := range visits {
		preFlagId := v.PrerequisiteFeatureFlagId
		if preFlagId != flag.Info.Id {
			item := e.dataStorage.Get(data.FeaturesCat, preFlagId)
			if len(item.Item.GetId()) > 0 {

				er := e.matchUserVariation(item.Item.(data.FeatureFlag), user, event)
				if utils.GetString(er.Index) != v.ValueOptionsVariationValue.VariationValue {
					ffp = v
					break
				}
			}
		}
	}
	log.Printf("ffp = %v", ffp)
	ret := data.NewEvalResultWithOption(flag.Info.VariationOptionWhenDisabled,
		model.EvaReasonPrerequisiteFailed,
		false,
		flag.Info.KeyName,
		flag.Info.Name)
	return &ret
}

func equalsClause(user model.FFCUser, clause data.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	value := clause.Value
	return len(pv) > 0 && pv == value
}

func containsClause(user model.FFCUser, clause data.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	value := clause.Value
	return len(pv) > 0 && strings.Contains(pv, value)

}

func oneOfClause(user model.FFCUser, clause data.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	var values []string
	err := json.Unmarshal([]byte(clause.Value), &values)
	if err != nil {
		log.Printf("oneOfClause to json error, error = %v", err)
		return false
	}
	if len(pv) == 0 {
		return false
	}
	for _, v := range values {
		if pv == v {
			return true
		}
	}
	return false
}
func startsWithClause(user model.FFCUser, clause data.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	value := clause.Value
	return len(pv) > 0 && strings.HasPrefix(pv, value)
}

func endsWithClause(user model.FFCUser, clause data.RuleItem) bool {
	pv := user.GetProperty(clause.Property)
	value := clause.Value
	return len(pv) > 0 && strings.HasSuffix(pv, value)

}
func trueClause(user model.FFCUser, clause data.RuleItem) bool {
	pv := user.GetProperty(clause.Property)
	if len(pv) > 0 && strings.ToLower(pv) == "true" {
		return true
	}
	return false
}
func falseClause(user model.FFCUser, clause data.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	if len(pv) > 0 && strings.ToLower(pv) == "false" {
		return true
	}
	return false
}

func matchRegExClause(user model.FFCUser, clause data.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	value := clause.Value
	reg, _ := regexp.Compile(value)
	ret := reg.Match([]byte(pv))
	return ret
}

func (e *Evaluator) inSegmentClause(user model.FFCUser, clause data.RuleItem) bool {

	pv := user.Key
	var lists []string
	err := json.Unmarshal([]byte(clause.Value), &lists)
	if err != nil {
		log.Printf("inSegmentClause to json error, error = %v", err)
		return false
	}

	for _, v := range lists {
		item := e.dataStorage.Get(data.SegmentsCat, v)
		if item == (data.Item{}) {
			return false
		}
		seg := item.Item.(*data.Segment)
		ret := seg.IsMatchUser(pv)
		if ret == nil {
			rules := seg.Rules
			for _, r := range rules {
				temp := e.ifUserMatchRule(user, r.RuleJsonContent)
				if temp {
					return true
				}
			}
		} else {
			return ret.Value
		}
	}
	return false
}
func thanClause(user model.FFCUser, clause data.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	value := clause.Value

	pvf := utils.GetFloat64(pv)
	valuef := utils.GetFloat64(value)
	switch clause.Operation {
	case model.EvaGeClause:
		return pvf >= valuef
	case model.EvaGtClause:
		return pvf > valuef
	case model.EvaLeClause:
		return pvf <= valuef
	case model.EvaLtClause:
		return pvf < valuef
	default:
		return false

	}
}

func (e *Evaluator) ifUserMatchRule(user model.FFCUser, clauses []data.RuleItem) bool {

	for _, v := range clauses {

		ret := e.ifUserMatchClause(user, v)
		if !ret {
			return false
		}
	}
	return true

}
func (e *Evaluator) ifUserMatchClause(user model.FFCUser, clause data.RuleItem) bool {

	var op string
	op = clause.Operation
	// segment hasn't any operation
	if len(op) == 0 {
		op = clause.Property
	}

	if strings.Contains(op, model.EvaThanClause) {
		return thanClause(user, clause)
	}
	if strings.Contains(op, model.EvaEqClause) {
		return equalsClause(user, clause)
	}
	if strings.Contains(op, model.EvaNeqClause) {
		return !equalsClause(user, clause)
	}
	if strings.Contains(op, model.EvaContainsClause) {
		return containsClause(user, clause)
	}

	if strings.Contains(op, model.EvaNotContainClause) {
		return !containsClause(user, clause)
	}

	if strings.Contains(op, model.EvaIsOneOfClause) {
		return oneOfClause(user, clause)
	}

	if strings.Contains(op, model.EvaNotOneOfClause) {
		return !oneOfClause(user, clause)
	}

	if strings.Contains(op, model.EvaStartsWithClause) {
		return startsWithClause(user, clause)
	}

	if strings.Contains(op, model.EvaEndsWithClause) {
		return endsWithClause(user, clause)
	}

	if strings.Contains(op, model.EvaIsTrueClause) {
		return trueClause(user, clause)
	}
	if strings.Contains(op, model.EvaIsFalseClause) {
		return falseClause(user, clause)
	}

	if strings.Contains(op, model.EvaMatchRegexClause) {
		return matchRegExClause(user, clause)
	}
	if strings.Contains(op, model.EvaNotMatchRegexClause) {
		return !matchRegExClause(user, clause)
	}

	if strings.Contains(op, model.EvaIsInSegmentClause) {
		return e.inSegmentClause(user, clause)
	}

	if strings.Contains(op, model.EvaNotInSegmentClause) {
		return !e.inSegmentClause(user, clause)
	}
	return false
}
