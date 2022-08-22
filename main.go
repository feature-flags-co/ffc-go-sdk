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
	//websocket()
	//
	ret := utils.PercentageOfKey("test-value")
	fmt.Println(ret)
	fmt.Print(utils.BuildToken("ad2sdfad="))

	//
	//dst := make([]byte, 0)
	//dst2 := make([]byte, 25, 25)
	//ascii85.Encode(dst, []byte("test-value"))
	//fmt.Println(dst)
	//ascii85.Decode(dst2, dst, false)
	//fmt.Println(string(dst2))

	//select {}
}
