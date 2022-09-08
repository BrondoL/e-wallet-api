package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

type JWTSConfig struct {
}

func NewJWTService(c *JWTSConfig) JWTService {
	return &jwtService{}
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))
var JWT_TTL, _ = strconv.Atoi(os.Getenv("JWT_TTL"))
var ISSUER = os.Getenv("JWT_ISSUER")

type idTokenClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
}

func (s *jwtService) GenerateToken(userID int) (string, error) {
	payload := idTokenClaims{}
	payload.ExpiresAt = &jwt.NumericDate{Time: time.Now().Add(time.Minute * time.Duration(JWT_TTL))}
	payload.IssuedAt = &jwt.NumericDate{Time: time.Now()}
	payload.Issuer = ISSUER
	payload.UserID = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
