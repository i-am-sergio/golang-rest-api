package routes

import (
	"github.com/i-am-sergio/golang-rest-api/controllers"
	"github.com/labstack/echo/v4"
)

// UserRoutes function
func UserRoutes(e *echo.Echo) {

	users := e.Group("/api/v1/users")

	users.GET("", controllers.GetUsers)
	users.GET("/:id", controllers.GetUserById)
	users.POST("", controllers.SaveUser)
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
}
