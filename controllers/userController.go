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

// get users
// @Summary get users
// @Description gets all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Router /users [get]
// @Param skip query int false "skip" default(0)
// @Param limit query int false "limit" default(10)
func GetUsers(context *gin.Context) {
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
	var user models.User
	var users []models.User
	userRepo := repositories.NewUserRepository(database.DB)
	usersRes, err := userRepo.GetAll(&user, skip, limit)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}
	if err := usersRes.Find(&users).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"users": users})
}

// get user
// @Summary get user
// @Description gets user by id
// @Param user_id path int true "User ID"
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{user_id} [get]
func GetUser(context *gin.Context) {
	var user models.User
	userId := context.Param("user_id")
	userRepo := repositories.NewUserRepository(database.DB)
	userEntity, err := userRepo.GetById(userId, &user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return
	}
	context.JSON(http.StatusOK, userEntity)
}

// post user
// @Summary post user
// @Description create a user
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Param  user  body      schemas.User  true  "user JSON"
// @Router /users [post]
func PostUser(context *gin.Context) {
	var user models.User
	json.NewDecoder(context.Request.Body).Decode(&user)
	userRepo := repositories.NewUserRepository(database.DB)
	createUser, err := userRepo.Create(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, "error creating user")
	}
	context.JSON(http.StatusOK, createUser)
}

// patch user
// @Summary patch user
// @Description update a user
// @Param user_id path int true "User ID"
// @Param  user  body      schemas.UserUpdate  true  "user JSON"
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{user_id} [patch]
func PatchUser(context *gin.Context) {
	var user models.User
	userId := context.Param("user_id")

	var updatedUser schemas.UserUpdate
	if err := context.ShouldBindJSON(&updatedUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userRepo := repositories.NewUserRepository(database.DB)
	userRes, err := userRepo.Update(userId, &user, &updatedUser)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}
	context.JSON(http.StatusOK, userRes)
}

// delete user
// @Summary delete user
// @Description delete a user
// @Param user_id path int true "User ID"
// @Tags users
// @Accept json
// @Produce json
// @Success 200
// @Router /users/{user_id} [delete]
func DeleteUser(context *gin.Context) {
	userId := context.Param("user_id")
	var user models.User
	userRepo := repositories.NewUserRepository(database.DB)
	userRes := userRepo.Delete(userId, &user)
	if userRes != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": "user successfully deleted from the database", "data": user})
}
