package route

import s "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/service"

type Router struct {
	userService s.UserService
	jwtService  s.JWTService
}

type RouterConfig struct {
	UserService s.UserService
	JWTService  s.JWTService
}

func NewRouter(c *RouterConfig) *Router {
	return &Router{
		userService: c.UserService,
		jwtService:  c.JWTService,
	}
}
