package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"student-career-platform/internal/database"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	os.Setenv("JWT_SECRET", "test-secret")

	err := database.Connect()

	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	router := gin.New()

	router.POST("/login", Login)

	requestBody := `{
		"email": "avinash@example.com",
		"password": "password123"
	}`

	request := httptest.NewRequest(
		http.MethodPost,
		"/login",
		strings.NewReader(requestBody),
	)

	request.Header.Set(
		"Content-Type",
		"application/json",
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

	var response map[string]interface{}

	err = json.Unmarshal(
		recorder.Body.Bytes(),
		&response,
	)

	if err != nil {
		t.Fatalf(
			"failed to parse response: %v",
			err,
		)
	}

	token, ok := response["token"].(string)

	if !ok || token == "" {
		t.Fatal("expected JWT token in response")
	}
}
