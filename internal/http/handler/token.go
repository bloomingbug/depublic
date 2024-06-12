package handler

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/validator"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
)

type TokenHandler struct {
	tokenService service.TokenService
}

func (h *TokenHandler) Generate(c echo.Context) error {
	input := new(binder.VerifyOTPRequest)
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, validator.ValidatorErrors(err)))
	}

	token, err := h.tokenService.GenerateTokenRegistration(c.Request().Context(), input.OTPCode, input.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "Success", token))
}

func NewTokenHandler(tokenService service.TokenService) TokenHandler {
	return TokenHandler{
		tokenService: tokenService,
	}
}
