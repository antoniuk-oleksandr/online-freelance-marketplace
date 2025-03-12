package dto

type OrderByIdResponse struct {
	Order            Order               `json:"order"`
	Freelance        Freelance           `json:"service"`
	ServiceQuestions []FreelanceQuestion `json:"freelanceQuestions"`
}
