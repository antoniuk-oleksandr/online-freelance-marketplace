package model

type HomeData struct {
	BestFreelancers []BestFreelancer `json:"best_freelancers" db:"best_freelancers"`
	KeyMetrics      KeyMetrics        `json:"key_metrics" db:"key_metrics"`
	BestFreelances  []BestFreelance   `json:"best_freelances" db:"fbest_freelances"`
	BestReviews     []BestReview      `json:"best_reviews" db:"best_reviews"`
}