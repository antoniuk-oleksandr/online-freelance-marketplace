package queries

const ChangeUserPasswordPrivateKeyByEmailQuery = `
UPDATE users
SET 
	password = $1, 
	private_key = $2
WHERE email = $3
`
