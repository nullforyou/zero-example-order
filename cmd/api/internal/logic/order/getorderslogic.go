package order

import (
	"context"
	"order/cmd/dao/model"

	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrdersLogic {
	return &GetOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrdersLogic) GetOrders(req *types.OrdersCollectionReq) (resp *[]types.OrdersCollectionResp, err error) {
	err = l.svcCtx.DbEngine.Model(model.Order{}).Preload("OrderDetail").Limit(req.PageSize).Offset((req.Page - 1) * req.PageSize).Find(&resp).Error
	return resp, nil
}
