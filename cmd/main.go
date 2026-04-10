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
	v1 := r.Group("/api/v1")

	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
		v1.POST("/me", http.AuthMiddleWare(), authHandler.Profile)
		v1.POST("/refreshtoken", authHandler.RefreshToken)
	}

	_ = r.Run(":8080")
}
