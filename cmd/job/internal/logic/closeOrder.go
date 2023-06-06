package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-base/utils/xerr"
	"google.golang.org/grpc/status"
	"greet-pb/order/types/order"
	"order/cmd/job/internal/svc"
	"order/cmd/jobtype"
)

type CloseOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseOrderHandler(svcCtx *svc.ServiceContext) *CloseOrderHandler {
	return &CloseOrderHandler{
		svcCtx:svcCtx,
	}
}

func (l *CloseOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	logx.WithContext(ctx).Debugf("执行【%s】", jobtype.DeferCloseOrderJob)
	var closeOrderPayload jobtype.DeferCloseOrderPayload
	if err := json.Unmarshal(t.Payload(), &closeOrderPayload); err != nil {
		logx.Errorf("执行【%s】时解析有效载荷时发生错误 payload:%s", jobtype.DeferCloseOrderJob, t.Payload())
		return nil
	}
	_, err := l.svcCtx.OrderRpc.CloseOrder(ctx, &order.CloseOrderReq{OrderSerialNumber: closeOrderPayload.OrderSerialNumber})
	if err == nil {
		logx.WithContext(ctx).Debugf("执行【%s】完成", jobtype.DeferCloseOrderJob)
	} else {
		causeErr := errors.Cause(err)
		if grpcStatus, ok := status.FromError(causeErr); ok { //grpc错误
			grpcCode := uint32(grpcStatus.Code())
			if grpcCode >= xerr.ErrorBusiness {
				logx.WithContext(ctx).Errorf("执行Job[%s]时rpc服务业务错误 err:%+v", jobtype.DeferCloseOrderJob, err)
				return nil
			} else {
				logx.WithContext(ctx).Errorf("执行Job[%s]时rpc服务不可预知错误 err:%+v", jobtype.DeferCloseOrderJob, err)
				return err
			}
		} else {
			//其他错误
			logx.WithContext(ctx).Errorf("执行【%s】时发生不可预知错误 err:%+v", jobtype.DeferCloseOrderJob, err)
			return err
		}
	}
	return nil
}
