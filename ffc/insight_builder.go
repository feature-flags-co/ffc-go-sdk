package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"math"
)

type InsightBuilder struct {
	MaxRetryTimes int
	RetryInterval int
	HttpConfig    HttpConfig
	EventUri      string
	Capacity      float64
}

func NewInsightBuilder() *InsightBuilder {
	return &InsightBuilder{}
}

func (i *InsightBuilder) CreateInsightProcessor(context Context) InsightProcessor {

	eventSender := i.createInsightEventSender(context)

	flushInterval := model.InsightDefaultFlushInterval.Milliseconds()
	insightConfig := NewInsightConfig(eventSender, i.EventUri, flushInterval, int(math.Max(i.Capacity,
		model.InsightDefaultCapacity)))

	insight := NewInsight(insightConfig)
	return &insight
}

func (i *InsightBuilder) createInsightEventSender(context Context) *InsightEvent {
	return NewInsightEvent(context.HttpConfig, i.MaxRetryTimes, i.RetryInterval)
}

func (i *InsightBuilder) SetEventUri(eventUri string) *InsightBuilder {
	i.EventUri = eventUri
	return i
}

func (i *InsightBuilder) SetCapacity(capacity float64) *InsightBuilder {
	i.Capacity = capacity
	return i
}

func (i *InsightBuilder) SetRetryInterval(retryInterval int) *InsightBuilder {
	i.RetryInterval = retryInterval
	return i
}

func (i *InsightBuilder) SetMaxRetryTimes(maxRetryTimes int) *InsightBuilder {
	i.MaxRetryTimes = maxRetryTimes
	return i
}
