package model

import "time"

type Review struct {
	ID        int64            `json:"id" db:"review_id"`
	Content   string           `json:"content" db:"content"`
	Rating    int64            `json:"rating" db:"rating"`
	CreatedAt time.Time        `json:"createdAt" db:"created_at"`
	EndedAt   time.Time        `json:"endedAt" db:"ended_at"`
	Customer  *Customer        `json:"customer" db:"customer"`
	Freelance *ReviewFreelance `json:"service" db:"service"`
}
