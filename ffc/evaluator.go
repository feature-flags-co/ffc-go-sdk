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
	FeatureFlag data.FeatureFlag
	Segment     data.Segment
}

func NewEvaluator(featureFlag data.FeatureFlag, segment data.Segment) Evaluator {
	return Evaluator{
		FeatureFlag: featureFlag,
		Segment:     segment,
	}
}

func (e *Evaluator) Evaluate(flag data.FeatureFlag, user model.FFCUser, event data.Event) *data.EvalResult {
	if len(user.UserName) == 0 || len(flag.Id) == 0 {
		return nil
	}
	return matchUserVariation(flag, user, event)
}

func matchUserVariation(flag data.FeatureFlag, user model.FFCUser, event data.Event) *data.EvalResult {

	// return a value when flag is off or not match prerequisite rule
	var er *data.EvalResult
	er = matchFeatureFlagDisabledUserVariation(flag, user, event)
	if er != nil {
		return er
	}

	//return the value of target user
	er = matchTargetedUserVariation(flag, user)
	if er != nil {
		return er
	}

	//return the value of matched rule
	er = matchConditionedUserVariation(flag, user)
	if er != nil {
		return er
	}
	//get value from default rule
	er = matchDefaultUserVariation(flag, user)
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

func matchDefaultUserVariation(flag data.FeatureFlag, user model.FFCUser) *data.EvalResult {

	return nil
}

func matchConditionedUserVariation(flag data.FeatureFlag, user model.FFCUser) *data.EvalResult {
	return nil

}
func matchTargetedUserVariation(flag data.FeatureFlag, user model.FFCUser) *data.EvalResult {
	return nil

}
func matchFeatureFlagDisabledUserVariation(flag data.FeatureFlag, user model.FFCUser, event data.Event) *data.EvalResult {

	return nil
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
		item := data.GetDataStorage().Get(data.SegmentsCat, v)
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
