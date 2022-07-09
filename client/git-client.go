package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"client/clickhouse"
	"client/models"

	"github.com/nats-io/nats.go"
)

var (
	token    string = os.Getenv("NATS_TOKEN")
	natsurl  string = os.Getenv("NATS_ADDRESS")
	dbAdress string = os.Getenv("DB_ADDRESS")
	dbPort   string = os.Getenv("DB_PORT")
	url      string = fmt.Sprintf("tcp://%s:%s?debug=true", dbAdress, dbPort)
)

func main() {
	// Connect to NATS
	var sc *nats.Conn
	var err error
	for i := 0; i < 5; i++ {
		nc, err := nats.Connect(natsurl, nats.Name("Github metrics"), nats.Token(token))
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
	js, err := sc.JetStream()
	log.Print(js)
	checkErr(err)
	stream, err := js.StreamInfo("GITMETRICS")
	checkErr(err)
	log.Println(stream)
	// Create durable consumer monitor
	js.Subscribe("GITMETRICS.bridge", func(msg *nats.Msg) {
		msg.Ack()
		var metrics models.Gitevent
		err := json.Unmarshal(msg.Data, &metrics)
		if err != nil {
			log.Fatal(err)
		}
		//Get clickhouse connection and insert event
		clickhouse.InsertEvent(url, metrics)

	}, nats.Durable("GIT_CONSUMER"), nats.ManualAck())

	runtime.Goexit()

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
