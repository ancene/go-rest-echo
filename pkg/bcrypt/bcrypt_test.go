package bcrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	hash, err := Hash("password")
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestCompare(t *testing.T) {
	isValid := Compare("password", "$2a$14$chAbAunvqIMLCD4bWE1g0O1Nz4QvMUL6L5xmqSEWrjlS.Lva.tDwS")
	assert.True(t, isValid)
}
