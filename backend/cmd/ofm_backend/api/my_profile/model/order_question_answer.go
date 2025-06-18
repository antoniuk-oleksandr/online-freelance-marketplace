package model

type OrderQuestionAnswer struct {
	Question string `json:"question" db:"question"`
	Answer   string `json:"answer" db:"answer"`
}
