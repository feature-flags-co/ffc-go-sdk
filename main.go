package main

import (
	"flag"
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/ffc"
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
	streamingBuilder := ffc.NewStreamingBuilder().NewStreamingURI("wss://api-dev.featureflag.co")

	config := ffc.DefaultFFCConfigBuilder().
		SetOffline(false).
		UpdateProcessorFactory(streamingBuilder).
		Build()
	client := ffc.NewClient(envSecret, config)

	fmt.Println(client)
	tags := client.GetAllUserTags()
	fmt.Println(tags)

}
func main() {
	websocket()
	fmt.Print(utils.BuildToken("ad2sdfad="))
	select {}
}
