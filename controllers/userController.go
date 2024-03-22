package controllers

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/i-am-sergio/golang-rest-api/db"
	"github.com/i-am-sergio/golang-rest-api/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	// Lógica para obtener los usuarios
	return c.String(http.StatusOK, "Get all users")
}

// Handler para la ruta "/users/:id"
func GetUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	var user models.User
	db.DB.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

// Handler para la ruta "/users"
func SaveUser(c echo.Context) error {

	if c.Request().Body == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Empty request body"})
	}

	newUser := new(models.User)

	// Decode JSON to struct
	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	// Verify if the JSON is empty
	if reflect.DeepEqual(*newUser, models.User{}) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON format"})
	}
	// save the new user
	if result := db.DB.Create(&newUser); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}
	// returh the new user and 201 status code
	return c.JSON(http.StatusCreated, newUser)
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
