package models

import "time"

type Skill struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	Name        string    `gorm:"not null" json:"name"`
	Proficiency int       `gorm:"not null" json:"proficiency"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateSkillRequest struct {
	Name        string `json:"name" binding:"required"`
	Proficiency int    `json:"proficiency" binding:"required,min=0,max=100"`
}

type UpdateSkillRequest struct {
	Name        string `json:"name" binding:"required"`
	Proficiency int    `json:"proficiency" binding:"required,min=0,max=100"`
}