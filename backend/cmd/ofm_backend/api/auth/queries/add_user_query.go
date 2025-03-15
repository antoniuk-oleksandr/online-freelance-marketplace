package queries

var AddUserQuery = `
INSERT INTO users
(first_name, surname, email, username, password)
VALUES (:first_name, :surname, :email, :username, :password)
`
