package hellowire

import (
	"context"
)

// Usecase is a usecase of user
type Usecase struct{}

// New is a constructor of Usecase
func New() *Usecase {
	return &Usecase{}
}

// FindAll is a method to find all users
func (u *Usecase) FindAll(ctx context.Context) string {
	return "Hello Wire!"
}
