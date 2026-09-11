package handlers

import (
	"os"
	"testing"

	"golang.org/x/crypto/bcrypt"

	"student-career-platform/internal/database"
	"student-career-platform/internal/models"
)

func TestMain(m *testing.M) {
	err := database.Connect()
	if err != nil {
		panic(err)
	}

	err = database.DB.AutoMigrate(
		&models.User{},
		&models.Profile{},
		&models.Project{},
		&models.Skill{},
		&models.Application{},
	)
	if err != nil {
		panic(err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)
	if err != nil {
		panic(err)
	}

	testUser := models.User{
		Name:         "Avinash",
		Username:     "avinash",
		Email:        "avinash@example.com",
		PasswordHash: string(passwordHash),
	}

	database.DB.FirstOrCreate(
		&testUser,
		models.User{Email: "avinash@example.com"},
	)

	code := m.Run()

	os.Exit(code)
}
