package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/bloomingbug/depublic/internal/http/middleware"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/bloomingbug/depublic/pkg/route"
	"github.com/labstack/echo/v4"
)

type Server struct {
	*echo.Echo
}

func NewServer(publicRoutes, privateRoutes []*route.Route, secretKey string) *Server {
	e := echo.New()

	mw := middleware.NewMiddleware(secretKey)

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
			endpoint.Add(r.Method, r.Path, r.Handler, mw.For(r.Middleware))
		}
	}

	return &Server{e}
}

func (s *Server) Run() {
	runServer(s)
	gracefulShutdown(s)
}

func runServer(s *Server) {
	go func() {
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

	go func() {
		if err := s.Shutdown(ctx); err != nil {
			s.Logger.Fatal("Server Shutdown: ", err)
		}
	}()
}
