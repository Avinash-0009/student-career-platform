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

func TestUpdateProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Profile Test User",
		Username:     fmt.Sprintf("profile-%d", timestamp),
		Email:        fmt.Sprintf("profile-%d@example.com", timestamp),
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

	router.PUT(
		"/profile",
		middleware.AuthRequired(),
		UpdateProfile,
	)

	requestBody := `{
		"university": "GD Goenka University",
		"degree": "B.Tech Computer Science",
		"graduation_year": 2028,
		"cgpa": 8.0,
		"github": "https://github.com/test",
		"linkedin": "https://linkedin.com/in/test",
		"portfolio": "https://example.com"
	}`

	request := httptest.NewRequest(
		http.MethodPut,
		"/profile",
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

	var profile models.Profile

	result = database.DB.
		Where("user_id = ?", user.ID).
		First(&profile)

	if result.Error != nil {
		t.Fatalf(
			"failed to find profile: %v",
			result.Error,
		)
	}

	if profile.University != "GD Goenka University" {
		t.Fatalf(
			"expected university GD Goenka University, got %s",
			profile.University,
		)
	}

	if profile.Degree != "B.Tech Computer Science" {
		t.Fatalf(
			"expected degree B.Tech Computer Science, got %s",
			profile.Degree,
		)
	}

	if profile.GraduationYear != 2028 {
		t.Fatalf(
			"expected graduation year 2028, got %d",
			profile.GraduationYear,
		)
	}

	if profile.CGPA != 8.0 {
		t.Fatalf(
			"expected CGPA 8.0, got %f",
			profile.CGPA,
		)
	}
}
func TestGetProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	os.Setenv("JWT_SECRET", "test-secret")

	timestamp := time.Now().UnixNano()

	user := models.User{
		Name:         "Get Profile User",
		Username:     fmt.Sprintf("get-profile-%d", timestamp),
		Email:        fmt.Sprintf("get-profile-%d@example.com", timestamp),
		PasswordHash: "test-password",
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		t.Fatalf("failed to create test user: %v", result.Error)
	}

	profile := models.Profile{
		UserID:         user.ID,
		University:     "GD Goenka University",
		Degree:         "B.Tech Computer Science",
		GraduationYear: 2028,
		CGPA:           8.0,
		GitHub:         "https://github.com/test",
		LinkedIn:       "https://linkedin.com/in/test",
		Portfolio:      "https://example.com",
	}

	result = database.DB.Create(&profile)

	if result.Error != nil {
		t.Fatalf("failed to create test profile: %v", result.Error)
	}

	token, err := services.GenerateToken(user.ID)

	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	router := gin.New()

	router.GET(
		"/profile",
		middleware.AuthRequired(),
		GetProfile,
	)

	request := httptest.NewRequest(
		http.MethodGet,
		"/profile",
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

	body := recorder.Body.String()

	if !strings.Contains(body, "GD Goenka University") {
		t.Fatalf(
			"expected university in response, got %s",
			body,
		)
	}

	if !strings.Contains(body, "B.Tech Computer Science") {
		t.Fatalf(
			"expected degree in response, got %s",
			body,
		)
	}
}
