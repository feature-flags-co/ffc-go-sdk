package model

type FlagState struct {
	BasicFlagState
	Data *EvalDetail
}

// newFlagState The object provides a standard return responding the request of getting a flag value from a client sdk
// @param success
// @param message  the reason without flag value
// @param data  a flag value with reason
// @Return a FlagState
func newFlagState(success bool, message string, date *EvalDetail) *FlagState {

	return &FlagState{
		BasicFlagState: BasicFlagState{
			Message: message,
			Success: success,
		},
		Data: date,
	}
}

// OfFlagState build a good flag stat
// @param data  a flag value with reason
// @Return a FlagState
func OfFlagState(data *EvalDetail) *FlagState {
	var reason string
	if data.IsSuccess() {
		reason = "OK"
	} else {
		reason = data.Reason
	}
	return newFlagState(data.IsSuccess(), reason, data)
}

// Empty build a flag state without flag value
// @Param message the reason without flag value
// @Return a FlagState
func Empty(message string) *FlagState {
	return newFlagState(false, message, nil)
}
