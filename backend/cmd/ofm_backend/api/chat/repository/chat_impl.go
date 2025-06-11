package repository

import (
	"ofm_backend/cmd/ofm_backend/api/chat/model"
	"ofm_backend/cmd/ofm_backend/api/chat/queries"
	"ofm_backend/cmd/ofm_backend/enum"

	"github.com/jmoiron/sqlx"
)

type chatRepository struct {
	db *sqlx.DB
}

func NewChatRepository(db *sqlx.DB) ChatRepository {
	return &chatRepository{db: db}
}

func (c *chatRepository) SaveMessage(message model.ChatMessageToReceive) (int, error) {
	var messageId int
	rows, err := c.db.Query(
		queries.SaveChatMessageQuery, message.SenderId, message.Content,
		message.ContentIV, message.SentAt, message.OrderId, enum.Sent,
	)
	if err != nil {
		return messageId, err
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&messageId); err != nil {
			return messageId, err
		}
	}

	return messageId, nil
}
