package http

import (
	"lotesaleagent/auth"
	"lotesaleagent/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService auth.AuthService
}

func (authHandler *AuthHandler) Register(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := authHandler.AuthService.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.WrapError()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})

}
func (authHandler *AuthHandler) Login(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := authHandler.AuthService.Login(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.WrapError()})
		return
	}
	c.JSON(http.StatusCreated, token)
}
