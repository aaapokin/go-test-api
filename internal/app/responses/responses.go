package responses

import "github.com/gin-gonic/gin"

type message struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func JSON(c *gin.Context, stataus_code int, data interface{}) {
	var message message
	message.Code = stataus_code
	message.Data = data
	c.JSON(stataus_code, message)
	c.Abort()
}
