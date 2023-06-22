package routes

import (
	controller "go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RolesRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/roles", controller.Getroles)
	incomingRoutes.GET("/roles/:role_id", controller.Getrole)
	incomingRoutes.POST("/roles", controller.Postrole)
	incomingRoutes.PATCH("/roles/:role_id", controller.Patchrole)
	incomingRoutes.DELETE("/roles/:role_id", controller.Deleterole)
}
