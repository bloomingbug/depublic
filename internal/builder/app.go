package builder

import (
	"github.com/bloomingbug/depublic/internal/http/handler"
	"github.com/bloomingbug/depublic/internal/http/router"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/bloomingbug/depublic/pkg/route"
	"github.com/bloomingbug/depublic/pkg/scheduler"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, jwtToken jwt_token.JwtToken, scheduler scheduler.Scheduler) []*route.Route {
	handlers := make(map[string]interface{})
	helloHandler := handler.NewHelloHandler()
	handlers["hello"] = &helloHandler

	otpRepository := repository.NewOneTimePasswordRepository(db)
	otpService := service.NewOneTimePasswordService(otpRepository, scheduler)
	otpHandler := handler.NewOneTimePasswordHandler(otpService)
	handlers["otp"] = &otpHandler

	tokenRepository := repository.NewTokenRepository(db)
	userRepository := repository.NewUserRepository(db)

	tokenService := service.NewTokenService(otpRepository, tokenRepository, userRepository, scheduler)
	tokenHandler := handler.NewTokenHandler(tokenService)
	handlers["token"] = &tokenHandler

	userService := service.NewUserService(tokenRepository, userRepository, jwtToken)
	userHandler := handler.NewUserHandler(userService)
	handlers["user"] = &userHandler

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)
	handlers["event"] = eventHandler

	return router.AppPublicRoutes(handlers)
}

func BuildAppPrivateRoutes(db *gorm.DB, redisDB *redis.Pool) []*route.Route {
	handlers := make(map[string]interface{})

	helloHandler := handler.NewHelloHandler()
	handlers["hello"] = &helloHandler

	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)
	handlers["event"] = &eventHandler

	return router.AppPrivateRoutes(handlers)
}

func BuildLocationHandler(db *gorm.DB) handler.LocationHandler {
	locationRepo := repository.NewLocationRepository(db)
	locationService := service.NewLocationService(locationRepo)
	locationHandler := handler.NewLocationHandler(locationService)
	return locationHandler
}
