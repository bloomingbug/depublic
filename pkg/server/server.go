package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/bloomingbug/depublic/pkg/route"
	"github.com/bloomingbug/depublic/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
}

func NewServer(publicRoutes, privateRoutes []*route.Route, secretKey string) *Server {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.Success(http.StatusOK, "API Depublic", nil))
	})

	endpoint := e.Group("/api")

	if len(publicRoutes) > 0 {
		for _, r := range publicRoutes {
			endpoint.Add(r.Method, r.Path, r.Handler)
		}
	}

	if len(privateRoutes) > 0 {
		for _, r := range privateRoutes {
			endpoint.Add(r.Method, r.Path, r.Handler, JWTProtection(secretKey))
		}
	}

	return &Server{e}
}

func (s *Server) Run() {
	runServer(s)
	gracefulShutdown(s)
}

func runServer(s *Server) {
	go func () {
		err := s.Start(":8080")
		log.Fatal(err)
	}()
}

func gracefulShutdown(s *Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	go func(){
		if err := s.Shutdown(ctx); err != nil {
			s.Logger.Fatal("Server Shutdown: ", err)
		}
	}()
}

func JWTProtection(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, response.Error(http.StatusUnauthorized, "Login required to access this resource"))
		},
	})
}