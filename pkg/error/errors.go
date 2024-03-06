package core

import "errors"

var (
	// ErrUserNotFound is returned when the user is not found
	ErrUserNotFound = errors.New("user not found")
)
