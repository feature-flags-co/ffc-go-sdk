package data

import "github.com/feature-flags-co/ffc-go-sdk/model"

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
	return NewEvalResult("", model.EvaNoEvalRes, reason, false, keyName, name)
}
func ErrorWithDefaultValue(defaultValue string, reason string, keyName string, name string) EvalResult {
	return NewEvalResult(defaultValue, model.EvaNoEvalRes, reason, false, keyName, name)
}

