package user

import (
	"context"
	userDomain "github.com/hum2/backend/internal/domain/user"
	idVo "github.com/hum2/backend/internal/domain/user/id"
	"github.com/hum2/backend/internal/infrastructure/db/ent"
)

// Repository is a repository of user
type Repository struct {
	dbHandler ent.Handler
}

// New is a constructor of Repository
func New(dbHandler ent.Handler) userDomain.Repository {
	return &Repository{
		dbHandler: dbHandler,
	}
}

// FindAll is a method to find all users
func (r *Repository) FindAll(ctx context.Context) ([]*userDomain.User, error) {
	client := r.dbHandler.GetClient()

	users, err := client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	userDomains := make([]*userDomain.User, len(users))
	for i, u := range users {
		userDomains[i] = userDomain.NewFromRepository(u.ID, u.Name, u.Birthday)
	}

	return userDomains, nil
}

// FindByID is a method to find user by id
func (r *Repository) FindByID(ctx context.Context, userID idVo.UserID) (*userDomain.User, error) {
	client := r.dbHandler.GetClient()

	u, err := client.User.Get(ctx, userID.UUID())
	if err != nil {
		return nil, err
	}

	return userDomain.NewFromRepository(u.ID, u.Name, u.Birthday), nil
}

// Create is a method to create user
func (r *Repository) Create(ctx context.Context, user *userDomain.User) error {
	client := r.dbHandler.GetClient()

	_, err := client.User.Create().
		SetID(user.ID().UUID()).
		SetName(user.Name()).
		SetBirthday(user.Birthday()).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Update is a method to update user
func (r *Repository) Update(ctx context.Context, user *userDomain.User) error {
	client := r.dbHandler.GetClient()

	u, err := client.User.Get(ctx, user.ID().UUID())
	if err != nil {
		return err
	}
	_, err = u.Update().
		SetName(user.Name()).
		SetBirthday(user.Birthday()).
		Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Delete is a method to delete user by id
func (r *Repository) Delete(ctx context.Context, userID idVo.UserID) error {
	client := r.dbHandler.GetClient()

	return client.User.DeleteOneID(userID.UUID()).Exec(ctx)
}
