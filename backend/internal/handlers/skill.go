package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
	"student-career-platform/internal/models"
)
func CreateSkill(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var request models.CreateSkillRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	skill := models.Skill{
		UserID:      userID.(uint),
		Name:        request.Name,
		Proficiency: request.Proficiency,
	}

	result := database.DB.Create(&skill)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create skill",
		})
		return
	}

	c.JSON(http.StatusCreated, skill)
}
func GetSkills(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	var skills []models.Skill

	result := database.DB.
		Where("user_id = ?", userID).
		Order("proficiency DESC").
		Find(&skills)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch skills",
		})
		return
	}

	c.JSON(http.StatusOK, skills)
}
func UpdateSkill(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	skillID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid skill id",
		})
		return
	}

	var skill models.Skill

	result := database.DB.
		Where("id = ? AND user_id = ?", skillID, userID).
		First(&skill)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "skill not found",
		})
		return
	}

	var request models.UpdateSkillRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	skill.Name = request.Name
	skill.Proficiency = request.Proficiency

	if err := database.DB.Save(&skill).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to update skill",
		})
		return
	}

	c.JSON(http.StatusOK, skill)
}
func DeleteSkill(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not authenticated",
		})
		return
	}

	skillID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid skill id",
		})
		return
	}

	result := database.DB.
		Where("id = ? AND user_id = ?", skillID, userID).
		Delete(&models.Skill{})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete skill",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "skill not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "skill deleted successfully",
	})
}