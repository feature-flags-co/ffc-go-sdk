package ffc

import (
	"encoding/json"
	"github.com/feature-flags-co/ffc-go-sdk/common"
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"log"
	"strings"
)

type Evaluator struct {
	FeatureFlag model.FeatureFlag
	Segment     model.Segment
}

func NewEvaluator(featureFlag model.FeatureFlag, segment model.Segment) Evaluator {
	return Evaluator{
		FeatureFlag: featureFlag,
		Segment:     segment,
	}
}

func (e *Evaluator) Evaluate(flag model.FeatureFlag, user common.FFCUser, event model.Event) *model.EvalResult {
	if len(user.UserName) == 0 || len(flag.Id) == 0 {
		return nil
	}
	return matchUserVariation(flag, user, event)
}

func matchUserVariation(flag model.FeatureFlag, user common.FFCUser, event model.Event) *model.EvalResult {

	// return a value when flag is off or not match prerequisite rule
	var er *model.EvalResult
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
				event.Add(model.OfFlagEventVariation(flag.Info.KeyName, er))
			}
		}
	}()
	return er
}

func matchDefaultUserVariation(flag model.FeatureFlag, user common.FFCUser) *model.EvalResult {

	return nil
}

func matchConditionedUserVariation(flag model.FeatureFlag, user common.FFCUser) *model.EvalResult {
	return nil

}
func matchTargetedUserVariation(flag model.FeatureFlag, user common.FFCUser) *model.EvalResult {
	return nil

}
func matchFeatureFlagDisabledUserVariation(flag model.FeatureFlag, user common.FFCUser, event model.Event) *model.EvalResult {

	return nil
}

func equalsClause(user common.FFCUser, clause model.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	value := clause.Value
	return len(pv) > 0 && pv == value
}

func containsClause(user common.FFCUser, clause model.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	value := clause.Value
	return len(pv) > 0 && strings.Contains(pv, value)

}

func oneOfClause(user common.FFCUser, clause model.RuleItem) bool {

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
func startsWithClause(user common.FFCUser, clause model.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	value := clause.Value
	return len(pv) > 0 && strings.HasPrefix(pv, value)
}

func endsWithClause(user common.FFCUser, clause model.RuleItem) bool {
	pv := user.GetProperty(clause.Property)
	value := clause.Value
	return len(pv) > 0 && strings.HasSuffix(pv, value)

}
func trueClause(user common.FFCUser, clause model.RuleItem) bool {
	pv := user.GetProperty(clause.Property)
	if len(pv) > 0 && strings.ToLower(pv) == "true" {
		return true
	}
	return false
}
func falseClause(user common.FFCUser, clause model.RuleItem) bool {

	pv := user.GetProperty(clause.Property)
	if len(pv) > 0 && strings.ToLower(pv) == "false" {
		return true
	}
	return false
}

func matchRegExClause(user common.FFCUser, clause model.RuleItem) bool {

	// TODO
	return false
}

func inSegmentClause(user common.FFCUser, clause model.RuleItem) bool {

	// TODO
	return false
}

func thanClause(user common.FFCUser, clause model.RuleItem) bool {

	//pv := user.GetProperty(clause.Property)
	//clauseValue := clause.Value

	// TODO

	return false
}

func ifUserMatchClause(user common.FFCUser, clause model.RuleItem) bool {

	var op string
	op = clause.Operation
	// segment hasn't any operation
	if len(op) == 0 {
		op = clause.Property
	}

	if strings.Contains(op, common.EvaThanClause) {
		return thanClause(user, clause)
	}
	if strings.Contains(op, common.EvaEqClause) {
		return equalsClause(user, clause)
	}
	if strings.Contains(op, common.EvaNeqClause) {
		return !equalsClause(user, clause)
	}
	if strings.Contains(op, common.EvaContainsClause) {
		return containsClause(user, clause)
	}

	if strings.Contains(op, common.EvaNotContainClause) {
		return !containsClause(user, clause)
	}

	if strings.Contains(op, common.EvaIsOneOfClause) {
		return oneOfClause(user, clause)
	}

	if strings.Contains(op, common.EvaNotOneOfClause) {
		return !oneOfClause(user, clause)
	}

	if strings.Contains(op, common.EvaStartsWithClause) {
		return startsWithClause(user, clause)
	}

	if strings.Contains(op, common.EvaEndsWithClause) {
		return endsWithClause(user, clause)
	}

	if strings.Contains(op, common.EvaIsTrueClause) {
		return trueClause(user, clause)
	}
	if strings.Contains(op, common.EvaIsFalseClause) {
		return falseClause(user, clause)
	}

	if strings.Contains(op, common.EvaMatchRegexClause) {
		return matchRegExClause(user, clause)
	}
	if strings.Contains(op, common.EvaNotMatchRegexClause) {
		return !matchRegExClause(user, clause)
	}

	if strings.Contains(op, common.EvaIsInSegmentClause) {
		return inSegmentClause(user, clause)
	}

	if strings.Contains(op, common.EvaNotInSegmentClause) {
		return !inSegmentClause(user, clause)
	}
	return false
}
