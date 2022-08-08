package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kube-tarian/git-bridge/agent/pkg/clients"
	"github.com/kube-tarian/git-bridge/agent/pkg/config"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Config *config.Config
	server *http.Server
	conn   *clients.NATSContext
}

func New(conf *config.Config, conn *clients.NATSContext) *Application {
	app := &Application{
		Config: conf,
		conn:   conn,
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

func (app *Application) Routes() *gin.Engine {
	router := gin.New()

	router.POST("/github", app.githubHandler)
	router.POST("/gitlab", app.gitlabHandler)
	router.POST("/bitbucket", app.bitBucketHandler)
	router.POST("/azure", app.azureHandler)
	router.POST("/gitea", app.giteaHandler)
	return router
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
