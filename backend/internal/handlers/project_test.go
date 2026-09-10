package handlers

import (
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

func TestCreateProject(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Project Test User",
		Username:     fmt.Sprintf("project-%d", timestamp),
		Email:        fmt.Sprintf("project-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.POST(
		"/projects",
		middleware.AuthRequired(),
		CreateProject,
	)

	requestBody := `{
		"name": "Test Project",
		"description": "A project created during testing",
		"technologies": "Go, React, PostgreSQL",
		"github_url": "https://github.com/test/project",
		"live_url": "https://example.com",
		"status": "Completed"
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/projects",
		strings.NewReader(requestBody),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
	)

	request.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf(
			"expected status 201, got %d: %s",
			recorder.Code,
			recorder.Body.String(),
		)
	}

	var project models.Project

	result = database.DB.
		Where("user_id = ?", user.ID).
		First(&project)

	if result.Error != nil {
		t.Fatalf(
			"failed to find created project: %v",
			result.Error,
		)
	}

	if project.Name != "Test Project" {
		t.Fatalf(
			"expected project name Test Project, got %s",
			project.Name,
		)
	}

	if project.Status != "Completed" {
		t.Fatalf(
			"expected status Completed, got %s",
			project.Status,
		)
	}
}
func TestUpdateProject(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Update Project User",
		Username:     fmt.Sprintf("update-project-%d", timestamp),
		Email:        fmt.Sprintf("update-project-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	project := models.Project{
		UserID:        user.ID,
		Name:          "Old Project Name",
		Description:   "Old description",
		Technologies:  "Go",
		Status:        "In Progress",
	}

	result = database.DB.Create(&project)

	if result.Error != nil {
		t.Fatalf("failed to create test project: %v", result.Error)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.PUT(
		"/projects/:id",
		middleware.AuthRequired(),
		UpdateProject,
	)

	requestBody := `{
		"name": "Updated Project Name",
		"description": "Updated description",
		"technologies": "Go, React, PostgreSQL",
		"github_url": "https://github.com/test/updated-project",
		"live_url": "https://example.com/updated",
		"status": "Completed"
	}`

	request := httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/projects/%d", project.ID),
		strings.NewReader(requestBody),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
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

	var updatedProject models.Project

	result = database.DB.
		First(&updatedProject, project.ID)

	if result.Error != nil {
		t.Fatalf(
			"failed to find updated project: %v",
			result.Error,
		)
	}

	if updatedProject.Name != "Updated Project Name" {
		t.Fatalf(
			"expected updated name, got %s",
			updatedProject.Name,
		)
	}

	if updatedProject.Status != "Completed" {
		t.Fatalf(
			"expected status Completed, got %s",
			updatedProject.Status,
		)
	}
}
func TestDeleteProject(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Delete Project User",
		Username:     fmt.Sprintf("delete-project-%d", timestamp),
		Email:        fmt.Sprintf("delete-project-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	project := models.Project{
		UserID:       user.ID,
		Name:         "Project To Delete",
		Description:  "This project will be deleted",
		Technologies: "Go",
		Status:       "Completed",
	}

	result = database.DB.Create(&project)

	if result.Error != nil {
		t.Fatalf("failed to create test project: %v", result.Error)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.DELETE(
		"/projects/:id",
		middleware.AuthRequired(),
		DeleteProject,
	)

	request := httptest.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/projects/%d", project.ID),
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

	var deletedProject models.Project

	result = database.DB.
		First(&deletedProject, project.ID)

	if result.Error == nil {
		t.Fatal("expected project to be deleted, but it still exists")
	}
}