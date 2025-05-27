export type SignInData = {
    usernameOrEmail: string,
    password: string
}

export type SignUpData = {
    firstName: string,
    surname: string,
    username: string,
    password: string,
    email: string,
    privateKey: Uint8Array
    publicKey: Uint8Array
    privateKeyIV: Uint8Array
    privateKeySalt: Uint8Array
}
