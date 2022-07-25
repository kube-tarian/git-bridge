package publish

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type jsModel struct {
	js           nats.JetStreamContext
	eventSubject string
}

type Models struct {
	JS jsModel
}

//NewModels returns a nats js pool
func NewModels(js nats.JetStreamContext, subject string) Models {
	return Models{
		JS: jsModel{
			js:           js,
			eventSubject: subject,
		},
	}
}

func (m *jsModel) Publish(value any) {
	metricsJson, _ := json.Marshal(value)
	_, err := m.js.Publish(m.eventSubject, metricsJson)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(metricsJson))
	log.Printf("Metrics with eventSubject:%s has been published\n", m.eventSubject)
}
