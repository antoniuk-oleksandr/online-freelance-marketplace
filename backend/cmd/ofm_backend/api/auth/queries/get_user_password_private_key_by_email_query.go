package queries

const GetUserPasswordPrivateKeyByEmailQuery = `
SELECT 
password, private_key 
FROM users
WHERE email = $1
`
