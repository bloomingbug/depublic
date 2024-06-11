package pkg

import (
	"log"

	"github.com/bloomingbug/depublic/configs"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type scheduler struct {
	rdb *redis.Pool
	cfg configs.NamespaceConfig
}

func (s *scheduler) RunScheduler(tasks []entity.Task) {
	var enqueuer = work.NewEnqueuer(s.cfg.Namespace, s.rdb)

	for _, t := range tasks {
		_, err := enqueuer.Enqueue(t.TaskName, work.Q{"email_address": t.Email, "user_id": t.UserID})
		if err != nil {
			log.Fatal(err)
		}
	}
}

type Scheduler interface {
	RunScheduler(tasks []entity.Task)
}

func NewScheduler(rdb *redis.Pool, cfg configs.NamespaceConfig) Scheduler {
	return &scheduler{
		rdb: rdb,
		cfg: cfg,
	}
}
