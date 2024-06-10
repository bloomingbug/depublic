package builder

import (
	"github.com/bloomingbug/depublic/internal/http/handler"
	"github.com/bloomingbug/depublic/internal/http/router"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/bloomingbug/depublic/pkg/route"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, jwtToken jwt_token.JwtToken) []*route.Route {
	return router.AppPublicRoutes(handler.NewHelloHandler())
}

func BuildAppPrivateRoutes(db *gorm.DB, redisDB *redis.Client) []*route.Route {
	return router.AppPrivateRoutes(handler.NewHelloHandler())
}
