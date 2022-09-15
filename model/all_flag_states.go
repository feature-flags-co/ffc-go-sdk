package model

import (
	"encoding/json"
	"log"
)

// AllFlagState The object provides a standard return responding the request of getting all flag values from a client sdk
type AllFlagState struct {
	BasicFlagState
	Cache map[string]EvalDetail
}

// NewAllFlagStates init AllFlagState Object
// @Param  success
// @Param  message   the reason without flag value
// @Param  data  a flag value with reason
// @Return the AllFlagState Object
func NewAllFlagStates(success bool, message string, data []EvalDetail) AllFlagState {

	var msg string
	if success {
		msg = "OK"
	} else {
		msg = message
	}

	return AllFlagState{
		BasicFlagState: BasicFlagState{
			Success: success,
			Message: msg,
		},
		Cache: initData(data),
	}
}

// initData change EvalDetail list to map
// @Param EvalDetail list
func initData(data []EvalDetail) map[string]EvalDetail {

	dataMap := make(map[string]EvalDetail, 0)
	if data == nil {
		return dataMap
	}
	for _, detail := range data {
		dataMap[detail.KeyName] = detail
	}
	return dataMap
}

// Empty build a AllFlagStates without flag value
// @Param  the reason without flag value
// @Return a AllFlagStates
func (a *AllFlagState) Empty(message string) AllFlagState {
	return NewAllFlagStates(false, message, nil)
}

// Of build a AllFlagStates
// @Param success true if the last request is successful
// @Param message the reason
// @Param data    all flag values
// @Param <T>     String/Boolean/Numeric Type
// @Return a AllFlagStates
func (a *AllFlagState) Of(success bool, message string, data []EvalDetail) AllFlagState {
	return NewAllFlagStates(success, message, data)
}

// Get return a detail of a given flag key name
// @Param flagKeyName flag key name
// @Return an {@link EvalDetail}
func (a *AllFlagState) Get(flagKeyName string) EvalDetail {
	return a.Cache[flagKeyName]

}

// FromJson build a AllFlagStates from json
// @Param  a string json
// @Return a AllFlagStates
func (a *AllFlagState) FromJson(jsonStr string) AllFlagState {

	var allFlagState AllFlagState
	err := json.Unmarshal([]byte(jsonStr), &allFlagState)
	if err != nil {
		log.Fatalf("convert json string to AllFlagState object error, error = %v", err)
		return AllFlagState{}
	}
	return allFlagState

}
