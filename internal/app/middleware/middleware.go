package middleware

import (
	"app/internal/app"
	"app/internal/app/responses"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenValue string
		bearerToken := c.Request.Header.Get("Authorization")
		if len(strings.Split(bearerToken, " ")) == 2 {
			tokenValue = strings.Split(bearerToken, " ")[1]
		}
		token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
			//exp проверка по этому полю зашито в пакете

			// claims := token.Claims.(jwt.MapClaims)
			// exp := claims["exp"].(float64)
			// exp2 := int64(exp)
			// if exp2 < time.Now().Add(time.Second*time.Duration(app.Config.JwtTimeSeconds)).Unix() {
			// 	return nil, errors.New("token expired")
			// }
			return app.Config.JwtKey, nil
		})
		if err != nil || token == nil || !token.Valid {
			responses.Unauthorize(c)
			// c.Abort()
			return
		}
		c.Next()
	}
}
