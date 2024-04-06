package responses

import (
	"app/internal/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Jwt(c *gin.Context, token string) {
	tokenModel := &models.Token{Token: token}
	JSON(c, http.StatusOK, tokenModel)
}
