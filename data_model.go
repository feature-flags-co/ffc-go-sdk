package main

type StreamingMessage struct {
	MessageType string `json:"messageType"`
}

type All struct {
	StreamingMessage
	Data Data
}
type FeatureFlagBasicInfo struct {
}

type FeatureFlagPrerequisite struct {
}

type TargetRule struct {
}

type TargetIndividuals struct {
}

type VariationOption struct {
}

type Segment struct {
	Id         string
	IsArchived bool
	Timestamp  int64
	Included   []string
	excluded   []string
	rules      []TargetRule
}

type TimestampUserTag struct {
	Id         string
	IsArchived bool
	Timestamp  int64
}
type FeatureFlag struct {
	Id                  string                    `json:"id"`
	IsArchived          bool                      `json:"isArchived"`
	Timestamp           int64                     `json:"timestamp"`
	ExptIncludeAllRules bool                      `json:"exptIncludeAllRules"`
	Info                FeatureFlagBasicInfo      `json:"ff"`
	Prerequisites       []FeatureFlagPrerequisite `json:"ffp"`
	Rules               []TargetRule              `json:"fftuwmtr"`
	Targets             []TargetIndividuals       `json:"targetIndividuals"`
	Variations          []VariationOption         `json:"variationOptions"`
}

type Data struct {
	EventType    string
	FeatureFlags []FeatureFlag
	segments     []Segment
	UserTags     []TimestampUserTag
	Timestamp    int64
}
