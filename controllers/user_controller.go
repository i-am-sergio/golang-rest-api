package controllers

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/i-am-sergio/golang-rest-api/db"
	"github.com/i-am-sergio/golang-rest-api/models"
	"github.com/labstack/echo/v4"
)

var errorMessage string = "Internal Server Error"
var invalidIdMessage string = "Invalid ID"

// Handler for the route GET "/users"
func GetUsers(c echo.Context) error {
	var users []models.User
	result := db.DB.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusOK, users)
}

// Handler for the route GET "/users/:id"
func GetUserById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidIdMessage})
	}
	var user models.User
	db.DB.First(&user, id)
	return c.JSON(http.StatusOK, user)
}

// Handler for the route POST "/users"
func SaveUser(c echo.Context) error {
	var newUser models.User
	err := c.Bind(&newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	// validate the user
	if err := validator.New().Struct(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	// save the new user
	result := db.DB.Create(&newUser)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusCreated, newUser)
}

// Handler para la ruta PUT "/users/:id"
func UpdateUser(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": invalidIdMessage})
	}
	var user models.User
	// Find the user with the ID provided
	userFound := db.DB.First(&user, id)
	if userFound.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	// Bind the updated user data from the request
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}
	// Update user in the database
	result := db.DB.Model(&user).Updates(user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusOK, user)
}

// Handler para la ruta DELETE "/users/:id"
func DeleteUser(c echo.Context) error {
	paramId := c.Param("id")
	id, err := strconv.ParseUint(paramId, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	var user models.User
	// Find the user with the ID provided
	userFound := db.DB.First(&user, id)
	if userFound.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	// Delete user from the database
	result := db.DB.Delete(&user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": errorMessage})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
