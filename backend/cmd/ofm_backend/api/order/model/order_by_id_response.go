package model

type OrderByIdResponse struct {
	Order            Order               `json:"order" db:"order"`
	Service          Freelance           `json:"service" db:"service"`
	ServiceQuestions []FreelanceQuestion `json:"service_questions" db:"service_questions"`
}
