package handler

import (
	"fmt"
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
	var input binder.RegisterRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, form_validator.ValidatorErrors(err)))
	}

	fmt.Println(input.Name)

	birthdate, err := time.Parse("2006-01-02", input.Birthdate)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid birthdate format"})
	}

	userNew := entity.NewUser(input.Name, input.Email, input.Password, input.Phone, input.Address, input.Avatar, &birthdate, entity.Gender(input.Gender), entity.Role(entity.Buyer))

	user, err := h.userService.UserRegistration(c, input.Token, input.Email, userNew)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "Success", echo.Map{
		"name":  user.Name,
		"email": user.Email,
	}))
}

func (h *UserHandler) Login(c echo.Context) error {
	var input binder.LoginRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, form_validator.ValidatorErrors(err)))
	}

	token, err := h.userService.Login(c.Request().Context(), input.Email, input.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Success(http.StatusOK, true, "login success", echo.Map{"token": token}))
}

func NewUserHandler(userService service.UserService) UserHandler {
	return UserHandler{userService}
}
