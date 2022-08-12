package ffc

import (
	"testing"
	"time"
)

func TestPing(t *testing.T) {

	tt := timePtr(time.Now())
	PingOrDataSync(tt, "")
}

func timePtr(t time.Time) *time.Time {
	return &t
}
