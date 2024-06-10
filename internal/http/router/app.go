package router

import (
	"net/http"

	"github.com/bloomingbug/depublic/internal/http/handler"
	"github.com/bloomingbug/depublic/pkg/route"
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
			Method:     http.MethodGet,
			Path:       "/private",
			Handler:    h.Say,
			Middleware: "login",
		},
	}
}
