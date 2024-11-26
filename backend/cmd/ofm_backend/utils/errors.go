package utils

import "errors"

var (
	ErrNoFound              = errors.New("Not found")
	ErrTempTokenExpired     = errors.New("Temporary token has expired")
	ErrMailSend             = errors.New("gomail: could not send email 1: gomail: invalid address \"test11\": mail: missing '@' or angle-addr")
	ErrUsernameIsTaken      = errors.New("Username is already taken")
	ErrEmailIsTaken         = errors.New("Email is already taken")
	ErrInvalidRequestBody   = errors.New("Invalid request body")
	ErrUsernameNotAvailable = errors.New("Username is not available")
	ErrUnexpectedError      = errors.New("An unexpected error occurred")
	ErrEmailDoesNotExist    = errors.New("Email does not exist")
	ErrUsernameDoesNotExist = errors.New("Username does not exist")
	ErrBlacklistedToken     = errors.New("Token is blacklisted")
)
