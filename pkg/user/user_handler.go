package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	// User defines the type for json operations.
	User struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Cellphone string `json:"cellphone"`
	}

	handler struct{ db map[int]*User }
)

var (
	// Handler handles operation methods.
	h handler
)

//*----------*//
//* Handlers *//
//*----------*//

// CreateOne User
func (h *handler) CreateOne(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

// GetOne User
func (h *handler) GetOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := h.db[id]

	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

// UpdateOne User
func (h *handler) UpdateOne(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	h.db[id].Name = u.Name

	return c.JSON(http.StatusOK, h.db[id])
}

// DeleteOne User
func (h *handler) DeleteOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(h.db, id)

	return c.NoContent(http.StatusNoContent)
}

// GetAll Users
func (h *handler) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, h.db)
}
