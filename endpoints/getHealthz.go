package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
