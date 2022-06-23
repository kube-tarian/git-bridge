package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/vijeyash1/gitevent/clickhouse"
	"github.com/vijeyash1/gitevent/models"
)

//to read the token from env variables
var token string = os.Getenv("NATS_TOKEN")

var natsurl string = os.Getenv("NATS_ADDRESS")

var dbAdress string = os.Getenv("DB_ADDRESS")

var dbPort string = os.Getenv("DB_PORT")

var url string = fmt.Sprintf("tcp://%s:%s?debug=true", dbAdress, dbPort)

func main() {
	// Connect to NATS

	nc, err := nats.Connect(natsurl, nats.Name("Git metrics"), nats.Token(token))
	checkErr(err)
	log.Println(nc)
	js, err := nc.JetStream()
	log.Print(js)
	checkErr(err)

	stream, err := js.StreamInfo("GITMETRICS")
	checkErr(err)
	log.Println(stream)

	//Get clickhouse connection
	connection, err := clickhouse.GetClickHouseConnection(url)
	if err != nil {
		log.Fatal(err)
	}

	//Create schema
	clickhouse.CreateGitSchema(connection)

	// Create durable consumer monitor
	js.Subscribe("GITMETRICS.git", func(msg *nats.Msg) {
		msg.Ack()
		var metrics models.Gitevent
		err := json.Unmarshal(msg.Data, &metrics)
		if err != nil {
			log.Fatal(err)
		}

		// Insert event
		clickhouse.InsertGitEvent(connection, metrics)
		log.Println()
	}, nats.Durable("EVENTS_CONSUMER"), nats.ManualAck())

	runtime.Goexit()

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
