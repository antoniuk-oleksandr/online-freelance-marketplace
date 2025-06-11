package queries

var SaveChatMessageQuery = `
INSERT INTO chat_messages (sender_id, content, content_iv, sent_at, order_id, type)
VALUES ($1, $2, $3, TO_TIMESTAMP($4::BIGINT / 1000), $5, $6)
RETURNING chat_message_id
`