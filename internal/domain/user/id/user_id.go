package id

import "github.com/google/uuid"

type UserID uuid.UUID

// New is a constructor of UserID.
func New() UserID {
	return UserID(uuid.New())
}

// NewFromRepository is a constructor of UserID from repository.
func NewFromRepository(id uuid.UUID) UserID {
	return UserID(id)
}

// NewFromInput is a constructor of UserID from input.
// TODO validation
func NewFromInput(id string) UserID {
	return UserID(uuid.MustParse(id))
}

// UUID is a getter of uuid.UUID.
func (u UserID) UUID() uuid.UUID {
	return uuid.UUID(u)
}
