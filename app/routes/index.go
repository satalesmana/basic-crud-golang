package routes

import (
	"app-basic-crud/app/controller/users"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) *gin.Engine {
	userEndpoints := users.NewUserHandler()

	routers := route.Group("/users")
	{
		routers.GET("/", userEndpoints.GetAll)       //get all data
		routers.POST("/", userEndpoints.AddNew)      // add new
		routers.GET("/:id", userEndpoints.FindBy)    // find by id
		routers.PUT("/:id", userEndpoints.Edit)      // find by id
		routers.DELETE("/:id", userEndpoints.Delete) // find by id
	}

	return route
}
