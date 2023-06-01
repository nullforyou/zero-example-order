package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-common/utils/response"
	"go-common/utils/xerr"
	"google.golang.org/grpc/status"
	"net/http"
	"reflect"

	"order/cmd/api/internal/config"
	"order/cmd/api/internal/handler"
	"order/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/orders-api-dev.yaml", "the config file")

func main() {
	flag.Parse()

	logx.DisableStat()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

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
			causeErr := errors.Cause(err)
			if grpcStatus, ok := status.FromError(causeErr); ok { // grpc err错误

				grpcStatus.WithDetails()

				grpcCode := uint32(grpcStatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
				return http.StatusInternalServerError, nil
			} else {
				return http.StatusInternalServerError, nil
			}
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
