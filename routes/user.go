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
	incomingRoutes.GET("/users/roles", controller.GetRelationsUsersRoles)
	incomingRoutes.POST("/users/:user_id/roles/:role_id", controller.CreateRelationUserRole)
	incomingRoutes.DELETE("/users/:user_id/roles/:role_id", controller.DeleteRelationUserRole)
}
