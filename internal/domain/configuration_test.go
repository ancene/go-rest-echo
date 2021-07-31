package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	conf := NewConfiguration()
	assert.NotNil(t, conf)

	// app
	assert.Empty(t, conf.App.Host)
	assert.Equal(t, 8080, conf.App.Port)
	assert.Empty(t, conf.App.Name)
	assert.Empty(t, conf.App.Timezone)

	// pg
	assert.Empty(t, conf.Postgres.Host)
	assert.Equal(t, 5432, conf.Postgres.Port)
	assert.Empty(t, conf.Postgres.Username)
	assert.Empty(t, conf.Postgres.Password)
	assert.Empty(t, conf.Postgres.Database)

	// mysql
	assert.Empty(t, conf.Mysql.Host)
	assert.Equal(t, 3306, conf.Mysql.Port)
	assert.Empty(t, conf.Mysql.Username)
	assert.Empty(t, conf.Mysql.Password)
	assert.Empty(t, conf.Mysql.Database)
}
