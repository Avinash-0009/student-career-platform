package models

type ApplicationStatus string

const (
	StatusApplied    ApplicationStatus = "Applied"
	StatusAssessment ApplicationStatus = "Assessment"
	StatusInterview  ApplicationStatus = "Interview"
	StatusOffer      ApplicationStatus = "Offer"
	StatusRejected   ApplicationStatus = "Rejected"
)

func IsValidStatus(status ApplicationStatus) bool {
	switch status {
	case StatusApplied,
		StatusAssessment,
		StatusInterview,
		StatusOffer,
		StatusRejected:
		return true
	default:
		return false
	}
}
func CanTransition(from, to ApplicationStatus) bool {
	if from == to {
		return true
	}

	switch from {
	case StatusApplied:
		return to == StatusAssessment ||
			to == StatusInterview ||
			to == StatusRejected

	case StatusAssessment:
		return to == StatusInterview ||
			to == StatusRejected

	case StatusInterview:
		return to == StatusOffer ||
			to == StatusRejected

	case StatusOffer:
		return false

	case StatusRejected:
		return false
	}

	return false
}