package model

import "github.com/feature-flags-co/ffc-go-sdk/common"

type EvalResult struct {
	Index            int
	Value            string
	Reason           string
	SendToExperiment bool
	KeyName          string
	Name             string
}

func NewEvalResult(value string, index int, reason string, sendToExperiment bool, keyName string, name string) EvalResult {

	return EvalResult{
		Index:            index,
		Value:            value,
		Reason:           reason,
		SendToExperiment: sendToExperiment,
		KeyName:          keyName,
		Name:             name,
	}
}
func Error(reason string, keyName string, name string) EvalResult {
	return NewEvalResult("", common.EvaNoEvalRes, reason, false, keyName, name)
}
func ErrorWithDefaultValue(defaultValue string, reason string, keyName string, name string) EvalResult {
	return NewEvalResult(defaultValue, common.EvaNoEvalRes, reason, false, keyName, name)
}

