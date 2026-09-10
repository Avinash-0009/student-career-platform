package services

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret")

	userID := uint(123)

	tokenString, err := GenerateToken(userID)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if tokenString == "" {
		t.Fatal("expected token, got empty string")
	}

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return []byte("test-secret"), nil
		},
	)

	if err != nil {
		t.Fatalf("failed to parse token: %v", err)
	}

	if !token.Valid {
		t.Fatal("expected token to be valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		t.Fatal("expected JWT claims")
	}

	tokenUserID, ok := claims["user_id"].(float64)

	if !ok {
		t.Fatal("expected user_id claim")
	}

	if uint(tokenUserID) != userID {
		t.Fatalf(
			"expected user_id %d, got %v",
			userID,
			tokenUserID,
		)
	}
}