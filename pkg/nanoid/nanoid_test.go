package nanoid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNnanoID(t *testing.T) {
	nid := NanoID()
	assert.NotEmpty(t, nid)
	assert.Equal(t, 11, len(nid))

	nid = NanoID(100)
	assert.NotEmpty(t, nid)
	assert.Equal(t, 100, len(nid))
}
