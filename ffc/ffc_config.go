package ffc

import (
	"net/http"
	"time"
)

const (
	ConfigDefaultBaseUri        = "https://api.featureflag.co"
	ConfigDefaultEventsUri      = "https://api.featureflag.co"
	ConfigDefaultStartWaitTime  = time.Duration(time.Second)
	HttpConfigDefaultConnTime   = time.Duration(time.Second * 10)
	HttpConfigDefaultSocketTime = time.Duration(time.Second * 15)
)

var ffcConfig *FFCConfig
var ffcConfigBuilder *FFCConfigBuilder

type HttpConfig struct {
	ConnectTime time.Duration
	SocketTime  time.Duration
	Headers     map[string]string
}

type FFCConfig struct {
	StartWaitTime    time.Duration
	OffLine          bool
	HttpConfig       HttpConfig
	StreamingBuilder *StreamingBuilder
}

type BasicConfig struct {
	EnvSecret string
	OffLine   bool
}

func DefaultFFCConfig() *FFCConfig {
	if ffcConfig != nil {
		return ffcConfig
	} else {
		ffb := FFCConfigBuilder{}
		return ffb.Build()
	}
}

func DefaultFFCConfigBuilder() *FFCConfigBuilder {
	if ffcConfigBuilder != nil {
		return ffcConfigBuilder
	} else {
		ffb := FFCConfigBuilder{}
		return &ffb
	}
}

// FFCConfigBuilder build data for ffcconfig object
type FFCConfigBuilder struct {
	StartWaitTime    time.Duration
	StreamingBuilder *StreamingBuilder
	Offline          bool
}

func (c *FFCConfigBuilder) Build() *FFCConfig {
	ffcConfig := FFCConfig{
		HttpConfig: HttpConfig{
			ConnectTime: HttpConfigDefaultConnTime,
			SocketTime:  HttpConfigDefaultSocketTime,
		},
		StreamingBuilder: c.StreamingBuilder,
	}
	return &ffcConfig
}

func (c *FFCConfigBuilder) UpdateProcessorFactory(streamingBuilder *StreamingBuilder) *FFCConfigBuilder {
	c.StreamingBuilder = streamingBuilder
	return c
}

func (c *FFCConfigBuilder) insightProcessorFactory() *FFCConfigBuilder {
	return c
}

func (c *FFCConfigBuilder) SetOffline(offline bool) *FFCConfigBuilder {
	c.Offline = offline
	return c
}

func HeaderBuilderFor(httpConfig HttpConfig) http.Header {

	header := http.Header{}
	for k, v := range httpConfig.Headers {
		header.Add(k, v)
	}
	return header
}
