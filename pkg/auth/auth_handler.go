package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	// admin is reponsible for API authentication.
	auth map[string]string // map[username]password
)

var (
	// map[username]password
	admin = map[string]string{
		"macapa":  "PASSmacapa",
		"varejao": "PASSvarejao",
	}

	// IsLoggedIn is used as a JWT middleware.
	IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	})
)

//*----------*//
//* Handlers *//
//*----------*//

// Login route.
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if password != admin[username] {
		return echo.ErrUnauthorized
	}

	jwt, err := generateToken(username)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": jwt,
	})
}

// Accessible is an unauthenticated route.
func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Restricted group is for authorized users.
func Restricted(c echo.Context) error {
	user := c.Get("username").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["username"].(string)

	return c.String(http.StatusOK, "Welcome "+name+"!")
}

//*---------*//
//* Private *//
//*---------*//

func generateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "error", err
	}

	return t, nil
}
