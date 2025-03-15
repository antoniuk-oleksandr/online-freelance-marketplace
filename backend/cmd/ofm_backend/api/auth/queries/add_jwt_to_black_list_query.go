package queries

const AddJWTToBlackListQuery = `
INSERT INTO blacklisted_tokens
(token)
VALUES ($1)
`
