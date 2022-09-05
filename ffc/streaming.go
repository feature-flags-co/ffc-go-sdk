package ffc

import (
	"encoding/json"
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/data"
	"github.com/feature-flags-co/ffc-go-sdk/model"
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

var socketConn *websocket.Conn

func NewStreaming(config Context, streamingURI string) *Streaming {
	return &Streaming{
		BasicConfig:  config.BasicConfig,
		HttpConfig:   config.HttpConfig,
		StreamingURL: strings.TrimRight(streamingURI, "/") + model.DefaultStreamingPath,
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
	syncMessage := data.NewDataSyncMessage(timestamp, msgType)
	msg, _ := json.Marshal(syncMessage)
	log.Printf("ping message:%v", string(msg))
	if socketConn != nil {
		err := socketConn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Ping write error :", err)
			return
		}
	}
}

func (s *Streaming) Start() bool {
	go s.Connect()
	return true
}
func (s *Streaming) IsInitialized() bool {
	return data.GetDataStorage().IsInitialized()
}

func (s *Streaming) Connect() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// get env secret from basic config
	envSecret := s.BasicConfig.EnvSecret
	token := utils.BuildToken(envSecret)

	// build wss request url
	path := fmt.Sprintf(s.StreamingURL+model.AuthParams, token)
	log.Printf("connecting: %s", path)

	// build request headers
	headers := utils.HeaderBuilderFor(s.HttpConfig.Headers)

	// setup web socket connection
	c, rsp, err := websocket.DefaultDialer.Dial(path, headers)
	socketConn = c

	// send data sync message
	PingOrDataSync(nil, model.MsgTypeDataSync)

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

	ticker := time.NewTicker(model.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:

			// send ping message to websocket server
			log.Printf("send ping msg %v", t)
			PingOrDataSync(&t, model.MsgTypePing)
			PingOrDataSync(&t, model.MsgTypeDataSync)
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

func processDateAsync(all data.All) bool {

	eventType := all.EventType
	version := all.Timestamp
	dataMap := all.ToStorageType()

	// init all data to data storage map
	if model.EventTypeFullOps == eventType {
		data.GetDataStorage().Initialization(dataMap, version)
	} else if model.EventTypePatchOps == eventType {

		// update part data to data storage
		for k, v := range dataMap {
			for k1, v1 := range v {
				data.GetDataStorage().Upsert(k, k1, v1, version)
			}
		}
	}
	return true
}

// ProcessMessage receive message from web socket and convert to all object.
// @Param message the data receive from socket
func ProcessMessage(message string) {
	var msgModel data.StreamingMessage
	err := json.Unmarshal([]byte(message), &msgModel)
	if err != nil {
		log.Fatalf("process message to StreamingMessage object error, error = %v", err)
		return
	}

	// process data sync message
	if model.MsgTypeDataSync == msgModel.MessageType {
		var all data.All
		err = json.Unmarshal([]byte(message), &all)
		if err != nil {
			log.Fatalf("process message to All object error, error = %v", err)
			return
		}
		go processDateAsync(all)
	}
}
