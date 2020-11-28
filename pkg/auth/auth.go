package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Login route
func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "rob" || password != "pike" {
		return echo.ErrUnauthorized // Throws unauthorized error
	}

	token := jwt.New(jwt.SigningMethodHS256) // Create token

	claims := token.Claims.(jwt.MapClaims) // Set claims
	claims["name"] = "Rob Pike"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// Unauthenticated route
func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Restricted group
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
