package ffc

import "github.com/feature-flags-co/ffc-go-sdk/utils"

type InsightEvent struct {
	MaxRetryTimes int
	RetryInterval int
	HttpConfig    HttpConfig
}

func NewInsightEvent(config HttpConfig, maxRetryTimes int, retryInterval int) *InsightEvent {

	return &InsightEvent{
		MaxRetryTimes: maxRetryTimes,
		RetryInterval: retryInterval,
		HttpConfig:    config,
	}
}
func (i *InsightEvent) SendEvent(eventUrl string, json string) {
	utils.PostJson(eventUrl, json, i.MaxRetryTimes, i.RetryInterval)
}
