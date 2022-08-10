package datamodel

// Category Represents a separated namespace of storable data items.
// The SDK passes instances of this type to the data store to specify whether it is referring to
// a feature flag, a user segment, etc

var (
	FeatuersCat Category
	SegmentsCat Category
	UserTagsCat Category
)

type Category struct {
	Name            string
	PollingApiUrl   string
	StreamingApiUrl string
}

type Item struct {
	item TimestampData
}

func init() {
	FeatuersCat = NewCategory("featureFlags", "/api/public/sdk/latest-feature-flags", "/streaming")
	SegmentsCat = NewCategory("segments", "/api/public/sdk/latest-feature-flags", "/streaming")
	UserTagsCat = NewCategory("userTags", "/api/public/sdk/latest-feature-flags", "/streaming")
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
