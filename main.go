package main

import (
	"net/http"
	"os"

	"go-backend/database"
	"go-backend/models"
	"go-backend/routes"

	docs "go-backend/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db, error := database.DBinstance()
	if error == nil {
		db.AutoMigrate(models.User{})
		db.AutoMigrate(models.Project{})
		db.AutoMigrate(models.Role{})
	}

	docs.SwaggerInfo.Title = os.Getenv("TITLE")
	docs.SwaggerInfo.Description = os.Getenv("DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("VERSION")

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.HealthRoutes(router)
	routes.UserRoutes(router)

	http.ListenAndServe(":8000", router)
}
