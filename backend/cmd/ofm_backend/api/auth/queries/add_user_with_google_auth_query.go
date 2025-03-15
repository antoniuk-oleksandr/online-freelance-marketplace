package queries

const AddUserWithGoogleAuthQuery = `
INSERT INTO users
(username, email, password, first_name, surname, avatar_id)
VALUES ($1, $2, $3, $4, $5, $6)
`
