package queries

const GetEmailByUsernameIfExistsQuery = `
SELECT email 
FROM users 
WHERE username = $1
`