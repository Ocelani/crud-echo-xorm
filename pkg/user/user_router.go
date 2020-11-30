package user

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Routes defines the URL routing for User operations.
func Routes(e *echo.Echo) {
	r := e.Group("/users")
	r.Use(middleware.JWT([]byte("secret")))

	r.POST("", h.CreateOne)
	r.GET("", h.GetAll)
	r.GET("/:id", h.GetOne)
	r.PUT("/:id", h.UpdateOne)
	r.DELETE("/:id", h.DeleteOne)
}
