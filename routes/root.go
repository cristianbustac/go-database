package routes

import (
	controller "go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func HealthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/health", controller.HealthCheck)
}
