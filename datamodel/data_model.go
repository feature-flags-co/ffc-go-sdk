package datamodel

import "github.com/feature-flags-co/ffc-go-sdk/common"

type StreamingMessage struct {
	MessageType string `json:"messageType"`
}

type DataSyncMessage struct {
	StreamingMessage
	Data InternalData `json:"data"`
}
type InternalData struct {
	Timestamp int64 `json:"timestamp"`
}

type All struct {
	MessageType string `json:"messageType"`
	Data        `json:"data"`
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

type FeatureFlagBasicInfo struct {
	Id                                            string                             `json:"id"`
	Name                                          string                             `json:"name"`
	Type                                          int                                `json:"type"`
	KeyName                                       string                             `json:"keyName"`
	Status                                        string                             `json:"status"`
	IsDefaultRulePercentageRolloutsIncludedInExpt bool                               `json:"IsDefaultRulePercentageRolloutsIncludedInExpt"`
	LastUpdatedTime                               string                             `json:"lastUpdatedTime"`
	DefaultRulePercentageRollouts                 []VariationOptionPercentageRollout `json:"defaultRulePercentageRollouts"`
	VariationOptionWhenDisabled                   VariationOption                    `json:"variationOptionWhenDisabled"`
}

type VariationOptionPercentageRollout struct {
	ExptRollout       float64         `json:"exptRollout"`
	RolloutPercentage []float64       `json:"rolloutPercentage"`
	ValueOption       VariationOption `json:"valueOption"`
}
type VariationOption struct {
	LocalId        int64  `json:"localId"`
	DisplayOrder   int64  `json:"displayOrder"`
	VariationValue string `json:"variationValue"`
}

type FeatureFlagPrerequisite struct {
	PrerequisiteFeatureFlagId  string          `json:"prerequisiteFeatureFlagId"`
	ValueOptionsVariationValue VariationOption `json:"valueOptionsVariationValue"`
}

type TargetRule struct {
	RuleId                          string                             `json:"ruleId"`
	RuleName                        string                             `json:"ruleName"`
	IsIncludedInExpt                bool                               `json:"isIncludedInExpt"`
	RuleJsonContent                 []RuleItem                         `json:"ruleJsonContent"`
	ValueOptionsVariationRuleValues []VariationOptionPercentageRollout `json:"valueOptionsVariationRuleValues"`
}
type RuleItem struct {
	Property  string `json:"property"`
	Operation string `json:"operation"`
	Value     string `json:"value"`
}
type TargetIndividuals struct {
	Individuals []FeatureFlagTargetIndividualUser `json:"individuals"`
	ValueOption VariationOption                   `json:"valueOption"`
}

type FeatureFlagTargetIndividualUser struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	KeyId string `json:"keyId"`
	Email string `json:"email"`
}

type Segment struct {
	Id         string       `json:"id"`
	IsArchived bool         `json:"isArchived"`
	Timestamp  int64        `json:"timestamp"`
	Included   []string     `json:"included"`
	Excluded   []string     `json:"excluded"`
	Rules      []TargetRule `json:"rules"`
}

type TimestampUserTag struct {
	common.UserTag
	Id         string `json:"id"`
	IsArchived bool   `json:"isArchived"`
	Timestamp  int64  `json:"timestamp"`
}

type Data struct {
	EventType    string             `json:"eventType"`
	FeatureFlags []FeatureFlag      `json:"featureFlags"`
	Segments     []Segment          `json:"segments"`
	UserTags     []TimestampUserTag `json:"userTags"`
	Timestamp    int64              `json:"timestamp"`
}


type TimestampData interface {
	GetId() string
	IsArchived() bool
	GetTimestamp() int64
	GetType() int
}


func NewDataSyncMessage(timestamp int64, msgType string) DataSyncMessage {

	var data InternalData
	if timestamp == 0 {
		data = InternalData{}
	} else {
		data = InternalData{
			Timestamp: timestamp,
		}
	}
	syncMessage := DataSyncMessage{
		Data: data,
		StreamingMessage: StreamingMessage{
			MessageType: msgType,
		},
	}
	return syncMessage

}
