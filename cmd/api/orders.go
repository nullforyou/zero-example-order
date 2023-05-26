package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-common/utils/response"
	"go-common/utils/xerr"
	"net/http"
	"reflect"

	"order/cmd/api/internal/config"
	"order/cmd/api/internal/handler"
	"order/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/orders-api.yaml", "the config file")

func main() {
	flag.Parse()

	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		logx.Debugf("进入错误处理程序",  reflect.TypeOf(err))
		switch e := err.(type) {
		case *xerr.BusinessError:
			return http.StatusBadRequest, &response.Body{Code: e.GetErrCode(), Message: e.GetErrMsg()}
		default:
			fmt.Println("type:unkown,value:", err)
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
