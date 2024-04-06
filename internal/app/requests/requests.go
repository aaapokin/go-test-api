package requests

import (
	"app/internal/app/models"
	"app/internal/app/responses"

	"github.com/gin-gonic/gin"
)

type RequestInterface interface {
	validate() error
	GetDTO() interface{}
}

func badResponse(context *gin.Context, message *models.Message) {
	responses.BadRequestByMessage(context, message)
}

func badResponseByError(context *gin.Context, err error) {
	message := &models.Message{Error: err.Error()}
	badResponse(context, message)
}
