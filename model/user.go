package model

import "errors"

var (
	// ErrNameRequired enforces user to have a name
	ErrNameRequired = errors.New("Name is required")

	// ErrPasswordRequired enforces user to have a password
	ErrPasswordRequired = errors.New("Password is required")

	// ErrPasswordLen enforces password length to be at least 3 chars
	ErrPasswordLen = errors.New("Password must be at least 3 chars")
)

// User represents an user of the system
type User struct {
	ID       int64
	Name     string
	Password string
	Active   bool
}

// IsValid validates user
func (u *User) IsValid() error {
	if len(u.Name) == 0 {
		return ErrNameRequired
	}

	if len(u.Password) == 0 {
		return ErrPasswordRequired
	}

	if len(u.Password) < 3 {
		return ErrPasswordLen
	}

	return nil
}
