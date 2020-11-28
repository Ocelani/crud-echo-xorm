package user

import "github.com/labstack/echo/v4"

// Routes defines the URL routing for User operations.
func Routes(e *echo.Echo) {
	e.GET("/users", users.getAll)
	e.GET("/users/:id", users.getOne)
	e.POST("/users", users.createOne)
	e.PUT("/users/:id", users.updateOne)
	e.DELETE("/users/:id", users.deleteOne)
}
