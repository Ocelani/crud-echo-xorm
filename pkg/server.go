package pkg

import (
	"github.com/Ocelani/mercafacil/pkg/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server provides the application services.
func Server() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start("localhost:1323"))

	user.Routes(e)
}
