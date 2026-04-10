package http

import (
	"fmt"
	"lotesaleagent/model"
	"lotesaleagent/model/token"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		authHeader := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(
				http.StatusUnauthorized, model.Response{
					Status:  "failed",
					Message: fmt.Sprintf("Authorization header format must be Bearer {token}"),
				})
			c.Abort()
			return
		}
		accessToken := strings.TrimSpace(authHeader[7:])
		var valid, payload = token.VerifyAndDecode(accessToken)
		if valid == false {
			c.JSON(http.StatusUnauthorized, model.Response{
				Status:  "failed",
				Message: fmt.Sprintf("Invalid token"),
			})
			c.Abort()
			return
		}
		if payload["exp"].(float64) < float64(time.Now().Unix()) {
			c.JSON(http.StatusUnauthorized, model.Response{
				Status:  "failed",
				Message: fmt.Sprintf("token expired"),
			})
			c.Abort()
			return
		}
		userId := payload["id"].(string)
		c.Set("userId", userId)
		c.Next()
	}
}
