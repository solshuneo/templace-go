package main

import (
	"lotesaleagent/auth"
	"lotesaleagent/internal/repository/gsql"
	"lotesaleagent/internal/rest/http"
	"lotesaleagent/model"

	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	userRepo := gsql.NewGormUserRepository(db)
	authService := auth.Service{
		UserInterface: userRepo,
	}
	authService.Register(&model.User{
		Username: "dbUser",
		Password: "dbPass",
	})
	authHandler := http.AuthHandler{
		AuthService: authService,
	}
	r := gin.Default()
	r.Use(http.AuthMiddleware())
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	_ = r.Run(":8080")
}
