package route

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/handler"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) Transaction(route *gin.RouterGroup, h *handler.Handler) {
	route.Use(middleware.AuthMiddleware(r.jwtService, r.userService))
	route.GET("/transactions", h.GetTransactions)
	route.POST("/top-up", h.TopUp)
	route.POST("/transfer", h.Transfer)
}
