package route

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/handler"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) User(route *gin.RouterGroup, h *handler.Handler) {
	route.GET("/profiles", middleware.AuthMiddleware(r.jwtService, r.userService), h.Profile)
}
