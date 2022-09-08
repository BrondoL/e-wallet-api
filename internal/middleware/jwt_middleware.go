package middleware

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	s "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/service"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(jwtService s.JWTService, userService s.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, "unrecognized token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) != 2 {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, "unrecognized token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		encodedToken := arrayToken[1]
		token, err := jwtService.ValidateToken(encodedToken)
		if err != nil {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, "not a valid bearer token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(payload["user_id"].(float64))

		params := &dto.UserRequestParams{}
		params.UserID = userID
		user, err := userService.GetUser(params)
		if err != nil {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
