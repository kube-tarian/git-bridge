package publish

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/vijeyash1/gitevent/models"
)

type jsModel struct {
	js nats.JetStreamContext
}

type Models struct {
	JS jsModel
}

//NewModels returns a nats js pool
func NewModels(js nats.JetStreamContext) Models {
	return Models{
		JS: jsModel{
			js: js,
		},
	}
}

//GitPublish method gets the composed data and marshal it and publish it to the Nats jetstream
func (m *jsModel) GitPublish(d *models.Gitevent) {
	metricsJson, _ := json.Marshal(d)
	_, err := m.js.Publish("GITMETRICS.git", metricsJson)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(metricsJson))
	log.Printf("Metrics with eventSubject:%s has been published\n", "GITMETRICS.git")
}
