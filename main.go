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
	envSecret := "NWM4LTAzODgtNCUyMDIyMDcwNzE0MzUzN19fMTc3X18yMDZfXzQxNl9fZGVmYXVsdF8zNDY2Yw=="
	//streamingBuilder := ffc.NewStreamingBuilder().NewStreamingURI("wss://api-dev.featureflag.co")
	//
	//insightBuilder := ffc.NewInsightBuilder().SetEventUri("https://api-dev.featureflag.co")

	config := ffc.NewConfigBuilder().
		SetOffline(false).
		//SetUpdateProcessorFactory(streamingBuilder).
		//SetInsightProcessorFactory(insightBuilder).
		Build()
	client = ffc.NewClient(envSecret, config)
	fmt.Println(client)

	//ffcUser = model.NewFFUserBuilder().
	//	UserName("zttt").
	//	Key("zttt").
	//	Country("country").
	//	Email("email").
	//	Custom("key", "value").Build()
}
func main() {
	websocket()
	fmt.Print(utils.BuildToken("ad2sdfad="))
	httpServer()

}

func httpServer() {
	http.HandleFunc("/health", health)

	http.HandleFunc("/index", index)
	http.HandleFunc("/metric", TestTrackMetricWithValue)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, request *http.Request) {
	ffcUser := model.NewFFUserBuilder().
		UserName("zttt").
		Key("zttt").
		Country("country").
		Email("email").
		Custom("key", "value").Build()
	client.GetAllLatestFlagsVariations(ffcUser)
	//client.IntVariation("featureD", user, 0)
}

func TestTrackMetricWithValue(w http.ResponseWriter, request *http.Request) {

	ffcUser := model.NewFFUserBuilder().
		UserName("zttt").
		Key("zttt").
		Country("country").
		Email("email").
		Custom("key", "value").Build()

	var eventName string
	var eventValue string
	values := request.URL.Query()
	eventName = values.Get("ename")
	eventValue = values.Get("evalue")
	client.TrackMetricWithValue(ffcUser, eventName, utils.GetFloat64(eventValue))
}

func health(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "ok")
}
