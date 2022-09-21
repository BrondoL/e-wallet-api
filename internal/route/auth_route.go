package route

import (
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func (r *Router) Auth(route *gin.RouterGroup, h *handler.Handler) {
	route.POST("/sign-up", h.Register)
	route.POST("/sign-in", h.Login)
	route.POST("/forgot-password", h.ForgotPassword)
	route.POST("/reset-password", h.ResetPassword)
}
