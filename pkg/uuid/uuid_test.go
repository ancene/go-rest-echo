package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	uid := UUID()
	assert.NotEmpty(t, uid)
	assert.Equal(t, 36, len(uid))
}
