package connection

import (
	"path/filepath"
	"testing"

	"github.com/ancene/go-rest-echo/internal/domain"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewMysqlConnection(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		envPath, err := filepath.Abs("../../../.env.test")
		assert.NoError(t, err)
		assert.NoError(t, godotenv.Load(envPath))

		conf := domain.NewConfiguration()
		db, err := NewMysqlConnection(conf)
		assert.NoError(t, err)
		assert.NotNil(t, db)
	})

	t.Run("error configuration nil", func(t *testing.T) {
		db, err := NewMysqlConnection(nil)
		assert.Error(t, err)
		assert.Nil(t, db)
	})

	t.Run("error can't connect to database", func(t *testing.T) {
		conf := &domain.Configuration{}
		db, err := NewMysqlConnection(conf)
		assert.Error(t, err)
		assert.Nil(t, db)
	})
}
