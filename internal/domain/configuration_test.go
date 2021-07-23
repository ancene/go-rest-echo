package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	conf := NewConfiguration()
	assert.NotNil(t, conf)
}
