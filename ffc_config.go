package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/streaming"
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
	StreamingBuilder *streaming.StreamingBuilder
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
		return ffb.build()
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
	StreamingBuilder *streaming.StreamingBuilder
	Offline          bool
}

func (c *FFCConfigBuilder) build() *FFCConfig {
	ffcConfig := FFCConfig{
		HttpConfig: HttpConfig{
			ConnectTime: HttpConfigDefaultConnTime,
			SocketTime:  HttpConfigDefaultSocketTime,
		},
		StreamingBuilder: c.StreamingBuilder,
	}
	return &ffcConfig
}

func (c *FFCConfigBuilder) updateProcessorFactory(streamingBuilder *streaming.StreamingBuilder) *FFCConfigBuilder {
	c.StreamingBuilder = streamingBuilder
	return c
}

func (c *FFCConfigBuilder) insightProcessorFactory() *FFCConfigBuilder {
	return c
}

func (c *FFCConfigBuilder) offline(offline bool) *FFCConfigBuilder {
	c.Offline = offline
	return c
}
