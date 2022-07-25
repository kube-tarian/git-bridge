package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	router := gin.Default()
	router.POST("/gitea", app.giteaHandler)
	router.POST("/azure", app.azureHandler)
	router.POST("/bitbucket", app.bitBucketHandler)
	router.POST("/gitlab", app.gitlabHandler)
	router.POST("/github", app.githubHandler)
	return router
}
