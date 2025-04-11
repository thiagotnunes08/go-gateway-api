package domain

import "errors"

var (
	ErrAccountNotFound  = errors.New("account not found")
	ErrDuplicatedAPIKey = errors.New("account already exists")
)
