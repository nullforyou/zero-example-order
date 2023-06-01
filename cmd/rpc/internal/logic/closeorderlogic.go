package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-zero-base/utils/xerr"
	"gorm.io/gorm"
	"greet-pb/order/types/order"
	"order/cmd/business"
	"order/cmd/dao/model"
	"order/cmd/dao/query"
	"order/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CloseOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCloseOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseOrderLogic {
	return &CloseOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CloseOrder 关闭订单
func (l *CloseOrderLogic) CloseOrder(in *order.CloseOrderReq) (*order.CloseOrderReply, error) {
	query.SetDefault(l.svcCtx.DbEngine)
	orderDao := query.Order
	orderModel, err := orderDao.WithContext(context.Background()).Where(orderDao.OrderSerialNumber.Eq(in.GetOrderSerialNumber())).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("订单不存在"))
	}
	if orderModel.OrderStatus == business.WAIT_PAYMENT_STATE {
		orderUpdate := model.Order{}
		orderUpdate.OrderStatus = business.IS_CLOSED
		orderUpdate.OrderStatusName = business.OrderStateZhCN(orderUpdate.OrderStatus)
		l.svcCtx.DbEngine.Model(orderModel).Updates(orderUpdate)
	}
	return &order.CloseOrderReply{OrderSerialNumber: orderModel.OrderSerialNumber}, nil
}
