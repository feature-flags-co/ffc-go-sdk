package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/common"
	"github.com/feature-flags-co/ffc-go-sdk/model"
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

func (e *Evaluator) Evaluate(flag model.FeatureFlag, user common.FFCUser, event model.Event) model.EvalResult {

	// TODO will finish this code
	return model.EvalResult{}
}
