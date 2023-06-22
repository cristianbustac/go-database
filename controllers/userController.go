package controller

import (
	"encoding/json"
	"go-backend/database"
	"go-backend/models"
	"go-backend/schemas"
	"net/http"

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
func GetUsers(context *gin.Context) {
	var users []*models.User
	users_res := database.DB.Find(&users).Offset(0).Limit(10)
	if users_res.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}
	context.JSON(http.StatusOK, users)
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
	userId := context.Param("user_id")
	var user []*models.User
	user_res := database.DB.First(&user, userId)
	if user_res.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return
	}
	context.JSON(http.StatusOK, user)
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
	createUser := database.DB.Create(&user)
	err := createUser.Error
	if err != nil {
		context.JSON(http.StatusBadRequest, "error creating user")
	}
	context.JSON(http.StatusOK, &user)
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
	userDb := database.DB.First(&user, userId)
	if userDb.Error != nil {
		context.JSON(http.StatusBadRequest, "error getting the user")
	}
	if err := context.ShouldBindJSON(&updatedUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_res := database.DB.Model(&user).Updates(&updatedUser)
	if user_res.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}
	database.DB.First(&user, userId)
	context.JSON(http.StatusOK, user)
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
	var user []*models.User
	user_res := database.DB.First(&user, userId)
	if user_res.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user"})
		return
	}
	database.DB.Delete(&user)
	context.JSON(http.StatusOK, gin.H{"success": "user successfully deleted from the database", "data": user})
}
