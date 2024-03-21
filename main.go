package main

import (
	"net/http"

	"github.com/i-am-sergio/golang-rest-api/controllers"
	"github.com/i-am-sergio/golang-rest-api/db"
	"github.com/labstack/echo/v4"
)

func concat(strings ...string) string {
	var result string
	for _, s := range strings {
		result += s
	}
	return result
}

func main() {

	db.DBConnection()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routeTask := "/tasks"

	e.GET(concat(routeTask), controllers.GetTasks)
	e.GET(concat(routeTask, "/:id"), controllers.GetTask)
	e.POST(concat(routeTask), controllers.SaveTask)
	e.PUT(concat(routeTask, "/:id"), controllers.UpdateTask)
	e.DELETE(concat(routeTask, "/:id"), controllers.DeleteTask)

	routeUser := "/users"

	e.GET(concat(routeUser), controllers.GetUsers)
	e.GET(concat(routeUser, "/:id"), controllers.GetUser)
	e.POST(concat(routeUser), controllers.SaveUser)
	e.PUT(concat(routeUser, "/:id"), controllers.UpdateUser)
	e.DELETE(concat(routeUser, "/:id"), controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":5000"))

}
