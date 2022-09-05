package ffc

import (
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"time"
)

type StreamingBuilder struct {
	StreamingURI    string
	FirstRetryDelay time.Duration
	MaxRetryTimes   int64
}

func NewStreamingBuilder() *StreamingBuilder {
	builder := StreamingBuilder{}
	return &builder
}

func (s *StreamingBuilder) CreateUpdateProcessor(context Context) UpdateProcessor {

	// set default StreamingURI
	if len(s.StreamingURI) == 0 {
		s.StreamingURI = model.ConfigDefaultBaseUri
	}
	stream := NewStreaming(context, s.StreamingURI)
	stream.Start()
	return stream
}

func (s *StreamingBuilder) NewDefaultStreamingURI() *StreamingBuilder {
	s.StreamingURI = model.ConfigDefaultBaseUri
	return s
}

func (s *StreamingBuilder) NewStreamingURI(uri string) *StreamingBuilder {
	s.StreamingURI = uri
	return s
}

func (s *StreamingBuilder) SetFirstRetryDelay(duration time.Duration) *StreamingBuilder {
	s.FirstRetryDelay = duration
	return s
}

func (s *StreamingBuilder) SetMaxRetryTimes(maxRetryTimes int64) *StreamingBuilder {
	s.MaxRetryTimes = maxRetryTimes
	return s
}
