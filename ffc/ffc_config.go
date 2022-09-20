package ffc

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
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
	ConnectTime     time.Duration
	SocketTime      time.Duration
	Headers         map[string]string
	TLSClientConfig *tls.Config
	Proxy           func(*http.Request) (*url.URL, error)
}

type Config struct {
	StartWaitTime           time.Duration
	OffLine                 bool
	HttpConfig              HttpConfig
	UpdateProcessorFactory  UpdateProcessorFactory
	InsightProcessorFactory InsightProcessorFactory
	HttpConfigFactory       HttpConfigFactory
}

type BasicConfig struct {
	EnvSecret string
	OffLine   bool
}

func newConfig(builder *ConfigBuilder) *Config {

	// build process factory
	var updateProcessorFactory UpdateProcessorFactory
	if builder.Offline {
		// offline mode
		updateProcessorFactory = &NullUpdateProcessorFactory{}
	} else {
		// Online mode
		if builder.UpdateProcessorFactory == nil {
			updateProcessorFactory = StreamingBuilderFactory()
		} else {
			updateProcessorFactory = builder.UpdateProcessorFactory
		}
	}

	// build insight factory
	var insightProcessorFactory InsightProcessorFactory
	if builder.InsightProcessorFactory == nil {
		insightProcessorFactory = InsightBuilderFactory()
	} else {
		insightProcessorFactory = builder.InsightProcessorFactory
	}

	// build http factory
	var httpConfigFactory HttpConfigFactory
	if builder.HttpConfigFactory == nil {
		httpConfigFactory = HttpConfigBuilderFactory()
	} else {
		httpConfigFactory = builder.HttpConfigFactory
	}
	ffcConfig := Config{
		HttpConfig: HttpConfig{
			ConnectTime: HttpConfigDefaultConnTime,
			SocketTime:  HttpConfigDefaultSocketTime,
		},
		UpdateProcessorFactory:  updateProcessorFactory,
		InsightProcessorFactory: insightProcessorFactory,
		HttpConfigFactory:       httpConfigFactory,
		StartWaitTime:           builder.StartWaitTime,
		OffLine:                 builder.Offline,
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
	StartWaitTime           time.Duration
	UpdateProcessorFactory  UpdateProcessorFactory
	InsightProcessorFactory InsightProcessorFactory
	HttpConfigFactory       HttpConfigFactory
	Offline                 bool
}

func (c *ConfigBuilder) Build() *Config {
	return newConfig(c)
}

func (c *ConfigBuilder) SetUpdateProcessorFactory(streamingBuilder *StreamingBuilder) *ConfigBuilder {
	c.UpdateProcessorFactory = streamingBuilder
	return c
}

func (c *ConfigBuilder) SetInsightProcessorFactory(insightProcessorFactory InsightProcessorFactory) *ConfigBuilder {
	c.InsightProcessorFactory = insightProcessorFactory
	return c
}

func (c *ConfigBuilder) SetHttpConfigFactory(httpConfigFactory HttpConfigFactory) *ConfigBuilder {
	c.HttpConfigFactory = httpConfigFactory
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
