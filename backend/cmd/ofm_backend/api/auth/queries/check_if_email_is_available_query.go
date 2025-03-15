package queries

const CheckIfEmailIsAvailableQuery = `
SELECT EXISTS
(SELECT  * FROM users WHERE email = $1)
`