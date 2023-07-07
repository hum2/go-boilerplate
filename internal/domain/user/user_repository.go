//go:generate goimports mock_$GOFILE
//go:generate gofmt -w mock_$GOFILE
package user

import (
	"context"
	idVo "github.com/hum2/backend/internal/domain/user/id"
)

// Repository is a repository interface of user.
type Repository interface {
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, userID idVo.UserID) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID idVo.UserID) error
}
