package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"

	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/form_validator"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService service.UserService
}

func (h *UserHandler) Registration(c echo.Context) error {
	req := new(binder.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			form_validator.ValidatorErrors(err),
		))
	}

	var birthdate *time.Time

	if req.Birthdate != nil {
		parsedDate, err := time.Parse("2006-01-02", *req.Birthdate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid birthdate format"})
		}
		birthdate = &parsedDate
	} else {
		birthdate = nil
	}

	userNew := entity.NewUser(req.Name, req.Email, req.Password, req.Phone, req.Address, req.Avatar, birthdate, req.Gender, entity.Buyer)

	user, err := h.userService.UserRegistration(c, req.Token, userNew)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "Success", echo.Map{
		"name":  user.Name,
		"email": user.Email,
	}))
}

func (h *UserHandler) Login(c echo.Context) error {
	req := new(binder.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			form_validator.ValidatorErrors(err)))
	}

	token, err := h.userService.Login(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "login success", echo.Map{"token": token}))
}

func (h *UserHandler) ResetPassword(c echo.Context) error {
	req := new(binder.ResetPasswordRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(
			http.StatusBadRequest,
			false,
			form_validator.ValidatorErrors(err),
		))
	}
	err := h.userService.ChangePassword(c.Request().Context(), req.Token, req.Password)
	if err != nil {
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, response.Error(
				http.StatusInternalServerError,
				false,
				err.Error()))
		}
	}

	return c.JSON(http.StatusOK, response.Success(
		http.StatusOK,
		false,
		"success",
		nil))
}

func (h *UserHandler) Profile(c echo.Context) error {
	userAuth, _ := c.Get("user").(*jwt.Token)

	user, err := h.userService.GetProfile(c.Request().Context(), userAuth)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "berhasil mendapatkan profile user", user))
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService}
}
