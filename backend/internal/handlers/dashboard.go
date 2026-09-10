package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
	"student-career-platform/internal/models"
)

func GetDashboard(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var stats models.DashboardStats

	database.DB.Model(&models.Application{}).
		Where("user_id = ?", userID).
		Count(&stats.TotalApplications)

	database.DB.Model(&models.Application{}).
		Where("user_id = ? AND status = ?", userID, "Applied").
		Count(&stats.Applied)

	database.DB.Model(&models.Application{}).
		Where("user_id = ? AND status = ?", userID, "Assessment").
		Count(&stats.Assessment)

	database.DB.Model(&models.Application{}).
		Where("user_id = ? AND status = ?", userID, "Interview").
		Count(&stats.Interview)

	database.DB.Model(&models.Application{}).
		Where("user_id = ? AND status = ?", userID, "Offer").
		Count(&stats.Offer)

	database.DB.Model(&models.Application{}).
		Where("user_id = ? AND status = ?", userID, "Rejected").
		Count(&stats.Rejected)

	database.DB.Model(&models.Project{}).
		Where("user_id = ?", userID).
		Count(&stats.Projects)

	database.DB.Model(&models.Skill{}).
		Where("user_id = ?", userID).
		Count(&stats.Skills)

	c.JSON(http.StatusOK, stats)
}