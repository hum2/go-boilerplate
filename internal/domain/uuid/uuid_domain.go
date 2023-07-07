package uuid

import (
	"github.com/google/uuid"
)

// UUID is a domain model of uuid.
type UUID string

// New is a constructor of UUID.
func New() (UUID, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	uuid := u.String()
	return UUID(uuid), nil
}
