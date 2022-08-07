package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/common"
)

type FFCClient struct {
	Offline bool
}

func NewFFCClient(envSecret string, config *FFCConfig) FFCClient {

	basicConfig := BasicConfig{OffLine: config.OffLine, EnvSecret: envSecret}
	contextConfig := Context{BasicConfig: basicConfig, HttpConfig: config.HttpConfig}

	stream := NewStreaming(contextConfig, config.StreamingBuilder.StreamingURI)
	stream.Connect()

	return FFCClient{Offline: config.OffLine}
}

func (f *FFCClient) GetAllUserTags() []common.UserTag {
	return []common.UserTag{}
}