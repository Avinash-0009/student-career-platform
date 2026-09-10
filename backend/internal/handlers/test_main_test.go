package handlers

import (
	"os"
	"testing"

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

	code := m.Run()

	os.Exit(code)
}