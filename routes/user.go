package routes

import (
	controller "go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers)
	incomingRoutes.GET("/users/:user_id", controller.GetUser)
	incomingRoutes.POST("/users", controller.PostUser)
	incomingRoutes.PATCH("/users/:user_id", controller.PatchUser)
	incomingRoutes.DELETE("/users/:user_id", controller.DeleteUser)
}
