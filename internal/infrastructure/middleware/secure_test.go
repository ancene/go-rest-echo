package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSecure(t *testing.T) {
	mid := Secure()
	assert.NotNil(t, mid)

	// NOTE: this should test with server, but i'm lazy
}
