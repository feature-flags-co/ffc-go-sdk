package ffc

import (
	"crypto/tls"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"net/http"
	"net/url"
	"time"
)

type HttpConfigBuilder struct {
	Proxy           func(*http.Request) (*url.URL, error)
	ConnectTime     time.Duration
	SocketTime      time.Duration
	TLSClientConfig *tls.Config
}

func NewHttpConfigBuilder() *HttpConfigBuilder {
	builder := HttpConfigBuilder{
	}
	return &builder
}

func (h *HttpConfigBuilder) CreateHttpConfig(config BasicConfig) HttpConfig {
	return HttpConfig{
		ConnectTime:     h.ConnectTime,
		SocketTime:      h.SocketTime,
		TLSClientConfig: h.TLSClientConfig,
		Headers:         utils.DefaultHeaders(config.EnvSecret),
		Proxy:           h.Proxy,
	}
}

func (h *HttpConfigBuilder) SetConnectTime(connectTime time.Duration) *HttpConfigBuilder {
	h.ConnectTime = connectTime
	return h
}

func (h *HttpConfigBuilder) SetSocketTime(socketTime time.Duration) *HttpConfigBuilder {
	h.SocketTime = socketTime
	return h
}

func (h *HttpConfigBuilder) SetHttpProxy(proxyHost string, proxyPort int) *HttpConfigBuilder {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(proxyHost)
	}
	h.Proxy = proxy
	return h
}

func (h *HttpConfigBuilder) SetSSlConfig(config *tls.Config) *HttpConfigBuilder {
	h.TLSClientConfig = config
	return h
}
