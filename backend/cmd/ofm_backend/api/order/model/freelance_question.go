package model

type FreelanceQuestion struct {
	Id      int64  `json:"service_question_id" db:"service_question_id"`
	Content string `json:"content" db:"content"`
}
