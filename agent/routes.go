package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/github", app.githubHandler)
	router.HandlerFunc(http.MethodPost, "/gitlab", app.gitlabHandler)
	router.HandlerFunc(http.MethodPost, "/bitbucket", app.bitBucketHandler)
	router.HandlerFunc(http.MethodPost, "/azure", app.azureHandler)
	router.HandlerFunc(http.MethodPost, "/gitea", app.giteaHandler)
	return router
}
