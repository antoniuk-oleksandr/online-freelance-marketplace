package model

type OrderAnswer struct {
	ServiceQuestionId int    `db:"service_question_id"`
	Content           string `db:"content"`
	OrderId           int    `db:"order_id"`
}
