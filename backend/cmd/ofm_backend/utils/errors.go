package utils

import "errors"

var ErrUserAlreadyExists = errors.New("User already exists with this username")

var ErrNoFound = errors.New("Not found")

var ErrTempTokenExpired = errors.New("Temporary token has expired")