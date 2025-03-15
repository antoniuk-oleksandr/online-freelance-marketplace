package queries

const ChangeUserPasswordByEmailQuery = `
UPDATE users
SET password = $1
WHERE email = $2
`
