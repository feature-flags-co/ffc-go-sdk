package ffc

import "github.com/feature-flags-co/ffc-go-sdk/common"

type FFCClient struct {
}

func NewFFCClient(envSecret string, config FFCConfig) FFCClient {
	return FFCClient{}
}

func (f *FFCClient) getAllUserTags() []common.UserTag {
	return []common.UserTag{}
}
