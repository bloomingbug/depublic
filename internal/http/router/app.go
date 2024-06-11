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
	allRoles = []string{Administrator, Buyer}
	// onlyAdmin = []string{Administrator}
	// onlyBuyer = []string{Buyer}
	onlyGuest = []string{Guest}
)

func AppPublicRoutes(h map[string]interface{}) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/public",
			Handler: h["hello"].(*handler.HelloHandler).Say,
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
		{
			Method:  http.MethodPost,
			Path:    "/request-otp",
			Handler: h["otp"].(*handler.OneTimePasswordHandler).Generate,
			Roles:   onlyGuest,
		},
	}
}
