export enum ResponseErrorEnum {
    InvalidToken = "Invalid token",
    ExpiredToken = "Token is expired",
    WrongToken = "Wrong token",
    InvalidCredentials = "Invalid credentials",
    InvalidRequestBody = "Invalid request body",
    InvalidServiceID = "Invalid freelance service ID",
    ServiceNotFound = "Freelance service not found",
    InvalidUserID = "Invalid user ID",
    UserNotFound = "User not found",
    UserNotConfirmed = "User not confirmed",
    CouldNotCreateAccount = "Could not create account",
    UnexpectedError = "An unexpected error occurred",
    InvalidEmail = "Invalid email address",
    UsernameIsTaken = "Username is already taken",
    EmailIsTaken = "Email is already taken",
    EmailDoesNotExist = "Email does not exist",
    UsernameDoesNotExist = "Username does not exist",
    BlacklistedToken = "Token is blacklisted",
    InvalidCursor = "Cursor is invalid",
    ErrTooManyRequests = "Too many requests",
    ErrEncryption = "Encryption error",
    ErrInvalidAccessTokens = "Invalid access tokens",
}

export type ResponseError = {
    error: ResponseErrorEnum,
}