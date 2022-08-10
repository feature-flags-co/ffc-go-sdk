package common

import "time"

const (

	// message type
	MsgTypeDataSync = "data-sync"
	MsgTypePing     = "ping"
	MsgTypeDataPong = "pong"

	// PingInterval web socket ping interval
	PingInterval = time.Duration(time.Second * 10)

	// event type
	EventTypeFullOps  = "full"
	EventTypePatchOps = "patch"

	AuthParams           = "?token=%s&type=server&version=2"
	DefaultStreamingPath = "/streaming"

	FFCFeatureFlag     = 100
	FFCArchivedVdata   = 200
	FFCPersistentVdata = 300
	FFCSegment         = 400
	FFCUserTag         = 500
)
