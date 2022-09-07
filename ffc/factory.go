package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/data"
)

type UpdateProcessor interface {

	// Start Starts the client update processing.
	// @return completion status indicates the client has been initialized.
	Start() bool

	// IsInitialized Returns true once the client has been initialized and will never return false again.
	// @return true if the client has been initialized
	IsInitialized() bool
}

type UpdateProcessorFactory interface {

	// CreateUpdateProcessor Creates an implementation instance.
	// @param context allows access to the client configuration
	// @return an {@link UpdateProcessor}
	CreateUpdateProcessor(context Context) UpdateProcessor
}

type InsightProcessor interface {

	// Send  Records an event asynchronously.
	Send(event data.Event)

	// Flush Specifies that any buffered events should be sent as soon as possible, rather than waiting
	// for the next flush interval. This method is asynchronous, so events still may not be sent
	// until a later time.
	Flush()
}

type InsightProcessorFactory interface {

	// CreateInsightProcessor creates an implementation of {@link InsightProcessor}
	CreateInsightProcessor(context Context) InsightProcessor
}

func StreamingBuilderFactory() *StreamingBuilder {
	return NewStreamingBuilder()
}

type NullUpdateProcessorFactory struct {
}

func (n *NullUpdateProcessorFactory) CreateUpdateProcessor(context Context) UpdateProcessor {
	return &NullUpdateProcessor{}
}

type NullUpdateProcessor struct {
}

func (n *NullUpdateProcessor) Start() bool {
	return true
}
func (n *NullUpdateProcessor) IsInitialized() bool {
	return data.GetDataStorage().IsInitialized()
}
