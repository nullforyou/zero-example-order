package logic

import (
	"context"

	"greet-pb/order/types/order"
	"order/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentSuccessOrderTccCancelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaymentSuccessOrderTccCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccessOrderTccCancelLogic {
	return &PaymentSuccessOrderTccCancelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PaymentSuccessOrderTccCancel 订单支付成功TccCancel
func (l *PaymentSuccessOrderTccCancelLogic) PaymentSuccessOrderTccCancel(in *order.PaymentSuccessTccReq) (*order.PaymentSuccessTccReply, error) {
	// todo: add your logic here and delete this line
	logx.WithContext(l.ctx).Info("进入 PaymentSuccessOrderTccCancel")
	return &order.PaymentSuccessTccReply{}, nil
}
