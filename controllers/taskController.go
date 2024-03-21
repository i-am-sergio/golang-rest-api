package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {
	// Lógica para obtener los usuarios
	return c.String(http.StatusOK, "Get all tasks")
}

// Handler para la ruta "/tasks/:id"
func GetTask(c echo.Context) error {
	id := c.Param("id")
	// Lógica para obtener el usuario con el ID proporcionado
	return c.String(http.StatusOK, "Get task with ID "+id)
}

// Handler para la ruta "/tasks"
func SaveTask(c echo.Context) error {
	// Lógica para guardar el usuario
	return c.String(http.StatusCreated, "Save task")
}

// Handler para la ruta "/tasks/:id"
func UpdateTask(c echo.Context) error {
	id := c.Param("id")
	// Lógica para actualizar el usuario con el ID proporcionado
	return c.String(http.StatusOK, "Update task with ID "+id)
}

// Handler para la ruta "/tasks/:id"
func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	// Lógica para eliminar el usuario con el ID proporcionado
	return c.String(http.StatusOK, "Delete task with ID "+id)
}
