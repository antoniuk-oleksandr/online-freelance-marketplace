package queries

const GetUserPasswordQuery = `
SELECT
username, password 
FROM users 
WHERE username = $1 OR email = $1
`