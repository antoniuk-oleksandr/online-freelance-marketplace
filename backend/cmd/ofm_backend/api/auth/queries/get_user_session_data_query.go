package queries

const GetUserSessionDataQuery = `
WITH chat_partners_data AS (
    SELECT DISTINCT
        U.user_id,
        U.public_key
    FROM orders O
    JOIN users U ON (
        (O.customer_id = $1 AND U.user_id = O.freelancer_id) OR
        (O.freelancer_id = $1 AND U.user_id = O.customer_id)
    )
    WHERE
        (O.freelancer_id = $1 OR O.customer_id = $1)
    ORDER BY user_id
),
user_data AS (
    SELECT master_key
    FROM users U
    WHERE user_id = $1
)

SELECT
(
    SELECT json_agg(
        json_build_object(
            'user_id', CPD.user_id,
            'public_key', encode(CPD.public_key, 'base64')
        )
    ) FROM chat_partners_data CPD
) as chat_partners,
(
    SELECT * FROM user_data
)
`
