package repository

import "github.com/jmoiron/sqlx"

type chatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) ChatRepository {
	return &chatRepository{db: db}
}
