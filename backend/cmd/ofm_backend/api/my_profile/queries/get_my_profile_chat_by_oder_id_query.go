package queries

var GetMyProfileChatByOrderIdQuery = `
WITH chat_partner_data AS (
    SELECT
        U.user_id, U.username,
        F.name as avatar,
        TO_CHAR(U.last_seen_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') AS last_seen_at
    FROM orders O
    INNER JOIN users U ON U.user_id = O.freelancer_id
    LEFT JOIN files F ON f.file_id = U.avatar_id
    WHERE O.order_id = $1
),
messages_data AS (
    SELECT
        CM.chat_message_id, CM.sender_id, CM.content,
        TO_CHAR(CM.sent_at AT TIME ZONE 'UTC', 'YYYY-MM-DD"T"HH24:MI:SS.US"Z"') AS sent_at,
        CM.type,
        COALESCE(json_agg(F.name) FILTER (WHERE F.name IS NOT NULL), '[]') AS files
    FROM orders O
    LEFT JOIN chat_messages CM ON CM.order_id = O.order_id
    LEFT JOIN chat_message_files CMF ON CMF.chat_message_id = CM.chat_message_id
    LEFT JOIN files F ON F.file_id = CMF.file_id
    WHERE O.order_id = $1
    GROUP BY CM.chat_message_id, CM.sender_id, CM.content, CM.sent_at, CM.type
)

SELECT
(
SELECT json_build_object(
    'user_id', CPD.user_id,
    'username', CPD.username,
    'avatar', COALESCE(CPD.avatar, ''),
    'last_seen_at', CPD.last_seen_at
) as chat_partner
FROM chat_partner_data CPD
),
(
SELECT json_agg(
    json_build_object(
        'chat_message_id', MD.chat_message_id,
        'sender_id', MD.sender_id,
        'content', MD.content,
        'sent_at', MD.sent_at,
        'files', MD.files,
        'type', MD.type
    )
) as messages
FROM messages_data MD
)

`