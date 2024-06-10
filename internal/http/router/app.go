package router

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/http/handler"
	"github.com/bloomingbug/depublic/pkg/route"
)

const (
	Administrator = "Administrator"
	Buyer         = "Buyer"
	Guest         = "Guest"
)

var (
	allRoles  = []string{Administrator, Buyer}
	onlyAdmin = []string{Administrator}
	onlyBuyer = []string{Buyer}
	onlyGuest = []string{Guest}
)

func AppPublicRoutes(h handler.HelloHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/public",
			Handler: h.Say,
		},
	}
}

func AppPrivateRoutes(h handler.HelloHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/private",
			Handler: h.Say,
			Roles:   allRoles,
		},
	}
}
