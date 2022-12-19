package task

import (
	"WakaTImeGo/config"
	"fmt"
	"github.com/hibiken/asynq"
	log "github.com/sirupsen/logrus"
)

/**
 * @Author: Kim
 * @Description:
 * @File:  start
 * @Date: 12/13/2022 8:28 PM
 */

var taskClient *asynq.Client

func InitTaskService() {
	redisHost, redisPasswd, redisPort, redisDb := config.GetRedisConfig()

	taskServer := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: redisHost + ":" + fmt.Sprint(redisPort),
			// Omit if no password is required
			Password: redisPasswd,
			// Use a dedicated db number for asynq.
			// By default, Redis offers 16 databases (0..15)
			DB: redisDb,
		},

		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 1,
			// Optionally specify multiple queues with different priority.
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			// See the godoc for other configuration options
		},
	)

	// mux maps a type to a handler
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeDurationDelivery, HandleDurationDeliveryTask)

	taskClient = asynq.NewClient(asynq.RedisClientOpt{
		Addr: redisHost + ":" + fmt.Sprint(redisPort),
		// Omit if no password is required
		Password: redisPasswd,
		// Use a dedicated db number for asynq.
		// By default, Redis offers 16 databases (0..15)
		DB: redisDb,
	})

	if err := taskServer.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

func GetTaskClient() *asynq.Client {
	return taskClient
}
