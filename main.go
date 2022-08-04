package ffc

import (
	"flag"
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/streaming"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func websocket() {
	envSecret := "ZDMzLTY3NDEtNCUyMDIxMTAxNzIxNTYyNV9fMzZfXzQ2X185OF9fZGVmYXVsdF80ODEwNA=="
	streamingBuilder := streaming.NewStreamingBuilder().
		NewStreamingURI("wss://api-dev.featureflag.co")

	config := DefaultFFCConfigBuilder().
		offline(false).
		updateProcessorFactory(streamingBuilder).build()
	client := NewFFCClient(envSecret, config)

	tags := client.getAllUserTags()
	fmt.Println(tags)

}
func main() {
	fmt.Print(utils.BuildToken("ad2sdfad="))
}
