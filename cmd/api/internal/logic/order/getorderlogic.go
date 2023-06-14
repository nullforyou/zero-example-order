package order

import (
	"context"
	"greet-pb/order/orderclient"
	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"
	"order/cmd/business"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderItemReq) (resp *types.OrderItemResp, err error) {

	order, err := l.svcCtx.OrderRpc.GetOrder(l.ctx, &orderclient.GetOrderReq{OrderSerialNumber: req.OrderSerialNumber})

	if err != nil {
		return nil, err
	}

	formatTimeStr := time.Unix(order.PaymentLimitTime, 0).Format(business.YYMMDDHHMMSS)

	return &types.OrderItemResp{
		OrderSerialNumber: order.OrderSerialNumber,
		OrderStatus:       order.OrderStatus,
		OrderAmount:       order.OrderAmount,
		GoodsNum:          order.GoodsNum,
		PaymentLimitTime:  formatTimeStr,
	}, nil
}
