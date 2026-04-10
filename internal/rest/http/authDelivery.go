package http

import (
	"fmt"
	"lotesaleagent/auth"
	"lotesaleagent/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService auth.Service
}

func (authHandler *AuthHandler) Register(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		var errWrap = model.NewError(err)
		fmt.Println(errWrap)
		c.JSON(http.StatusAccepted, model.Response{
			Status:  "failed",
			Message: "Body?",
		})
		return
	}
	err := authHandler.AuthService.Register(&user)
	if err != nil {
		fmt.Println(err.String())
		c.JSON(http.StatusAccepted, model.Response{
			Status:  "failed",
			Message: "Register failed",
		})
		return
	}
	c.JSON(http.StatusCreated, model.Response{
		Status:  "success",
		Message: "Register success",
	})

}
func (authHandler *AuthHandler) Login(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		var errWrap = model.NewError(err)
		fmt.Println(errWrap)
		c.JSON(http.StatusAccepted, model.Response{
			Status:  "failed",
			Message: "Body?",
		})
		return
	}
	token, err := authHandler.AuthService.Login(&user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusAccepted, model.Response{
			Status:  "failed",
			Message: "Login failed",
		})
		return
	}
	c.JSON(http.StatusCreated, token)
}
