package queries

const CheckIfUsernameIsAvailableQuery = `
SELECT EXISTS 
(SELECT  * FROM users WHERE username = $1)
`