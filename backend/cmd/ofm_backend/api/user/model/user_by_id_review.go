package model

import "time"

type UserByIdReview struct {
	ID           int64     `db:"id" json:"id"`
	Content      string    `db:"content" json:"content"`
	Rating       int       `db:"rating" json:"rating"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	EndedAt      time.Time `db:"ended_at" json:"ended_at"`
	UserID       int64     `db:"user_id" json:"user_id"`
	FirstName    string    `db:"first_name" json:"first_name"`
	Surname      string    `db:"surname" json:"surname"`
	Avatar       *string   `db:"avatar" json:"avatar"`
	ServiceID    int64     `db:"service_id" json:"service_id"`
	Price        float64   `db:"price" json:"price"`
	ServiceImage *string   `db:"service_image" json:"service_image"`
	Title        string    `db:"title" json:"title"`
}
