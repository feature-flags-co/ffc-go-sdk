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
	return false
}

func (f *DefaultEvent) Add(element interface{}) *Event {
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
	if len(f.User.UserName) > 0 && len(f.UserVariations) > 0 {
		return true
	}
	return false
}

func (f *FlagEvent) Add(element interface{}) Event {
	fev := element.(FlagEventVariation)
	f.UserVariations = append(f.UserVariations, fev)
	return f
}

func NewFlagEvent(user model.FFCUser) FlagEvent {
	event := DefaultEvent{
		User: user,
	}
	return FlagEvent{
		DefaultEvent:   event,
		UserVariations: make([]FlagEventVariation, 0),
	}
}

type Metric struct {
	Route        string
	Type         string
	EventName    string
	NumericValue float64
	AppType      string
}

func NewMetric(eventName string, numericValue float64) Metric {
	return Metric{
		Route:        "index/metric",
		Type:         "CustomEvent",
		AppType:      "javaserverside",
		EventName:    eventName,
		NumericValue: numericValue,
	}
}

type MetricEvent struct {
	DefaultEvent
	Metrics []Metric
}

func NewMetricEvent(user model.FFCUser) MetricEvent {
	event := DefaultEvent{
		User: user,
	}
	return MetricEvent{
		DefaultEvent: event,
		Metrics:      make([]Metric, 0),
	}
}

func (m *MetricEvent) IsSendEvent() bool {

	if len(m.User.UserName) > 0 && len(m.Metrics) > 0 {
		return true
	}
	return false
}

func (m *MetricEvent) Add(element interface{}) Event {

	event := element.(Metric)
	m.Metrics = append(m.Metrics, event)
	return m
}
