package controller

import (
	"encoding/json"
	"go-backend/database"
	"go-backend/database/repositories"
	"go-backend/models"
	"go-backend/schemas"
	"net/http"
	"strconv"

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
// @Param skip query int false "skip" default(0)
// @Param limit query int false "limit" default(10)
func Getroles(context *gin.Context) {
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
	var role models.Role
	var roles []*models.Role
	roleRepo := repositories.NewRoleRepository(database.DB)
	rolesRes, err := roleRepo.GetAll(&role, skip, limit)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve roles"})
		return
	}
	if err := rolesRes.Find(&roles).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}
	context.JSON(http.StatusOK, roles)
}

// get role
// @Summary get role
// @Description gets role by id
// @Param role_id path int true "Role ID"
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Router /roles/{role_id} [get]
func Getrole(context *gin.Context) {
	var role models.Role
	roleId := context.Param("role_id")
	roleRepo := repositories.NewRoleRepository(database.DB)
	roleEntity, err := roleRepo.GetById(roleId, &role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve role"})
		return
	}
	context.JSON(http.StatusOK, roleEntity)
}

// post role
// @Summary post role
// @Description create a role
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Param  role  body      schemas.Role  true  "role JSON"
// @Router /roles [post]
func Postrole(context *gin.Context) {
	var role models.Role
	json.NewDecoder(context.Request.Body).Decode(&role)
	roleRepo := repositories.NewRoleRepository(database.DB)
	createrole, err := roleRepo.Create(&role)
	if err != nil {
		context.JSON(http.StatusBadRequest, "error creating role")
	}
	context.JSON(http.StatusOK, createrole)
}

// patch role
// @Summary patch role
// @Description update a role
// @Param role_id path int true "Role ID"
// @Param  role  body      schemas.RoleUpdate  true  "role JSON"
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Router /roles/{role_id} [patch]
func Patchrole(context *gin.Context) {
	var role models.Role
	roleId := context.Param("role_id")

	var updatedrole schemas.RoleUpdate
	if err := context.ShouldBindJSON(&updatedrole); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	roleRepo := repositories.NewRoleRepository(database.DB)
	roleRes, err := roleRepo.Update(roleId, &role, &updatedrole)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update role"})
		return
	}
	context.JSON(http.StatusOK, roleRes)
}

// delete role
// @Summary delete role
// @Description delete a role
// @Param role_id path int true "Role ID"
// @Tags roles
// @Accept json
// @Produce json
// @Success 200
// @Router /roles/{role_id} [delete]
func Deleterole(context *gin.Context) {
	roleId := context.Param("role_id")
	var role models.Role
	roleRepo := repositories.NewRoleRepository(database.DB)
	roleRes := roleRepo.Delete(roleId, &role)
	if roleRes != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve role"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": "role successfully deleted from the database", "data": role})
}
