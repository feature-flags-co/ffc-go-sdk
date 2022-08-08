package datamodel

// Category Represents a separated namespace of storable data items.
// The SDK passes instances of this type to the data store to specify whether it is referring to
// a feature flag, a user segment, etc

var (
	Featuers Category
	Segments Category
	UserTags Category
)

type Category struct {
	Name            string
	PollingApiUrl   string
	StreamingApiUrl string
}

type Item struct {
	item TimestampUserTag
}

func init() {
	Featuers = NewCategory("featureFlags", "/api/public/sdk/latest-feature-flags", "/streaming")
	Segments = NewCategory("segments", "/api/public/sdk/latest-feature-flags", "/streaming")
	UserTags = NewCategory("userTags", "/api/public/sdk/latest-feature-flags", "/streaming")
}
func NewCategory(name string, pollingApiUrl string, streamingApiUrl string) Category {
	return Category{
		Name:            name,
		PollingApiUrl:   pollingApiUrl,
		StreamingApiUrl: streamingApiUrl,
	}
}

func OfCategory(name string) Category {
	return Category{
		Name:            name,
		PollingApiUrl:   "unknown",
		StreamingApiUrl: "unknown",
	}
}
