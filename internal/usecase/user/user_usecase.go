package user

import (
	"context"
	userDomain "github.com/hum2/backend/internal/domain/user"
	idVo "github.com/hum2/backend/internal/domain/user/id"
	"github.com/hum2/backend/internal/usecase/shared/transaction"
)

// Usecase is a usecase of user
type Usecase struct {
	tx   transaction.Handler
	repo userDomain.Repository
}

// New is a constructor of Usecase
func New(tx transaction.Handler, repo userDomain.Repository) *Usecase {
	return &Usecase{
		tx:   tx,
		repo: repo,
	}
}

// FindAll is a method to find all users
func (u *Usecase) FindAll(ctx context.Context) ([]*userDomain.User, error) {
	return u.repo.FindAll(ctx)
}

// FindByID is a method to find user by id
func (u *Usecase) FindByID(ctx context.Context, userID idVo.UserID) (*userDomain.User, error) {
	return u.repo.FindByID(ctx, userID)
}

// Create is a method to create user
func (u *Usecase) Create(ctx context.Context, name, birthday string) (*userDomain.User, error) {
	res, err := u.tx.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		user, err := userDomain.New(name, birthday)
		if err != nil {
			return nil, err
		}
		err = u.repo.Create(ctx, user)
		if err != nil {
			return nil, err
		}
		return user, nil
	})
	if err != nil {
		return nil, err
	}
	return res.(*userDomain.User), nil
}

// Update is a method to update user
func (u *Usecase) Update(ctx context.Context, user *userDomain.User) error {
	_, err := u.tx.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		return nil, u.repo.Update(ctx, user)
	})
	return err
}

// Delete is a method to delete user by id
func (u *Usecase) Delete(ctx context.Context, id idVo.UserID) error {
	_, err := u.tx.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		return nil, u.repo.Delete(ctx, id)
	})
	return err
}
