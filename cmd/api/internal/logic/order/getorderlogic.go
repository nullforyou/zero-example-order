package order

import (
	"context"
	"errors"
	"go-zero-base/utils/xerr"
	"gorm.io/gorm"
	"order/cmd/dao/model"

	"order/cmd/api/internal/svc"
	"order/cmd/api/internal/types"

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
	err = l.svcCtx.DbEngine.Model(model.Order{}).Where("order_serial_number = ?", req.OrderSerialNumber).First(&resp).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorNotFound), xerr.SetMsg("订单不存在"))
	}
	return resp, nil
}
