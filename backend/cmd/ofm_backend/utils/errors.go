package utils

import "errors"

var ErrNoFound = errors.New("Not found")

var ErrTempTokenExpired = errors.New("Temporary token has expired")

var ErrMailSend = errors.New("gomail: could not send email 1: gomail: invalid address \"test11\": mail: missing '@' or angle-addr")

var ErrUsernameIsTaken = errors.New("Username is already taken")

var ErrEmailIsTaken = errors.New("Email is already taken")