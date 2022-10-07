package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthz() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}
}
