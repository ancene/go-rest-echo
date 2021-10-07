package root_test

import (
	"path/filepath"
	"testing"

	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/ancene/go-rest-echo/internal/infrastructure/handler"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RootSuite struct {
	suite.Suite
	configuration *domain.Configuration
	handler       *handler.RootHanlder
}

// SetupSuite it will run first of test
func (rs *RootSuite) SetupSuite() {
	// setup suite
	// like create new handler, etc.
	rs.T().Log("start", 10)
}

// TearDownSuite it will run last of test
func (rs *RootSuite) TearDownSuite() {
	// tear down suite
	// like close resource
	rs.T().Log("end", 0)
}

// TestRootSuites is orkestrator for this suite test
// like func main(){} or func init(){}
func TestRootSuites(t *testing.T) {
	// skip this suite if testing command containe -short or
	// if you wan to only run unit-test
	if testing.Short() {
		t.Skip("Skip test for root suite")
	}

	// setup env
	envPath, err := filepath.Abs("../../.env.test")
	assert.NoError(t, err)
	assert.NoError(t, godotenv.Load(envPath))

	// setup configuration
	conf := domain.NewConfiguration()

	// setup handler
	hand := handler.NewRootHandler(conf)

	// setup suite
	rootSuite := &RootSuite{
		configuration: conf,
		handler:       hand,
	}

	// run suite test
	suite.Run(t, rootSuite)
}

// actual test of application behavior
func (rs *RootSuite) TestConfiguration() {
	rs.Assert().NotNil(rs.configuration)
}

// actual test of application behavior
func (rs *RootSuite) TestHandlerNotNil() {
	rs.Assert().NotNil(rs.handler)
}
