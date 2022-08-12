package ffc

import (
	"log"
	"time"
)

const (
	ConfigDefaultBaseUri        = "https://api.featureflag.co"
	ConfigDefaultEventsUri      = "https://api.featureflag.co"
	ConfigDefaultStartWaitTime  = time.Duration(time.Second)
	HttpConfigDefaultConnTime   = time.Duration(time.Second * 10)
	HttpConfigDefaultSocketTime = time.Duration(time.Second * 15)
)


func init()  {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
}

var ffcConfig *Config
var ffcConfigBuilder *ConfigBuilder

type HttpConfig struct {
	ConnectTime time.Duration
	SocketTime  time.Duration
	Headers     map[string]string
}

type Config struct {
	StartWaitTime    time.Duration
	OffLine          bool
	HttpConfig       HttpConfig
	StreamingBuilder *StreamingBuilder
}

type BasicConfig struct {
	EnvSecret string
	OffLine   bool
}

func DefaultFFCConfig() *Config {
	if ffcConfig != nil {
		return ffcConfig
	} else {
		ffb := ConfigBuilder{}
		return ffb.Build()
	}
}

func DefaultFFCConfigBuilder() *ConfigBuilder {
	if ffcConfigBuilder != nil {
		return ffcConfigBuilder
	} else {
		ffb := ConfigBuilder{}
		return &ffb
	}
}

// ConfigBuilder build data for config object
type ConfigBuilder struct {
	StartWaitTime    time.Duration
	StreamingBuilder *StreamingBuilder
	Offline          bool
}

func (c *ConfigBuilder) Build() *Config {
	ffcConfig := Config{
		HttpConfig: HttpConfig{
			ConnectTime: HttpConfigDefaultConnTime,
			SocketTime:  HttpConfigDefaultSocketTime,
		},
		StreamingBuilder: c.StreamingBuilder,
	}
	return &ffcConfig
}

func (c *ConfigBuilder) UpdateProcessorFactory(streamingBuilder *StreamingBuilder) *ConfigBuilder {
	c.StreamingBuilder = streamingBuilder
	return c
}

func (c *ConfigBuilder) insightProcessorFactory() *ConfigBuilder {
	return c
}

func (c *ConfigBuilder) SetOffline(offline bool) *ConfigBuilder {
	c.Offline = offline
	return c
}
