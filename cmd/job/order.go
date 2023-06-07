package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"order/cmd/job/internal/logic"
	"order/cmd/job/internal/svc"
	"os"

	"order/cmd/job/internal/config"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/order-job.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	logx.DisableStat()

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("== >> [%s] 运行作业错误 err:%+v", c.Name, err)
		os.Exit(1)
	}
}
