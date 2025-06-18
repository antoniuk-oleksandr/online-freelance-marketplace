export enum ResponseErrorEnum {
    // General Errors
    EnvVarNotSet = "Environment variable not set",
    NotFound = "Not found",
    InvalidRequestBody = "Invalid request body",
    UnexpectedError = "An unexpected error occurred",
    TooManyRequests = "Too many requests",
    ParsingError = "Parsing error",
    InvalidPathParam = "Invalid path parameter",
    InvalidToken = "Invalid token",
    ExpiredToken = "Token has expired",
    InvalidParameter = "Invalid parameter",
  
    // Auth
    TempTokenExpired = "Temporary token has expired",
    MailSend = "gomail: could not send email 1: gomail: invalid address \"test11\": mail: missing '@' or angle-addr",
    UsernameIsTaken = "Username is already taken",
    EmailIsTaken = "Email is already taken",
    UsernameNotAvailable = "Username is not available",
    EmailDoesNotExist = "Email does not exist",
    UsernameDoesNotExist = "Username does not exist",
    BlacklistedToken = "Token is blacklisted",
    Unauthorized = "Unauthorized",
    SessionCacheNotFound = "Session cache not found",
    InvalidSessionCache = "Invalid session cache",
    InvalidPublicKey = "Invalid public key",
  
    // User
    UserNotFound = "User not found",
  
    // Search
    InvalidCursor = "Cursor is invalid",
    InvalidPaymentPublicKey = "Invalid payment public key",
  
    // Payment Errors
    DecryptionFailed = "Failed to decrypt payment data",
    PaymentCreationFailed = "Failed to create payment",
    TransactionFailed = "Transaction failed",
    OrderCreationFailed = "Error while creating order",
    InvalidCardNumber = "Invalid card number",
    PaymentStatusUpdateFailed = "Failed to update payment status",
    PayPalPaymentFailed = "PayPal payment was unsuccessful",
  
    // Order Errors
    OrderNotFound = "Order not found",
    FreelanceNotFound = "Freelance not found",
    FreelanceQuestionsNotFound = "Freelance questions not found",
    AlreadySubmitted = "Order requirements have already been submitted",
  
    // File Errors
    FailedFileUploadRequest = "Failed to upload file",
  
    // My Profile Errors
    NoDataFound = "No data found",
    CompletingOrder = "Error while completing order",
    AddingDeliveryFiles = "Error while adding delivery files",
    OrderReviewNotFound = "Order review not found",
  
    // Chat Errors
    BadUserIdQueryFormat = "Bad user id query format",
  
    // Extras (from existing enum)
    InvalidCredentials = "Invalid credentials",
    InvalidServiceID = "Invalid freelance service ID",
    ServiceNotFound = "Freelance service not found",
    InvalidUserID = "Invalid user ID",
    UserNotConfirmed = "User not confirmed",
    CouldNotCreateAccount = "Could not create account",
    InvalidEmail = "Invalid email address",
    EncryptionError = "Encryption error",
    InvalidAccessTokens = "Invalid access tokens",
    InvalidSharedSecret = "Invalid shared secret",
  }
  
  export type ResponseError = {
    error: ResponseErrorEnum
  }
  