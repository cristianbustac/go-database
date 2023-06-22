package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppHealth struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Version     string `json:"Version"`
}

// PingExample godoc
// @Summary health check
// @Description do ping
// @Tags health check
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /health [get]
func HealthCheck(context *gin.Context) {
	data := AppHealth{
		Title:       "database golang",
		Description: "this is a test database by cristian bustamante",
		Version:     "0.0.0",
	}

	context.JSON(http.StatusOK, data)
}
