package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	User struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Cellphone string `json:"cellphone"`
	}
	handler struct {
		db map[int]*User
	}
)

var (
	users handler
	seq   = 1
)

//*----------*//
//* Handlers *//
//*----------*//

// create User
func (users *handler) createOne(c echo.Context) error {
	u := &User{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

// getOne User
func (users *handler) getOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := users.db[id]
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

// update User
func (users *handler) updateOne(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users.db[id].Name = u.Name
	return c.JSON(http.StatusOK, users.db[id])
}

// deleteOne User
func (users *handler) deleteOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(users.db, id)
	return c.NoContent(http.StatusNoContent)
}

// getAll Users
func (users *handler) getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, users.db)
}
