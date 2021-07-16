package main

import (
	"github.com/johnkhw/ecommerce/routes"
)

/* func resourceController(m models, r *gin.Engine) {
	group := r.Group("/" + m.name)
	{
		group.GET("/", func(c *gin.Context) {

		})
		//Get by Primary Key
		group.GET("/:id", func(c *gin.Context) {

		})
		//Create New
		group.POST("/", func(c *gin.Context) {

		})
		//Update by Primary Key
		group.POST("/", func(c *gin.Context) {

		})
		//Delete by Primary Key
		group.DELETE("/:id", func(c *gin.Context) {

		})
	}
} */

func main() {
	r := routes.RouteInit()

	r.Run(":3001") // listen and serve on 0.0.0.0:8080
}
