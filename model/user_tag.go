package model

type UserTag struct {
	RequestProperty string `json:"requestProperty"`
	Source          string `json:"source"`
	UserProperty    string `json:"userProperty"`
}

var (
	HEADER       = "header"
	QUERY_STRING = "querystring"
	COOKIE       = "cookie"
	POST_BODY    = "body"
)

func NewUserTag(requestProperty string, source string, userProperty string) *UserTag {
	return &UserTag{
		RequestProperty: requestProperty,
		Source:          source,
		UserProperty:    userProperty,
	}
}

func (u *UserTag) Of(requestProperty string, source string, userProperty string) *UserTag {
	return NewUserTag(requestProperty, source, userProperty)
}
