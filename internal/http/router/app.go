package router

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/http/handler"
	"github.com/bloomingbug/depublic/pkg/route"
)

const (
	Administrator = "Administrator"
	Buyer         = "Buyer"
)

var (
	allRoles = []string{Administrator, Buyer}
	// onlyAdmin = []string{Administrator}
	// onlyBuyer = []string{Buyer}
)

func AppPublicRoutes(h map[string]interface{}) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/public",
			Handler: h["hello"].(*handler.HelloHandler).Say,
		},
		{
			Method:  http.MethodPost,
			Path:    "/request-otp",
			Handler: h["otp"].(*handler.OneTimePasswordHandler).Generate,
		},
		{
			Method:  http.MethodPost,
			Path:    "/verify-otp",
			Handler: h["token"].(*handler.TokenHandler).Generate,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/register",
			Handler: h["user"].(*handler.UserHandler).Registration,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: h["user"].(*handler.UserHandler).Login,
		},
	}
}

func AppPrivateRoutes(h map[string]interface{}) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/private",
			Handler: h["hello"].(*handler.HelloHandler).Say,
			Roles:   allRoles,
		},
	}
}
