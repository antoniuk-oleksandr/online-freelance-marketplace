package queries

const CheckIfTokenBlacklistedQuery = `
SELECT EXISTS 
(SELECT token FROM blacklisted_tokens WHERE token = $1)
`