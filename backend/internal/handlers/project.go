package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
	"student-career-platform/internal/models"
)
func CreateProject(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var request models.CreateProjectRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	project := models.Project{
		UserID:       userID.(uint),
		Name:         request.Name,
		Description:  request.Description,
		Technologies: request.Technologies,
		GitHubURL:    request.GitHubURL,
		LiveURL:      request.LiveURL,
		Status:       request.Status,
	}

	result := database.DB.Create(&project)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create project",
		})
		return
	}

	c.JSON(http.StatusCreated, project)
}
func GetProjects(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var projects []models.Project

	result := database.DB.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&projects)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch projects",
		})
		return
	}

	c.JSON(http.StatusOK, projects)
}
func GetProject(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	projectID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid project id",
		})
		return
	}

	var project models.Project

	result := database.DB.
		Where("id = ? AND user_id = ?", projectID, userID).
		First(&project)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "project not found",
		})
		return
	}

	c.JSON(http.StatusOK, project)
}
func UpdateProject(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	projectID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid project id",
		})
		return
	}

	var project models.Project

	result := database.DB.
		Where("id = ? AND user_id = ?", projectID, userID).
		First(&project)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "project not found",
		})
		return
	}

	var request models.UpdateProjectRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	project.Name = request.Name
	project.Description = request.Description
	project.Technologies = request.Technologies
	project.GitHubURL = request.GitHubURL
	project.LiveURL = request.LiveURL
	project.Status = request.Status

	if err := database.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update project",
		})
		return
	}

	c.JSON(http.StatusOK, project)
}
func DeleteProject(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	projectID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid project id",
		})
		return
	}

	result := database.DB.
		Where("id = ? AND user_id = ?", projectID, userID).
		Delete(&models.Project{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete project",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "project not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "project deleted successfully",
	})
}