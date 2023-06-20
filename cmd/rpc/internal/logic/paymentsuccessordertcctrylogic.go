package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmcli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"greet-pb/order/types/order"
	"order/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentSuccessOrderTccTryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaymentSuccessOrderTccTryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccessOrderTccTryLogic {
	return &PaymentSuccessOrderTccTryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PaymentSuccessOrderTccTry 订单支付成功TccTry
func (l *PaymentSuccessOrderTccTryLogic) PaymentSuccessOrderTccTry(in *order.PaymentSuccessTccReq) (*order.PaymentSuccessTccReply, error) {
	// todo: add your logic here and delete this line
	logx.WithContext(l.ctx).Info("进入 PaymentSuccessOrderTccTry")
	return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	return &order.PaymentSuccessTccReply{}, nil
}
