package queries

const AddUserWithGoogleAuthQuery = `
INSERT INTO users (
    email,
    username,
    password,
    first_name,
    surname,
    avatar_id,
    private_key,
    public_key,
    private_key_iv,
    private_key_salt,
    master_key
) VALUES (
    :email,
    :username,
    :password,
    :first_name,
    :surname,
    :avatar_id,
    :private_key,
    :public_key,
    :private_key_iv,
    :private_key_salt,
    :master_key
) RETURNING user_id;

`
