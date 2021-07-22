package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	mid := Logger()
	assert.NotNil(t, mid)

	// NOTE: this should test with server, but i'm lazy
}
