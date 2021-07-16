package routes

import (
	"net/http"

	"github.com/johnkhw/ecommerce/app/http/controllers"

	"github.com/gin-gonic/gin"
)

func RouteInit() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	user := r.Group("/user")
	{
		controllers.ResourceController(user, controllers.NewUserController())
		user.GET("/test", controllers.NewUserController().Test)
	}

	return r
}
