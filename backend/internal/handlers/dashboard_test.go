package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
	"student-career-platform/internal/middleware"
	"student-career-platform/internal/models"
	"student-career-platform/internal/services"
)

func TestGetDashboard(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Dashboard Test User",
		Username:     fmt.Sprintf("dashboard-%d", timestamp),
		Email:        fmt.Sprintf("dashboard-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	// Create applications
	applications := []models.Application{
		{
			UserID:   user.ID,
			Company:  "Google",
			JobTitle: "Software Engineer Intern",
			Status:   "Applied",
		},
		{
			UserID:   user.ID,
			Company:  "Microsoft",
			JobTitle: "Software Engineer Intern",
			Status:   "Interview",
		},
		{
			UserID:   user.ID,
			Company:  "Amazon",
			JobTitle: "Software Engineer Intern",
			Status:   "Rejected",
		},
	}

	result = database.DB.Create(&applications)

	if result.Error != nil {
		t.Fatalf(
			"failed to create test applications: %v",
			result.Error,
		)
	}

	// Create projects
	project := models.Project{
		UserID:       user.ID,
		Name:         "Dashboard Test Project",
		Description:  "Test project",
		Technologies: "Go, React",
		Status:       "Completed",
	}

	result = database.DB.Create(&project)

	if result.Error != nil {
		t.Fatalf(
			"failed to create test project: %v",
			result.Error,
		)
	}

	// Create skills
	skill := models.Skill{
		UserID:      user.ID,
		Name:        "Go",
		Proficiency: 80,
	}

	result = database.DB.Create(&skill)

	if result.Error != nil {
		t.Fatalf(
			"failed to create test skill: %v",
			result.Error,
		)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.GET(
		"/dashboard",
		middleware.AuthRequired(),
		GetDashboard,
	)

	request := httptest.NewRequest(
		http.MethodGet,
		"/dashboard",
		nil,
	)

	request.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d: %s",
			recorder.Code,
			recorder.Body.String(),
		)
	}

	var stats models.DashboardStats

	err = json.Unmarshal(
		recorder.Body.Bytes(),
		&stats,
	)

	if err != nil {
		t.Fatalf(
			"failed to parse dashboard response: %v",
			err,
		)
	}

	if stats.TotalApplications != 3 {
		t.Fatalf(
			"expected 3 applications, got %d",
			stats.TotalApplications,
		)
	}

	if stats.Applied != 1 {
		t.Fatalf(
			"expected 1 applied application, got %d",
			stats.Applied,
		)
	}

	if stats.Interview != 1 {
		t.Fatalf(
			"expected 1 interview application, got %d",
			stats.Interview,
		)
	}

	if stats.Rejected != 1 {
		t.Fatalf(
			"expected 1 rejected application, got %d",
			stats.Rejected,
		)
	}

	if stats.Projects != 1 {
		t.Fatalf(
			"expected 1 project, got %d",
			stats.Projects,
		)
	}

	if stats.Skills != 1 {
		t.Fatalf(
			"expected 1 skill, got %d",
			stats.Skills,
		)
	}

	if !strings.Contains(
		recorder.Body.String(),
		`"total_applications":3`,
	) {
		t.Fatalf(
			"expected dashboard JSON response, got %s",
			recorder.Body.String(),
		)
	}
}