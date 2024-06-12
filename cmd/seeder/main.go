package main

import (
	"github.com/bloomingbug/depublic/configs"
	"github.com/bloomingbug/depublic/db/seeds"
	"github.com/bloomingbug/depublic/pkg/postgres"
)

func main() {
	cfg, err := configs.NewConfig(".env")
	checkError(err)

	postgres, err := postgres.InitProgres(&cfg.Postgres)
	checkError(err)

	seeds.CreateUserSeeds(postgres)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
