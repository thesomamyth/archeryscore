package main

import (
	"archeryscore/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", endpoints.GetHealthz())
	r.GET("/healthz", endpoints.GetHealthz())
	r.GET("/scores", endpoints.GetScores())
	r.Run()
}
