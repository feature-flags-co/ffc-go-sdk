package ffc

type Context struct {
	HttpConfig  HttpConfig
	BasicConfig BasicConfig
}

// NewContext create Context object
// @Param envSecret
// @Param config
// @Return a Context object
func NewContext(envSecret string, config *Config) Context {
	basicConfig := BasicConfig{OffLine: config.OffLine, EnvSecret: envSecret}
	httpConfig := config.HttpConfigFactory.CreateHttpConfig(basicConfig)
	return Context{
		BasicConfig: basicConfig,
		HttpConfig:  httpConfig,
	}
}
