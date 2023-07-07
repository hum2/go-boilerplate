package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	uuid, err := New()
	assert.NoError(t, err)
	assert.NotNil(t, uuid)
}
