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

type DeleteOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteOrder 删除订单
func (l *DeleteOrderLogic) DeleteOrder(in *order.DeleteOrderReq) (*order.DeleteOrderReply, error) {

	query.SetDefault(l.svcCtx.DbEngine)
	orderDao := query.Order
	orderModel, err := orderDao.WithContext(context.Background()).Where(orderDao.OrderSerialNumber.Eq(in.GetOrderSerialNumber())).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode("ErrorOrderNotExists"))
	}
	if orderModel.OrderStatus == business.CANCELLED_STATE || orderModel.OrderStatus == business.IS_CLOSED || orderModel.OrderStatus == business.SETTLED_STATE {
		orderUpdate := model.Order{}
		orderUpdate.IsUserDelete = business.IS_SURE
		nowTime := time.Now()
		orderUpdate.UserDeleteTime = &nowTime
		l.svcCtx.DbEngine.Model(orderModel).Updates(orderUpdate)
	}

	return &order.DeleteOrderReply{}, nil
}
