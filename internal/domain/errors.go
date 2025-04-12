package domain

import "errors"

var (
	ErrAccountNotFound   = errors.New("account not found")
	ErrDuplicatedAPIKey  = errors.New("account already exists")
	ErrInvalidAmout      = errors.New("invalid amout")
	ErrInvalidStatus     = errors.New("invalid status")
	ErrInvoiceNotFound   = errors.New("invoice not found")
	ErrUnauthorizedAcces = errors.New("access invalid")
)
