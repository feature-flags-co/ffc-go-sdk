package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/data"
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"log"
)

const (
	FLAGS = iota
	FLUSH
	SHUTDOWN
	METRICS
)

type InsightConfig struct {
	sender        InsightEventSender
	EventUrl      string
	FlushInterval int64
	Capacity      int
}

func NewInsightConfig(sender InsightEventSender, baseUri string, flushInterval int64, capacity int) InsightConfig {

	var uri string
	if len(baseUri) == 0 {
		uri = model.InsightDefaultEventURI
	} else {
		uri = baseUri
	}
	return InsightConfig{
		sender:        sender,
		EventUrl:      uri + model.InsightEventPath,
		FlushInterval: flushInterval,
		Capacity:      capacity,
	}
}

type Insight struct {
	InsightConfig InsightConfig
	queue         *utils.Queue
}

func NewInsight(config InsightConfig) Insight {

	return Insight{
		InsightConfig: config,
		queue:         utils.NewQueue(),
	}
}
func (i *Insight) Send(event data.Event) {

	switch event.(type) {
	case *data.FlagEvent:
		i.putEventAsync(FLAGS, event)

	case *data.MetricEvent:
		i.putEventAsync(METRICS, event)

	default:
		log.Printf("ignore event type: %v; returning default value", event)
	}
	return
}

func (i *Insight) Flush() {
	return
}

func (i *Insight) putEventAsync(insightType uint, event data.Event) {
	i.queue.Push(event)
}
