package handler

import (
	"net/http"

	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/labstack/echo/v4"
)

type HelloHandler struct{}

func (h *HelloHandler) Say(c echo.Context) error {
	return c.JSON(http.StatusOK, response.Success(http.StatusOK, "Hello", nil))
}

func NewHelloHandler() HelloHandler {
	return HelloHandler{}
}
