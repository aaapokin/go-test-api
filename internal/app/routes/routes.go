package routes

import (
	"app/internal/app/handlers"
	"app/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetupRouter() *gin.Engine {
	router = gin.Default()
	apiV1Group := router.Group("/api/v1")
	{
		apiV1Group.GET("ping", handlers.Ping)
		apiV1Group.GET("ping2", middleware.JwtMiddleware(), handlers.Ping)
		apiV1Group.POST("auth", handlers.Auth)
	}
	return router
}
