package models

type DashboardStats struct {
	TotalApplications int64 `json:"total_applications"`
	Applied           int64 `json:"applied"`
	Assessment        int64 `json:"assessment"`
	Interview         int64 `json:"interview"`
	Offer             int64 `json:"offer"`
	Rejected          int64 `json:"rejected"`
	Projects          int64 `json:"projects"`
	Skills            int64 `json:"skills"`
}