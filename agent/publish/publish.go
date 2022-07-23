package publish

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kube-tarian/git-bridge/github"
	"github.com/kube-tarian/git-bridge/models"
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

//GitPublish method gets the composed data and marshal it and publish it to the Nats jetstream
func (m *jsModel) GitPublish(d *models.Gitevent) {
	metricsJson, _ := json.Marshal(d)
	_, err := m.js.Publish(m.eventSubject, metricsJson)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(metricsJson))
	log.Printf("Metrics with eventSubject:%s has been published\n", m.eventSubject)
}

func (m *jsModel) Samplegithubpublish(d interface{}) {

	switch value := d.(type) {
	case *github.PushPayload:
		metricsJson, _ := json.Marshal(value)
		_, err := m.js.Publish(m.eventSubject, metricsJson)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(metricsJson))
		log.Printf("Metrics with eventSubject:%s has been published\n", m.eventSubject)
	}
}
