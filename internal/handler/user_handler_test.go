package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/mocks"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_Profile(t *testing.T) {
	userService := mocks.NewUserService(t)
	authService := mocks.NewAuthService(t)
	walletService := mocks.NewWalletService(t)
	transactionService := mocks.NewTransactionService(t)
	jwtService := mocks.NewJWTService(t)

	h := NewHandler(&HandlerConfig{
		UserService:        userService,
		AuthService:        authService,
		WalletService:      walletService,
		TransactionService: transactionService,
		JWTService:         jwtService,
	})

	t.Run("test success get user details", func(t *testing.T) {
		walletService.Mock.On("GetWalletByUserId", &dto.WalletRequestBody{UserID: 1}).Return(&model.Wallet{ID: 1, UserID: 1, Number: "100001", Balance: 0}, nil).Once()

		r := gin.Default()
		endpoint := "/profiles"
		r.GET(endpoint, MiddlewareMockUser, h.Profile)
		formattedUser := &dto.UserDetailResponse{}

		req, _ := http.NewRequest(
			http.MethodGet,
			endpoint,
			MakeRequestBody(formattedUser),
		)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var response utils.SResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		expectedResponse := utils.SResponse{
			Meta: utils.Meta{
				Message: "show profile success",
				Code:    http.StatusOK,
				Status:  "success",
			},
			Data: map[string]interface{}{
				"id":    float64(1),
				"name":  "nabil",
				"email": "nabil@user.com",
				"wallet": map[string]interface{}{
					"id":      float64(1),
					"number":  "100001",
					"balance": float64(0),
				},
			},
		}

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, expectedResponse, response)
	})

	t.Run("test error get user details", func(t *testing.T) {
		walletService.Mock.On("GetWalletByUserId", &dto.WalletRequestBody{UserID: 1}).Return(&model.Wallet{}, errors.New("something went wrong")).Once()

		r := gin.Default()
		endpoint := "/profiles"
		r.GET(endpoint, MiddlewareMockUser, h.Profile)
		formattedUser := &dto.UserDetailResponse{}

		req, _ := http.NewRequest(
			http.MethodGet,
			endpoint,
			MakeRequestBody(formattedUser),
		)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var response utils.EResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		expectedResponse := utils.EResponse{
			Meta: utils.Meta{
				Message: "show profile failed",
				Code:    http.StatusInternalServerError,
				Status:  "error",
			},
			Error: "something went wrong",
		}

		assert.Equal(t, 500, w.Code)
		assert.Equal(t, expectedResponse, response)
	})
}
