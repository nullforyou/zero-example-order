package logic

import (
	"context"

	"greet-pb/order/types/order"
	"order/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentSuccessOrderTccConfirmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaymentSuccessOrderTccConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccessOrderTccConfirmLogic {
	return &PaymentSuccessOrderTccConfirmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// PaymentSuccessOrderTccConfirm 订单支付成功TccConfirm
func (l *PaymentSuccessOrderTccConfirmLogic) PaymentSuccessOrderTccConfirm(in *order.PaymentSuccessTccReq) (*order.PaymentSuccessTccReply, error) {
	// todo: add your logic here and delete this line
	logx.WithContext(l.ctx).Info("进入 PaymentSuccessOrderTccConfirm")
	return &order.PaymentSuccessTccReply{}, nil
}
