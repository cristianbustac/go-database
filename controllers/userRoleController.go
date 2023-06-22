package controller

import (
	"go-backend/database"
	"go-backend/database/repositories"
	"go-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get relation
// @Summary get relation users roles
// @Description gets all users roles
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Router /users/roles [get]
// @Param skip query int false "skip" default(0)
// @Param limit query int false "limit" default(10)
func GetRelationsUsersRoles(context *gin.Context) {
	var user models.User
	var users []*models.User
	skip, err := strconv.Atoi(context.DefaultQuery("skip", "0"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for skip"})
		return
	}

	limit, err := strconv.Atoi(context.DefaultQuery("limit", "10"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid value for limit"})
		return
	}
	relaRepo := repositories.NewUsersRolesRepository(database.DB)
	relaRes, err := relaRepo.GetAllRelations(&user, skip, limit, "Roles")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	relaRes.Preload("Roles").Find(&users)
	context.JSON(http.StatusOK, users)
}

// post relation user role
// @Summary post relation user role
// @Description post relation user role
// @Param role_id path int true "Role ID"
// @Param user_id path int true "User ID"
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{user_id}/roles/{role_id} [post]
func CreateRelationUserRole(context *gin.Context) {
	var role models.Role
	var user models.User
	roleId := context.Param("role_id")
	userId := context.Param("user_id")
	relaRepo := repositories.NewUsersRolesRepository(database.DB)
	err := relaRepo.CreateRelation(userId, roleId, &user, &role, "Roles")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve role"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": "relation successfully created"})
}

// delete relation user role
// @Summary delete relation user role
// @Description delete relation user role
// @Param role_id path int true "Role ID"
// @Param user_id path int true "User ID"
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{user_id}/roles/{role_id} [delete]
func DeleteRelationUserRole(context *gin.Context) {
	var role models.Role
	var user models.User
	roleId := context.Param("role_id")
	userId := context.Param("user_id")

	relaRepo := repositories.NewUsersRolesRepository(database.DB)
	relaRes := relaRepo.DeleteRelation(userId, roleId, &user, &role, "Roles")
	if relaRes != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve role"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": "relation successfully deleted from the database"})
}
