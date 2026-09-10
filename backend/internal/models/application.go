package models

import "time"

type Application struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	UserID      uint       `gorm:"not null;index" json:"user_id"`
	Company     string     `gorm:"not null" json:"company"`
	JobTitle    string     `gorm:"not null" json:"job_title"`
	JobURL      string     `json:"job_url"`
	Status      string     `gorm:"not null" json:"status"`
	AppliedDate *time.Time `json:"applied_date"`
	Deadline    *time.Time `json:"deadline"`
	Notes       string     `json:"notes"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type CreateApplicationRequest struct {
	Company     string     `json:"company" binding:"required"`
	JobTitle    string     `json:"job_title" binding:"required"`
	JobURL      string     `json:"job_url"`
	Status      string     `json:"status" binding:"required"`
	AppliedDate *time.Time `json:"applied_date"`
	Deadline    *time.Time `json:"deadline"`
	Notes       string     `json:"notes"`
}

type UpdateApplicationRequest struct {
	Company     string     `json:"company" binding:"required"`
	JobTitle    string     `json:"job_title" binding:"required"`
	JobURL      string     `json:"job_url"`
	Status      string     `json:"status" binding:"required"`
	AppliedDate *time.Time `json:"applied_date"`
	Deadline    *time.Time `json:"deadline"`
	Notes       string     `json:"notes"`
}