package models

import "testing"

func TestIsValidStatus(t *testing.T) {
	tests := []struct {
		name   string
		status ApplicationStatus
		valid  bool
	}{
		{
			name:   "Applied is valid",
			status: StatusApplied,
			valid:  true,
		},
		{
			name:   "Assessment is valid",
			status: StatusAssessment,
			valid:  true,
		},
		{
			name:   "Interview is valid",
			status: StatusInterview,
			valid:  true,
		},
		{
			name:   "Offer is valid",
			status: StatusOffer,
			valid:  true,
		},
		{
			name:   "Rejected is valid",
			status: StatusRejected,
			valid:  true,
		},
		{
			name:   "Invalid status",
			status: ApplicationStatus("Banana"),
			valid:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidStatus(tt.status)

			if result != tt.valid {
				t.Fatalf(
					"expected %v, got %v",
					tt.valid,
					result,
				)
			}
		})
	}
}

func TestCanTransition(t *testing.T) {
	tests := []struct {
		name string
		from ApplicationStatus
		to   ApplicationStatus
		want bool
	}{
		{
			name: "Applied to Assessment",
			from: StatusApplied,
			to:   StatusAssessment,
			want: true,
		},
		{
			name: "Assessment to Interview",
			from: StatusAssessment,
			to:   StatusInterview,
			want: true,
		},
		{
			name: "Interview to Offer",
			from: StatusInterview,
			to:   StatusOffer,
			want: true,
		},
		{
			name: "Interview to Rejected",
			from: StatusInterview,
			to:   StatusRejected,
			want: true,
		},
		{
			name: "Assessment to Applied",
			from: StatusAssessment,
			to:   StatusApplied,
			want: false,
		},
		{
			name: "Assessment to Offer",
			from: StatusAssessment,
			to:   StatusOffer,
			want: false,
		},
		{
			name: "Offer to Interview",
			from: StatusOffer,
			to:   StatusInterview,
			want: false,
		},
		{
			name: "Rejected to Interview",
			from: StatusRejected,
			to:   StatusInterview,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanTransition(tt.from, tt.to)

			if result != tt.want {
				t.Fatalf(
					"expected %v, got %v",
					tt.want,
					result,
				)
			}
		})
	}
}