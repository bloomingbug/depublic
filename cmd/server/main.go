package main

import (
	"github.com/bloomingbug/depublic/configs"
	"github.com/bloomingbug/depublic/internal/builder"
	"github.com/bloomingbug/depublic/pkg/cache"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/bloomingbug/depublic/pkg/postgres"
	"github.com/bloomingbug/depublic/pkg/scheduler"
	"github.com/bloomingbug/depublic/pkg/server"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	postgres, err := postgres.InitProgres(&cfg.Postgres)
	checkError(err)

	redis := cache.InitCache(&cfg.Redis)

	jwtToken := jwt_token.NewJwtToken(cfg.JWT.SecretKey)
	scheduler := scheduler.NewScheduler(redis, cfg.Namespace)

	publicRoutes := builder.BuildAppPublicRoutes(postgres, jwtToken)
	privateRoutes := builder.BuildAppPrivateRoutes(postgres, redis, scheduler)

	srv := server.NewServer(publicRoutes, privateRoutes, cfg.JWT.SecretKey)
	srv.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
