package application

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func (app *Application) giteaHandler(c *gin.Context) {
	repo := "Gitea"
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error Reading Request Body")
	}
	app.conn.Publish(jsonData, repo)
}

func (app *Application) azureHandler(c *gin.Context) {
	repo := "Azure"
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error Reading Request Body")
	}
	app.conn.Publish(jsonData, repo)
}

//githubHandler handles the github webhooks post requests.
func (app *Application) githubHandler(c *gin.Context) {
	repo := "Github"
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error Reading Request Body")
	}
	app.conn.Publish(jsonData, repo)
}

//gitlabHandler handles the github webhooks post requests.
func (app *Application) gitlabHandler(c *gin.Context) {
	repo := "Gitlab"
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error Reading Request Body")
	}
	app.conn.Publish(jsonData, repo)
}

//bitBucketHandler handles the github webhooks post requests.
func (app *Application) bitBucketHandler(c *gin.Context) {
	repo := "BitBucket"
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Error Reading Request Body")
	}
	app.conn.Publish(jsonData, repo)
}
