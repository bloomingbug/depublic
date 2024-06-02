package main

import (
	"github.com/bloomingbug/depublic/config"
)

func main() {
	_, err := config.NewConfig(".env")
	checkError(err)

	// postgres, err := postgres.InitProgres(&cfg.Postgres)
	// checkError(err)

	// redis := cache.InitCache(&cfg.Redis)

	// token := token.NewTokenUseCase(cfg.JWT.SecretKey)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}