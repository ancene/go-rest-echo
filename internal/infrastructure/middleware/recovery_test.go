package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	mid := Recovery()
	assert.NotNil(t, mid)

	// NOTE: this should test with server, but i'm lazy
}
