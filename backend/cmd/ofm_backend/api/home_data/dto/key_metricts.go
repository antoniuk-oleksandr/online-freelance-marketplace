package dto

type KeyMetrics struct {
	FreelancesAvailable int     `json:"freelancesAvailable"`
	ProjectsCompleted   int     `json:"projectsCompleted"`
	AvgRating           float64 `json:"avgRating"`
}
