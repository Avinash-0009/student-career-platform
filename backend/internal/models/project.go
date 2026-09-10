package models

import "time"

type Project struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null;index" json:"user_id"`
	Name         string    `gorm:"not null" json:"name"`
	Description  string    `json:"description"`
	Technologies string    `json:"technologies"`
	GitHubURL    string    `json:"github_url"`
	LiveURL      string    `json:"live_url"`
	Status       string    `gorm:"not null" json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateProjectRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	Technologies string `json:"technologies"`
	GitHubURL    string `json:"github_url"`
	LiveURL      string `json:"live_url"`
	Status       string `json:"status" binding:"required"`
}

type UpdateProjectRequest struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	Technologies string `json:"technologies"`
	GitHubURL    string `json:"github_url"`
	LiveURL      string `json:"live_url"`
	Status       string `json:"status" binding:"required"`
}