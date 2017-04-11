package model

import (
	"errors"

	"github.com/mathsalmi/goarch/util"
)

// Errors
var (
	ErrNameRequired     = errors.New("Name is required")
	ErrPasswordRequired = errors.New("Password is required")
	ErrPasswordLen      = errors.New("Password must be at least 3 chars")
	ErrUserNotFound     = errors.New(`User could not be found`)
)

// User represents an user of the system
type User struct {
	ID       int64  `xorm:"pk 'id'"`
	Name     string `xorm:"'username'"`
	Password string `xorm:"'password'"`
	Active   bool   `xorm:"'active'"`
}

// IsValid validates user
func (u *User) IsValid() *util.AppError {
	err := util.NewError()

	if len(u.Name) == 0 {
		err.Append(ErrNameRequired)
	}

	if len(u.Password) == 0 {
		err.Append(ErrPasswordRequired)
	}

	if len(u.Password) < 3 {
		err.Append(ErrPasswordLen)
	}

	return err.Value()
}
