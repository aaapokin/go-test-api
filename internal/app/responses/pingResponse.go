package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	JSON(c, http.StatusOK, gin.H{"message": "pong"})
}
