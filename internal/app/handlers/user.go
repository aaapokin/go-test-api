package handlers

import (
	"app/internal/app"
	"app/internal/app/models"
	"app/internal/app/requests"
	"app/internal/app/responses"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	model := &models.User{}
	request := requests.NewUserLoginRequest(c, model)

	userInDB, error := request.GetDTO()
	if error != nil {
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) // Дополнительные действия (в формате мапы) для шифрования
	claims["exp"] = time.Now().Add(time.Second * time.Duration(app.Config.JwtTimeSeconds)).Unix()

	claims["admin"] = true
	claims["login"] = userInDB.Login
	tokenString, err := token.SignedString(app.Config.JwtKey)
	if err != nil {
		responses.BadRequestByError(c, err)
		return
	}
	responses.Jwt(c, tokenString)
}
