package data

import (
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"time"
)

type Event interface {
	IsSendEvent() bool
	Add(element interface{}) Event
}

type NullEvent struct {
}

func (n *NullEvent) IsSendEvent() bool {
	return false
}
func (n *NullEvent) Add(element interface{}) Event {
	return nil
}

type DefaultEvent struct {
	User model.FFCUser
}

func (f *DefaultEvent) IsSendEvent() bool {
	// TODO
	return false
}

func (f *DefaultEvent) Add(element interface{}) *Event {
	// TODO
	return nil
}

type FlagEventVariation struct {
	FeatureFlagKeyName string
	Timestamp          int64
	Variation          *EvalResult
}

func NewFlagEventVariation(featureFlagKeyName string, variation *EvalResult) FlagEventVariation {

	return FlagEventVariation{
		FeatureFlagKeyName: featureFlagKeyName,
		Timestamp:          time.Now().UnixNano() / 1e6,
		Variation:          variation,
	}
}

type FlagEvent struct {
	DefaultEvent
	UserVariations []FlagEventVariation
}

func (f *FlagEvent) IsSendEvent() bool {
	// TODO
	return false
}

func (f *FlagEvent) Add(element interface{}) Event {
	// TODO
	return nil
}

func NewFlagEvent(user model.FFCUser) FlagEvent {
	event := DefaultEvent{
		User: user,
	}
	return FlagEvent{
		DefaultEvent: event,
	}

}

type MetricEvent struct {
}

func (m *MetricEvent) IsSendEvent() bool {
	// TODO
	return false
}

func (m *MetricEvent) Add(element interface{}) Event {
	// TODO
	return nil
}
