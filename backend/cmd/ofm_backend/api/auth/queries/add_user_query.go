package queries

var AddUserQuery = `
INSERT INTO users
(first_name, surname, email, username, password, public_key, private_key)
VALUES (:first_name, :surname, :email, :username, :password, :public_key, :private_key)
`
