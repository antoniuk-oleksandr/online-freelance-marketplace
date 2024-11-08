package utils

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists with this username")

var ErrNoFound = errors.New("Not found")
