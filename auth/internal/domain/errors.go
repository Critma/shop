package domain

import "fmt"

var (
	ErrNotFound           = fmt.Errorf("not found")
	ErrAlreadyExists      = fmt.Errorf("already exists")
	ErrInvalidCredentials = fmt.Errorf("invalid credentials")
)
