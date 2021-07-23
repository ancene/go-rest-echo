package infrastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	srv := NewServer(nil, nil)
	assert.NotNil(t, srv)
}
