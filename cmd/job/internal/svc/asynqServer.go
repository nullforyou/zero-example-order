package svc

import (
	"github.com/hibiken/asynq"
	"order/cmd/job/internal/config"
)

func newAsynqServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr:c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			IsFailure: func(err error) bool {
				//fmt.Printf("== >> asynq服务器执行任务失败 err : %+v \n", err)
				return true
			},
			Concurrency: 20, //最大并发进程任务数
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)
}