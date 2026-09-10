package models

type Profile struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	UserID         uint    `gorm:"unique;not null" json:"user_id"`
	University     string  `json:"university"`
	Degree         string  `json:"degree"`
	GraduationYear int     `json:"graduation_year"`
	CGPA           float64 `json:"cgpa"`
	GitHub         string  `json:"github"`
	LinkedIn       string  `json:"linkedin"`
	Portfolio      string  `json:"portfolio"`
}

type UpdateProfileRequest struct {
	University     string  `json:"university"`
	Degree         string  `json:"degree"`
	GraduationYear int     `json:"graduation_year"`
	CGPA           float64 `json:"cgpa"`
	GitHub         string  `json:"github"`
	LinkedIn       string  `json:"linkedin"`
	Portfolio      string  `json:"portfolio"`
}