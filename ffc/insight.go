package ffc

import "github.com/feature-flags-co/ffc-go-sdk/data"

type Insight struct {
}

func (i *Insight) Send(event data.Event) {
	return
}

func (i *Insight) Flush() {
	return
}
