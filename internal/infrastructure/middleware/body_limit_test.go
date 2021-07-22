package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBodyLimit(t *testing.T) {
	mid := BodyLimit()
	assert.NotNil(t, mid)

	// NOTE: this should test with server, but i'm lazy
}
