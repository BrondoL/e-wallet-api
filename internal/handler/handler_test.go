package handler

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		c *HandlerConfig
	}
	tests := []struct {
		name string
		args args
		want *Handler
	}{
		{
			name: "Test new handler",
			args: args{
				c: &HandlerConfig{
					UserService:        mocks.NewUserService(t),
					AuthService:        mocks.NewAuthService(t),
					WalletService:      mocks.NewWalletService(t),
					TransactionService: mocks.NewTransactionService(t),
					JWTService:         mocks.NewJWTService(t),
				},
			},
			want: NewHandler(&HandlerConfig{
				UserService:        mocks.NewUserService(t),
				AuthService:        mocks.NewAuthService(t),
				WalletService:      mocks.NewWalletService(t),
				TransactionService: mocks.NewTransactionService(t),
				JWTService:         mocks.NewJWTService(t),
			}),
		},
		{
			name: "Test nill handler",
			args: args{
				c: &HandlerConfig{},
			},
			want: NewHandler(&HandlerConfig{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewHandler(tt.args.c), "NewHandler(%v)", tt.args.c)
		})
	}
}
