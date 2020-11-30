package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON  = `{"id":1,"name":"Rob Pike","cellphone":"5541954122723"}`
	usersTest = handler{
		db: map[int]*User{
			1: {1, "Rob Pike", "5541954122723"},
		},
	}
)

func Test_Create(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, usersTest.CreateOne(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}

func Test_GetOne(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, usersTest.GetOne(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}
