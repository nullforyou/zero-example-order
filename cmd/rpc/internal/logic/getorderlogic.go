package logic

import (
	"context"
	"github.com/pkg/errors"
	"go-common/tool"
	"go-zero-base/utils/xerr"
	"gorm.io/gorm"
	"order/cmd/dao/model"

	"greet-pb/order/types/order"
	"order/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetOrder 查询订单
func (l *GetOrderLogic) GetOrder(in *order.GetOrderReq) (*order.GetOrderReply, error) {
	orderModel := model.Order{}
	err := l.svcCtx.DbEngine.Model(model.Order{}).Where("order_serial_number = ?", in.OrderSerialNumber).First(&orderModel).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorNotFound), xerr.SetMsg("订单不存在"))
	}

	orderOut := order.GetOrderReply{
		ID:                orderModel.ID,
		OrderSerialNumber: orderModel.OrderSerialNumber,
		MemberID:          orderModel.MemberID,
		Client:            orderModel.Client,
		OrderStatus:       orderModel.OrderStatus,
		OrderStatusName:   orderModel.OrderStatusName,
		OrderAmount:       tool.Float64ToString(orderModel.OrderAmount, 2),
		MemberNickname:    orderModel.MemberNickname,
		MemberMobile:      orderModel.MemberMobile,
		PaymentStatus:     orderModel.PaymentStatus,
		GoodsNum:          orderModel.GoodsNum,
		PaymentLimitTime:  orderModel.PaymentLimitTime.Unix(),
	}
	if orderModel.PaymentSn != "" {
		orderOut.PaymentAmount = tool.Float64ToString(orderModel.PaymentAmount, 2)
		orderOut.PaymentSn = orderModel.PaymentSn
		orderOut.PaymentType = orderModel.PaymentType
		orderOut.PaymentTime = orderModel.PaymentTime.Unix()
	}

	return &orderOut, nil
}
