package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Register(c *gin.Context) {
	input := &dto.RegisterRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("register failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.CreateUser(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("register failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	inputWallet := &dto.WalletRequestBody{}
	inputWallet.UserID = int(newUser.ID)
	newWallet, err := h.walletService.CreateWallet(inputWallet)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("register failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	token, err := h.jwtService.GenerateToken(int(newUser.ID))
	if err != nil {
		response := utils.ErrorResponse("register failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(newUser, newWallet, token)
	response := utils.SuccessResponse("register success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) Login(c *gin.Context) {
	input := &dto.LoginRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("login failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.authService.Attempt(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("login failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	inputWallet := &dto.WalletRequestBody{}
	inputWallet.UserID = int(loggedinUser.ID)
	wallet, err := h.walletService.GetWalletByUserId(inputWallet)
	if err != nil {
		response := utils.ErrorResponse("login failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := h.jwtService.GenerateToken(int(loggedinUser.ID))
	if err != nil {
		response := utils.ErrorResponse("login failed", http.StatusInternalServerError, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formattedLogin := dto.FormatLogin(loggedinUser, wallet, token)
	response := utils.SuccessResponse("login success", http.StatusOK, formattedLogin)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	input := &dto.ForgotPasswordRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("forgot password failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	passwordReset, err := h.authService.ForgotPass(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("forgot password failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedPasswordReset := dto.FormatForgotPassword(passwordReset)
	response := utils.SuccessResponse("forgot password success", http.StatusOK, formattedPasswordReset)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) ResetPassword(c *gin.Context) {
	input := &dto.ResetPasswordRequestBody{}

	err := c.ShouldBindJSON(input)
	if err != nil {
		errors := utils.FormatValidationError(err)
		response := utils.ErrorResponse("reset password failed", http.StatusUnprocessableEntity, errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	passwordReset, err := h.authService.ResetPass(input)
	if err != nil {
		statusCode := utils.GetStatusCode(err)
		response := utils.ErrorResponse("reset password failed", statusCode, err.Error())
		c.JSON(statusCode, response)
		return
	}

	formattedUser := dto.FormatUser(&passwordReset.User)
	response := utils.SuccessResponse("reset password success", http.StatusOK, formattedUser)
	c.JSON(http.StatusOK, response)
}
