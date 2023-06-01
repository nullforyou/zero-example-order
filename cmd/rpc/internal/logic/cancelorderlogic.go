package logic

import (
	"context"
	"errors"
	"go-common/utils/xerr"
	"gorm.io/gorm"
	"order/cmd/business"
	"order/cmd/dao/model"
	"order/cmd/dao/query"
	"time"

	"greet-pb/order/types/order"
	"order/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CancelOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelOrderLogic {
	return &CancelOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CancelOrder 取消订单
func (l *CancelOrderLogic) CancelOrder(in *order.CancelOrderReq) (*order.CancelOrderReply, error) {
	query.SetDefault(l.svcCtx.DbEngine)
	orderDao := query.Order
	orderModel, err := orderDao.WithContext(context.Background()).Where(orderDao.OrderSerialNumber.Eq(in.GetOrderSerialNumber())).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode("ErrorOrderNotExists"))
	}
	if orderModel.OrderStatus == business.WAIT_PAYMENT_STATE {
		updateOrder := model.Order{}
		updateOrder.OrderStatus = business.CANCELLED_STATE
		updateOrder.OrderStatusName = business.OrderStateZhCN(updateOrder.OrderStatus)
		updateOrder.CancelOperator = &in.CancelOperator
		updateOrder.CancelCause = &in.CancelCause
		nowTime := time.Now()
		updateOrder.CancelTime = &nowTime
		l.svcCtx.DbEngine.Model(orderModel).Updates(updateOrder)
	}

	return &order.CancelOrderReply{OrderSerialNumber: orderModel.OrderSerialNumber}, nil
}
