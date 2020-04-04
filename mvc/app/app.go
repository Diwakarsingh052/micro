package app

import (
	"github.com/gin-gonic/gin"
)

// this app.go starts the web server

var router *gin.Engine

func init() {
	router = gin.Default()
}

func StartApp() {

	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
