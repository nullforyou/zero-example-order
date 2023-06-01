package order

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-common/tool"
	orderRpc "greet-pb/order/types/order"
	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderReq) (resp *types.CreateOrderResp, err error) {
	userId, err := tool.GetUidFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	rpcCreateOrderReq := orderRpc.CreateOrderReq{UserId: userId, TotalPrice: tool.Float64ToString(req.TotalPrice, 2)}
	rpcCreateOrderInfo := orderRpc.CreateOrderInfo{
		Client:               req.OrderInfo.Client,
		AppointmentStartTime: req.OrderInfo.AppointmentStartTime,
		AppointmentEndTime:   req.OrderInfo.AppointmentEndTime,
		SenderAddressId:      req.OrderInfo.SenderAddressId,
		ReceiveAddressId:     req.OrderInfo.ReceiveAddressId,
		Remark:               req.OrderInfo.Remark,
	}

	for _, orderGoodsReq := range req.OrderInfo.Goods {
		orderGoods := orderRpc.CreateOrderGoods{
			GoodsId:   orderGoodsReq.GoodsId,
			GoodsName: orderGoodsReq.GoodsName,
			Num:       orderGoodsReq.Num,
		}
		rpcCreateOrderInfo.Goods = append(rpcCreateOrderInfo.Goods, &orderGoods)
	}
	rpcCreateOrderReq.OrderInfo = &rpcCreateOrderInfo
	// 使用order rpc
	order, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &rpcCreateOrderReq)

	if err != nil {
		return nil, err
	}
	return &types.CreateOrderResp{OrderSerialNumber: order.OrderSerialNumber}, nil
}
