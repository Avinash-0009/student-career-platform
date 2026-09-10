package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
	"student-career-platform/internal/models"
)
func CreateApplication(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var request models.CreateApplicationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	status := models.ApplicationStatus(request.Status)

    if !models.IsValidStatus(status) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": "invalid application status",
	})
	return
    }

	application := models.Application{
		UserID:      userID.(uint),
		Company:     request.Company,
		JobTitle:    request.JobTitle,
		JobURL:      request.JobURL,
		Status:      string(status),
		AppliedDate: request.AppliedDate,
		Deadline:    request.Deadline,
		Notes:       request.Notes,
	}

	result := database.DB.Create(&application)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create application",
		})
		return
	}

	c.JSON(http.StatusCreated, application)
}
func GetApplications(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var applications []models.Application

	result := database.DB.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&applications)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch applications",
		})
		return
	}

	c.JSON(http.StatusOK, applications)
}
func GetApplication(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	applicationID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid application id",
		})
		return
	}

	var application models.Application

	result := database.DB.
		Where("id = ? AND user_id = ?", applicationID, userID).
		First(&application)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "application not found",
		})
		return
	}

	c.JSON(http.StatusOK, application)
}
func UpdateApplication(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	applicationID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid application id",
		})
		return
	}

	var application models.Application

	result := database.DB.
		Where("id = ? AND user_id = ?", applicationID, userID).
		First(&application)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "application not found",
		})
		return
	}

	var request models.UpdateApplicationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newStatus := models.ApplicationStatus(request.Status)

if !models.IsValidStatus(newStatus) {
    c.JSON(http.StatusBadRequest, gin.H{
        "error": "invalid application status",
    })
    return
}

currentStatus := models.ApplicationStatus(application.Status)

if !models.CanTransition(currentStatus, newStatus) {
    c.JSON(http.StatusBadRequest, gin.H{
        "error": "invalid status transition",
    })
    return
}

	application.Company = request.Company
	application.JobTitle = request.JobTitle
	application.JobURL = request.JobURL
	application.Status = string(newStatus)
	application.AppliedDate = request.AppliedDate
	application.Deadline = request.Deadline
	application.Notes = request.Notes

	if err := database.DB.Save(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update application",
		})
		return
	}

	c.JSON(http.StatusOK, application)
}
func DeleteApplication(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	applicationID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid application id",
		})
		return
	}

	result := database.DB.
		Where("id = ? AND user_id = ?", applicationID, userID).
		Delete(&models.Application{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete application",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "application not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "application deleted successfully",
	})
}