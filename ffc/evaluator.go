package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/common"
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"log"
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
