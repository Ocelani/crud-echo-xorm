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
	userJSON = `{"id":1,"name":"Jon Snow","cellphone":"5541954122723"}`
	mockDB   = handler{
		db: map[int]*User{
			1: {1, "Jon Snow", "5541954122723"},
		},
	}
	e = echo.New()
)

// var userJSON = func() (users []*user) {
// 	f, err := os.Open("./json/contacts-macapa.json")
// 	dec := json.NewDecoder(f)
// 	for dec.More() {
// 		var u user
// 		if err = dec.Decode(&u); err != nil {
// 			log.Fatalln(err)
// 		}
// 		users = append(users, &u)
// 	}
// 	return
// }

func Test_Create(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, users.createOne(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}

func Test_GetOne(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, users.getOne(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, strings.ReplaceAll(rec.Body.String(), "\n", ""))
	}
}
