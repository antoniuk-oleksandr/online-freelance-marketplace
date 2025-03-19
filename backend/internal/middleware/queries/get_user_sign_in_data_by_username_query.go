package queries

const GetUserSignInDataByUsernameQuery = `
SELECT U.user_id, COALESCE(F.name, '') as avatar
FROM users U
LEFT JOIN files F ON F.file_id = U.avatar_id
WHERE U.username = $1;
`