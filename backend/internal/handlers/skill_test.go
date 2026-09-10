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

func TestCreateSkill(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Skill Test User",
		Username:     fmt.Sprintf("skill-%d", timestamp),
		Email:        fmt.Sprintf("skill-%d@example.com", timestamp),
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
		"/skills",
		middleware.AuthRequired(),
		CreateSkill,
	)

	requestBody := `{
		"name": "Go",
		"proficiency": 80
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/skills",
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

	var skill models.Skill

	result = database.DB.
		Where("user_id = ?", user.ID).
		First(&skill)

	if result.Error != nil {
		t.Fatalf(
			"failed to find created skill: %v",
			result.Error,
		)
	}

	if skill.Name != "Go" {
		t.Fatalf(
			"expected skill Go, got %s",
			skill.Name,
		)
	}

	if skill.Proficiency != 80 {
		t.Fatalf(
			"expected proficiency 80, got %d",
			skill.Proficiency,
		)
	}
}
func TestUpdateSkill(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Update Skill User",
		Username:     fmt.Sprintf("update-skill-%d", timestamp),
		Email:        fmt.Sprintf("update-skill-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	skill := models.Skill{
		UserID:      user.ID,
		Name:        "Go",
		Proficiency: 60,
	}

	result = database.DB.Create(&skill)

	if result.Error != nil {
		t.Fatalf("failed to create test skill: %v", result.Error)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.PUT(
		"/skills/:id",
		middleware.AuthRequired(),
		UpdateSkill,
	)

	requestBody := `{
		"name": "Go",
		"proficiency": 90
	}`

	request := httptest.NewRequest(
		http.MethodPut,
		fmt.Sprintf("/skills/%d", skill.ID),
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

	var updatedSkill models.Skill

	result = database.DB.
		First(&updatedSkill, skill.ID)

	if result.Error != nil {
		t.Fatalf(
			"failed to find updated skill: %v",
			result.Error,
		)
	}

	if updatedSkill.Proficiency != 90 {
		t.Fatalf(
			"expected proficiency 90, got %d",
			updatedSkill.Proficiency,
		)
	}
}
func TestDeleteSkill(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Delete Skill User",
		Username:     fmt.Sprintf("delete-skill-%d", timestamp),
		Email:        fmt.Sprintf("delete-skill-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	skill := models.Skill{
		UserID:      user.ID,
		Name:        "Docker",
		Proficiency: 75,
	}

	result = database.DB.Create(&skill)

	if result.Error != nil {
		t.Fatalf("failed to create test skill: %v", result.Error)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.DELETE(
		"/skills/:id",
		middleware.AuthRequired(),
		DeleteSkill,
	)

	request := httptest.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/skills/%d", skill.ID),
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

	var deletedSkill models.Skill

	result = database.DB.
		First(&deletedSkill, skill.ID)

	if result.Error == nil {
		t.Fatal("expected skill to be deleted, but it still exists")
	}
}