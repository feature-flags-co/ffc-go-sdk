package common

import "time"

const (

	// message type
	MsgTypeDataSync = "data-sync"
	MsgTypePing     = "ping"
	MsgTypeDataPong = "pong"

	// PingInterval web socket ping interval
	PingInterval = time.Duration(time.Second * 10)

	// event type
	EventTypeFullOps  = "full"
	EventTypePatchOps = "patch"

	AuthParams           = "?token=%s&type=server&version=2"
	DefaultStreamingPath = "/streaming"

	FFCFeatureFlag     = 100
	FFCArchivedVdata   = 200
	FFCPersistentVdata = 300
	FFCSegment         = 400
	FFCUserTag         = 500

	// evaluator relate
	EvaNoEvalRes                = -1
	EvaReasonUserNotSpecified   = "user not specified"
	EvaReasonFlagOff            = "flag off"
	EvaReasonPrerequisiteFailed = "prerequisite failed"
	EvaReasonTargetMatch        = "target match"
	EvaReasonRuleMatch          = "rule match"
	EvaReasonFallthrough        = "fall through all rules"
	EvaReasonClientNotReady     = "client not ready"
	EvaReasonFlagNotFound       = "flag not found"
	EvaReasonWrongType          = "wrong type"
	EvaReasonError              = "error in evaluation"
	EvaFlagKeyUnknown           = "flag key unknown"
	EvaFlagNameUnknown          = "flag name unknown"
	EvaFlagValueUnknown         = "flag value unknown"
	EvaThanClause               = "Than"
	EvaGeClause                 = "BiggerEqualThan"
	EvaGtClause                 = "BiggerThan"
	EvaLeClause                 = "LessEqualThan"
	EvaLtClause                 = "LessThan"
	EvaEqClause                 = "Equal"
	EvaNeqClause                = "NotEqual"
	EvaContainsClause           = "Contains"
	EvaNotContainClause         = "NotContain"
	EvaIsOneOfClause            = "IsOneOf"
	EvaNotOneOfClause           = "NotOneOf"
	EvaStartsWithClause         = "StartsWith"
	EvaEndsWithClause           = "EndsWith"
	EvaIsTrueClause             = "IsTrue"
	EvaIsFalseClause            = "IsFalse"
	EvaMatchRegexClause         = "MatchRegex"
	EvaNotMatchRegexClause      = "NotMatchRegex"
	EvaIsInSegmentClause        = "User is in segment"
	EvaNotInSegmentClause       = "User is not in segment"
	EvaFlagDisableStats         = "Disabled"
	EvaFlagEnableStats          = "Enabled"
)
