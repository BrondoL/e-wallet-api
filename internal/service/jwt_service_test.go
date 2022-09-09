package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJWTService(t *testing.T) {
	var jwtService = NewJWTService(&JWTSConfig{})

	assert.NotNil(t, jwtService)
}
