package pkg

import (
	"github.com/Ocelani/mercafacil/internal"
	"github.com/Ocelani/mercafacil/pkg/auth"
	"github.com/Ocelani/mercafacil/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Server provides the application services.
func Server() {
	internal.Databases()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	user.Routes(e)
	auth.Routes(e)

	e.Logger.Fatal(e.Start("localhost:1323"))
}
