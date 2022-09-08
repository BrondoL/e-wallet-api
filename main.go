package main

import (
	"fmt"
	"os"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/config"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/handler"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/repository"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/route"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.GetConn()

	userRepository := repository.NewUserRepository(&repository.URConfig{DB: db})
	passwordResetRepository := repository.NewPasswordResetRepository(&repository.PRConfig{DB: db})

	userService := service.NewUserService(&service.USConfig{UserRepository: userRepository})
	authService := service.NewAuthService(&service.ASConfig{UserRepository: userRepository, PasswordResetRepository: passwordResetRepository})
	jwtService := service.NewJWTService(&service.JWTSConfig{})

	h := handler.NewHandler(&handler.HandlerConfig{
		UserService: userService,
		AuthService: authService,
		JWTService:  jwtService,
	})

	routes := route.NewRouter(&route.RouterConfig{})

	router := gin.Default()
	router.NoRoute(h.NoRoute)

	version := os.Getenv("API_VERSION")
	api := router.Group(fmt.Sprintf("/api/%s", version))

	routes.Auth(api, h)

	router.Run(":8000")
}
