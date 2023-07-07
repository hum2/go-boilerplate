package user

import (
	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
	idVo "github.com/hum2/backend/internal/domain/user/id"
	"time"
)

const (
	minNameLength = 1
	maxNameLength = 100
)

var ErrMaxNameLength = errors.New("name length must be between 1 and 100")

// User is a domain model of user.
type User struct {
	id       idVo.UserID
	name     string
	birthday time.Time
}

// New is a constructor of User.
func New(name, birthday string) (*User, error) {
	if err := validateName(name); err != nil {
		return nil, ErrMaxNameLength
	}
	t, err := time.Parse(time.DateOnly, birthday)
	if err != nil {
		return nil, err
	}
	return &User{
		id:       idVo.New(),
		name:     name,
		birthday: t,
	}, nil
}

// NewFromRepository is a constructor of User from repository.
func NewFromRepository(id uuid.UUID, name string, birthday time.Time) *User {
	return &User{
		id:       idVo.NewFromRepository(id),
		name:     name,
		birthday: birthday,
	}
}

// ID is a getter of idVo.
func (u *User) ID() idVo.UserID {
	return u.id
}

// Name is a getter of name.
func (u *User) Name() string {
	return u.name
}

// Birthday is a getter of birthday.
func (u *User) Birthday() time.Time {
	return u.birthday
}

// BirthdayString is a getter of birthday as string.
func (u *User) BirthdayString() string {
	return u.birthday.Format(time.DateOnly)
}

func validateName(name string) error {
	if len(name) < minNameLength || len(name) > maxNameLength {
		return ErrMaxNameLength
	}
	return nil
}
