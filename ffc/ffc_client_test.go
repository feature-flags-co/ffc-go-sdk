package ffc

import (
	"fmt"
	"github.com/feature-flags-co/ffc-go-sdk/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var client Client
var ffcUser model.FFCUser

func twebsocket() {
	//envSecret := "ZDMzLTY3NDEtNCUyMDIxMTAxNzIxNTYyNV9fMzZfXzQ2X185OF9fZGVmYXVsdF80ODEwNA=="
	envSecret := "NWM4LTAzODgtNCUyMDIyMDcwNzE0MzUzN19fMTc3X18yMDZfXzQxNl9fZGVmYXVsdF8zNDY2Yw=="
	//streamingBuilder := ffc.NewStreamingBuilder().NewStreamingURI("wss://api-dev.featureflag.co")

	config := NewConfigBuilder().
		SetOffline(false).
		//SetUpdateProcessorFactory(streamingBuilder).
		Build()
	client = NewClient(envSecret, config)
	fmt.Println(client)

	ffcUser = model.NewFFUserBuilder().
		UserName("userName").
		Country("country").
		Email("email").
		Custom("key", "value").Build()
	select {
	}
}

func TestMain(m *testing.M) {
	//go twebsocket()
	m.Run()
}

func Test_GetAllUserTags(t *testing.T) {
	go twebsocket()
	time.Sleep(5000)
	tags := client.GetAllUserTags()
	assert.EqualValues(t, 2, len(tags))
}
