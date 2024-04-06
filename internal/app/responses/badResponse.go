package responses

import (
	"app/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestByMessage(c *gin.Context, message *models.Message) {
	JSON(c, http.StatusBadRequest, message)
}

func BadRequestByError(c *gin.Context, error error) {
	JSON(c, http.StatusBadRequest, models.Message{Error: error.Error()})
}

func Unauthorize(c *gin.Context) {
	JSON(c, http.StatusUnauthorized, models.Message{Error: "Invalid Token"})
}
