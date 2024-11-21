export enum ResponseErrorEnum {
    InvalidToken = "Invalid token",
    ExpiredToken = "Token is expired",
    WrongToken = "Wrong token",
    InvalidCredentials = "Invalid credentials",
    InvalidRequestBody = "Invalid request body",
    InvalidServiceID = "Invalid freelance service ID",
    ServiceNotFound = "Freelance service not found",
    InvalidUserID = "Invalid user ID",
    UserAlreadyExists = "User already exists",
    UserNotFound = "User not found",
    UserNotConfirmed = "User not confirmed",
    CouldNotCreateAccount = "Could not create account",
    UnexpectedError = "An unexpected error occurred",
}

export type ResponseError = {
    error?: ResponseErrorEnum,
    message?: ResponseErrorEnum,
}