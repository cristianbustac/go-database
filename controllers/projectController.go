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

// get projects
// @Summary get projects
// @Description gets all projects
// @Tags projects
// @Accept json
// @Produce json
// @Success 200
// @Router /projects [get]
// @Param skip query int false "skip" default(0)
// @Param limit query int false "limit" default(10)
func Getprojects(context *gin.Context) {
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
	var project models.Project
	var projects []*models.Project
	projectRepo := repositories.NewProjectRepository(database.DB)
	projectsRes, err := projectRepo.GetAll(&project, skip, limit)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve projects"})
		return
	}
	if err := projectsRes.Find(&projects).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}
	context.JSON(http.StatusOK, projects)
}

// get project
// @Summary get project
// @Description gets project by id
// @Param project_id path int true "Project ID"
// @Tags projects
// @Accept json
// @Produce json
// @Success 200
// @Router /projects/{project_id} [get]
func Getproject(context *gin.Context) {
	var project models.Project
	projectId := context.Param("project_id")
	projectRepo := repositories.NewProjectRepository(database.DB)
	projectEntity, err := projectRepo.GetById(projectId, &project)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve project"})
		return
	}
	context.JSON(http.StatusOK, projectEntity)
}

// post project
// @Summary post project
// @Description create a project
// @Tags projects
// @Accept json
// @Produce json
// @Success 200
// @Param  project  body      schemas.Project  true  "project JSON"
// @Router /projects [post]
func Postproject(context *gin.Context) {
	var project models.Project
	json.NewDecoder(context.Request.Body).Decode(&project)
	projectRepo := repositories.NewProjectRepository(database.DB)
	createproject, err := projectRepo.Create(&project)
	if err != nil {
		context.JSON(http.StatusBadRequest, "error creating project")
	}
	context.JSON(http.StatusOK, createproject)
}

// patch project
// @Summary patch project
// @Description update a project
// @Param project_id path int true "Project ID"
// @Param  project  body      schemas.ProjectUpdate  true  "project JSON"
// @Tags projects
// @Accept json
// @Produce json
// @Success 200
// @Router /projects/{project_id} [patch]
func Patchproject(context *gin.Context) {
	var project models.Project
	projectId := context.Param("project_id")

	var updatedproject schemas.ProjectUpdate
	if err := context.ShouldBindJSON(&updatedproject); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	projectRepo := repositories.NewProjectRepository(database.DB)
	projectRes, err := projectRepo.Update(projectId, &project, &updatedproject)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update project"})
		return
	}
	context.JSON(http.StatusOK, projectRes)
}

// delete project
// @Summary delete project
// @Description delete a project
// @Param project_id path int true "Project ID"
// @Tags projects
// @Accept json
// @Produce json
// @Success 200
// @Router /projects/{project_id} [delete]
func Deleteproject(context *gin.Context) {
	projectId := context.Param("project_id")
	var project models.Project
	projectRepo := repositories.NewProjectRepository(database.DB)
	projectRes := projectRepo.Delete(projectId, &project)
	if projectRes != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve project"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"success": "project successfully deleted from the database", "data": project})
}
