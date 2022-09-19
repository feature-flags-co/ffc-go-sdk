package ffc

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


}
