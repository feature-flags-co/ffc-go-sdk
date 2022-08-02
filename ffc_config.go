package ffc

import "time"

const (
	ConfigDefaultBaseUri        = "https://api.featureflag.co"
	ConfigDefaultEventsUri      = "https://api.featureflag.co"
	ConfigDefaultStartWaitTime  = time.Duration(time.Second)
	HttpConfigDefaultConnTime   = time.Duration(time.Second * 10)
	HttpConfigDefaultSocketTime = time.Duration(time.Second * 15)
)

var ffcConfig *FFCConfig

type HttpConfig struct {
	ConnectTime time.Duration
	SocketTime  time.Duration
	Headers     map[string]string
}

type FFCConfig struct {
	StartWaitTime time.Duration
	OffLine       bool
	HttpConfig    HttpConfig
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

type FFCConfigBuilder struct {
	StartWaitTime time.Duration
}

func (c *FFCConfigBuilder) build() *FFCConfig {
	ffcConfig := FFCConfig{
		HttpConfig: HttpConfig{
			ConnectTime: HttpConfigDefaultConnTime,
			SocketTime:  HttpConfigDefaultSocketTime,
		},
	}
	return &ffcConfig
}

func (c *FFCConfigBuilder) updateProcessorFactory() *FFCConfigBuilder {
	return c
}

func (c *FFCConfigBuilder) insightProcessorFactory() *FFCConfigBuilder {
	return c
}
