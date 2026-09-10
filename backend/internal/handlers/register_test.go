package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
)

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	router := gin.New()

	router.POST("/register", Register)

	email := fmt.Sprintf(
		"test-%d@example.com",
		time.Now().UnixNano(),
	)

	requestBody := fmt.Sprintf(`{
		"name": "Test User",
		"username": "testuser-%d",
		"email": "%s",
		"password": "testpassword123"
	}`, time.Now().UnixNano(), email)

	request := httptest.NewRequest(
		http.MethodPost,
		"/register",
		strings.NewReader(requestBody),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
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
}