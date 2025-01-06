package dto

type HomeData struct {
	BestFreelancers []BestFreelancer `json:"bestFreelancers"`
	BestFreelances  []BestFreelance  `json:"bestServices"`
	KeyMetrics      KeyMetrics       `json:"keyMetrics"`
	BestReviews     []BestReview     `json:"bestReviews"`
}
