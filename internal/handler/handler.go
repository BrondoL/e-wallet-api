package handler

import s "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/service"

type Handler struct {
	userService s.UserService
	authService s.AuthService
	jwtService  s.JWTService
}

type HandlerConfig struct {
	UserService s.UserService
	AuthService s.AuthService
	JWTService  s.JWTService
}

func NewHandler(c *HandlerConfig) *Handler {
	return &Handler{
		userService: c.UserService,
		authService: c.AuthService,
		jwtService:  c.JWTService,
	}
}
