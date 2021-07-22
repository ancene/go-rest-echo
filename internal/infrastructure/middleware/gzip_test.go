package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGzip(t *testing.T) {
	mid := Gzip()
	assert.NotNil(t, mid)

	// NOTE: this should test with server, but i'm lazy
}
