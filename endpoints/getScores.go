package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetScores() gin.HandlerFunc {
	scores := []string{"10", "9", "9"}
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"scores": scores,
		})
	}
}
