package ffc

import (
	"encoding/json"
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/common"
	"github.com/feature-flags-co/ffc-go-sdk/datamodel"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

type Streaming struct {
	BasicConfig  BasicConfig
	HttpConfig   HttpConfig
	StreamingURL string
}

var sockectConn *websocket.Conn

func NewStreaming(config Context, streamingURI string) *Streaming {
	return &Streaming{
		BasicConfig:  config.BasicConfig,
		HttpConfig:   config.HttpConfig,
		StreamingURL: strings.TrimRight(streamingURI, "/") + common.DefaultStreamingPath,
	}
}

// PingOrDataSync websocket ping
func PingOrDataSync(stime *time.Time, msgType string) {

	var timestamp int64
	if stime == nil {
		timestamp = 0
	} else {
		timestamp = stime.UnixNano() / 1e6
	}
	syncMessage := datamodel.NewDataSyncMessage(timestamp, msgType)
	msg, _ := json.Marshal(syncMessage)
	log.Printf("ping message:%v", string(msg))
	if sockectConn != nil {
		err := sockectConn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Ping write error :", err)
			return
		}
	}
}

func (s *Streaming) Connect() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// get env secret from basic config
	envSecret := s.BasicConfig.EnvSecret
	token := utils.BuildToken(envSecret)

	// build wss request url
	path := fmt.Sprintf(s.StreamingURL+common.AuthParams, token)
	log.Printf("connecting: %s", path)

	// build request headers
	headers := HeaderBuilderFor(s.HttpConfig)

	// setup web socket connection
	c, rsp, err := websocket.DefaultDialer.Dial(path, headers)
	sockectConn = c

	// send data sync message
	PingOrDataSync(nil, common.MsgTypeDataSync)

	if err != nil {
		log.Fatal("dial error=", err, " rsp=", rsp)
	}
	log.Printf("connected: %s,%v", path, rsp)
	defer c.Close()
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
			ProcessMessage(string(message))
		}
	}()

	ticker := time.NewTicker(common.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:

			// send ping message to websocket server
			log.Printf("send ping msg %v", t)
			PingOrDataSync(&t, common.MsgTypePing)
			PingOrDataSync(&t, common.MsgTypeDataSync)
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}

func processDateAsync(data datamodel.All) bool {

	eventType := data.EventType
	version := data.Timestamp

	dataMap := data.ToStorageType()

	// init all data to data storage map
	if common.EventTypeFullOps == eventType {
		datamodel.GetDataStorage().Initialization(dataMap, version)
	} else if common.EventTypePatchOps == eventType {

		// update part data to data storage
		for k, v := range dataMap {
			for k1, v1 := range v {
				datamodel.GetDataStorage().Upsert(k, k1, v1, version)
			}
		}
	}
	return true
}

// ProcessMessage receive message from web socket and convert to all object.
// @Param message the data receive from socket
func ProcessMessage(message string) {
	var msgModel datamodel.StreamingMessage
	err := json.Unmarshal([]byte(message), &msgModel)
	if err != nil {
		log.Fatalf("process message to StreamingMessage object error, error = %v", err)
		return
	}

	// process data sync message
	if common.MsgTypeDataSync == msgModel.MessageType {
		var all datamodel.All
		err = json.Unmarshal([]byte(message), &all)
		if err != nil {
			log.Fatalf("process message to All object error, error = %v", err)
			return
		}
		go processDateAsync(all)
	}
}
