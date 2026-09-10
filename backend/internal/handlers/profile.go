package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
	"student-career-platform/internal/models"
)
func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var profile models.Profile

	result := database.DB.Where("user_id = ?", userID).First(&profile)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "profile not found",
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}
func UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var request models.UpdateProfileRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var profile models.Profile

	result := database.DB.Where("user_id = ?", userID).First(&profile)

	if result.Error != nil {
		profile = models.Profile{
			UserID: userID.(uint),
		}
	}

	profile.University = request.University
	profile.Degree = request.Degree
	profile.GraduationYear = request.GraduationYear
	profile.CGPA = request.CGPA
	profile.GitHub = request.GitHub
	profile.LinkedIn = request.LinkedIn
	profile.Portfolio = request.Portfolio

	result = database.DB.Save(&profile)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save profile",
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}