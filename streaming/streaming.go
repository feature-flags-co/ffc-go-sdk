package streaming

import (
	"encoding/json"
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/datamodel"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"time"
)

const (
	StreamingFullOps     = "full"
	StreamingPatchOps    = "patch"
	AuthParams           = "?token=%s&type=server&version=2"
	DefaultStreamingPath = "/streaming"
)

func processDateAsync(data datamodel.All) {

	eventType := data.EventType
	//version := data.Timestamp
	if StreamingFullOps == eventType {

	}

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
	if datamodel.StreamingMsgTypeDataSync == msgModel.MessageType {
		var all datamodel.All
		err = json.Unmarshal([]byte(message), &all)
		if err != nil {
			log.Fatalf("process message to All object error, error = %v", err)
			return
		}
		go processDateAsync(all)
	}
}

func connect() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	envSecret := ""
	token := utils.BuildToken(envSecret)
	streamingURI := ""
	streamingURL := strings.TrimRight(streamingURI, "/") + DefaultStreamingPath
	path := fmt.Sprintf(streamingURL+AuthParams, token)
	u := url.URL{Scheme: "ws", Host: "", Path: path}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
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
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write:", err)
				return
			}
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
