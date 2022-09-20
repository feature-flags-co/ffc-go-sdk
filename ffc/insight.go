package ffc

import (
	"encoding/json"
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

	if event.IsSendEvent() {
		i.queue.Push(event)

		// TODO  add a go routine to send data to ffc server
		events := make([]map[string]interface{}, 0)
		events = append(events, serializeFlagEvent(event))
		jsonData, err := json.Marshal(events)
		if err != nil {
			log.Printf("envet marshal error, error: %v", err)
		} else {
			i.InsightConfig.sender.SendEvent(i.InsightConfig.EventUrl, string(jsonData))
		}
	}

}

func serializeFlagEvent(event data.Event) map[string]interface{} {

	flagEvent := event.(*data.FlagEvent)
	user := flagEvent.User
	jsonMap := serializeUser(user)

	aList := make([]map[string]interface{}, 0)
	if flagEvent.UserVariations != nil && len(flagEvent.UserVariations) > 0 {
		for _, v := range flagEvent.UserVariations {

			vMap := make(map[string]interface{}, 0)
			vMap["featureFlagKeyName"] = v.FeatureFlagKeyName
			vMap["sendToExperiment"] = v.Variation.SendToExperiment
			vMap["timestamp"] = v.Timestamp
			vMap["localId"] = v.Variation.Index
			vMap["variationValue"] = v.Variation.Value
			vMap["reason"] = v.Variation.Reason

			rMap := make(map[string]interface{}, 0)
			rMap["variation"] = vMap
			aList = append(aList, rMap)
		}
		jsonMap["userVariations"] = aList
	}
	return jsonMap
}

func serializeUser(user model.FFCUser) map[string]interface{} {

	dataMap := make(map[string]interface{}, 0)
	dataMap["userName"] = user.UserName
	dataMap["email"] = user.Email
	dataMap["country"] = user.Country
	dataMap["keyId"] = user.Key

	aList := make([]map[string]string, 0)
	if user.Custom != nil {

		for k, v := range user.Custom {
			cMap := make(map[string]string, 0)
			cMap["name"] = k
			cMap["value"] = v
			aList = append(aList, cMap)
		}
		dataMap["customizedProperties"] = aList
	}
	retMap := make(map[string]interface{}, 0)
	retMap["user"] = dataMap
	return retMap
}

func serializeMetric(event data.Event) map[string]interface{} {

	metricEvent := event.(*data.MetricEvent)
	user := metricEvent.User
	jsonMap := serializeUser(user)

	aList := make([]map[string]interface{}, 0)
	if metricEvent.Metrics != nil && len(metricEvent.Metrics) > 0 {

		for _, v := range metricEvent.Metrics {
			cMap := make(map[string]interface{}, 0)
			cMap["route"] = v.Route
			cMap["type"] = v.Type
			cMap["eventName"] = v.EventName
			cMap["numericValue"] = v.NumericValue
			cMap["appType"] = v.AppType
			aList = append(aList, cMap)
		}
		jsonMap["metrics"] = aList
	}
	return jsonMap
}
