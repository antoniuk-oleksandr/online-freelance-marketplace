package queries

var AddUserQuery = `
INSERT INTO users
(first_name, surname, email, username, password, public_key, private_key, private_key_iv, private_key_salt, master_key)
VALUES (
:first_name, :surname, :email, :username,
:password, :public_key, :private_key,
:private_key_iv, :private_key_salt, :master_key
)
`
