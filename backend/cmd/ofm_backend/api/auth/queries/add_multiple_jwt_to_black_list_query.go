package queries

const AddMultipleJWTToBlackListQuery = `
INSERT INTO blacklisted_tokens
(token)
VALUES (:token)
`