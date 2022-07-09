package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kube-tarian/git-bridge/publish"

	"github.com/nats-io/nats.go"
)

var (
	token   string = os.Getenv("NATS_TOKEN")
	natsurl string = os.Getenv("NATS_ADDRESS")
)

//constant variables to use with nats stream and
// nats publishing
const (
	streamName     = "GITMETRICS"
	streamSubjects = "GITMETRICS.*"
	eventSubject   = "GITMETRICS.git"
)

// var token string = "UfmrJOYwYCCsgQvxvcfJ3BdI6c8WBbnD"
// var natsurl string = "nats://localhost:4222"

//config will have the configuration details
type config struct {
	port      int
	nats      string
	natstoken string
}

type application struct {
	config  config
	publish publish.Models
}

func main() {
	cfg := config{
		port:      5001,
		nats:      natsurl,
		natstoken: token,
	}

	js := openJS(cfg)
	app := &application{
		config:  cfg,
		publish: publish.NewModels(js, eventSubject),
	}
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("Starting server on port", cfg.port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}
func openJS(cfg config) nats.JetStreamContext {

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
	// Creates JetStreamContext
	js, err := sc.JetStream()
	if err != nil {
		log.Fatal(err)
	}
	// Creates stream
	err = createStream(js)
	if err != nil {
		log.Fatal(err)
	}
	return js

}

// createStream creates a stream by using JetStreamContext
func createStream(js nats.JetStreamContext) error {
	// Check if the METRICS stream already exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	log.Printf("Retrieved stream %s", fmt.Sprintf("%v", stream))
	if err != nil {
		log.Printf("Error getting stream %s", err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
