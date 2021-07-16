package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func UserController() Controller {
	return &User{}
}
func (user *User) Index(c *gin.Context) {
	c.String(http.StatusOK, "User Index")
}
func (user *User) Create(c *gin.Context) {
	c.String(http.StatusOK, "User Create")
}
func (user *User) Store(c *gin.Context) {
	c.String(http.StatusOK, "User Store")
}
func (user *User) Show(c *gin.Context) {
	c.String(http.StatusOK, "User Show "+c.Param("id"))
}
func (user *User) Edit(c *gin.Context) {
	c.String(http.StatusOK, "User Edit "+c.Param("id"))
}
func (user *User) Update(c *gin.Context) {
	c.String(http.StatusOK, "User Update "+c.Param("id"))
}
func (user *User) Destroy(c *gin.Context) {
	c.String(http.StatusOK, "User Destroy "+c.Param("id"))
}
