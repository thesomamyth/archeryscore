package main

import (
	"archeryscore/endpoints"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthz", endpoints.GetHealthz)
	r.Run()
}
