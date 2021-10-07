package connection

import (
	"path/filepath"
	"testing"

	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestNewPostgresConnection(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		envPath, err := filepath.Abs("../../../.env.test")
		assert.NoError(t, err)
		assert.NoError(t, godotenv.Load(envPath))

		conf := domain.NewConfiguration()
		db, err := NewPostgresConnection(conf)
		assert.NoError(t, err)
		assert.NotNil(t, db)
	})

	t.Run("error configuration nil", func(t *testing.T) {
		db, err := NewPostgresConnection(nil)
		assert.Error(t, err)
		assert.Nil(t, db)
	})

	t.Run("error failed open to connection", func(t *testing.T) {
		conf := &domain.Configuration{}
		conf.Postgres.Host = "\"/\\"
		conf.Postgres.Port = 9_223_371_999_999_999_999
		conf.Postgres.Username = "=/==/="
		conf.Postgres.Password = "=/==/="
		conf.Postgres.Database = "=/==/= sslmode=true"
		db, err := NewPostgresConnection(conf)
		assert.Error(t, err)
		assert.Nil(t, db)
	})

	t.Run("error can't connect to database", func(t *testing.T) {
		conf := &domain.Configuration{}
		db, err := NewPostgresConnection(conf)
		assert.Error(t, err)
		assert.Nil(t, db)
	})
}
