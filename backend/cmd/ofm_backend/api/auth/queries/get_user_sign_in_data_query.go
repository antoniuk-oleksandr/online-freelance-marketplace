package queries

const GetUserSignInDataQuery = `
WITH user_data AS (
    SELECT
        U.user_id,
        U.username,
        COALESCE(F.name, '') AS avatar,
        encode(U.private_key, 'base64') AS private_key,
        encode(U.private_key_iv, 'base64') AS private_key_iv,
        encode(U.private_key_salt, 'base64') AS private_key_salt,
        encode(U.master_key, 'base64') AS master_key
    FROM users U
    LEFT JOIN files F ON F.file_id = U.avatar_id
    WHERE U.username = $1 OR U.email = $1
    LIMIT 1
),
chat_partners_data AS (
    SELECT DISTINCT
        U.user_id,
        U.public_key
    FROM orders O
    JOIN user_data UD ON TRUE
    JOIN users U ON (
        (O.customer_id = UD.user_id AND U.user_id = O.freelancer_id) OR
        (O.freelancer_id = UD.user_id AND U.user_id = O.customer_id)
    )
    WHERE O.freelancer_id = UD.user_id OR O.customer_id = UD.user_id
)

SELECT
(
    SELECT json_build_object(
        'user_id', AD.user_id,
        'username', AD.username,
        'avatar', AD.avatar,
        'private_key', AD.private_key,
        'private_key_iv', AD.private_key_iv,
        'private_key_salt', AD.private_key_salt,
        'master_key', AD.master_key
    ) FROM user_data AS AD
) AS user_data,
(
    SELECT json_agg(
        json_build_object(
            'user_id', CPD.user_id,
            'public_key', encode(CPD.public_key, 'base64')
        )
    ) FROM chat_partners_data CPD
) AS chat_partners;

`
