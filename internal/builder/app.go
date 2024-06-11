package builder

import (
	"github.com/bloomingbug/depublic/internal/http/handler"
	"github.com/bloomingbug/depublic/internal/http/router"
	"github.com/bloomingbug/depublic/internal/repository"
	"github.com/bloomingbug/depublic/internal/services"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/bloomingbug/depublic/pkg/route"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, jwtToken jwt_token.JwtToken) []*route.Route {
	handlers := make(map[string]interface{})
	helloHandler := handler.NewHelloHandler()
	handlers["hello"] = &helloHandler
	return router.AppPublicRoutes(handlers)
}

func BuildAppPrivateRoutes(db *gorm.DB, redisDB *redis.Client) []*route.Route {
	handlers := make(map[string]interface{})

	otpRepository := repository.NewOneTimePasswordRepository(db)
	otpService := services.NewOneTimePasswordService(otpRepository)
	otpHandler := handler.NewOneTimePasswordHandler(otpService)
	handlers["otp"] = &otpHandler

	helloHandler := handler.NewHelloHandler()
	handlers["hello"] = &helloHandler
	return router.AppPrivateRoutes(handlers)
}
