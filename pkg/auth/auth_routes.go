package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Routes defines the URL routing for Authentication middleware.
func Routes(e *echo.Echo) {
	e.POST("/login", login)     // Login route
	e.GET("/", accessible)      // Unauthenticated route
	r := e.Group("/restricted") // Restricted group

	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)
}
