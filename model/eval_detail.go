package model

import (
	"encoding/json"
	"log"
)

// EvalDetail An object combines the result of a flag evaluation with an explanation of how it was calculated.
// This object contains the details of evaluation of feature flag.
type EvalDetail struct {
	Variation interface{}
	Id        int64
	Reason    string
	Name      string
	KeyName   string
}

// OfEvalDetail build method, this method is only for internal use
// @Param variation
// @Param id
// @Param reason
// @Param keyName
// @Param name
// @Return  an EvalDetail
func OfEvalDetail(variation interface{}, id int64, reason string, keyName string, name string) *EvalDetail {

	return &EvalDetail{
		Variation: variation,
		Id:        id,
		Reason:    reason,
		Name:      name,
		KeyName:   keyName,
	}
}

// fromJson  build the method from a json string, this method is only for internal use
// @Param json string
// @Return a EvalDetail object
func (e *EvalDetail) fromJson(jsonStr string) EvalDetail {
	var evalDetail EvalDetail
	err := json.Unmarshal([]byte(jsonStr), &evalDetail)
	if err != nil {
		log.Fatalf("convert json string to AllFlagState object error, error = %v", err)
		return EvalDetail{}
	}
	return evalDetail
}

// jsonfy object converted to json string
// @Param evalDetail EvalDetail object
// @Return json string
func (e *EvalDetail) jsonfy(evalDetail EvalDetail) string {
	data, err := json.Marshal(evalDetail)
	if err != nil {
		log.Printf("evaldetail to json err, err = %v", err)
		return ""
	}
	return string(data)
}

// ToFlagState get FlagState from EvalDetail
func (e *EvalDetail) ToFlagState() FlagState {
	return OfFlagState(e)

}

// IsSuccess Returns true if the flag evaluation returned a good value,
// false if the default value returned
func (e *EvalDetail) IsSuccess() bool {
	return e.Id > 0
}
