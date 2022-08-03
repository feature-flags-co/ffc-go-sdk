package streaming

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

func (s *StreamingBuilder) newStreamingURI(uri string) *StreamingBuilder {
	s.StreamingURI = uri
	return s
}

func (s *StreamingBuilder) firstRetryDelay(duration time.Duration) *StreamingBuilder {
	s.firstRetryDelay(duration)
	return s
}

func (s *StreamingBuilder) maxRetryTimes(maxRetryTimes int64) *StreamingBuilder {
	s.MaxRetryTimes = maxRetryTimes
	return s
}
