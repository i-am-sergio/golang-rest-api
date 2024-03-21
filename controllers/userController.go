package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	// Lógica para obtener los usuarios
	return c.String(http.StatusOK, "Get all users")
}

// Handler para la ruta "/users/:id"
func GetUser(c echo.Context) error {
	id := c.Param("id")
	// Lógica para obtener el usuario con el ID proporcionado
	return c.String(http.StatusOK, "Get user with ID "+id)
}

// Handler para la ruta "/users"
func SaveUser(c echo.Context) error {
	// Lógica para guardar el usuario
	return c.String(http.StatusCreated, "Save user")
}

// Handler para la ruta "/users/:id"
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	// Lógica para actualizar el usuario con el ID proporcionado
	return c.String(http.StatusOK, "Update user with ID "+id)
}

// Handler para la ruta "/users/:id"
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	// Lógica para eliminar el usuario con el ID proporcionado
	return c.String(http.StatusOK, "Delete user with ID "+id)
}
