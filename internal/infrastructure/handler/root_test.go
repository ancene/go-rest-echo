package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/ancene/go-rest-echo/internal/domain"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewRootHanlder(t *testing.T) {
	conf := &domain.Configuration{}
	hand := NewRootHandler(conf)
	assert.NotNil(t, hand)
}

func TestRoot_Home(t *testing.T) {
	// setup test server environment
	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	exp := "{\"success\":true,\"message\":\"welcome to ancene\",\"data\":[],\"error\":null}\n"
	conf := &domain.Configuration{}
	hand := NewRootHandler(conf)
	err := hand.Home(c)

	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, exp, rec.Body.String())
}

func TestRoot_Health(t *testing.T) {
	// setup test server environment
	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// testing
	exp := "{\"success\":true,\"message\":\" is healhty\",\"data\":{\"connection\":{\"mysql\":\"good\",\"postgres\":\"good\"},\"http\":\"good\",\"memory\":\"good\"},\"error\":null}\n"
	conf := &domain.Configuration{}
	hand := NewRootHandler(conf)
	err := hand.Health(c)

	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, exp, rec.Body.String())
}
