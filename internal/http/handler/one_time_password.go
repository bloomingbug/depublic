package handler

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/validator"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
)

type OneTimePasswordHandler struct {
	otpService service.OneTimePasswordService
}

func (h *OneTimePasswordHandler) Generate(c echo.Context) error {
	input := new(binder.GenerateOTPRequest)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, validator.ValidatorErrors(err)))
	}

	otp, err := h.otpService.Generate(c.Request().Context(), input.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "berhasil mengirim otp ke email", map[string]string{
		"email": otp.Email,
	}))
}

func NewOneTimePasswordHandler(otpService service.OneTimePasswordService) OneTimePasswordHandler {
	return OneTimePasswordHandler{otpService: otpService}
}
