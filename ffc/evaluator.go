package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/common"
	"github.com/feature-flags-co/ffc-go-sdk/datamodel"
)

type Evaluator struct {
	FeatureFlag datamodel.FeatureFlag
	Segment     datamodel.Segment
}

func NewEvaluator(featureFlag datamodel.FeatureFlag, segment datamodel.Segment) Evaluator {
	return Evaluator{
		FeatureFlag: featureFlag,
		Segment:     segment,
	}
}

func (e *Evaluator) Evaluate(flag datamodel.FeatureFlag, user common.FFCUser, event datamodel.Event) datamodel.EvalResult {

	// TODO will finish this code
	return datamodel.EvalResult{}
}
