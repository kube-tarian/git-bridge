package application

import (
	"context"
	"errors"
	"fmt"
	"github.com/kube-tarian/git-bridge/client/pkg/clickhouse"
	"github.com/kube-tarian/git-bridge/client/pkg/clients"
	"github.com/kube-tarian/git-bridge/client/pkg/config"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Application struct {
	Config   *config.Config
	server   *http.Server
	conn     *clients.NATSContext
	dbClient *clickhouse.DBClient
}

func New(conf *config.Config, conn *clients.NATSContext, dbClient *clickhouse.DBClient) *Application {
	log.Println("Initializing Application")
	app := &Application{
		Config:   conf,
		conn:     conn,
		dbClient: dbClient,
	}

	app.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return app
}

func (app *Application) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.status)
	return router
}

func (app *Application) status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Client is working"))
}

func (app *Application) Start() {
	log.Println("Starting server on port", app.Config.Port)
	if err := app.server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server closed, readon: %v", err)
	}
}

func (app *Application) Close() {
	log.Printf("Closing the service gracefully")
	app.conn.Close()

	if err := app.server.Shutdown(context.Background()); err != nil {
		log.Printf("Could not close the service gracefully: %v", err)
	}
}
