package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

type IUserController interface {
	IResourceController
	Test(*gin.Context)
}

func NewUserController() IUserController {
	return &UserController{}
}
func (user *UserController) Index(c *gin.Context) {
	c.String(http.StatusOK, "User Index")
}
func (user *UserController) Create(c *gin.Context) {
	c.String(http.StatusOK, "User Create")
}
func (user *UserController) Store(c *gin.Context) {
	c.String(http.StatusOK, "User Store")
}
func (user *UserController) Show(c *gin.Context) {
	c.String(http.StatusOK, "User Show "+c.Param("id"))
}
func (user *UserController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "User Edit "+c.Param("id"))
}
func (user *UserController) Update(c *gin.Context) {
	c.String(http.StatusOK, "User Update "+c.Param("id"))
}
func (user *UserController) Destroy(c *gin.Context) {
	c.String(http.StatusOK, "User Destroy "+c.Param("id"))
}
func (user *UserController) Test(c *gin.Context) {
	c.String(http.StatusOK, "This is test")
}
