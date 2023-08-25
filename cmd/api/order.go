package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-base/utils/response"
	"go-zero-base/utils/xerr"
	"net/http"
	"order/cmd/api/internal/config"
	"order/cmd/api/internal/handler"
	"order/cmd/api/internal/svc"
)

var configFile = flag.String("f", "etc/order-api.yaml", "the config file")

func main() {
	flag.Parse()

	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		//JWT验证失败自定义处理
		response.Response(r, w, nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorTokenExpire)))
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
