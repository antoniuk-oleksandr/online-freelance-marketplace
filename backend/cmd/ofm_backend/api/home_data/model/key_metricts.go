package model

type KeyMetrics struct {
	FreelancesAvailable int     `json:"freelances_available" db:"freelances_available"`
	ProjectsCompleted   int     `json:"projects_completed" db:"projects_completed"`
	AvgRating           float64 `json:"avg_rating" db:"avg_rating"`
}
