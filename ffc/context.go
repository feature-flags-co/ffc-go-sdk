package ffc

type Context struct {
	HttpConfig  HttpConfig
	BasicConfig BasicConfig
}

// NewContext create Context object
// @Param httpConfig
// @Param basicConfig
// @Return a Context object
func NewContext(httpConfig HttpConfig, basicConfig BasicConfig) Context {
	return Context{
		BasicConfig: basicConfig,
		HttpConfig:  httpConfig,
	}
}
