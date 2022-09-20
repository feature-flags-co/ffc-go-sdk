package ffc

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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
	i.PostJson(eventUrl, json, i.MaxRetryTimes, i.RetryInterval)
}

func (i *InsightEvent) PostJson(url string, jsonData string, retryTimes int, retryInterval int) string {
	return i.PostJsonWithHeaders(url, jsonData, retryTimes, retryInterval)
}

func (i *InsightEvent) PostJsonWithHeaders(url string, jsonData string, retryTimes int,
	retryInterval int) string {

	request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Fatalf("http new request error, error = %v", err)
		return ""
	}
	if request == nil {
		log.Fatalf("http new request error, request is nil")
		return ""
	}
	if len(i.HttpConfig.Headers) > 0 {
		for k, v := range i.HttpConfig.Headers {
			request.Header.Set(k, v)
		}
	}

	var tLSClientConfig *tls.Config
	if i.HttpConfig.TLSClientConfig == nil {
		tLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		tLSClientConfig = i.HttpConfig.TLSClientConfig
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tLSClientConfig,
			Proxy:           i.HttpConfig.Proxy,
		},
	}

	for i := 0; i <= retryTimes; i++ {
		if i > 0 {
			time.Sleep(time.Duration(time.Duration(retryInterval) * time.Millisecond))
		}
		var doErr error
		response, doErr := client.Do(request)
		if doErr != nil {
			log.Fatalf("http do request error, error = %v", err)
		} else {
			body, _ := ioutil.ReadAll(response.Body)
			fmt.Println("response Body:", string(body))
			response.Body.Close()
			return string(body)
		}
	}
	return ""
}
