package handlers

import (
	"app/internal/app/responses"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	responses.Ping(c)
}
