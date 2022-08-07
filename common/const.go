package common

import "time"

const (

	// PingInterval web socket ping interval
	PingInterval = time.Duration(time.Second * 10)

	StreamingFullOps     = "full"
	StreamingPatchOps    = "patch"
	AuthParams           = "?token=%s&type=server&version=2"
	DefaultStreamingPath = "/streaming"
)
