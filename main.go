package main

import (
	"net/http"

	"github.com/i-am-sergio/golang-rest-api/db"
	"github.com/i-am-sergio/golang-rest-api/models"
	"github.com/i-am-sergio/golang-rest-api/routes"
	"github.com/labstack/echo/v4"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(&models.User{})

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.UserRoutes(e)

	e.Logger.Fatal(e.Start(":5000"))

}
