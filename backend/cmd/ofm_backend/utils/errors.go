package utils

import "errors"

var (
	// General Errors
	ErrEnvVarNotSet       = errors.New("Environment variable not set")
	ErrNotFound           = errors.New("Not found")
	ErrInvalidRequestBody = errors.New("Invalid request body")
	ErrUnexpectedError    = errors.New("An unexpected error occurred")
	ErrTooManyRequests    = errors.New("Too many requests")
	ErrParsingError       = errors.New("Parsing error")

	// Auth
	ErrTempTokenExpired     = errors.New("Temporary token has expired")
	ErrMailSend             = errors.New("gomail: could not send email 1: gomail: invalid address \"test11\": mail: missing '@' or angle-addr")
	ErrUsernameIsTaken      = errors.New("Username is already taken")
	ErrEmailIsTaken         = errors.New("Email is already taken")
	ErrUsernameNotAvailable = errors.New("Username is not available")
	ErrEmailDoesNotExist    = errors.New("Email does not exist")
	ErrUsernameDoesNotExist = errors.New("Username does not exist")
	ErrBlacklistedToken     = errors.New("Token is blacklisted")

	// User
	ErrUserNotFound = errors.New("User not found")

	//Search
	ErrInvalidCursor           = errors.New("Cursor is invalid")
	ErrInvalidPaymentPublicKey = errors.New("Invalid payment public key")

	// Payment Errors
	ErrDecryptionFailed          = errors.New("Failed to decrypt payment data")
	ErrPaymentCreationFailed     = errors.New("Failed to create payment")
	ErrTransactionFailed         = errors.New("Transaction failed")
	ErrOrderCreationFailed       = errors.New("Error while creating order")
	ErrInvalidCardNumber         = errors.New("Invalid card number")
	ErrPaymentStatusUpdateFailed = errors.New("Failed to update payment status")
	ErrPayPalPaymentFailed       = errors.New("PayPal payment was unsuccessful")

	// Order Errors
	ErrOrderNotFound              = errors.New("Order not found")
	ErrFreelanceNotFound          = errors.New("Freelance not found")
	ErrFreelanceQuestionsNotFound = errors.New("Freelance questions not found")
	ErrAlreadySubmitted           = errors.New("Order requirements have already been submitted")

	// File Errors
	ErrFailedFileUploadRequest = errors.New("Failed to upload file")
)

type ErrorResponse struct {
	Error string `json:"error"`
}
