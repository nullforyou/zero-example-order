package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hibiken/asynq"
	"github.com/shopspring/decimal"
	"go-zero-base/utils/xerr"
	"gorm.io/gorm"
	"greet-pb/user/types/user"
	"order/cmd/business"
	"order/cmd/dao/model"
	"order/cmd/dao/query"
	"order/cmd/jobtype"
	"strconv"
	"time"

	"greet-pb/order/types/order"
	"order/cmd/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order.CreateOrderReq) (*order.CreateOrderReply, error) {
	query.SetDefault(l.svcCtx.DbEngine)
	ctx := context.Background()
	orderDAL := query.Goods
	//所选商品的总价格
	goodsTotalAmount := decimal.NewFromFloat(0)

	orderModel := model.Order{}
	for _, createGoods := range in.OrderInfo.Goods {
		goodsModel, err := orderDAL.WithContext(ctx).Where(orderDAL.ID.Eq(createGoods.GoodsId)).First()
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("商品不存在"))
		}

		orderGoods := model.OrderGoods{
			CategoryID:   goodsModel.Category.CategoryID,
			CategoryName: goodsModel.Category.CategoryName,
			GoodsID:      goodsModel.ID,
			GoodsName:    goodsModel.GoodsName,
			GoodsPicture: goodsModel.GoodsPicture,
			GoodsPrice:   goodsModel.GoodsPrice,
			GoodsNum:     createGoods.Num,
		}
		orderModel.Goods = append(orderModel.Goods, orderGoods)

		goodsTotalAmount = goodsTotalAmount.Add(decimal.NewFromFloat(goodsModel.GoodsPrice).Mul(decimal.NewFromInt(createGoods.Num)))
	}

	//验证订单金额
	totalPrice, _ := decimal.NewFromString(in.TotalPrice)
	if !goodsTotalAmount.Equal(totalPrice) {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("订单价格错误"))
	}

	orderModel.MemberID = 1
	orderModel.Client = in.OrderInfo.Client
	orderModel.OrderStatus = business.WAIT_PAYMENT_STATE
	orderModel.OrderStatusName = business.OrderStateZhCN(orderModel.OrderStatus)
	orderModel.OrderAmount, _ = strconv.ParseFloat(goodsTotalAmount.String(), 64)
	orderModel.OrderSerialNumber = business.GenerateOrderNumber(time.Now())

	// 使用user rpc
	userInfo, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdReq{
		Id: in.UserId,
	})
	orderModel.MemberMobile = userInfo.Mobile
	orderModel.MemberNickname = userInfo.Nikename

	limitTime, _ := time.ParseDuration(business.PAYMENT_LIMIT_TIME)
	paymentLimitTime := time.Now().Add(limitTime)
	orderModel.PaymentLimitTime = &paymentLimitTime

	orderDetail := model.OrderDetail{}

	appointmentStartTime, _ := time.ParseInLocation(business.YYMMDDHHMM, in.OrderInfo.AppointmentStartTime, time.Local)
	orderDetail.AppointmentStartTime = appointmentStartTime

	appointmentEndTime, _ := time.ParseInLocation(business.YYMMDDHHMM, in.OrderInfo.AppointmentEndTime, time.Local)
	orderDetail.AppointmentEndTime = appointmentEndTime

	orderDetail.Remark = in.OrderInfo.Remark

	addressDal := query.Address
	senderAddressModel, err := addressDal.WithContext(ctx).Where(addressDal.ID.Eq(in.OrderInfo.SenderAddressId)).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("发货地址不存在"))
	}
	orderDetail.SenderName = senderAddressModel.ContactName
	orderDetail.SenderMobile = senderAddressModel.ContactMobile
	orderDetail.SenderProvince = senderAddressModel.ProvinceName
	orderDetail.SenderCity = senderAddressModel.CityName
	orderDetail.SenderCounty = senderAddressModel.CountyName
	orderDetail.SenderAddress = senderAddressModel.DetailedAddress

	receiveAddressModel, err := addressDal.WithContext(ctx).Where(addressDal.ID.Eq(in.OrderInfo.ReceiveAddressId)).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, xerr.NewBusinessError(xerr.SetCode(xerr.ErrorBusiness), xerr.SetMsg("收货地址不存在"))
	}
	orderDetail.ReceiveName = receiveAddressModel.ContactName
	orderDetail.ReceiveMobile = receiveAddressModel.ContactMobile
	orderDetail.ReceiveProvince = receiveAddressModel.ProvinceName
	orderDetail.ReceiveCity = receiveAddressModel.CityName
	orderDetail.ReceiveCounty = receiveAddressModel.CountyName
	orderDetail.ReceiveAddress = receiveAddressModel.DetailedAddress
	orderDetail.PackFee = 0
	orderDetail.IsPacking = 0

	orderModel.OrderDetail = orderDetail

	result := l.svcCtx.DbEngine.Create(&orderModel)

	//创建延迟关闭订单任务

	payload, err := json.Marshal(jobtype.DeferCloseOrderPayload{OrderSerialNumber: orderModel.OrderSerialNumber})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("创建演示订单任务的载荷时错误:%+v, sn:%s", err, orderModel.OrderSerialNumber)
	}
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseOrderJob, payload), asynq.ProcessAt(paymentLimitTime))
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseOrderJob, payload))
	if err != nil {
		logx.WithContext(l.ctx).Errorf("创建延时关闭订单任务时错误:%+v, sn:%s", err, orderModel.OrderSerialNumber)
	}

	return &order.CreateOrderReply{OrderSerialNumber: orderModel.OrderSerialNumber}, result.Error
}
