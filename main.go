package main

import (
	"flag"
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/ffc"
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"github.com/feature-flags-co/ffc-go-sdk/utils"
	"io"
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

var client ffc.Client

func websocket() {
	//envSecret := "ZDMzLTY3NDEtNCUyMDIxMTAxNzIxNTYyNV9fMzZfXzQ2X185OF9fZGVmYXVsdF80ODEwNA=="
	envSecret := "NWM4LTAzODgtNCUyMDIyMDcwNzE0MzUzN19fMTc3X18yMDZfXzQxNl9fZGVmYXVsdF8zNDY2Yw=="
	//streamingBuilder := ffc.NewStreamingBuilder().NewStreamingURI("wss://api-dev.featureflag.co")

	config := ffc.NewConfigBuilder().
		SetOffline(false).
		//SetUpdateProcessorFactory(streamingBuilder).
		Build()
	client = ffc.NewClient(envSecret, config)
	fmt.Println(client)

	//ffcUser := model.NewFFUserBuilder().
	//	UserName("userName").
	//	Country("country").
	//	Email("email").
	//	Custom("key", "value").Build()
	//
	//flagtStatue := client.VariationDetail("featureFlagKey", ffcUser, "defaultValue")
	//userTags := client.GetAllLatestFlagsVariations(ffcUser)
}
func main() {
	websocket()
	fmt.Print(utils.BuildToken("ad2sdfad="))
	httpServer()

}

func httpServer() {
	http.HandleFunc("/health", health)

	http.HandleFunc("/index", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, request *http.Request) {
	user := model.FFCUser{
		UserName: "test",
	}
	client.IntVariation("featureD", user, 0)
}

func health(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "ok")
}
