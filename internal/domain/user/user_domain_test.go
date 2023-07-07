package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	user, err := New(1, "test")
	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestUser_ID(t *testing.T) {
	user := &User{
		id:   1,
		name: "test",
	}
	assert.Equal(t, uint(1), user.ID())
}

func TestUser_Name(t *testing.T) {
	user := &User{
		id:   1,
		name: "test",
	}
	assert.Equal(t, "test", user.Name())
}
