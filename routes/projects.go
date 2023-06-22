package routes

import (
	controller "go-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ProjectsRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/projects", controller.Getprojects)
	incomingRoutes.GET("/projects/:project_id", controller.Getproject)
	incomingRoutes.POST("/projects", controller.Postproject)
	incomingRoutes.PATCH("/projects/:project_id", controller.Patchproject)
	incomingRoutes.DELETE("/projects/:project_id", controller.Deleteproject)
}
