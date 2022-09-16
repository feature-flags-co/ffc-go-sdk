package data

import (
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"strings"
)

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

func (a *All) ToStorageType() map[Category]map[string]Item {
	return a.Data.ToStorageType()
}

func (a *All) IsProcessData() bool {
	return model.MsgTypeDataSync == strings.ToLower(a.MessageType) &&
		len(a.Data.EventType) > 0 &&
		(model.EventTypeFullOps == strings.ToLower(a.Data.EventType) ||
			model.EventTypePatchOps == strings.ToLower(a.Data.EventType))
}

type TimestampData interface {
	GetId() string
	Archived() bool
	GetTimestamp() int64
	GetType() int
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

func (ff FeatureFlag) GetId() string {
	return ff.Id
}

func (ff FeatureFlag) Archived() bool {
	return ff.IsArchived
}

func (ff FeatureFlag) GetTimestamp() int64 {
	return ff.Timestamp
}

func (ff FeatureFlag) GetType() int {
	return model.FFCFeatureFlag
}

func (ff FeatureFlag) ToArchivedTimestampData() TimestampData {

	adata := ArchivedTimestampData{
		Id:         ff.Id,
		Timestamp:  ff.Timestamp,
		IsArchived: ff.IsArchived,
	}
	return &adata
}

type ArchivedTimestampData struct {
	Id         string `json:"id"`
	IsArchived bool   `json:"isArchived"`
	Timestamp  int64  `json:"timestamp"`
}

func (a *ArchivedTimestampData) GetId() string {
	return a.Id
}

func (a *ArchivedTimestampData) Archived() bool {
	return a.IsArchived
}

func (a *ArchivedTimestampData) GetTimestamp() int64 {
	return a.Timestamp
}

func (a *ArchivedTimestampData) GetType() int {
	return model.FFCArchivedVdata
}

type Segment struct {
	Id         string       `json:"id"`
	IsArchived bool         `json:"isArchived"`
	Timestamp  int64        `json:"timestamp"`
	Included   []string     `json:"included"`
	Excluded   []string     `json:"excluded"`
	Rules      []TargetRule `json:"rules"`
}

func (s *Segment) GetId() string {
	return s.Id
}

func (s *Segment) Archived() bool {
	return s.IsArchived
}

func (s *Segment) GetTimestamp() int64 {
	return s.Timestamp
}

func (s *Segment) GetType() int {
	return model.FFCSegment
}

func (s *Segment) ToArchivedTimestampData() TimestampData {

	adata := ArchivedTimestampData{
		Id:         s.Id,
		Timestamp:  s.Timestamp,
		IsArchived: s.IsArchived,
	}
	return &adata
}
type TimestampUserTag struct {
	model.UserTag
	Id         string `json:"id"`
	IsArchived bool   `json:"isArchived"`
	Timestamp  int64  `json:"timestamp"`
}

func (t TimestampUserTag) GetId() string {
	return t.Id
}

func (t TimestampUserTag) Archived() bool {
	return t.IsArchived
}

func (t TimestampUserTag) GetTimestamp() int64 {
	return t.Timestamp
}

func (t TimestampUserTag) GetType() int {
	return model.FFCSegment
}

func (t TimestampUserTag) ToArchivedTimestampData() TimestampData {

	aData := ArchivedTimestampData{
		Id:         t.Id,
		Timestamp:  t.Timestamp,
		IsArchived: t.IsArchived,
	}
	return &aData
}

type FeatureFlagBasicInfo struct {
	Id                                            string                             `json:"id"`
	Name                                          string                             `json:"name"`
	Type                                          int                                `json:"type"`
	KeyName                                       string                             `json:"keyName"`
	Status                                        string                             `json:"status"`
	IsDefaultRulePercentageRolloutsIncludedInExpt bool                               `json:"isDefaultRulePercentageRolloutsIncludedInExpt"`
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

func (t *TargetIndividuals) IsTargeted(userKeyId string) bool {

	if len(t.Individuals) == 0 {
		return false
	}
	for _, v := range t.Individuals {
		if v.KeyId == userKeyId {
			return true
		}
	}
	return false
}

type FeatureFlagTargetIndividualUser struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	KeyId string `json:"keyId"`
	Email string `json:"email"`
}

type Data struct {
	EventType    string             `json:"eventType"`
	FeatureFlags []FeatureFlag      `json:"featureFlags"`
	Segments     []*Segment         `json:"segments"`
	UserTags     []TimestampUserTag `json:"userTags"`
	Timestamp    int64
}

func (d *Data) ToStorageType() map[Category]map[string]Item {

	// feature flags
	featureFlags := d.FeatureFlags
	featureFlagsMap := make(map[string]Item)
	if len(featureFlags) > 0 {
		for _, v := range featureFlags {
			var timestampData TimestampData
			if v.IsArchived {
				timestampData = v.ToArchivedTimestampData()
			} else {
				timestampData = v
			}
			item := Item{
				Item: timestampData,
			}
			featureFlagsMap[timestampData.GetId()] = item

		}
	}

	// segments
	segments := d.Segments
	segmentsMap := make(map[string]Item)
	if len(segments) > 0 {

		for _, v := range segments {
			var timestampData TimestampData
			if v.IsArchived {
				timestampData = v.ToArchivedTimestampData()
			} else {
				timestampData = v
			}
			item := Item{
				Item: timestampData,
			}
			segmentsMap[timestampData.GetId()] = item
		}
	}

	// user tags
	userTags := d.UserTags
	userTagsMap := make(map[string]Item)
	if len(userTags) > 0 {
		for _, v := range userTags {
			var timestampData TimestampData
			if v.IsArchived {
				timestampData = v.ToArchivedTimestampData()
			} else {
				timestampData = v
			}
			item := Item{
				Item: timestampData,
			}
			userTagsMap[timestampData.GetId()] = item
		}
	}
	dataMap := make(map[Category]map[string]Item, 0)
	dataMap[FeaturesCat] = featureFlagsMap
	dataMap[SegmentsCat] = segmentsMap
	dataMap[UserTagsCat] = userTagsMap

	return dataMap
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
