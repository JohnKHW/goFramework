package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Index(*gin.Context)
	Create(*gin.Context)
	Store(*gin.Context)
	Show(*gin.Context)
	Edit(*gin.Context)
	Update(*gin.Context)
	Destroy(*gin.Context)
}

func ResourceController(g *gin.RouterGroup, c Controller) {
	//Index
	g.GET("/", c.Index)
	//Create New
	g.GET("/create", c.Create)
	//Store New
	g.POST("/", c.Store)
	//Show by Primary Key
	g.GET("/:id", c.Show)
	//Edit by Primary Key
	g.GET("/:id/edit", c.Edit)
	//Update by Primary Key
	g.PUT("/:id", c.Update)
	//Delete by Primary Key
	g.DELETE("/:id", c.Destroy)
}
