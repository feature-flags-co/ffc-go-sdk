package ffc

import (
	"log"
	"time"
)

const (
	ConfigDefaultStartWaitTime  = time.Duration(time.Second)
	HttpConfigDefaultConnTime   = time.Duration(time.Second * 10)
	HttpConfigDefaultSocketTime = time.Duration(time.Second * 15)
)

func init() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
}

var ffcConfigBuilder *ConfigBuilder

type HttpConfig struct {
	ConnectTime time.Duration
	SocketTime  time.Duration
	Headers     map[string]string
}

type Config struct {
	StartWaitTime          time.Duration
	OffLine                bool
	HttpConfig             HttpConfig
	UpdateProcessorFactory UpdateProcessorFactory
}

type BasicConfig struct {
	EnvSecret string
	OffLine   bool
}

func newConfig(builder *ConfigBuilder) *Config {

	var updateProcessorFactory UpdateProcessorFactory
	if builder.Offline {

		// offline mode
		// TODO
	} else {

		// Online mode
		updateProcessorFactory = StreamingBuilderFactory()
	}

	ffcConfig := Config{
		HttpConfig: HttpConfig{
			ConnectTime: HttpConfigDefaultConnTime,
			SocketTime:  HttpConfigDefaultSocketTime,
		},
		UpdateProcessorFactory: updateProcessorFactory,
		StartWaitTime:          builder.StartWaitTime,
	}
	return &ffcConfig
}

func NewConfigBuilder() *ConfigBuilder {
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
	return newConfig(c)
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

func (c *ConfigBuilder) SetStartWaitTime(startWaitTime time.Duration) *ConfigBuilder {
	c.StartWaitTime = startWaitTime
	return c
}
