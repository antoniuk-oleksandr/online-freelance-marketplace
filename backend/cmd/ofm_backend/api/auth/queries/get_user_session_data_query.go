package queries

const GetUserSessionDataQuery = `
SELECT master_key
FROM users U
WHERE user_id = $1
`
