package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORS(t *testing.T) {
	mid := CORS()
	assert.NotNil(t, mid)

	// NOTE: this should test with server, but i'm lazy
}
