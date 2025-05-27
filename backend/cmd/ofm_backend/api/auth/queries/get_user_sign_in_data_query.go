package queries

const GetUserSignInDataQuery = `
SELECT
    U.user_id,
    U.username,
    COALESCE(F.name , '')AS avatar,
    U.private_key,
    U.private_key_iv,
    U.private_key_salt,
    U.master_key
FROM users U
LEFT JOIN files F ON F.file_id = U.avatar_id
WHERE U.username = $1 OR U.email = $1
`
