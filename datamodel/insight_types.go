package datamodel

import "github.com/feature-flags-co/ffc-go-sdk/common"

type Event interface {
	IsSendEvent() bool
	Add(element interface{}) Event
}

type DefaultEvent struct {
	User common.FFCUser
}

func (f *DefaultEvent) IsSendEvent() bool {
	// TODO
	return false
}

func (f *DefaultEvent) Add(element interface{}) *Event {
	// TODO
	return nil
}

type FlagEvent struct {
	DefaultEvent
	UserVariations []FlagEventVariation
}
type FlagEventVariation struct {
	FeatureFlagKeyName string
	Timestamp          int64
}

func (f *FlagEvent) IsSendEvent() bool {
	// TODO
	return false
}

func (f *FlagEvent) Add(element interface{}) Event {
	// TODO
	return nil
}

func OfFlagEvent(user common.FFCUser) FlagEvent {
	event := DefaultEvent{
		User: user,
	}
	return FlagEvent{
		DefaultEvent: event,
	}

}
