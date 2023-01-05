package main

import (
	"github.com/gin-gonic/gin"
	ap "github.com/theMisbahnator/Reverbify/controllers"
	"github.com/theMisbahnator/Reverbify/initializers"
)

// import "github.com/gin-gonic/gin"

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()
	r.POST("/reverb-song", ap.Init_audio_processing)
	r.GET("/test", ap.Test)
	r.Run() // listen and serve on 0.0.0.0:8080
}
