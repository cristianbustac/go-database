package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// get roles
// @Summary get roles
// @Description gets all roles
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Router /roles [get]
func Getroles(context *gin.Context) {
	context.JSON(http.StatusOK, "Get roles")
}

// get role
// @Summary get role
// @Description gets role by id
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Router /roles/{role_id} [get]
func Getrole(context *gin.Context) {
	context.JSON(http.StatusOK, "get role")
}

// post role
// @Summary post role
// @Description create a role
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Router /roles [post]
func Postrole(context *gin.Context) {
	context.JSON(http.StatusOK, "Post role")
}

// patch role
// @Summary patch role
// @Description update a role
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Router /roles/{role_id} [patch]
func Patchrole(context *gin.Context) {
	context.JSON(http.StatusOK, "patch role")
}

// delete role
// @Summary delete role
// @Description delete a role
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Router /roles/{role_id} [delete]
func Deleterole(context *gin.Context) {
	context.JSON(http.StatusOK, "delete role")
}
