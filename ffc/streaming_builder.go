package ffc

import "time"

type StreamingBuilder struct {
	StreamingURI    string
	FirstRetryDelay time.Duration
	MaxRetryTimes   int64
}

func NewStreamingBuilder() *StreamingBuilder {
	builder := StreamingBuilder{
	}
	return &builder
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