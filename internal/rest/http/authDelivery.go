package http

import (
	"fmt"
	"net/http"
	"template-go/auth"
	"template-go/model"

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
	c.JSON(http.StatusCreated, model.Response{
		Status:  "success",
		Message: "Login success",
		Content: token,
	})
}

func (authHandler *AuthHandler) Profile(c *gin.Context) {
	userId := c.GetString("userId")
	foundUser, err := authHandler.AuthService.GetProfileById(userId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: "Get profile failed",
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "Get profile success",
		Content: foundUser,
	})
}

func (authHandler *AuthHandler) RefreshToken(c *gin.Context) {
	authRefresh := c.Request.Header.Get("X-Fresh-Refresh-Token")
	token, err := authHandler.AuthService.RefreshToken(authRefresh)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusAccepted, model.Response{
			Status:  "failed",
			Message: "Refresh Token failed",
		})
	}
	c.JSON(http.StatusOK, model.Response{
		Status: "success",

		Message: "Refresh Token success",
		Content: token,
	})
}
