package connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMysqlConnection(t *testing.T) {
	db, err := NewMysqlConnection()
	assert.Nil(t, db)
	assert.Nil(t, err)
}
