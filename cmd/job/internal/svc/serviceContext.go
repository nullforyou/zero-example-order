package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
	"greet-pb/order/orderclient"
	"order/cmd/job/internal/config"
)

type ServiceContext struct {
	Config config.Config
	AsynqServer *asynq.Server
	OrderRpc orderclient.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		AsynqServer: newAsynqServer(c),
		OrderRpc: orderclient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}