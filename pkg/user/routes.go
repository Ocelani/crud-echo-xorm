package user

import "github.com/labstack/echo/v4"

// Routes defines the URL routing for User operations.
func Routes(e *echo.Echo) {
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
}
