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
	"student-career-platform/internal/models"
	"student-career-platform/internal/services"
	"student-career-platform/internal/middleware"
)

func TestCreateApplication(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	// Create a unique test user
	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Application Test User",
		Username:     fmt.Sprintf("apptest-%d", timestamp),
		Email:        fmt.Sprintf("apptest-%d@example.com", timestamp),
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
		"/applications",
		middleware.AuthRequired(),
		CreateApplication,
	)

	requestBody := `{
		"company": "Google",
		"job_title": "Software Engineer Intern",
		"job_url": "https://example.com/job",
		"status": "Applied",
		"notes": "Application test"
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/applications",
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

	// Verify that the application was actually saved
	var application models.Application

	result = database.DB.
		Where("user_id = ?", user.ID).
		First(&application)

	if result.Error != nil {
		t.Fatalf(
			"failed to find created application: %v",
			result.Error,
		)
	}

	if application.Company != "Google" {
		t.Fatalf(
			"expected company Google, got %s",
			application.Company,
		)
	}

	if application.Status != "Applied" {
		t.Fatalf(
			"expected status Applied, got %s",
			application.Status,
		)
	}
}
func TestUpdateApplicationInvalidTransition(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Transition Test User",
		Username:     fmt.Sprintf("transition-%d", timestamp),
		Email:        fmt.Sprintf("transition-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	application := models.Application{
		UserID:  user.ID,
		Company: "Test Company",
		JobTitle: "Test Intern",
		Status:  "Assessment",
	}

	result = database.DB.Create(&application)

	if result.Error != nil {
		t.Fatalf(
			"failed to create test application: %v",
			result.Error,
		)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.PUT(
		"/applications/:id",
		middleware.AuthRequired(),
		UpdateApplication,
	)

	requestBody := `{
		"company": "Test Company",
		"job_title": "Test Intern",
		"status": "Offer"
	}`

	request := httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/applications/%d", application.ID),
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

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status 400, got %d: %s",
			recorder.Code,
			recorder.Body.String(),
		)
	}
}
func TestUpdateApplicationValidTransition(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Valid Transition User",
		Username:     fmt.Sprintf("valid-transition-%d", timestamp),
		Email:        fmt.Sprintf("valid-transition-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	application := models.Application{
		UserID:   user.ID,
		Company:  "Test Company",
		JobTitle: "Software Engineer Intern",
		Status:   "Applied",
	}

	result = database.DB.Create(&application)

	if result.Error != nil {
		t.Fatalf(
			"failed to create test application: %v",
			result.Error,
		)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.PUT(
		"/applications/:id",
		middleware.AuthRequired(),
		UpdateApplication,
	)

	requestBody := `{
		"company": "Test Company",
		"job_title": "Software Engineer Intern",
		"status": "Assessment"
	}`

	request := httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/applications/%d", application.ID),
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

	var updatedApplication models.Application

	result = database.DB.
		First(&updatedApplication, application.ID)

	if result.Error != nil {
		t.Fatalf(
			"failed to find updated application: %v",
			result.Error,
		)
	}

	if updatedApplication.Status != "Assessment" {
		t.Fatalf(
			"expected status Assessment, got %s",
			updatedApplication.Status,
		)
	}
}