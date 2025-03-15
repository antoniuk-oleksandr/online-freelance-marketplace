package queries

const GetUsernameByEmailIfExistsQuery = `
SELECT username 
FROM users 
WHERE email = $1
`