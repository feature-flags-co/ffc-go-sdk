package common

import (
	"encoding/json"
	"log"
)

// BasicFlagState the abstract class of feature flag state, which contains 2 property:
// success and message
// this class and his subclasses are used to communicate between saas/server-side sdk and client side sdk
type BasicFlagState struct {
	Success bool   //if the last evaluation is successful
	Message string //the message of flag state
}

// NewBasicFlagState
// @Param success: if the last evaluation is successful
// @Param message: the message of flag state
// @Return a flag state
func NewBasicFlagState(success bool, message string) BasicFlagState {

	return BasicFlagState{
		Success: success,
		Message: message,
	}
}

// Jsonfy  object converted to json string
// @Return BasicFlagState to json string
func (b *BasicFlagState) Jsonfy() string {

	data, err := json.Marshal(b)

	if err != nil {
		log.Printf("basic flag state to json err, err = %v", err)
		return ""
	}
	return string(data)
}
