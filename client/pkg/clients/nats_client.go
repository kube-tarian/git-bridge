package clients

import (
	"encoding/json"
	"log"
	"time"

	"github.com/kube-tarian/git-bridge/client/pkg/clickhouse"
	"github.com/kube-tarian/git-bridge/client/pkg/config"

	"github.com/nats-io/nats.go"
)

//constant variables to use with nats stream and
// nats publishing
const (
	streamSubjects string = "GITMETRICS.*"
	eventSubject   string = "GITMETRICS.git"
	eventConsumer  string = "Git-Consumer"
)

type NATSContext struct {
	conf     *config.Config
	conn     *nats.Conn
	stream   nats.JetStreamContext
	dbClient *clickhouse.DBClient
}

func NewNATSContext(conf *config.Config, dbClient *clickhouse.DBClient) (*NATSContext, error) {
	log.Println("Waiting before connecting to NATS at:", conf.NatsAddress)
	time.Sleep(1 * time.Second)

	conn, err := nats.Connect(conf.NatsAddress, nats.Name("Github metrics"), nats.Token(conf.NatsToken))
	if err != nil {
		return nil, err
	}

	ctx := &NATSContext{
		conf:     conf,
		conn:     conn,
		dbClient: dbClient,
	}

	stream, err := ctx.CreateStream()
	if err != nil {
		ctx.conn.Close()
		return nil, err
	}

	ctx.stream = stream
	ctx.Subscribe(eventSubject, eventConsumer, dbClient)

	_, err = stream.StreamInfo("GITMETRICS")
	if err != nil {
		return nil, err
	}

	return ctx, nil
}

func (n *NATSContext) CreateStream() (nats.JetStreamContext, error) {
	// Creates JetStreamContext
	stream, err := n.conn.JetStream()
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (n *NATSContext) Close() {
	n.conn.Close()
}

type pubData struct {
	Metrics json.RawMessage `json:"metrics"`
	Repo    string          `json:"repo"`
}

func (n *NATSContext) Subscribe(subject string, consumer string, conn *clickhouse.DBClient) {
	n.stream.Subscribe(subject, func(msg *nats.Msg) {
		msg.Ack()
		repo := msg.Header.Get("repo")
		metrics := &pubData{
			Metrics: json.RawMessage(msg.Data),
			Repo:    repo,
		}
		// metrics := &models.Gitevent{}
		data, err := json.Marshal(metrics)
		if err != nil {
			log.Fatal(err)

		}

		conn.InsertEvent(string(data))
		log.Println("Inserted metrics:", string(data))

	}, nats.Durable(consumer), nats.ManualAck())

}
