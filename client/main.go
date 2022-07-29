package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/kube-tarian/git-bridge/clickhouse"
	"github.com/kube-tarian/git-bridge/models"

	"github.com/nats-io/nats.go"
)

var (
	eventSubject  string = "GITMETRICS.git"
	eventConsumer string = "Git-Consumer"
	token         string = os.Getenv("NATS_TOKEN")
	natsurl       string = os.Getenv("NATS_ADDRESS")
	dbAdress      string = os.Getenv("DB_ADDRESS")
	dbPort        string = os.Getenv("DB_PORT")
	url           string = fmt.Sprintf("tcp://%s:%s?debug=true", dbAdress, dbPort)
)


type config struct {
	nats      string
	natstoken string
}

type jsPool struct {
	js nats.JetStreamContext
}

func NewJsPool(js nats.JetStreamContext) jsPool {
	return jsPool{
		js: js,
	}
}

func main() {

	cfg := config{
		nats:      natsurl,
		natstoken: token,
	}
	clients, err := clickhouse.Initialize(url)
	if err != nil {
		log.Fatal(err)
	}

	js := cfg.openJS()
	log.Print(js)
	pool := NewJsPool(js)
	pool.gitSubscriber(eventSubject, eventConsumer, clients)
	stream, err := js.StreamInfo("GITMETRICS")
	checkErr(err)
	log.Println(stream)
	// Create durable consumer monitor

	runtime.Goexit()

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c config) openJS() nats.JetStreamContext {
	// Connect to NATS
	var sc *nats.Conn
	var err error
	for i := 0; i < 5; i++ {
		nc, err := nats.Connect(c.nats, nats.Name("Github metrics"), nats.Token(c.natstoken))
		if err == nil {
			sc = nc
			break
		}

		fmt.Println("Waiting before connecting to NATS at:", natsurl)
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		log.Fatal("Error establishing connection to NATS:", err)
	}
	// Creates JetStreamContext
	js, err := sc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	return js

}

func (c jsPool) gitSubscriber(subject string, consumer string, conn *clickhouse.DBClient) {
	c.js.Subscribe(subject, func(msg *nats.Msg) {
		msg.Ack()
		metrics := &models.Gitevent{}
		err := json.Unmarshal(msg.Data, metrics)
		if err != nil {
			log.Fatal(err)
		}

		conn.InsertEvent(metrics)
		log.Println("Inserted metrics:", metrics)

	}, nats.Durable(consumer), nats.ManualAck())

}
